all:
	cd web && yarn && yarn build && cd ..
	statik -src=./web/build 
	go build -o app main.go

.ONESHELL:
statik:
	cd web && yarn && yarn build && cd ..
	statik -src=./web/build 

dev:
	watcher -startcmd -cmd="go run main.go"

clean:
	rm app

setup:
	go get github.com/rakyll/statik
	go get github.com/canthefason/go-watcher
	go get github.com/mitchellh/gox


crossplatform:
	gox -output="build/{{.Dir}}_{{.OS}}_{{.Arch}}"


.PHONY: all clean statik dev crossplatform