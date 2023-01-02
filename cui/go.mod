module github.com/aokuyama/go-study-layered/cui

go 1.20

replace github.com/aokuyama/go-study-layered/packages/domain => ../packages/domain

replace github.com/aokuyama/go-study-layered/packages/app => ../packages/app

replace github.com/aokuyama/go-study-layered/packages/infra => ../packages/infra

require (
	github.com/aokuyama/go-study-layered/packages/app v0.0.0-00010101000000-000000000000
	github.com/aokuyama/go-study-layered/packages/infra v0.0.0-00010101000000-000000000000
)

require github.com/aokuyama/go-study-layered/packages/domain v0.0.0-00010101000000-000000000000 // indirect
