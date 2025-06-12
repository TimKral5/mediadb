package auth

import (
	"fmt"

	"github.com/go-ldap/ldap/v3"
)

type LDAPConnection struct {
	conn    *ldap.Conn
	baseDN string
	bindDN string
	bindPw  string
}

func NewLDAPConnection(url string, baseDN string, bindDN string, bindPw string) (*LDAPConnection, error) {
	conn, err := ldap.DialURL(url)

	if err != nil {
		return nil, err
	}

	ctx := &LDAPConnection{
		conn,
		baseDN,
		bindDN,
		bindPw,
	}

	return ctx, nil
}

func (self *LDAPConnection) Disconnect() {
	self.conn.Close()
}

func (self *LDAPConnection) getUserDN(username string) (string, error) {
	err := self.conn.Bind(self.bindDN, self.bindPw)

	if err != nil {
		return "", err
	}

	req := ldap.NewSearchRequest(
		self.baseDN,
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
	userDN, err := self.getUserDN(creds.Username)
	
	if err != nil {
		return false, err
	}

	err = self.conn.Bind(userDN, creds.Password)
	
	if err != nil {
		return false, err
	}

	return true, nil
}

func (self *LDAPConnection) GetPermissions(creds Credentials) (int, error) {
	req := ldap.NewSearchRequest(
		self.baseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0, 0, false,
		fmt.Sprintf("(uid=%s)", ldap.EscapeFilter(creds.Username)),
		[]string{"dn", "cn"},
		nil,
	)

	res, err := self.conn.Search(req)

	if err != nil {
		return 0, err
	}

	return len(res.Entries), nil
}

