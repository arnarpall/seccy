docker_build('arnar.io/seccy', '.', dockerfile='Dockerfile')

k8s_yaml(helm('deployment/helm/seccy', 
  name='seccy', 
  namespace='seccy', 
  values=['deployment/helm/seccy/values-dev.yaml']))

k8s_resource('seccy', port_forwards=4040)