module shippy-service-email

go 1.14

replace github.com/lucas-clemente/quic-go => github.com/lucas-clemente/quic-go v0.19.3

require (
	github.com/Mikhgit/shippy/shippy-service-user v0.0.0-20201207110131-0fed603bda17
	github.com/lucas-clemente/quic-go v0.19.3 // indirect
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/nats/v2 v2.9.1
)
