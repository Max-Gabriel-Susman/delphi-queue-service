mod:
	go mod tidy 
	go mod vendor 

build:
	docker build --tag delphi-model-service .

# docker build --tag brometheus/delphi-training-service:v0.1.0 .
# docker build --tag brometheus/delphi-training-service:v0.2.0 .

run: 
	docker run \
		-p 50054:50054 \
		brometheus/delphi-training-service:v0.0.0

push: 
	docker push brometheus/delphi-model-service:tagname

update:
	docker build --tag brometheus/delphi-training-service:v0.0.0 .
	docker push brometheus/delphi-training-service:v0.0.0

# docker push brometheus/delphi-training-service:v0.1.0

## https://medium.com/@shradhasehgal/get-started-with-grpc-in-c-36f1f39367f4
## actually let's do this one: https://googlecloudrobotics.github.io/core/how-to/deploying-grpc-service.html
## actually let's try cloning this and aping it if it works: https://github.com/npclaudiu/grpc-cpp-docker

## run the c++ code 
# 
# $ mkdir -p cmake/build
# $ cd cmake/build
# $ cmake ../..


pre-build: # resolves failed to solve with frontend dockerfile.v0: failed to read dockerfile: error from sender: resolve : lstat server: no such file or directory
	export DOCKER_BUILDKIT=0
	export COMPOSE_DOCKER_CLI_BUILD=0