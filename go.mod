module github.com/ammmnia/tools

go 1.21

toolchain go1.21.8

require (
	github.com/IBM/sarama v1.43.0
	github.com/aliyun/aliyun-oss-go-sdk v3.0.2+incompatible
	github.com/aws/aws-sdk-go-v2 v1.23.1
	github.com/aws/aws-sdk-go-v2/config v1.25.4
	github.com/aws/aws-sdk-go-v2/credentials v1.16.3
	github.com/gin-gonic/gin v1.9.1
	github.com/go-zookeeper/zk v1.0.3
	github.com/golang-jwt/jwt/v4 v4.5.0
	github.com/google/uuid v1.6.0
	github.com/jinzhu/copier v0.4.0
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lestrrat-go/strftime v1.0.6
	github.com/magefile/mage v1.15.0
	github.com/minio/minio-go/v7 v7.0.69
	github.com/pkg/errors v0.9.1
	github.com/qiniu/go-sdk/v7 v7.18.2
	github.com/redis/go-redis/v9 v9.2.1
	github.com/shirou/gopsutil v3.21.11+incompatible
	github.com/stretchr/testify v1.8.4
	github.com/tencentyun/cos-go-sdk-v5 v0.7.47
	go.etcd.io/etcd/client/v3 v3.5.13
	go.mongodb.org/mongo-driver v1.12.0
	go.uber.org/zap v1.24.0
	golang.org/x/image v0.15.0
	google.golang.org/grpc v1.62.1
	google.golang.org/protobuf v1.33.0
	gopkg.in/src-d/go-git.v4 v4.13.1
	gopkg.in/yaml.v2 v2.4.0
	gopkg.in/yaml.v3 v3.0.1
	gorm.io/gorm v1.25.8
)

require (
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.5.1 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.14.5 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.2.4 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.5.4 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.7.1 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.2.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.10.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.2.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.10.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.16.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3 v1.43.1
	github.com/aws/aws-sdk-go-v2/service/sso v1.17.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.20.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.25.4 // indirect
	github.com/aws/smithy-go v1.17.0 // indirect
	github.com/bytedance/sonic v1.9.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/clbanning/mxj v1.8.4 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/eapache/go-resiliency v1.6.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20230731223053-c322873962e3 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/emirpasic/gods v1.12.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.14.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.7.6 // indirect
	github.com/jcmturner/gokrb5/v8 v8.4.4 // indirect
	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jonboulle/clockwork v0.4.0
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kevinburke/ssh_config v0.0.0-20190725054713-01f96b0aa0cd // indirect
	github.com/klauspost/compress v1.17.7 // indirect
	github.com/klauspost/cpuid/v2 v2.2.6 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/minio/md5-simd v1.1.2 // indirect
	github.com/minio/sha256-simd v1.0.1 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe // indirect
	github.com/mozillazg/go-httpheader v0.2.1 // indirect
	github.com/pelletier/go-toml/v2 v2.0.8 // indirect
	github.com/pierrec/lz4/v4 v4.1.21 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	github.com/rs/xid v1.5.0 // indirect
	github.com/sergi/go-diff v1.0.0 // indirect
	github.com/src-d/gcfg v1.4.0 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	github.com/tklauser/go-sysconf v0.3.13 // indirect
	github.com/tklauser/numcpus v0.7.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	github.com/xanzy/ssh-agent v0.2.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	go.etcd.io/etcd/api/v3 v3.5.13 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.13 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/arch v0.3.0 // indirect
	golang.org/x/crypto v0.21.0 // indirect
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	google.golang.org/genproto v0.0.0-20240123012728-ef4313101c80 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240123012728-ef4313101c80 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240318140521-94a12d6c2237 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/src-d/go-billy.v4 v4.3.2 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
)
