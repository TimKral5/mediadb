help:
	@echo 'Avaiblable Commands:'
	@echo '   build     Removes and then builds the container'
	@echo '   launch    Builds and then launches the containers'
	@echo '   stop      Stops the running containers'

build:
	-podman rmi localhost/mdb:latest
	podman build -t mdb ./src

launch:
	-podman rmi localhost/mdb:latest
	podman-compose up -d

stop:
	-podman-compose down

relaunch: stop launch
