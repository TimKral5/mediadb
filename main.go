package main

import (
	"mediadb/internals"
	"mediadb/utils"
	_ "embed"
)

//go:embed api-docs/html/index.html
var openapiDocumentation string

func main() {
	log := utils.NewLogger()
	log.Info("Launching MediaDB v0.1.0-alpha...")

	env, err := internals.LoadEnvironment()

	if err != nil {
		log.Error(err)
		return
	}

	prog := internals.NewProgram().
		AttachLogger(log).
		AttachEnvironment(env).
		AttachDocumentation(openapiDocumentation)

	log.Info("Connecting to LDAP...")
	err = prog.ConnectToLdap()

	if err != nil {
		log.Error(err)
		return
	} else {
		log.Info("Connected to LDAP")
	}

	log.Info("Initializing LDAP session...")
	err = prog.InitializeLdap()

	if err != nil {
		log.Error(err)
		return
	} else {
		log.Info("Initialized LDAP session")
	}

	log.Info("Connecting to MongoDB...")
	err = prog.ConnectToMongo()

	if err != nil {
		log.Error(err)
		return
	} else {
		log.Info("Connected to MongoDB")
	}

	log.Info("Launching HTTP server...")
	prog.LaunchHttpServer()
}

