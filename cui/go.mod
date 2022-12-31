module github.com/aokuyama/go-study-layered/cui

go 1.20

replace github.com/aokuyama/go-study-layered/src/domain => ../src/domain

replace github.com/aokuyama/go-study-layered/src/app => ../src/app

replace github.com/aokuyama/go-study-layered/src/infra => ../src/infra

require (
	github.com/aokuyama/go-study-layered/src/app v0.0.0-00010101000000-000000000000
	github.com/aokuyama/go-study-layered/src/infra v0.0.0-00010101000000-000000000000
)

require github.com/aokuyama/go-study-layered/src/domain v0.0.0-00010101000000-000000000000 // indirect
