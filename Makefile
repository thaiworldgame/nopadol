SERVICE=n9erp
REGISTRY=gcr.io/mrtomyum
COMMIT_SHA=$(shell git rev-parse HEAD)

default:
	# `make deploy` build and deploy to production

dev:
	go run main.go

clean:
	rm -f n9erp

build:
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o n9erp -a -ldflags '-w -s' main.go

docker: clean build
	docker build -t $(REGISTRY)/$(SERVICE):$(COMMIT_SHA) .
	docker push $(REGISTRY)/$(SERVICE):$(COMMIT_SHA)

test-api:
	cotton -u http://localhost:9999 -d tests/cotton/