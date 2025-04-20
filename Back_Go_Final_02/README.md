# Proyecto Go Backend - Persona y Propiedad

Este proyecto es un backend simple en Go que gestiona personas y propiedades, documentado con Swagger y listo para ejecutar con Docker y Redis.

## 📦 Estructura del Proyecto

- `cmd/main.go`: punto de entrada principal
- `controller/`: controladores HTTP
- `model/entities/`: modelos de datos
- `services/`: interfaces de servicio
- `services/implementation/`: implementación en memoria
- `routes/`: rutas de API
- `docs/`: documentación Swagger generada
- `Dockerfile`: configuración de la imagen backend
- `docker-compose.yml`: para levantar el backend y Redis

---

## 🚀 Requisitos Previos

- [Go](https://golang.org/dl/) >= 1.21
- [Docker](https://www.docker.com/)
- [swag](https://github.com/swaggo/swag)

Instalar swag (si no lo tienes):

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

---

## 🛠️ Configuración Local

### 1. Clona o descomprime el proyecto
```bash
cd go-backend-fixed
```

### 2. Genera Swagger Docs
```bash
swag init -g cmd/main.go --parseDependency --parseInternal
```

### 3. Ejecuta localmente (opcional)
```bash
go run cmd/main.go
```

Abre en navegador: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

---

## 🐳 Despliegue con Docker

### 1. Build de la imagen
```bash
docker build -t go-backend .
```

### 2. Run manual (sin Redis)
```bash
docker run -p 8080:8080 go-backend
```

---

## 📦 Despliegue con Docker Compose (incluye Redis)

### 1. Ejecutar todo con:
```bash
docker-compose up --build
```

Esto levanta el backend en `http://localhost:8080` y Redis en `localhost:6379`.

---

## 🧠 Notas
- Redis no está integrado aún en la lógica del servicio, pero el contenedor está preparado para que lo uses.
- Puedes modificar `services/implementation` para usar Redis en lugar de slices en memoria.
