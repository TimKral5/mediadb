package internals

import (
	"mediadb/auth"
	"mediadb/db"
	"mediadb/utils"
)

type Program struct {
	ldapConn  *auth.LDAPConnection
	mongoConn *db.MongoConnection
	log       utils.Logger
	env       *Environment
	documentation string
}

func NewProgram() *Program {
	prog := &Program{}
	return prog
}

func (self *Program) AttachLogger(log utils.Logger) *Program {
	self.log = log
	return self
}

func (self *Program) AttachEnvironment(env *Environment) *Program {
	self.env = env
	return self
}

func (self *Program) AttachDocumentation(doc string) *Program {
	self.documentation = doc
	return self
}

func (self *Program) ConnectToLdap() error {
	conn, err := auth.NewLDAPConnection(self.env.LdapConfig)
	self.ldapConn = conn
	return err
}

func (self *Program) ConnectToMongo() error {
	conn, err := db.NewMongoConnection(self.env.MongoConfig)
	self.mongoConn = conn

	if err != nil {
		return err
	}

	err = conn.Ping()
	return err
}

func (self *Program) InitializeLdap() error {
	i := auth.NewLdapInitializer(self.log)
	err := i.InitializeLdapGroups(self.ldapConn)
	return err
}

func (self *Program) LaunchHttpServer() {
	server := NewHttpServer(self, self.env.HttpConfig)
	server.LaunchHttpServer()
}
