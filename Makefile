HASH:=$(shell git rev-parse --short HEAD)
APPNAME=theapp

dockerbuild:
	docker build -t $(APPNAME):latest .

dockerrun: dockerbuild
	docker run --rm -it -p 8001:8001 -t theapp:latest

dockerpush:
	echo "$(DOCKER_PASSWORD)" | docker login -u "$(DOCKER_USERNAME)" --password-stdin
	docker tag $(APPNAME) $(USER_NAME)/$(APPNAME):latest
	docker tag $(APPNAME) $(USER_NAME)/$(APPNAME):$(HASH)
	docker push $(USER_NAME)/$(APPNAME):latest
	docker push $(USER_NAME)/$(APPNAME):$(SHA)