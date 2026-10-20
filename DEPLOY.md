# Despliegue de business-manager

## URLs

- Local: `http://localhost:8083`
- Público: `https://business-manager.edoutput.dev`
- Tenant de ejemplo: `https://cliente.edoutput.dev`
- API interna: `http://business-manager-gateway:3001`
- Edge: `http://business-manager-edge:80`

## Componentes

- `business-manager-edge`: Nginx (público)
- `business-manager-web`: Next.js 14
- `business-manager-gateway`: API Gateway Go
- `business-manager-{auth,company,customer,order,product,file,whatsapp}-svc`: microservicios Go
- 5 PostgreSQL + 1 MongoDB
- 5 contenedores de migración (`migrate/migrate`)

## Prueba local

```bash
docker network inspect shared-edge >/dev/null 2>&1 || docker network create shared-edge
cp .env.example .env
# Edita todos los change-me y las credenciales de AWS/Twilio.
./deploy/bootstrap.sh

# Si ya tienes un .env real y solo quieres probar con el ejemplo:
APP_ENV_FILE=.env.example DATA_ROOT="$PWD/.local-data" \
  docker compose --env-file .env.example up -d --build
```

En el servidor se usa el `.env` real:

```bash
docker compose up -d --build
```

Verificar:

```bash
curl -f http://localhost:8083/healthz
curl -f http://localhost:8083/api/roles
```

## Servidor

```bash
sudo mkdir -p /opt/docker-projects /opt/docker-projects-data
sudo chown -R "$USER":"$USER" /opt/docker-projects /opt/docker-projects-data

git clone https://github.com/edorguez/business-manager.git /opt/docker-projects/business-manager
cd /opt/docker-projects/business-manager

cp .env.example .env
nano .env
./deploy/bootstrap.sh
docker compose up -d --build
```

## Variables críticas

- `BASE_DOMAIN=edoutput.dev`
- `TENANT_BASE_DOMAIN=edoutput.dev`
- `RESERVED_SUBDOMAINS=www,tourism-agency,gps-medical-assistance,business-manager`
- `JWT_SECRET_KEY`: usar un secreto largo y aleatorio.
- Las URLs de bases y servicios ya apuntan a los nombres `business-manager-*`.

> Si cambias una contraseña de base de datos, actualiza también su URL en
> `*_DB_SOURCE_DOCKER_CONTAINER`; si la contraseña tiene caracteres especiales,
> debe ir URL-encoded.
>
> Si reutilizas el `.env` antiguo del proyecto, actualiza los hosts internos:
> `auth-db` → `business-manager-auth-db`, `company-db` → `business-manager-company-db`,
> `customer-db` → `business-manager-customer-db`, `order-db` → `business-manager-order-db`,
> `whatsapp-db` → `business-manager-whatsapp-db`, `product-db` → `business-manager-product-db`,
> y los `*_SVC_URL` a `business-manager-*-svc`. Es más seguro partir de `.env.example`.

## Persistencia

```text
/opt/docker-projects-data/business-manager/postgres/auth
/opt/docker-projects-data/business-manager/postgres/company
/opt/docker-projects-data/business-manager/postgres/customer
/opt/docker-projects-data/business-manager/postgres/order
/opt/docker-projects-data/business-manager/postgres/whatsapp
/opt/docker-projects-data/business-manager/mongo
```

## Migraciones

Cada servicio PostgreSQL tiene un contenedor `business-manager-migrate-*` que aplica `up` antes de arrancar el microservicio. Si una migración falla, el servicio dependiente no arranca.

## Cloudflare

El túnel y el `config.yml` se crean directamente en el servidor. Ver `tourism-agency/DEPLOY.md` para el bloque compartido.
