docker_build('arnar.io/seccy', '.', dockerfile='Dockerfile')
k8s_yaml(helm('deployment/helm/seccy'))
k8s_resource('seccy', port_forwards=4040)