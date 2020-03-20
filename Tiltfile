docker_build('arnar.io/seccy-service', '.', dockerfile='cmd/seccy-service/Dockerfile')
docker_build('arnar.io/seccy-rest-api', '.', dockerfile='cmd/seccy-rest-api/Dockerfile')

k8s_yaml(helm('deployment/helm/seccy-service',
  name='seccy-service',
  namespace='seccy', 
  values=['deployment/helm/seccy-service/values-dev.yaml']))

k8s_yaml(helm('deployment/helm/seccy-rest-api',
  name='seccy-rest-api',
  namespace='seccy'))

k8s_resource('seccy-service', port_forwards=4040)
k8s_resource('seccy-rest-api', port_forwards=8080)