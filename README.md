# {{ .Name }}

## Table of Contents
- [Overview](#overview)
- [Configuration](#configuration)
  - [Asset registration](#asset-registration)
  - [Resource definition](#resource-definition)
- [Functionality](#functionality)
- [Installation from source](#installation-from-source)
- [Additional notes](#additional-notes)
- [Contributing](#contributing)

## Overview
{{ .Name }} is a template repository which wraps the [Sensu Plugins Go Library][2].
To use this project as a template, click the "Use this template" button from the main project page.

When writing or updating a plugin's README from this template, review the Sensu Community
[plugin README style guide][3] for content suggestions and guidance.

## Configuration

### Asset registration

Assets are the best way to make use of this plugin. If you're not using an asset, please consider
doing so! If you're using sensuctl 5.13 with Sensu Backend 5.13 or later, you can use the following
command to add the asset:

```
sensuctl asset add {{ .GithubUser }}/{{ .GithubProject }}
```

If you're using an earlier version of sensuctl, you can find the asset on the [Bonsai Asset Index]([https://bonsai.sensu.io/assets/{{ .GithubUser }}/{{ .GithubProject }}](https://bonsai.sensu.io/assets/{{ .GithubUser }}/{{ .GithubProject }})).

### Resource definition

```yml
---
type: Handler
api_version: core/v2
metadata:
  name: {{ .GithubProject }}
  namespace: default
spec:
  command: {{ .GithubProject }} --example example_arg
  type: pipe
  runtime_assets:
  - {{ .GithubProject }}
```

## Functionality

After successfully creating a project from this template, update the `handlerConfigOptions` struct in
addition to the `executeHandler` and `checkArgs` functions in [main.go][7] to customize the behavior
of this handler plugin.

## Installation from source

The preferred way of installing and deploying this plugin is to use it as an Asset. If you would
like to compile and install the plugin from source or contribute to it, download the latest version
or create an executable script from this source.

From the local path of the {{ .GithubProject }} repository:

```
go build -o /usr/local/bin/{{ .GithubProject }} main.go
```

## Additional notes

To release a version of your project, simply tag the target sha with a semver release without a `v`
prefix (ex. `1.0.0`). This will trigger the [GitHub action][5] workflow to [build and release][4]
the plugin with goreleaser. Finalized releases can be found [here][6]. Register the asset with
[Bonsai][8] to share it with the community!

## Contributing

For more information about contributing to this plugin, see [Contributing][1].

[1]: https://github.com/sensu/sensu-go/blob/master/CONTRIBUTING.md
[2]: github.com/sensu-community/sensu-plugin-sdk
[3]: https://github.com/sensu-plugins/community/blob/master/PLUGIN_STYLEGUIDE.md
[4]: https://github.com/{{ .GithubUser }}/{{ .GithubProject }}/blob/master/.github/workflows/release.yml
[5]: https://github.com/{{ .GithubUser }}/{{ .GithubProject }}/actions
[6]: https://github.com/{{ .GithubUser }}/{{ .GithubProject }}/releases
[7]: https://github.com/{{ .GithubUser }}/{{ .GithubProject }}/blob/master/main.go
[8]: https://bonsai.sensu.io/