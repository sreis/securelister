codegen:
	apiserver-boot build generated

build:
	GOOS=linux GOARCH=amd64 apiserver-boot build executables --generate=false

container: build
	docker build -t sreis/securelister:latest .

kind.load-image:
	kind load docker-image sreis/securelister:latest
