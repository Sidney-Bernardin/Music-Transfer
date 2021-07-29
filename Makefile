docker:
	sudo docker build --build-arg API_KEY=${API_KEY} -t musictransfer .
	sudo docker run -p ${PORT}:8080 -it musictransfer .
