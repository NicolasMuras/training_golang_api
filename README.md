<h2><a id="user-content-tabla-de-contenido" class="anchor" aria-hidden="true" href="#tabla-de-contenido"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M7.775 3.275a.75.75 0 001.06 1.06l1.25-1.25a2 2 0 112.83 2.83l-2.5 2.5a2 2 0 01-2.83 0 .75.75 0 00-1.06 1.06 3.5 3.5 0 004.95 0l2.5-2.5a3.5 3.5 0 00-4.95-4.95l-1.25 1.25zm-4.69 9.64a2 2 0 010-2.83l2.5-2.5a2 2 0 012.83 0 .75.75 0 001.06-1.06 3.5 3.5 0 00-4.95 0l-2.5 2.5a3.5 3.5 0 004.95 4.95l1.25-1.25a.75.75 0 00-1.06-1.06l-1.25 1.25a2 2 0 01-2.83 0z"></path></svg></a>Tabla de contenido
</h2>
<ul>
  <li><a href="#requerimientos">Requerimientos</a></li>
  <li><a href="#implementacion-del-proyecto">Implementación del proyecto</a></li>
  <li><a href="#iniciar-aplicacion">Iniciar aplicación</a></li>
</ul>

<h2><a id="user-content-implementacion-del-proyecto" class="anchor" aria-hidden="true" href="#implementacion-del-proyecto"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M7.775 3.275a.75.75 0 001.06 1.06l1.25-1.25a2 2 0 112.83 2.83l-2.5 2.5a2 2 0 01-2.83 0 .75.75 0 00-1.06 1.06 3.5 3.5 0 004.95 0l2.5-2.5a3.5 3.5 0 00-4.95-4.95l-1.25 1.25zm-4.69 9.64a2 2 0 010-2.83l2.5-2.5a2 2 0 012.83 0 .75.75 0 001.06-1.06 3.5 3.5 0 00-4.95 0l-2.5 2.5a3.5 3.5 0 004.95 4.95l1.25-1.25a.75.75 0 00-1.06-1.06l-1.25 1.25a2 2 0 01-2.83 0z"></path></svg></a>Implementación del proyecto</h2>
<ul>
<li><strong>JWT</strong>: Libreria para el sistema de tokens.</li>
<li><strong>pgx</strong>: Libreria para la gestionar las conexiones a la DB y manejo de pool.</li>
<li><strong>Ecosistema Docker (docker, Dockerfile, Docker-Compose)</strong>:  A partir del Dockerfile en el raíz del proyecto se puede compilar la imagen que corre la REST API hecha en Django REST Framework, con todas sus dependencias y código fuente dentro. Con Docker-Compose se puede ejecutar la aplicación con un único comando, creando además un servidor de base de datos PostgreSQL.</li>
<li><strong>PostgreSQL</strong>: Sistema de gestion de base de datos elegido.</li>
</ul>

Empece consumiendo las apis externas para obtener datos y con la construccion de estructuras de datos para su almacenamiento, una vez tuve esto, fui probando con PostgreSQL en local para verificar que los datos se guardaban y se leian correctamente.

Posteriormente implemente un sistema de logeo sencillo en el cual obtenemos un token que nos sirve para consumir nuestro endpoint, se realiza una validacion del token y se entrega el array de canciones si todo esta ok.

Al final, cuando todo estaba funcionando me propuse a configurar Docker, hice un script para inicializar la db y quedo.

<h2><a id="user-content-requerimientos" class="anchor" aria-hidden="true" href="#requerimientos"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M7.775 3.275a.75.75 0 001.06 1.06l1.25-1.25a2 2 0 112.83 2.83l-2.5 2.5a2 2 0 01-2.83 0 .75.75 0 00-1.06 1.06 3.5 3.5 0 004.95 0l2.5-2.5a3.5 3.5 0 00-4.95-4.95l-1.25 1.25zm-4.69 9.64a2 2 0 010-2.83l2.5-2.5a2 2 0 012.83 0 .75.75 0 001.06-1.06 3.5 3.5 0 00-4.95 0l-2.5 2.5a3.5 3.5 0 004.95 4.95l1.25-1.25a.75.75 0 00-1.06-1.06l-1.25 1.25a2 2 0 01-2.83 0z"></path></svg></a>Requerimientos</h2>

<p>Para correr este proyecto se necesitan las siguientes dependencias:</p>
<ul>
<li><a href="https://docs.docker.com/get-docker/" rel="nofollow">Docker</a>.</li>
<li><a href="https://docs.docker.com/compose/install/" rel="nofollow">Docker compose</a>.</li>
</ul>
<p>Una vez instaladas las dependencias se puede correr la aplicación.</p>

<h2><a id="user-content-iniciar-aplicacion" class="anchor" aria-hidden="true" href="#iniciar-aplicacion"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M7.775 3.275a.75.75 0 001.06 1.06l1.25-1.25a2 2 0 112.83 2.83l-2.5 2.5a2 2 0 01-2.83 0 .75.75 0 00-1.06 1.06 3.5 3.5 0 004.95 0l2.5-2.5a3.5 3.5 0 00-4.95-4.95l-1.25 1.25zm-4.69 9.64a2 2 0 010-2.83l2.5-2.5a2 2 0 012.83 0 .75.75 0 001.06-1.06 3.5 3.5 0 00-4.95 0l-2.5 2.5a3.5 3.5 0 004.95 4.95l1.25-1.25a.75.75 0 00-1.06-1.06l-1.25 1.25a2 2 0 01-2.83 0z"></path></svg></a>Iniciar aplicación</h2>

Creamos los contenedores:

<pre><code>sudo docker-compose build
</code></pre>

Corremos los contenedores:

<pre><code>sudo docker-compose up
</code></pre>

<h2><a id="user-content-endpoints" class="anchor" aria-hidden="true" href="#endpoints"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M7.775 3.275a.75.75 0 001.06 1.06l1.25-1.25a2 2 0 112.83 2.83l-2.5 2.5a2 2 0 01-2.83 0 .75.75 0 00-1.06 1.06 3.5 3.5 0 004.95 0l2.5-2.5a3.5 3.5 0 00-4.95-4.95l-1.25 1.25zm-4.69 9.64a2 2 0 010-2.83l2.5-2.5a2 2 0 012.83 0 .75.75 0 001.06-1.06 3.5 3.5 0 00-4.95 0l-2.5 2.5a3.5 3.5 0 004.95 4.95l1.25-1.25a.75.75 0 00-1.06-1.06l-1.25 1.25a2 2 0 01-2.83 0z"></path></svg></a>Endpoints</h2> 

### `/login?username=nico&password=123456`

Nos proporciona un token, el cual utilizaremos para tener acceso a nuestra API.

#### Método: `GET`


#### Output

```json
"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJuYmYiOjE2NjU0MDMyMDAsInVzZXIiOiJ1c2VybmFtZSJ9.7SBdbA92ptvw9ve66Zb7QmZMzq4g2ZoOPgLui9JrZcs"

```

### `/search?filter=pink+floyd`

Nos devuelve las canciones cuyo nombre, artista o album coincidan.

#### Método: `GET`


#### Output

```json
{
    {
        "Id": 1067444892,
        "Name": "Wish You Were Here",
        "Artist": "Pink Floyd",
        "Duration": 304200,
        "Album": "A Foot In the Door: The Best of Pink Floyd",
        "Artwork": "https://is4-ssl.mzstatic.com/image/thumb/Music125/v4/fe/54/24/fe542475-2143-490f-62ae-855dd7b47086/886445636093.jpg/30x30bb.jpg",
        "Price": 1.29,
        "Origin": "apple"
    },
    {
        "Id": 1067444896,
        "Name": "Comfortably Numb",
        "Artist": "Pink Floyd",
        "Duration": 379213,
        "Album": "A Foot In the Door: The Best of Pink Floyd",
        "Artwork": "https://is4-ssl.mzstatic.com/image/thumb/Music125/v4/fe/54/24/fe542475-2143-490f-62ae-855dd7b47086/886445636093.jpg/30x30bb.jpg",
        "Price": 1.29,
        "Origin": "apple"
    }
}
```



<em>
  Nota: La parte de NGINX no quedo bien configurada, se lo debo.
</em><br>
