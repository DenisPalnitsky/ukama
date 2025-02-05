module github.com/ukama/ukama/services/cloud/foo

go 1.18

require (
	github.com/go-yaml/yaml v2.1.0+incompatible // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/ukama/ukama/services/common v0.0.0-20211015093708-cd6e230254b5 // indirect
	google.golang.org/grpc v1.41.0 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gorm.io/gorm v1.21.16 // indirect
)

replace github.com/ukama/ukama/services/common => ../../common

replace github.com/ukama/ukama/services/cloud/foo => ./

replace github.com/ukama/ukama/services/cloud/foo/mocks => ./mocks
