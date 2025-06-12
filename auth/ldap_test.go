package auth_test

import (
	"mediadb/auth"
	"testing"
)

var ldapUrl = "ldap://localhost:389"
var ldapBaseDN = "dc=example,dc=org"
var ldapBindDN = "cn=admin,dc=example,dc=org"
var ldapBindPw = "admin"

func TestValidateLogin(t *testing.T) {
	conn, err := auth.NewLDAPConnection(
		ldapUrl,
		ldapBaseDN,
		ldapBindDN,
		ldapBindPw,
	)

	if err != nil {
		t.Error(err)
		return
	}

	demoCreds := auth.Credentials {
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

func TestGetPermissions(t *testing.T) {
	conn, err := auth.NewLDAPConnection(
		ldapUrl,
		ldapBaseDN,
		ldapBindDN,
		ldapBindPw,
	)

	if err != nil {
		t.Error(err)
		return
	}

	demoCreds := auth.Credentials {
		Username: "demo",
		Password: "demo",
	}

	res, err := conn.ValidateLogin(demoCreds)

	if err != nil {
		t.Error(err)
		return
	}

	count, err := conn.GetPermissions(demoCreds)
	t.Log(count)

	if err != nil {
		t.Error(err)
		return
	}

	if !res {
		t.Error("Validation failed")
		return
	}
}

