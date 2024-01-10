# GATEWAY SERVICE

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

## Build docker image

```bash
    docker build -t jsflor/vc-gateway-service:latest .
```

## Push docker image

```bash
    docker push jsflor/vc-gateway-service:latest
```

## Set up k9s

```bash
    cd manifests && kubectl apply -f ./   
```
