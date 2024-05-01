
def main():
    print("""apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment-python
spec:
  replicas: 2
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
        - containerPort: 80""")

if __name__ == "__main__":
    main()