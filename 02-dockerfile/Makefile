export tag=v2.1
root:
	export ROOT=github.com/SimonXu666j/Cloud-Native/02-dockerfile

build:
	echo "building httpserver container"
	docker build -t simomxu/httpserver:${tag} .

push: build
	echo "pushing simomxu/httpserver"
	docker push simomxu/httpserver:${tag}

run_test: push
	echo "running httpserver container"
	docker run -d -p80:80 --name httpserver simomxu/httpserver:${tag}

check_container_ip: run_test
	echo "use nsenter command to check httpserver container ip"
	nsenter -t `docker inspect -f '{{.State.Pid}}' httpserver` -n ip a