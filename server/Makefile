.PHONY: init help run
DOMAIN=$(domain)

CURRENT_TIME= $(shell date +"%Y%m%d%H%M%S")

init:
	@echo "Creating an rsa 512 key pair"
	openssl genpkey -algorithm RSA -out ./auth/private.key -pkeyopt rsa_keygen_bits:2048
	openssl rsa -in ./auth/private.key -pubout -out ./shared/auth/public.key

help:
	@echo "usage example:"
	@echo " make run-server domain=xxx"

# example: make run domain=account
run:
	go run ./$(DOMAIN)/main.go

runall: 
	go run ./auth/main.go \
	| go run ./account/main.go \
	| go run ./page/main.go \
	| go run ./blob/main.go \
	| go run ./gateway/main.go

