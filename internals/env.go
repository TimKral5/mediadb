package internals

import (
	"mediadb/auth"
	"mediadb/db"
	"os"
)

type Environment struct {
	LdapConfig  *auth.LDAPConfig
	MongoConfig *db.MongoConfig
	HttpConfig  *HttpConfig
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
		GroupDN:   ldapGroupDN,
		UserDN:    ldapUserDN,
	}

	return config, nil
}

func loadMongoConfig() (*db.MongoConfig, error) {
	mongoUrl, isDefined := os.LookupEnv("MEDIADB_MONGO_URL")
	if !isDefined {
		return nil, &configError{
			"Environment variable undefined (MEDIADB_MONGO_URL)",
		}
	}

	config := &db.MongoConfig{
		Addr: mongoUrl,
	}

	return config, nil
}

func loadHttpConfig() (*HttpConfig, error) {
	httpAddress, isDefined := os.LookupEnv("MEDIADB_HTTP_ADDRESS")
	if !isDefined {
		return nil, &configError{
			"Environment variable undefined (MEDIADB_HTTP_ADDRESS)",
		}
	}

	config := &HttpConfig{
		Addr: httpAddress,
	}

	return config, nil
}

func LoadEnvironment() (*Environment, error) {
	ldapConfig, err := loadLdapConfig()

	if err != nil {
		return nil, err
	}

	mongoConfig, err := loadMongoConfig()

	if err != nil {
		return nil, err
	}

	httpConfig, err := loadHttpConfig()

	if err != nil {
		return nil, err
	}

	env := &Environment{
		LdapConfig:  ldapConfig,
		MongoConfig: mongoConfig,
		HttpConfig:  httpConfig,
	}

	return env, nil
}
