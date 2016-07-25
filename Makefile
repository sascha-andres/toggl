.PHONY: build install

build:
	CGO_ENABLED=0 GOOS=linux go build -a

docker: build
	docker build --no-cache -t briefbote/toggl:latest .
	docker push briefbote/toggl:latest
