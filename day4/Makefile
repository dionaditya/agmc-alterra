test:
	docker-compose -f docker-compose.test.yaml up --build --abort-on-container-exit
	docker-compose -f docker-compose.test.yaml down --volumes

run:
	docker-compose -f docker-compose.yaml up

build:
	docker-compose -f docker-compose.yaml down 
	docker-compose -f docker-compose.yaml build