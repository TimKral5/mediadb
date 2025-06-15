package auth_test

import (
	"mediadb/auth"
	"mediadb/internals"
	"testing"
)

var initialized = false
var ldapConfig *auth.LDAPConfig

func getConfig(t *testing.T) {
	if initialized {
		return
	}

	env, err := internals.LoadEnvironment()

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	ldapConfig = env.LdapConfig
}

func TestValidateLogin(t *testing.T) {
	getConfig(t)
	conn, err := auth.NewLDAPConnection(ldapConfig)

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
	getConfig(t)
	conn, err := auth.NewLDAPConnection(ldapConfig)

	if err != nil {
		t.Error(err)
		return
	}

	isMember, err := conn.IsUserGroupMember("demo", "test")

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
	getConfig(t)
	conn, err := auth.NewLDAPConnection(ldapConfig)

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
