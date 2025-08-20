help:
	echo "make c					runs client program"
	echo "make s					runs server program"

c:
	echo "Running Client"
	./znp-cs -r client

s:
	echo "Running Server"
	./znp-cs -r server

build:
	go build .
