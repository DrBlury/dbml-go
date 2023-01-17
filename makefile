IMAGE_NAME=dbml-go
PUBLIC_IMAGE_NAME=gitea.linuxcode.net/linuxcode/dbml-go

build:
	docker build -t $(IMAGE_NAME) .

example: build
	mkdir -p example/model
	docker run -it --rm -v $(PWD)/example:/app/data $(IMAGE_NAME) ./dbml-go-generator -f ./data/schema.dbml -p model -o ./data/model
