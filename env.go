package main

import (
	"mediadb/auth"
	"os"
)

type Environment struct {
	LdapConfig  *auth.LDAPConfig
	HttpAddress string
}

type configError struct {
	str string
}

func (self *configError) Error() string {
	return self.str
}

func loadLdapConfig() (*auth.LDAPConfig, error) {
	ldapAddress, isDefined := os.LookupEnv("MEDIADB_LDAP_ADDRESS")
	if !isDefined {
		return nil, &configError{
			"Environment variable undefined (MEDIADB_LDAP_ADDRESS)",
		}
	}

	ldapBaseDN, isDefined := os.LookupEnv("MEDIADB_BASE_DN")
	if !isDefined {
		return nil, &configError{
			"Environment variable undefined (MEDIADB_BASE_DN)",
		}
	}

	ldapBindDN, isDefined := os.LookupEnv("MEDIADB_BIND_DN")
	if !isDefined {
		return nil, &configError{
			"Environment variable undefined (MEDIADB_BIND_DN)",
		}
	}

	ldapBindPassword, isDefined := os.LookupEnv("MEDIADB_BIND_PASSWORD")
	if !isDefined {
		return nil, &configError{
			"Environment variable undefined (MEDIADB_BIND_PASSWORD)",
		}
	}

	ldapGroupDN, isDefined := os.LookupEnv("MEDIADB_GROUP_DN")
	if !isDefined {
		return nil, &configError{
			"Environment variable undefined (MEDIADB_GROUP_DN)",
		}
	}

	ldapUserDN, isDefined := os.LookupEnv("MEDIADB_USER_DN")
	if !isDefined {
		return nil, &configError{
			"Environment variable undefined (MEDIADB_USER_DN)",
		}
	}

	config := &auth.LDAPConfig{
		ServerURL: ldapAddress,
		BaseDN:    ldapBaseDN,
		BindDN:    ldapBindDN,
		BindPw:    ldapBindPassword,
		GroupDN:   ldapUserDN,
		UserDN:    ldapGroupDN,
	}

	return config, nil
}

func LoadEnvironment() (*Environment, error) {
	config, err := loadLdapConfig()

	if err != nil {
		return nil, err
	}

	httpAddress, isDefined := os.LookupEnv("MEDIADB_HTTP_ADDRESS")
	if !isDefined {
		return nil, &configError{
			"Environment variable undefined (MEDIADB_HTTP_ADDRESS)",
		}
	}

	env := &Environment{
		LdapConfig:  config,
		HttpAddress: httpAddress,
	}

	return env, nil
}
