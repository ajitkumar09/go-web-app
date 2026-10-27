# Go Web Application

This is a simple website written in Golang. It uses the `net/http` package to serve HTTP requests.

## Running the server

Install Go 1.22 or later, then run:

```bash
go run .
```

The server will start on port 8080. Open `http://localhost:8080/` for the home page, or visit `/courses`, `/about`, and `/contact`.

## Running tests

```bash
go test ./...
go vet ./...
```

## Building and running with Docker

```bash
docker build -t go-web-app:v1 .
docker run --rm -p 8080:8080 go-web-app:v1
```

The image uses a minimal, non-root runtime. The application health endpoint is available at `/healthz`.

## Deploying to Kubernetes

The manifests in `static/` define a two-replica Deployment, ClusterIP Service, and NGINX Ingress. Build or push the image to a registry accessible by your cluster and update the `image` in `static/deployment.yaml` as needed. For a local cluster, load the image into the cluster first.

```bash
kubectl apply -f static/
kubectl get deployments,services,ingresses
```

Install an NGINX Ingress Controller and point `goapp.local` to the ingress controller's address before opening the site.

## Continuous integration

GitHub Actions runs the Go tests, vet checks, and Docker build on pushes and pull requests.

## Looks like this

![Website](static/images/golang-website.png)
