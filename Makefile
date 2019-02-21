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


.PHONY: all clean statik dev