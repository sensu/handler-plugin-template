module github.com/{{ .GithubUser }}/{{ .GithubProject }}

go 1.14

require (
	github.com/sensu/sensu-plugin-sdk v0.13.1
	github.com/sensu/sensu-go/api/core/v2 v2.3.0
	github.com/sensu/sensu-go/types v0.3.0
)
