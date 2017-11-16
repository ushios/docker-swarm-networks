# docker-swarm-networks

### Start

```console
$ docker-compose build
$ docker swarm init
$ docker stack deploy -c docker-compose.yml swarm_network
$ docker service logs -f swarm_network_server
```

### Close

```console
$ docker stack rm swarm_network
```
