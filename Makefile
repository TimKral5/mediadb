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
	-${DOCKER} rmi localhost/mdb:latest
	${DCOMPOSE} up -d

stop:
	-${DOCKER} stop mdb_mdb_1
	-${DOCKER} stop mdb_mongodb_1
	-${DOCKER} stop mdb_prometheus_1
	-${DOCKER} stop mdb_loki_1
	-${DOCKER} stop mdb_grafana_1

rebuild:
	-${DOCKER} stop mdb_mdb_1 && ${DOCKER} rm mdb_mdb_1
	-${DOCKER} stop mdb_prometheus_1 && ${DOCKER} rm mdb_prometheus_1
	-${DOCKER} stop mdb_loki_1 && ${DOCKER} rm mdb_loki_1
	-${DOCKER} rmi localhost/mdb:latest
	${DCOMPOSE} up -d

kube-apply:
	kubectl apply -f deployment.yaml
kube-delete:
	kubectl delete -f deployment.yaml
