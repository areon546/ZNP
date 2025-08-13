help:
	echo "make client					runs client program"
	echo "make server					runs server program"

c:
	echo "Running Client"
	cd client && go run .

s:
	echo "Running Server"
	cd server && go run .
