DEPLOY_TIME := $(shell date -u +"%Y-%m-%dT%H-%M-%SZ_%s")
HASH:=$(shell git rev-parse --short HEAD)
APPNAME=theapp
REPO_NAME=jsburckhardt/ubiquitous-fortnight


dockerbuild:
	docker build -t $(APPNAME):latest .

dockerrun: dockerbuild
	docker run --rm -it -p 8001:8001 -t theapp:latest

tag:
	git tag "$(DEPLOY_TIME)_$(HASH)"
	git push https://${GH_TOKEN}@github.com/$(REPO_NAME) $(DEPLOY_TIME)_$(HASH)

dockerpush:
	echo "$(DOCKER_PASSWORD)" | docker login -u "$(DOCKER_USERNAME)" --password-stdin
	docker tag $(APPNAME) $(DOCKER_USERNAME)/$(APPNAME):latest
	docker tag $(APPNAME) $(DOCKER_USERNAME)/$(APPNAME):$(HASH)
	docker push $(DOCKER_USERNAME)/$(APPNAME):latest
	docker push $(DOCKER_USERNAME)/$(APPNAME):$(HASH)

#deploy: test dockerbuild docker tag dockerpush 