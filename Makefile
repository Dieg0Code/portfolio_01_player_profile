# Nombre del contenedor
CONTAINER_NAME=test_postgres_db

# Nombre de la imagen de Docker
IMAGE_NAME=postgres:latest

# Puerto en el que se expondrá la base de datos
PORT=5432

# Variables de entorno para la base de datos
POSTGRES_USER=test_user
POSTGRES_PASSWORD=test_password
POSTGRES_DB=test_db

# Objetivo para iniciar la base de datos de pruebas
start_db:
	docker run --name $(CONTAINER_NAME) -e POSTGRES_USER=$(POSTGRES_USER) -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -e POSTGRES_DB=$(POSTGRES_DB) -p $(PORT):5432 -d $(IMAGE_NAME)

# Objetivo para detener y eliminar el contenedor
stop_db:
	docker stop $(CONTAINER_NAME)
	docker rm $(CONTAINER_NAME)

# Objetivo para ver los logs del contenedor
logs_db:
	docker logs $(CONTAINER_NAME)
