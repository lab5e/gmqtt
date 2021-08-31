# lmqtt - a MQTT library

This library is based on the quite nice [gmqtt project](https://github.com/drmagice/gmqtt)
but with fewer features. If you are looking for a complete MQTT server in Go
with monitoring, a pre-built service and plugins to extend the features have a
look at it.

## Get Started

To get started with gmqtt, we need to compile it from the source code. Please ensure that you have a working
Go environment.

The following command will start gmqtt broker with default configuration.
The broker listens on 1883 for tcp server and 8883 for websocket server with `admin` and `prometheus` plugin loaded.

```bash
git clone https://github.com/lab5e/lmqtt
cd gmqtt/cmd/gmqttd
go run . start -c default_config.yml
```

## configuration

Gmqtt use `-c` flag to define configuration path. If not set, gmqtt reads `$HOME/gmqtt.yml` as default.  Here is a [sample configuration](https://github.com/lab5e/lmqtt/blob/master/cmd/gmqttd/default_config.yml).

## session persistence

Gmqtt uses memory to store session data by default and it is the recommended way because of the good performance.
But the session data will be lose after the broker restart. You can use redis as backend storage to prevent data
loss from restart:

```yaml
persistence:
  type: redis
  redis:
    # redis server address
    addr: "127.0.0.1:6379"
    # the maximum number of idle connections in the redis connection pool
    max_idle: 1000
    # the maximum number of connections allocated by the redis connection pool at a given time.
    # If zero, there is no limit on the number of connections in the pool.
    max_active: 0
    # the connection idle timeout, connection will be closed after remaining idle for this duration. If the value is zero, then idle connections are not closed
    idle_timeout: 240s
    password: ""
    # the number of the redis database
    database: 0
```

## Authentication

Gmqtt provides a simple username/password authentication mechanism. (Provided by [auth](https://github.com/lab5e/lmqtt/blob/master/plugin/auth) plugin).
It is not enabled in default configuration, you can change the configuration to enable it:

```yaml
# plugin loading orders
plugin_order:
  - auth
  - prometheus
  - admin
```

When auth plugin enabled, every clients need an account to get connected.You can add accounts through the HTTP API:

```bash
# Create: username = user1, password = user1pass
$ curl -X POST -d '{"password":"user1pass"}' 127.0.0.1:8083/v1/accounts/user1
{}
# Query
$ curl 127.0.0.1:8083/v1/accounts/user1
{"account":{"username":"user1","password":"20a0db53bc1881a7f739cd956b740039"}}
```

API Doc [swagger](https://github.com/lab5e/lmqtt/blob/master/plugin/auth/swagger)

## Docker

```bash
docker build -t gmqtt .
docker run -p 1883:1883 -p 8883:8883 -p 8082:8082 -p 8083:8083  -p 8084:8084  gmqtt
```

## Documentation

[godoc](https://www.godoc.org/github.com/lab5e/lmqtt)

## Hooks

Gmqtt implements the following hooks:

| Name | hooking point | possible usages  |
|------|------------|------------|
| OnAccept  | When accepts a TCP connection.(Not supported in websocket)| Connection rate limit, IP allow/block list. |
| OnStop  | When gmqtt stop |    |
| OnSubscribe  | When received a subscribe packet | Subscribe access control, modifies subscriptions. |
| OnSubscribed  | When subscribe succeed   |     |
| OnUnsubscribe  |  When received a unsubscribe packet | Unsubscribe access controls, modifies the topics that is going to unsubscribe.|
| OnUnsubscribed  | When unsubscribe succeed     |        |
| OnMsgArrived  | When received a publish packet  |  Publish access control, modifies message before delivery.|
| OnBasicAuth  | When received a connect packet without AuthMethod property | Authentication      |
| OnEnhancedAuth  | When received a connect packet with AuthMethod property (Only for v5 clients) | Authentication      |
| OnReAuth  | When received a auth packet (Only for v5 clients)        | Authentication      |
| OnConnected  | When the client connected succeed|      |
| OnSessionCreated  | When creates a new session       |         |
| OnSessionResumed  | When resumes from old session    |        |
| OnSessionTerminated  | When session terminated       |        |
| OnDelivered  | When a message is delivered to the client     |        |
| OnClosed  | When the client is closed  |        |
| OnMsgDropped  | When a message is dropped for some reasons|        |
| OnWillPublish | When the client is going to deliver a will message | Modify or drop the will message |
| OnWillPublished| When a will message has been delivered| |

See `/examples/hook` for details.

## How to write plugins

[How to write plugins](https://github.com/lab5e/lmqtt/blob/master/plugin/README.md)

## Contributing

Contributions are always welcome, see [Contribution Guide](https://github.com/lab5e/lmqtt/blob/master/CONTRIBUTING.md) for a complete contributing guide.

## Test

## Unit Test

```bash
go test -race ./...
```

## Integration Test

[paho.mqtt.testing](https://github.com/eclipse/paho.mqtt.testing).
