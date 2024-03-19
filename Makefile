.PHONY:build
build:
	docker build -t cron-web .

.PHONY:sh
sh:
	docker run \
	-v ./example.yml:/app/example.yml \
	-v ./hello.sh:/app/hello.sh \
	-it --rm \
	cron-web \
	/bin/sh


.PHONY:run
run:
	docker run \
    	-v ./example.yml:/app/example.yml \
    	-v ./hello.sh:/app/hello.sh \
    	-it --rm -d \
    	cron-web \
    	./main
