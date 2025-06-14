package auth_test

import (
	"mediadb/auth"
	"testing"
)

var ldapConfig = auth.LDAPConfig{
	ServerURL: "ldap://localhost:389",
	BaseDN:    "dc=example,dc=org",
	BindDN:    "cn=admin,dc=example,dc=org",
	BindPw:    "admin",
	GroupDN:   "ou=groups,dc=example,dc=org",
	UserDN:    "ou=users,dc=example,dc=org",
}

func TestValidateLogin(t *testing.T) {
	conn, err := auth.NewLDAPConnection(&ldapConfig)

	if err != nil {
		t.Error(err)
		return
	}

	demoCreds := auth.Credentials{
		Username: "demo",
		Password: "demo",
	}

	res, err := conn.ValidateLogin(demoCreds)

	if err != nil {
		t.Error(err)
		return
	}

	if !res {
		t.Error("Validation failed")
		return
	}
}

func TestIsUserGroupMember(t *testing.T) {
	conn, err := auth.NewLDAPConnection(&ldapConfig)

	if err != nil {
		t.Error(err)
		return
	}

	isMember, err := conn.IsUserGroupMember("demo", "test")
	t.Log(isMember)

	if err != nil {
		t.Error(err)
		return
	}

	if !isMember {
		t.Error("User is not member of group")
		return
	}
}

func TestCreateGroup(t *testing.T) {
	conn, err := auth.NewLDAPConnection(&ldapConfig)

	if err != nil {
		t.Error(err)
		return
	}

	err = conn.CreateGroup("mediadb_test_group")

	if err != nil {
		t.Error(err)
		return
	}

}
