
create:
	cp dbconfig.yml.template dbconfig.yml
build:
	env GOOS=windows env GOARCH=amd64 go build -o bin/main.exe main.go
	env GOOS=windows env GOARCH=amd64 go build -o bin/createDevice.exe cmd/createDevice.go
run:
	go run main.go
drop:
	go run cmd/dropDB.go

migrate:
	go run cmd/migrate.go

docker/push:
	make -f .circleci/ci.mk go/build
	make -f .circleci/ci.mk docker/build
	make -f .circleci/ci.mk docker/push
