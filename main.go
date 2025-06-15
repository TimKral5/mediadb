package main

import (
	"mediadb/internals"
	"mediadb/utils"
	"net/http"
)

func main() {
	log := utils.NewLogger()
	ctx := http.NewServeMux()
	log.Info("Launching MediaDB v0.1.0-alpha...")

	env, err := internals.LoadEnvironment()

	if err != nil {
		log.Error(err)
		return
	}

	prog := internals.NewProgram().
		AttachLogger(log).
		AttachEnvironment(env)

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

	log.Info("Launching HTTP server...")
	internals.LaunchHttpServer(log, ctx)
}

