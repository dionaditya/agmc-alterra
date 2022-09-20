test:
	docker-compose -f docker-compose.yaml down 
	docker-compose -f docker-compose.test.yaml up --build --abort-on-container-exit
	docker-compose -f docker-compose.test.yaml down --volumes

run:
	docker-compose -f docker-compose.dev.yaml down 
	docker-compose -f docker-compose.dev.yaml up

build:
	docker-compose -f docker-compose.dev.yaml down 
	docker-compose -f docker-compose.dev.yaml build

build-prod:
	docker-compose -f docker-compose.yaml down 
	docker-compose -f docker-compose.yaml build


run-prod:
	docker-compose -f docker-compose.yaml up