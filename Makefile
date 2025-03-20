DOCKER := docker
DCOMPOSE := docker-compose

help:
	@echo 'Avaiblable Commands:'
	@echo '   run           Sets up the docker-compose environment'
	@echo '   rm            Removes all containers'
	@echo '   reset         Resets all containers'
	@echo '   start         Starts up all containers'
	@echo '   stop          Stops all containers'
	@echo '   rebuild       Resets the environment excluding grafana'
	@echo '   kube-apply    Applies the Kubernetes configuration'
	@echo '   kube-delete   Cleans up the Kubenetes environment'

run:
	${DCOMPOSE} up -d
rm:
	${DCOMPOSE} down
reset: rm run

start:
	-${DOCKER} start mdb-mdb-1
	-${DOCKER} start mdb-mongodb-1
	-${DOCKER} start mdb-prometheus-1
	-${DOCKER} start mdb-loki-1
	-${DOCKER} start mdb-grafana-1

stop:
	-${DOCKER} stop mdb-mdb-1
	-${DOCKER} stop mdb-mongodb-1
	-${DOCKER} stop mdb-prometheus-1
	-${DOCKER} stop mdb-loki-1
	-${DOCKER} stop mdb-grafana-1

rebuild:
	-${DOCKER} stop mdb-mdb-1 && ${DOCKER} rm mdb-mdb-1
	-${DOCKER} stop mdb-mongodb-1 && ${DOCKER} rm mdb-mongodb-1
	-${DOCKER} stop mdb-prometheus-1 && ${DOCKER} rm mdb-prometheus-1
	-${DOCKER} stop mdb-loki-1 && ${DOCKER} rm mdb-loki-1
	-${DOCKER} rmi mdb:latest
	${DCOMPOSE} up -d

kube-apply:
	kubectl apply -f deployment.yaml
kube-delete:
	kubectl delete -f deployment.yaml
