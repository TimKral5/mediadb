package main

import (
	"mediadb/auth"
	"mediadb/utils"
)

type Program struct {
	ldapConn *auth.LDAPConnection
	log utils.Logger
	env *Environment
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

func (self *Program) ConnectToLdap() error {
	conn, err := auth.NewLDAPConnection(self.env.LdapConfig)
	self.ldapConn = conn
	return err
}

func (self *Program) InitializeLdap() error {
	i := auth.NewLdapInitializer(self.log)
	err := i.InitializeLdapGroups(self.ldapConn)
	return err
}

