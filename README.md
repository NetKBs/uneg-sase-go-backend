
# Backend - Pagina De Servicios Estudiantiles

Este repositorio representa la parte de backend de nuestro proyecto, que tiene como objetivo reemplazar el actual sitio web de servicios estudiantiles de la Universidad Nacional Experimental de Guayana, que no cumple con las necesidades de los estudiantes.


## Tecnologías

* Go
* Gin
* GORM
* SQLite (para desarrollo)


## API Documentación
Puedes encontrar la documentación de los endpoints generada usando swagger a través de la siguiente ruta:
```
  /swagger/index.html
```
## Variables De Entorno

Para ejecutar este proyecto, necesitarás agregar las siguientes variables de entorno en tu /config/.env

`PORT_SERVER` 

`SECRET_KEY`
## Ejecución

Clona el proyecto

```bash
  git clone https://github.com/NetKBs/uneg-sase-go-backend
```

Accede a la ruta del proyecto

```bash
  cd uneg-sase-go-backend
```

Crea las variables de entorno en tu .env

```bash
  mkdir config 
  touch ./config/.env
```

Instala las dependencias

```bash
  go mod download
```

Inicia el servidor
```bash
  go run main.go
```



