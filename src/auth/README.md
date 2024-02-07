# AUTH SERVICE

## Init database

```bash
    mysql -u root < init.sql
```

## Build docker image

```bash
    docker build -t jsflor/vc-auth-service:latest .
```

## Push docker image

```bash
    docker push jsflor/vc-auth-service:latest
```

## Set up k9s

```bash
    cd manifests && kubectl apply -f ./   
```
