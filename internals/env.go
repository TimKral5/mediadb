package internals

import (
	"context"
	"mediadb/auth"
	"mediadb/db"
	"os"
	"time"
)

type Environment struct {
	LdapConfig  *auth.LDAPConfig
	MongoConfig *db.MongoConfig
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	config := &db.MongoConfig{
		Addr: mongoUrl,
		Context: ctx,
		CancelContext: cancel,
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

	httpAddress, isDefined := os.LookupEnv("MEDIADB_HTTP_ADDRESS")
	if !isDefined {
		return nil, &configError{
			"Environment variable undefined (MEDIADB_HTTP_ADDRESS)",
		}
	}

	env := &Environment{
		LdapConfig:  ldapConfig,
		MongoConfig:  mongoConfig,
		HttpAddress: httpAddress,
	}

	return env, nil
}
