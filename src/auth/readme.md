# AUTH SERVICE

## Create virtual env

```python
    python3 -m venv venv
```

## Activate bash

```python
    source ./venv/bin/activate
```

## Check venv

```bash
    env | grep VIRTUAL
```

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
