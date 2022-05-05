module github.com/reapchain/reapchain

go 1.17

require (
	github.com/armon/go-metrics v0.3.10
	github.com/cosmos/go-bip39 v1.0.0
	github.com/ethereum/go-ethereum v1.10.16
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/onsi/ginkgo/v2 v2.1.3
	github.com/onsi/gomega v1.18.1
	github.com/pkg/errors v0.9.1
	github.com/rakyll/statik v0.1.7
	github.com/reapchain/cosmos-sdk v0.45.1-reapsdkv2
	github.com/reapchain/ethermint v0.12.0-reapethermintv2
	github.com/reapchain/ibc-go v0.3.0-reapibcv2
	github.com/reapchain/reapchain-core v0.1.5
	github.com/spf13/cast v1.4.1
	github.com/spf13/cobra v1.4.0
	github.com/stretchr/testify v1.7.1
	github.com/tendermint/tm-db v0.6.7
	go.opencensus.io v0.23.0
	google.golang.org/genproto v0.0.0-20220504150022-98cd25cafc72
	google.golang.org/grpc v1.46.0
	google.golang.org/protobuf v1.28.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/DataDog/zstd v1.4.8 // indirect
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/dgraph-io/badger/v2 v2.2007.3 // indirect
	github.com/dgraph-io/ristretto v0.1.0 // indirect
	github.com/golang/glog v1.0.0 // indirect
	github.com/google/btree v1.0.1 // indirect
	github.com/prometheus/tsdb v0.10.0 // indirect
	github.com/rjeczalik/notify v0.9.2 // indirect
	github.com/rs/zerolog v1.26.0 // indirect
	github.com/tklauser/go-sysconf v0.3.7 // indirect
	nhooyr.io/websocket v1.8.7 // indirect
)

replace (
	github.com/99designs/keyring => github.com/cosmos/keyring v1.1.7-0.20210622111912-ef00f8ac3d76
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	google.golang.org/grpc => google.golang.org/grpc v1.33.2

)
