$(V).SILENT:

help:
	echo "make client					runs client program"
	echo "make server					runs server program"

c:
	echo "Running client"
	make build -s
	./znp-cs -r client

s:
	echo "Running server"
	make build -s
	./znp-cs -r server

build:
	go build .
