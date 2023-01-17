# dbml-go

DBML-GO is a Go parser for DBML syntax. It can parse a DBML file and generate Go structs based on the table schema described in the DBML file.

The original project does not feature a suitable docker image. That means you need to install it on your PC. 

## Using the Docker Image

You can use the provided Docker image to run the dbml-gen-go-model command without having to install Go or the package dependencies.

docker run -it --rm -v $(PWD)/example:/app/data gitea.linuxcode.net/linuxcode/dbml-go ./dbml-go-generator -f ./data/schema.dbml -p model -o ./data/model

This command will run the container using the image 
```gitea.linuxcode.net/linuxcode/dbml-go```
and mount the example directory on host to the /app/data directory in the container. 

The command ```./dbml-go-generator -f ./data/schema.dbml -p model -o ./data/model`` will be executed inside the container, using the specified input parameters.

## Makefile

There's a makefile that contains useful commands to build and run the container

```make build```

This command will build the container using the Dockerfile in the current directory.

```make example```

This command will run the example command in the readme, it will build the container and run the command to generate the model.