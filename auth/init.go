package auth

import "mediadb/utils"

type LDAPInitializer struct {
	log utils.Logger
}

func NewLdapInitializer(log utils.Logger) *LDAPInitializer {
	init := &LDAPInitializer{
		log,
	}

	return init
}

// Tries to create a group and logs the result
func (self *LDAPInitializer) createGroup(conn *LDAPConnection, group string) error {
	self.log.Info("Checking group '", group, "'...")
	exists, err := conn.GroupExists(group)

	if err != nil {
		return err
	}

	if exists {
		self.log.Info("... already exists, nothing to do.")
		return nil
	}

	err = conn.CreateGroup(group)

	if err != nil {
		self.log.Error("... failed to create.")
		return err
	}

	self.log.Info("... successfully created.")

	return nil
}

func (self *LDAPInitializer) initMovieGroups(conn *LDAPConnection) error {
	err := self.createGroup(conn, "mediadb_create_movie")

	if err != nil {
		return err
	}

	err = self.createGroup(conn, "mediadb_get_movie")

	if err != nil {
		return err
	}

	err = self.createGroup(conn, "mediadb_update_movie")

	if err != nil {
		return err
	}

	err = self.createGroup(conn, "mediadb_delete_movie")

	if err != nil {
		return err
	}

	return nil
}

func (self *LDAPInitializer) InitializeLdapGroups(conn *LDAPConnection) error {
	err := self.initMovieGroups(conn)

	if err != nil {
		return err
	}

	return nil
}
