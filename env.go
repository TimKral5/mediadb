package main

import (
	"mediadb/auth"
	"mediadb/utils"
	"os"
)

type Environment struct {
	LdapConfig  *auth.LDAPConfig
	HttpAddress string
}

func loadLdapConfig(log utils.Logger) *auth.LDAPConfig {
	hasError := false

	ldapAddress, isDefined := os.LookupEnv("MEDIADB_LDAP_ADDRESS")
	if !isDefined {
		log.Error(
			"Environment variable undefined",
			"(MEDIADB_BIND_ADDRESS)",
		)

		hasError = true
	}

	ldapBaseDN, isDefined := os.LookupEnv("MEDIADB_BASE_DN")
	if !isDefined {
		log.Error(
			"Environment variable undefined",
			"(MEDIADB_BASE_DN)",
		)
		hasError = true
	}

	ldapBindDN, isDefined := os.LookupEnv("MEDIADB_BIND_DN")
	if !isDefined {
		log.Error(
			"Environment variable undefined",
			"(MEDIADB_BIND_DN)",
		)
		hasError = true
	}

	ldapBindPassword, isDefined := os.LookupEnv("MEDIADB_BIND_PASSWORD")
	if !isDefined {
		log.Error(
			"Environment variable undefined",
			"(MEDIADB_BIND_PASSWORD)",
		)
		hasError = true
	}

	ldapGroupDN, isDefined := os.LookupEnv("MEDIADB_GROUP_DN")
	if !isDefined {
		log.Error(
			"Environment variable undefined",
			"(MEDIADB_GROUP_DN)",
		)
		hasError = true
	}

	ldapUserDN, isDefined := os.LookupEnv("MEDIADB_USER_DN")
	if !isDefined {
		log.Error(
			"Environment variable undefined",
			"(MEDIADB_USER_DN)",
		)
		hasError = true
	}

	if hasError {
		return nil
	}

	config := &auth.LDAPConfig{
		ServerURL: ldapAddress,
		BaseDN:    ldapBaseDN,
		BindDN:    ldapBindDN,
		BindPw:    ldapBindPassword,
		GroupDN:   ldapUserDN,
		UserDN:    ldapGroupDN,
	}

	return config
}

func LoadEnvironment(log utils.Logger) *Environment {
	config := loadLdapConfig(log)

	if config == nil {
		return nil
	}

	httpAddress, isDefined := os.LookupEnv("MEDIADB_HTTP_ADDRESS")
	if !isDefined {
		log.Error(
			"Environment variable undefined",
			"(MEDIADB_HTTP_ADDRESS)",
		)
	}

	env := &Environment{
		LdapConfig:  config,
		HttpAddress: httpAddress,
	}

	return env
}
