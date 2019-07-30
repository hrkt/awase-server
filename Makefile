NAME     := awase
VERSION  := v0.0.1
REVISION := $(shell git rev-parse --short HEAD)

SRCS    := $(shell find . -type f -name '*.go')
LDFLAGS := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\" -extldflags \"-static\""

LABEL=$(NAME):latest


all: bin/$(NAME)

# dev prepare
devprepare:
	go get -u github.com/swaggo/swag/cmd/swag

# api
swaginit:
	swag init

depinit:
	dep init

depensure:
	dep ensure

bin/$(NAME): $(SRCS)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -a -tags $(NAME) -installsuffix $(NAME) $(LDFLAGS) -o bin/$(NAME)

run: 
	go run main.go models.go repository.go

release-run: bin/$(NAME)
	GIN_MODE=release ./bin/$(NAME)

check:
	curl http://localhost:8080/api/greeting
	echo ""

fmt:
	go fmt *.go

test:
	go test

# container

build-container:
	sudo docker build -t $(LABEL) --build-arg version=$(VERSION) --build-arg revision=$(REVISION) .

run-container:
	sudo docker run -p 8080:8080 $(LABEL)

clean-container:
	sudo docker image prune
	sudo docker container prune

check-container:
	curl http://localhost:8080/api/greeting

push-container:
	sudo docker push $(LABEL) 192.168.3.223:5443/$(LABEL)