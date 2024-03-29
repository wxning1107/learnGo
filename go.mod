module learnGoSource

go 1.14

require (
	github.com/evanphx/json-patch v4.9.0+incompatible // indirect
	github.com/go-redis/redis v6.15.8+incompatible
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/google/go-cmp v0.5.1 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/googleapis/gnostic v0.5.1 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/imdario/mergo v0.3.5 // indirect
	github.com/prometheus/client_golang v1.7.1
	github.com/prometheus/client_model v0.2.0
	github.com/prometheus/common v0.10.0
	github.com/rs/xid v1.2.1
	github.com/samuel/go-zookeeper v0.0.0-20190923202752-2cc03de413da
	github.com/stretchr/testify v1.5.1
	github.com/tjfoc/gmsm v1.4.1
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45 // indirect
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4 // indirect
	google.golang.org/appengine v1.5.0 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	k8s.io/api v0.17.0
	k8s.io/apimachinery v0.18.8
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/kube-openapi v0.0.0-20200811211545-daf3cbb84823 // indirect
	k8s.io/utils v0.0.0-20191114184206-e782cd3c129f
)

replace (
	k8s.io/api => k8s.io/api v0.0.0-20191004102349-159aefb8556b

	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20191004074956-c5d2f014d689

)
