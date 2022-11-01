first-run:
	docker-compose build && docker-compose up -d && docker-compose up techtask


run:
	docker-compose up techtask

