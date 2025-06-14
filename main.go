package main

import (
	"mediadb/utils"
	"net/http"
)

func main() {
	log := utils.NewLogger()
	ctx := http.NewServeMux()
	log.Info("Launching MediaDB v0.1.0-alpha...")

	env, err := LoadEnvironment()

	if err != nil {
		log.Error(err)
		return
	}

	prog := NewProgram().
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
	LaunchHttpServer(log, ctx)
}

