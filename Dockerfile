# Primera Etapa: Compilación de la aplicación
FROM golang:1.22.1 AS builder

# Etiqueta del Mantenedor
LABEL maintainer="Diego"

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos go.mod y go.sum y descargar las dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente desde el directorio actual al directorio de trabajo dentro del contenedor
COPY main.go .
COPY src/ src/

# Compilar la aplicación Go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


# Segunda Etapa: Imagen Final
FROM alpine:3.16

# Instalar certificados CA para permitir conexiones HTTPS
RUN apk --no-cache add ca-certificates

# Agregar un usuario no-root
RUN adduser -S -D -H -h /app appuser

# Cambiar al usuario no-root
USER appuser

# Establecer el directorio de trabajo como /app para el usuario no-root
WORKDIR /app

# Copiar el binario pre-compilado desde la primera etapa (builder) al directorio actual (/app) de la segunda etapa
COPY --from=builder /app/main .

# Exponer el puerto 8080 para el mundo exterior
EXPOSE 8080

# Comando para ejecutar el binario
CMD ["./main"]
