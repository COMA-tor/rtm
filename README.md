# RTM

RTM is a real time measurements platform that use `golang`, `mqtt`, `swagger` and `redis`.

## Usage

```sh
docker-compose up -d
```

This command start two containers:
    - Redis that store data (measurements and iata codes)
    - Grafana that allow data-visualization on `localhost:3000`

### Install and launcn agent runner

The agent runner allows to run a measurement agent directly in the cli. 

Install it by using this command: 

```sh
go install github.com/COMA-tor/rtm/agent/runner
```

Now the runner is installed and can be used directly in the cli.

```sh
runner -help
```

> The runner directly output in stderr.

#### Runner configuration

The runner can be configured with a yaml config file.

```yaml
broker_host: localhost
broker_port: 1883
client_id: test
topic: /topic/test/
qos: 2
sensor_type: wind speed
sensor_unit: °C
```

## Possible improvements

- Handle reload signals in agent runner to reload the sensor and the agent