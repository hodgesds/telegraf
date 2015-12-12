# ZeroMq Plugin
The ZeroMq plugin is able to collect data from a variety of ZeroMq sockets.

# Default configuration
The default configuration creates two ZeroMq sockets; one **PULL**
socket and one **SUB** socket. These two socket types can be
customized as needed.

```
[socket]

    [socket.PULL]
    host = "127.0.0.1"
    port = "5555"

    [socket.SUB]
    host  = "127.0.0.1"
    port  = "5556"
    topic = "metric"
```
