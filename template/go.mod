module github.com/ntons/tongo/template

go 1.14

// current release of etcd is not compatible well with go modules,
// the issue may be fixed in v3.5.z, using master instead

require (
	github.com/flosch/pongo2 v0.0.0-20200913210552-0d938eb266f3
	go.etcd.io/etcd/v3 v3.3.0-rc.0.0.20200930024832-ab4cc3caef3d
)
