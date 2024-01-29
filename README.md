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

### Route mp3converter.com and rabbitmq-manager.com to 127.0.0.1

Add a "127.0.0.1 mp3converter.com" and "127.0.0.1 rabbitmq-manager.com" in the hosts' list

```bash
    sudo vim /etc/hosts
```

Enable ingress functionality in minikube

```bash
    minikube addons enable ingress
```

Tunnel to 127.0.0.1

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
   brew install mysql && brew services start mysql
```

### Mongodb

Install mongo -> https://www.mongodb.com/docs/manual/tutorial/install-mongodb-on-os-x/

```bash
    brew tap mongodb/brew && brew update && brew install mongodb-community@7.0 && brew services start mongodb-community
```

## CURL

Get jwt credentials

```bash
    curl -X POST http://mp3converter.com/login -u test@mail.com:test123
```

Upload a video to convert

```bash
    curl -X POST -F 'file=@./video.mp4' -H 'Authorization: Bearer ${jwt}' http://mp3converter.com/upload
```

Download audio from video

```bash
    curl --output mp3_download.mp3 -X GET -H 'Authorization: Bearer ${jwt}' http://mp3converter.com/download?fid=${fid}
```
