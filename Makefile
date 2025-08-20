help:
	echo "make c					runs client program"
	echo "make s					runs server program"


c:
	echo "Running Client"
	./znp-cs -r client

s:
	echo "Running Server"
	make build
	./znp-cs -r server

m:
	make build
	./znp-cs -r m


build:
	go build .
