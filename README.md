# docker-swarm-networks
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fushios%2Fdocker-swarm-networks.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fushios%2Fdocker-swarm-networks?ref=badge_shield)


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
$ docker swarm leave -f
```


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fushios%2Fdocker-swarm-networks.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fushios%2Fdocker-swarm-networks?ref=badge_large)