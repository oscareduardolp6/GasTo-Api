# Usa la imagen base del Dev Container
FROM mcr.microsoft.com/devcontainers/go:1-1.22-bookworm

# Establece el directorio de trabajo para módulos Go
WORKDIR /app

# Copia los archivos go.mod y go.sum primero
COPY go.mod go.sum ./

# Descarga las dependencias del proyecto
RUN go mod download

# Copia el resto de los archivos del proyecto
COPY . .

# Establece el directorio de trabajo para compilar
WORKDIR /app/src

# Ajusta permisos para todos los archivos (opcional, ajustar según sea necesario)
RUN chmod -R 755 /app

# Compila la aplicación
RUN go build -o /app/myapp main.go

EXPOSE 8080

# Establece el comando de inicio
CMD ["/app/myapp"]
