module github.com/Mikhgit/shippy/shippy-service-consignment

go 1.14

replace github.com/lucas-clemente/quic-go => github.com/lucas-clemente/quic-go v0.19.3

require (
	github.com/Mikhgit/shippy/shippy-service-user v0.0.0-20201209195328-90089c28dc4b
	github.com/Mikhgit/shippy/shippy-service-vessel v0.0.0-20201210140420-7f16b516884e
	github.com/golang/protobuf v1.4.3
	github.com/micro/go-micro/v2 v2.9.1
	github.com/pkg/errors v0.9.1
	go.mongodb.org/mongo-driver v1.4.4
	google.golang.org/protobuf v1.23.0
)
