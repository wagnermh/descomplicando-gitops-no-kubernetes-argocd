package main

import "fmt"

func main() {
	// Valid Kubernetes deployment YAML
	yaml := `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.23.0
        ports:
        - containerPort: 80
`

	// Print the YAML to the standard output
	fmt.Printf("%s\n", yaml)
}
