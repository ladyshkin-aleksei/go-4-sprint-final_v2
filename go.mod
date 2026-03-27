module github.com/Yandex-Practicum/tracker

go 1.24.1

require (
		github.com/stretchr/testify v1.10.0
		spentcalories v0.0.0-00010101000000-000000000000
		daysteps v0.0.0-00010101000000-000000000000
)
require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	spentcalories v0.0.0-00010101000000-000000000000 // indirect
	daysteps v0.0.0-00010101000000-000000000000 // indirect
)

replace (
	spentcalories => /internal/spentcalories
	daysteps => /internal/daysteps
)