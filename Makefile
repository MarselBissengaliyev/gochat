build:
	docker build -t gou-talk .

run:
	docker run --name=gou-talk -p 80:8080 gou-talk

run-terminal:
	docker exec -it gou-talk bash
