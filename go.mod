module github.com/reapchain/mercury

go 1.17

require (
	contrib.go.opencensus.io/exporter/prometheus v0.4.0
	github.com/ethereum/go-ethereum v1.10.11
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/rakyll/statik v0.1.7
	github.com/reapchain/cosmos-sdk v0.44.5
	github.com/reapchain/ethermint v0.0.2
	github.com/reapchain/ibc-go v0.0.2
	github.com/reapchain/reapchain-core v0.1.2
	github.com/spf13/cast v1.4.1
	github.com/spf13/cobra v1.3.0
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tm-db v0.6.6
	go.opencensus.io v0.23.0
	google.golang.org/genproto v0.0.0-20211208223120-3a66f561d7aa
	google.golang.org/grpc v1.43.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/deckarep/golang-set v1.7.1 // indirect
	github.com/dgraph-io/badger/v2 v2.2007.3 // indirect
	github.com/dgraph-io/ristretto v0.1.0 // indirect
	github.com/golang/glog v1.0.0 // indirect
	github.com/prometheus/tsdb v0.10.0 // indirect
	github.com/rjeczalik/notify v0.9.2 // indirect
	nhooyr.io/websocket v1.8.7 // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace github.com/99designs/keyring => github.com/cosmos/keyring v1.1.7-0.20210622111912-ef00f8ac3d76
