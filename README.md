# Video converter app

Easily convert video to mp3 audio files

## Requirements

### Docker

Install docker locally -> https://www.docker.com/get-started/

```bash
    docker --version
```

### Kubectl

Install Kubernetes command line tools -> https://kubernetes.io/docs/tasks/tools/

```bash
    kubectl
```

### Minikube

Install Minikube -> https://minikube.sigs.k8s.io/docs/start/

```bash
    minikube start
```

### K9s

Install k9s -> https://github.com/derailed/k9s?tab=readme-ov-file#installation

```bash
    k9s
```

### Route mp3converter.com to 127.0.0.1

Add a "127.0.0.1 mp3converter.com" in the hosts' list

```bash
    sudo vim /etc/hosts
```

Enable ingress functionality in minikube

```bash
    minikube addons enable ingress
```

Tunnel mp3converter.com to 127.0.0.1

```bash
    minikube tunnel
```

### Python3

Install python3 -> https://www.python.org/downloads/

```bash
    python3 --version
```

### Mysql

Install mysql -> https://formulae.brew.sh/formula/mysql

```bash
    brew services start mysql
```
