{{- /* Ignore this text, until templating is ran via [sensu-plugin-tool](https://github.com/sensu/sensu-plugin-tool) the below badge links wiill not render */ -}}

[![Sensu Bonsai Asset](https://img.shields.io/badge/Bonsai-Download%20Me-brightgreen.svg?colorB=89C967&logo=sensu)](https://bonsai.sensu.io/assets/{{ .GithubUser }}/{{ .GithubProject }})
![goreleaser](https://github.com/{{ .GithubUser }}/{{ .GithubProject }}/workflows/goreleaser/badge.svg)
[![Go Test](https://github.com/{{ .GithubUser }}/{{ .GithubProject }}/workflows/Go%20Test/badge.svg)](https://github.com/{{ .GithubUser }}/{{ .GithubProject }}/actions?query=workflow%3A%22Go+Test%22)
[![goreleaser](https://github.com/{{ .GithubUser }}/{{ .GithubProject }}/workflows/goreleaser/badge.svg)](https://github.com/{{ .GithubUser }}/{{ .GithubProject }}/actions?query=workflow%3Agoreleaser)

# Handler Plugin Template

## Overview
handler-plugin-template is a template repository which wraps the [Sensu Plugin SDK][2].
To use this project as a template, click the "Use this template" button from the main project page.
Once the repository is created from this template, you can use the [Sensu Plugin Tool][9] to
populate the templated fields with the proper values.

## Functionality

After successfully creating a project from this template, update the `Config` struct with any
configuration options for the plugin, map those values as plugin options in the variable `options`,
and customize the `checkArgs` and `executeHandler` functions in [main.go][7].

When writing or updating a plugin's README from this template, review the Sensu Community
[plugin README style guide][3] for content suggestions and guidance. Remove everything
prior to `# {{ .Name }}` from the generated README file, and add additional context about the
plugin per the style guide.

## Releases with Github Actions

To release a version of your project, simply tag the target sha with a semver release without a `v`
prefix (ex. `1.0.0`). This will trigger the [GitHub action][5] workflow to [build and release][4]
the plugin with goreleaser. Register the asset with [Bonsai][8] to share it with the community!

***

# {{ .Name }}

## Table of Contents
- [Overview](#overview)
- [Files](#files)
- [Usage examples](#usage-examples)
- [Configuration](#configuration)
  - [Asset registration](#asset-registration)
  - [Handler definition](#handler-definition)
  - [Annotations](#annotations)
- [Installation from source](#installation-from-source)
- [Additional notes](#additional-notes)
- [Contributing](#contributing)

## Overview

The {{ .Name }} is a [Sensu Handler][6] that ...

## Files

## Usage examples

## Configuration

### Asset registration

[Sensu Assets][10] are the best way to make use of this plugin. If you're not using an asset, please
consider doing so! If you're using sensuctl 5.13 with Sensu Backend 5.13 or later, you can use the
following command to add the asset:

```
sensuctl asset add {{ .GithubUser }}/{{ .GithubProject }}
```

If you're using an earlier version of sensuctl, you can find the asset on the [Bonsai Asset Index][https://bonsai.sensu.io/assets/{{ .GithubUser }}/{{ .GithubProject }}].

### Handler definition

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
  - {{ .GithubUser}}/{{ .GithubProject }}
```

#### Proxy Support

This handler supports the use of the environment variables HTTP_PROXY,
HTTPS_PROXY, and NO_PROXY (or the lowercase versions thereof). HTTPS_PROXY takes
precedence over HTTP_PROXY for https requests.  The environment values may be
either a complete URL or a "host[:port]", in which case the "http" scheme is assumed.

### Annotations

All arguments for this handler are tunable on a per entity or check basis based on annotations.  The
annotations keyspace for this handler is `sensu.io/plugins/{{ .GithubProject }}/config`.

#### Examples

To change the example argument for a particular check, for that checks's metadata add the following:

```yml
type: CheckConfig
api_version: core/v2
metadata:
  annotations:
    sensu.io/plugins/{{ .GithubProject }}/config/example-argument: "Example change"
[...]
```

## Installation from source

The preferred way of installing and deploying this plugin is to use it as an Asset. If you would
like to compile and install the plugin from source or contribute to it, download the latest version
or create an executable script from this source.

From the local path of the {{ .GithubProject }} repository:

```
go build
```

## Additional notes

## Contributing

For more information about contributing to this plugin, see [Contributing][1].

[1]: https://github.com/sensu/sensu-go/blob/master/CONTRIBUTING.md
[2]: https://github.com/sensu/sensu-plugin-sdk
[3]: https://github.com/sensu-plugins/community/blob/master/PLUGIN_STYLEGUIDE.md
[4]: https://github.com/sensu/handler-plugin-template/blob/master/.github/workflows/release.yml
[5]: https://github.com/sensu/handler-plugin-template/actions
[6]: https://docs.sensu.io/sensu-go/latest/reference/handlers/
[7]: https://github.com/sensu/handler-plugin-template/blob/master/main.go
[8]: https://bonsai.sensu.io/
[9]: https://github.com/sensu/sensu-plugin-tool
[10]: https://docs.sensu.io/sensu-go/latest/reference/assets/
