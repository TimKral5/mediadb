package auth

import (
	"fmt"

	"github.com/go-ldap/ldap/v3"
	"slices"
)

type LDAPConfig struct {
	BaseDN  string
	BindDN  string
	BindPw  string
	GroupDN string
	UserDN  string
}

type LDAPConnection struct {
	conn   *ldap.Conn
	config LDAPConfig
}

func NewLDAPConnection(url string, config LDAPConfig) (*LDAPConnection, error) {
	conn, err := ldap.DialURL(url)

	if err != nil {
		return nil, err
	}

	ctx := &LDAPConnection{
		conn,
		config,
	}

	return ctx, nil
}

func (self *LDAPConnection) Disconnect() {
	self.conn.Close()
}

func (self *LDAPConnection) GetUserDN(username string) (string, error) {
	err := self.conn.Bind(self.config.BindDN, self.config.BindPw)

	if err != nil {
		return "", err
	}

	req := ldap.NewSearchRequest(
		self.config.UserDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0, 0, false,
		fmt.Sprintf("(uid=%s)", ldap.EscapeFilter(username)),
		[]string{"dn"},
		nil,
	)

	res, err := self.conn.Search(req)

	if err != nil {
		return "", err
	}

	if len(res.Entries) != 1 {
		return "", fmt.Errorf("User does not exist or multiple entries")
	}

	return res.Entries[0].DN, nil
}

func (self *LDAPConnection) ValidateLogin(creds Credentials) (bool, error) {
	userDN, err := self.GetUserDN(creds.Username)

	if err != nil {
		return false, err
	}

	err = self.conn.Bind(userDN, creds.Password)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (self *LDAPConnection) GetGroupMembers(group string) ([]string, error) {
	err := self.conn.Bind(self.config.BindDN, self.config.BindPw)

	if err != nil {
		return nil, err
	}

	req := ldap.NewSearchRequest(
		self.config.GroupDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0, 0, false,
		fmt.Sprintf("(cn=%s)", ldap.EscapeFilter(group)),
		[]string{"member", "dn"},
		nil,
	)

	res, err := self.conn.Search(req)

	if err != nil {
		return nil, err
	}

	if len(res.Entries) != 1 {
		return nil, fmt.Errorf("No or multiple entries found")
	}

	members := res.Entries[0].GetAttributeValues("member")
	return members, nil
}

func (self *LDAPConnection) IsUserGroupMember(username string, group string) (bool, error) {
	userDN, err := self.GetUserDN(username)

	if err != nil {
		return false, err
	}

	members, err := self.GetGroupMembers(group)

	if err != nil {
		return false, err
	}

	if slices.Contains(members, userDN) {
		return true, nil
	}

	return false, nil
}

func (self *LDAPConnection) GroupExists(group string) (bool, error) {
	err := self.conn.Bind(self.config.BindDN, self.config.BindPw)

	if err != nil {
		return false, err
	}

	req := ldap.NewSearchRequest(
		self.config.GroupDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0, 0, false,
		fmt.Sprintf("(cn=%s)", ldap.EscapeFilter(group)),
		[]string{"dn"},
		nil,
	)

	res, err := self.conn.Search(req)

	if err != nil {
		return false, err
	}

	if len(res.Entries) != 1 {
		return false, nil
	}

	return true, nil
}

func (self *LDAPConnection) CreateGroup(group string) error {
	newDN := fmt.Sprintf("cn=%s,%s", ldap.EscapeDN(group), self.config.GroupDN)
	err := self.conn.Bind(self.config.BindDN, self.config.BindPw)

	if err != nil {
		return err
	}	

	req := ldap.NewAddRequest(newDN, nil)
	req.Attribute("objectClass", []string{"groupOfNames"})
	req.Attribute("member", []string{ self.config.BindDN })

	err = self.conn.Add(req)

	if err != nil {
		return err
	}

	return nil
}

