# RTM

RTM is a real time measurements platform that use `golang`, `mqtt`, `swagger` and `redis`.

## Usage

```sh
docker-compose up -d
```

This command start two containers:
    - Redis that store data (measurements and iata codes)
    - Grafana that allow data-visualization on `localhost:3000`
