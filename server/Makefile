DOMAIN=$(domain)

CURRENT_TIME= $(shell date +"%Y%m%d%H%M%S")

help:
	@echo "usage example:"
	@echo " make run-server domain=xxx"

# example: make run domain=account
run:
	go run ./$(DOMAIN)/main.go