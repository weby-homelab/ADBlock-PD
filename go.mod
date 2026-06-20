module github.com/weby-homelab/adblock-pd

go 1.26.4

require (
	github.com/AdguardTeam/dnsproxy v0.81.4
	github.com/AdguardTeam/golibs v0.35.13
	github.com/AdguardTeam/urlfilter v0.23.2
	github.com/NYTimes/gziphandler v1.1.1
	github.com/ameshkov/dnscrypt/v2 v2.4.0
	github.com/bluele/gcache v0.0.2
	github.com/c2h5oh/datasize v0.0.0-20231215233829-aa82cc1e6500
	github.com/digineo/go-ipset/v2 v2.2.1
	github.com/fsnotify/fsnotify v1.10.1
	// TODO(e.burkov): This package is deprecated; find a new one or use our
	// own code for that.  Perhaps, use gopacket.
	github.com/go-ping/ping v1.2.0
	github.com/google/go-cmp v0.7.0
	github.com/google/gopacket v1.1.19
	github.com/google/renameio/v2 v2.0.2
	github.com/google/uuid v1.6.0
	github.com/insomniacslk/dhcp v0.0.0-20260407060928-11b94ed970f2
	github.com/kardianos/service v1.2.4
	github.com/mdlayher/ethernet v0.0.0-20220221185849-529eae5b6118
	github.com/mdlayher/netlink v1.11.0
	github.com/mdlayher/packet v1.1.2
	// TODO(a.garipov): This package is deprecated; Use gopacket.
	github.com/mdlayher/raw v0.1.0
	github.com/miekg/dns v1.1.72
	github.com/quic-go/quic-go v0.59.1
	github.com/stretchr/testify v1.11.1
	github.com/ti-mo/netfilter v0.5.3
	go.etcd.io/bbolt v1.4.3
	// TODO(e.burkov): Update to a stable tag.
	go.yaml.in/yaml/v4 v4.0.0-rc.4.0.20260405193028-802e24f4fbcc
	golang.org/x/crypto v0.53.0
	golang.org/x/exp v0.0.0-20260410095643-746e56fc9e2f
	golang.org/x/net v0.56.0
	golang.org/x/sys v0.46.0
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
	howett.net/plist v1.0.1
)

require (
	github.com/ameshkov/dnsstamps v1.0.3 // indirect
	github.com/beefsack/go-rate v0.0.0-20220214233405-116f4ca011a0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/josharian/native v1.1.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mdlayher/socket v0.6.0 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pierrec/lz4/v4 v4.1.26 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/quic-go/qpack v0.6.0 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/u-root/uio v0.0.0-20240224005618-d2acac8f3701 // indirect
	golang.org/x/mod v0.36.0 // indirect
	golang.org/x/sync v0.21.0 // indirect
	golang.org/x/text v0.38.0 // indirect
	golang.org/x/tools v0.45.0 // indirect
	gonum.org/v1/gonum v0.17.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
