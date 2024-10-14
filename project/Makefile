# BROKER_BINARY=brokerApp
BIO_DATA_BINARY=biodataApp
# PASS_BINARY= passApp
# MAIL_BINARY=mailerApp
# LOGGER_BINARY=loggerServiceApp


# ## up: starts all containers in the background without forcing build
# up:
# 	@echo "Starting Docker images..."
# 	docker-compose up -d
# 	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
# not runniing live in terminal (-d) 
# up_build_d: build_auth build_pass build_broker build_mail build_logger
# 	@echo "Stopping docker images (if running...)"
# 	docker-compose down
# 	@echo "Building (when required) and starting docker images..."
# 	docker-compose up --build -d
# 	@echo "Docker images built and started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_pass
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

# ## build_broker: builds the broker binary as a linux executable
# build_broker:
# 	@echo "Building broker binary..."
# 	cd ../broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
# 	@echo "Done!"

## build_auth: builds the auth binary as a linux executable
# build_auth:
# 	@echo "Building auth binary..."
# 	cd ../authentication-service && env GOOS=linux CGO_ENABLED=0 go build -o ${AUTH_BINARY} ./cmd/api
# 	@echo "Done!"

## build_pass: builds the auth binary as a linux executable
build_pass:
	@echo "Building auth binary..."
	cd ../biodata-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BIO_DATA_BINARY} ./cmd/api
	@echo "Done!"

# ## build_mail: builds the mail binary as a linux executable
# build_mail:
# 	@echo "Building mail binary..."
# 	cd ../mail-service && env GOOS=linux CGO_ENABLED=0 go build -o ${MAIL_BINARY} ./cmd/api
# 	@echo "Done!"

# ## build_logger: builds the logger binary as a linux executable
# build_logger:
# 	@echo "Building logger binary..."
# 	cd ../logger-service && env GOOS=linux CGO_ENABLED=0 go build -o ${LOGGER_BINARY} ./cmd/api
# 	@echo "Done!"
