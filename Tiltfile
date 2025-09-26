# Load the restart_process extension
load('ext://restart_process', 'docker_build_with_restart')

### K8s Config ###

# Uncomment to use secrets
# k8s_yaml('./infra/development/k8s/secrets.yaml')

k8s_yaml('./infra/development/k8s/app-config.yaml')

### End of K8s Config ###
### API Gateway ###

gateway_compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/gateway ./services/gateway'
# if os.name == 'nt':
#   gateway_compile_cmd = './infra/development/docker/api-gateway-build.bat'

local_resource(
  'gateway-compile',
  gateway_compile_cmd,
  deps=['./services/gateway', './shared'], labels="compiles")


docker_build_with_restart(
  'edorguez/gateway',
  '.',
  entrypoint=['/app/build/gateway'],
  dockerfile='./infra/development/docker/gateway.Dockerfile',
  only=[
    './build/gateway',
    './shared',
  ],
  live_update=[
    sync('./build', '/app/build'),
    sync('./shared', '/app/shared'),
  ],
)

k8s_yaml('./infra/development/k8s/gateway-deployment.yaml')
k8s_resource('gateway', port_forwards=3001,
             resource_deps=['gateway-compile'], labels="services")
### End of API Gateway ###
### Auth Service ###

auth_compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/auth-svc ./services/auth-svc/cmd/main.go'
# if os.name == 'nt':
#  trip_compile_cmd = './infra/development/docker/trip-build.bat'

local_resource(
  'auth-svc-compile',
  auth_compile_cmd,
  deps=['./services/auth-svc', './shared'], labels="compiles")

docker_build_with_restart(
  'edorguez/auth-svc',
  '.',
  entrypoint=['/app/build/auth-svc'],
  dockerfile='./infra/development/docker/auth-svc.Dockerfile',
  only=[
    './build/auth-svc',
    './shared',
  ],
  live_update=[
    sync('./build', '/app/build'),
    sync('./shared', '/app/shared'),
  ],
)

k8s_yaml('./infra/development/k8s/auth-svc-deployment.yaml')
k8s_resource('auth-svc', resource_deps=['auth-svc-compile'], labels="services")

### End of Auth Service ###
### Company Service ###

company_compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/company-svc ./services/company-svc'

# if os.name == 'nt':
#  driver_compile_cmd = './infra/development/docker/driver-build.bat'

local_resource(
  'company-svc-compile',
  company_compile_cmd,
  deps=['./services/company-svc', './shared'], labels="compiles")

docker_build_with_restart(
  'edorguez/company-svc',
  '.',
  entrypoint=['/app/build/company-svc'],
  dockerfile='./infra/development/docker/company-svc.Dockerfile',
  only=[
    './build/company-svc',
    './shared',
  ],
  live_update=[
    sync('./build', '/app/build'),
    sync('./shared', '/app/shared'),
  ],
)

k8s_yaml('./infra/development/k8s/company-svc-deployment.yaml')
k8s_resource('company-svc', resource_deps=['company-svc-compile'], labels="services")

### End of Company Service ###
#### Customer Service ###

customer_compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/customer-svc ./services/customer-svc'

# if os.name == 'nt':
#  driver_compile_cmd = './infra/development/docker/driver-build.bat'

local_resource(
  'customer-svc-compile',
  customer_compile_cmd,
  deps=['./services/customer-svc', './shared'], labels="compiles")

docker_build_with_restart(
  'edorguez/customer-svc',
  '.',
  entrypoint=['/app/build/customer-svc'],
  dockerfile='./infra/development/docker/customer-svc.Dockerfile',
  only=[
    './build/customer-svc',
    './shared',
  ],
  live_update=[
    sync('./build', '/app/build'),
    sync('./shared', '/app/shared'),
  ],
)

k8s_yaml('./infra/development/k8s/customer-svc-deployment.yaml')
k8s_resource('customer-svc', resource_deps=['customer-svc-compile'], labels="services")

### End of Customer Service ###
##### File Service ###

file_compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/file-svc ./services/file-svc'

# if os.name == 'nt':
#  driver_compile_cmd = './infra/development/docker/driver-build.bat'

local_resource(
  'file-svc-compile',
  file_compile_cmd,
  deps=['./services/file-svc', './shared'], labels="compiles")

docker_build_with_restart(
  'edorguez/file-svc',
  '.',
  entrypoint=['/app/build/file-svc'],
  dockerfile='./infra/development/docker/file-svc.Dockerfile',
  only=[
    './build/file-svc',
    './shared',
  ],
  live_update=[
    sync('./build', '/app/build'),
    sync('./shared', '/app/shared'),
  ],
)

k8s_yaml('./infra/development/k8s/file-svc-deployment.yaml')
k8s_resource('file-svc', resource_deps=['file-svc-compile'], labels="services")

### End of File Service ###
###### Order Service ###

order_compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/order-svc ./services/order-svc'

# if os.name == 'nt':
#  driver_compile_cmd = './infra/development/docker/driver-build.bat'

local_resource(
  'order-svc-compile',
  order_compile_cmd,
  deps=['./services/order-svc', './shared'], labels="compiles")

docker_build_with_restart(
  'edorguez/order-svc',
  '.',
  entrypoint=['/app/build/order-svc'],
  dockerfile='./infra/development/docker/order-svc.Dockerfile',
  only=[
    './build/order-svc',
    './shared',
  ],
  live_update=[
    sync('./build', '/app/build'),
    sync('./shared', '/app/shared'),
  ],
)

k8s_yaml('./infra/development/k8s/order-svc-deployment.yaml')
k8s_resource('order-svc', resource_deps=['order-svc-compile'], labels="services")

### End of Order Service ###
####### Product Service ###

product_compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/product-svc ./services/product-svc'

# if os.name == 'nt':
#  driver_compile_cmd = './infra/development/docker/driver-build.bat'

local_resource(
  'product-svc-compile',
  product_compile_cmd,
  deps=['./services/product-svc', './shared'], labels="compiles")

docker_build_with_restart(
  'edorguez/product-svc',
  '.',
  entrypoint=['/app/build/product-svc'],
  dockerfile='./infra/development/docker/product-svc.Dockerfile',
  only=[
    './build/product-svc',
    './shared',
  ],
  live_update=[
    sync('./build', '/app/build'),
    sync('./shared', '/app/shared'),
  ],
)

k8s_yaml('./infra/development/k8s/product-svc-deployment.yaml')
k8s_resource('product-svc', resource_deps=['product-svc-compile'], labels="services")

### End of Product Service ###
######## Whatsapp Service ###

whatsapp_compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/whatsapp-svc ./services/whatsapp-svc'

# if os.name == 'nt':
#  driver_compile_cmd = './infra/development/docker/driver-build.bat'

local_resource(
  'whatsapp-svc-compile',
  whatsapp_compile_cmd,
  deps=['./services/whatsapp-svc', './shared'], labels="compiles")

docker_build_with_restart(
  'edorguez/whatsapp-svc',
  '.',
  entrypoint=['/app/build/whatsapp-svc'],
  dockerfile='./infra/development/docker/whatsapp-svc.Dockerfile',
  only=[
    './build/whatsapp-svc',
    './shared',
  ],
  live_update=[
    sync('./build', '/app/build'),
    sync('./shared', '/app/shared'),
  ],
)

k8s_yaml('./infra/development/k8s/whatsapp-svc-deployment.yaml')
k8s_resource('whatsapp-svc', resource_deps=['whatsapp-svc-compile'], labels="services")

### End of Whatsapp Service ###
### Web Frontend ###

# docker_build(
#   'ride-sharing/web',
#   '.',
#   dockerfile='./infra/development/docker/web.Dockerfile',
# )

# k8s_yaml('./infra/development/k8s/web-deployment.yaml')
# k8s_resource('web', port_forwards=3000, labels="frontend")

### End of Web Frontend ###
