codegen:
	apiserver-boot build generated

build:
	apiserver-boot build executables --generate=false

container: build
	docker build -t sreis/securelister:latest .
