
# Wiselink Events

El proyecto es una aplicacion para administrar eventos, y usuarios que se subscriben a estos, el proyecto lo dividi en un back desarrollado con go y el front hecho con react, ambos proyectos los conecte bajo una misma red de contenedores docker, en el que se incluyen todas sus dependencias para un despliege facil. 
Los detalles y tecnologias usadas a continuacion:


# Back (Go)


Por parte del backend, realize una arquitecura por capas o "clean arquitecture", en el que dividi mi aplicacion en tres capaz:  
## Domain

- Es donde estableci toda la logica de negocia, y requerimientos que se me solicitaron en la prueba
- Aqui desarrolle las entidades que se ajustan a resolver el dominio del problema
- Tambien estableci las interfaces de los casos de uso que luego se implementaran en un capa exterior
- Y por ultimo los dtos (Data Transfer Objects), que son que me definen la interfaz por las cual comunicarme con mi frontend, diviendolo en input y output

## Application

- Es donde implemente los casos de uso, definidos en mi domain, tambien
- Defini las interfaces de mis repositorios, hice uso del patron repository, para evitar cualquier tipo de dependencia con la base de datos.
- Tambien hice mappers, para mapear mis entidades a dtos.

## Infrastructure

- Aqui es donde las dependencias del proyecto se inyectan, y se implementan todas aquellas abstracciones realizadas en las capas inferiores, mediante el principio de inversion de control o "Ioc"
- Tambien mapeo mis entidades del domain a entidades de las bases de datos, ya que en una arquitectura, se considera a la base de datos solo un detalle, que no es relevante al dominio. Hago uso de gorm como ORM.
- Todos los enpoint de mi API rest los coloque aqui (rest/), ya que la forma de consumir la aplicacion puede variar, (grapql, etc), y por ultimo la configuracion de mis variables de entorno, db, etc tambien estan ubicados en esta capa.
- Para la seguridad de la aplicacion use JWT, en el que dependiendo del endpoint, se accedia al recurso solo si contaba con los roles y permision requeridos.

## Extras

- En la raiz de mi proyecto se encuentra un script "wait-for-it.sh", que es el encargado de demorar el deploy de mi backend, hasta que mi base de datos este completamente inicializada, ya que la directiva depends-on, de mi docker-compose no logra sincronizarlos correctamente.

- Tambien se encuentra un directorio tests/ donde realize test unitarios de las partes fundamentales de la aplicacion.


# Front (React)

La aplicacion frontend fue desarrolada con ReactJS y TailwindCSS para estilar, tambien hice uso RxJS, para manejar eventos entre distintos componentes, y Context Provider para manejar el estado global de la aplicacion (usuarios, etc), para deployar , hago uso de gninx




## Como ejecutar en local

Clona el proyecto

```bash
  git clone https://github.com/franBarrientos/Eventos-Wiselink
```

Situate en el el

```bash
  cd Eventos-Wiselink
```

Ejecuta los contenedores

```bash
  docker compose up -d
```
Visita el siguiente link (app front)
```bash
  http://localhost:5175/
```



## 🔗 Documentacion de la API en Postman

Para probar los endpoints, el archivo de inicializacion de la base de datos, ya inserta dos usuarios, con sus respectivos roles

### Usuario con role ADMIN
```bash
  email:  admin@admin.com
  password  123
```


### Usuario con role USER
```bash
  email:  user@user.com
  password  123
```

[![portfolio](https://res.cloudinary.com/postman/image/upload/t_team_logo/v1629869194/team//w_10,h_10/2893aede23f01bfcbd2319326bc96a6ed0524eba759745ed6d73405a3a8b67a8)](https://documenter.getpostman.com/view/25276650/2s9YkhiQLY)



## FAQ
### Posibles errores
#### Question 1
####  /usr/bin/env:"nash\r': No such file or directory 

Para solucionar este error, solo seteamos git con la siguiente configuracion, 
``
git config --global core.autocrlf false
``

#### Question 2
####  Error al levantar el contenedor el back por primera vez

Solo nos ubicamos en el directorio del proyecto, donde se encuentra el docker-compose, y ejecutamos 
``
    docker compose start
``
este error se debe a que la base de datos aun no esta lista cuando se levanta el back.



