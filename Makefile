.PHONY: check clean test build deploy

test:
	go test -v ./...

clean:
	$(MAKE) -C connect clean
	$(MAKE) -C disconnect clean
	$(MAKE) -C publish clean

build: clean
	@echo "building handlers for aws lambda"
	sam build

build-ConnectFunction:
	@echo "building handler for aws lambda"
	$(MAKE) -C connect build

build-DisconnectFunction:
	@echo "building handler for aws lambda"
	$(MAKE) -C disconnect build

build-PublishFunction:
	@echo "building handler for aws lambda"
	$(MAKE) -C publish build

deploy:
	@echo "deploying infrastructure and code"
	sam deploy
