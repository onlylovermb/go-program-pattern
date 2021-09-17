# functional options
### Example
``` go
// 假如我们需要初始化一个server结构体如下所示，结构体中有必需配置（有默认值&无默认值）和非必需配置，如何设置server 的初始化函数？
type Server struct {
	// 必需配置，无默认值
	IP   string
	Port int
	// 必需配置，有默认值
	Protocol string
	Timeout  time.Duration
	Maxconns int
	// 可选配置
	TLS *tls.Config
}
```
### Demo
- Level 0：use demo with many NewServerXXX function
```go
// simple but too ugly
serverA, _ := NewServerDefault("127.0.0.1", 2001)
serverB, _ := NewServerWithProtocol("127.0.0.1", 2001, "udp")
serverC, _ := NewServerWithProtocolAndTLS("127.0.0.1", 2001, "udp", TLS(&tls.Config{}))
```

- Level 1：use demo with a config options
```go
// friendly but not perfect 
type Config struct{
	Protocol string
	Timeout  time.Duration
	Maxconns int
	TLS *tls.Config
}

type Server struct{
    IP string
    Port int
    Conf *Config
}

serverA, _ := NewServer("127.0.0.1", 2001, nil)
config := Config{
    	Protocol: "tcp",
		Timeout:  30 * time.Second,
		Maxconns: 10,
}
serverB, _ := NewServer("127.0.0.1", 2001, &config)
```

- Level 2：use demo with builder
```go
// complex and not suit for go

type ServerBuilder struct {
  Server
}

func (sb *ServerBuilder) Create(addr string, port int) *ServerBuilder {
  sb.Server.Addr = addr
  sb.Server.Port = port
  return sb
}

func (sb *ServerBuilder) WithProtocol(protocol string) *ServerBuilder {
  sb.Server.Protocol = protocol 
  return sb
}

func (sb *ServerBuilder) WithMaxConn( maxconn int) *ServerBuilder {
  sb.Server.MaxConns = maxconn
  return sb
}

func (sb *ServerBuilder) WithTimeOut( timeout time.Duration) *ServerBuilder {
  sb.Server.Timeout = timeout
  return sb
}

func (sb *ServerBuilder) WithTLS( tls *tls.Config) *ServerBuilder {
  sb.Server.TLS = tls
  return sb
}

func (sb *ServerBuilder) Build() (Server) {
  return  sb.Server
}


sb := ServerBuilder{}
server, err := sb.Create("127.0.0.1", 8080).
  WithProtocol("udp").
  WithMaxConn(1024).
  WithTimeOut(30*time.Second).
  Build()
```

- **Level 3：use demo with functional options**
```go
// maybe the best way
serverA, _ := NewServer("127.0.0.1", 2001)
serverB, _ := NewServer("127.0.0.1", 2001, Protocol("udp"))
serverC, _ := NewServer("127.0.0.1", 2001, Protocol("udp"), TLS(&tls.Config{}))
```