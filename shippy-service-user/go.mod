module github.com/Mikhgit/shippy/shippy-service-user

go 1.14

replace github.com/lucas-clemente/quic-go => github.com/lucas-clemente/quic-go v0.19.3

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/protobuf v1.4.3
	github.com/jmoiron/sqlx v1.2.0
	github.com/lib/pq v1.9.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/nats/v2 v2.9.1
	github.com/satori/go.uuid v1.2.0
	golang.org/x/crypto v0.0.0-20201203163018-be400aefbc4c
	google.golang.org/protobuf v1.23.0
)
