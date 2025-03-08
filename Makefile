run: build
	@./bin/api

build:
	@go build -o bin/api cmd/api/main.go

genkey:
	@openssl genpkey -algorithm RSA -out keys/private.pem -pkeyopt rsa_keygen_bits:2048 2>/dev/null
	@openssl rsa -pubout -in keys/private.pem -out keys/public.pem 2>/dev/null
	@echo "Keys generated"

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  run       Run the application"
	@echo "  build     Build the application"
	@echo "  genkey		 Generate private and public key"
	@echo "  help      Display this help message"

.DEFAULT_GOAL := help