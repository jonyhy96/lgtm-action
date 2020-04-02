# lgtm-action

[![Go Report Card](https://goreportcard.com/badge/github.com/jonyhy96/lgtm-action)](https://goreportcard.com/report/github.com/jonyhy96/lgtm-action)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Approve pull requests when numbers of `lgtm` comment reaches require times.

Suggestions and issues can be posted on the repositories [issues page](https://github.com/jonyhy96/lgtm-action/issues).

[Inputs](#Inputs)
* [token](#token)
* [times](#times)
* [owners](#owners)

[Example usage](#Example-usage)

## Inputs

### `token`

**Required** Github auth token use for make access to this pr.

### `times`

The number of owners that you want to get approved from.(default: 1)

### `owners`

The owners file containing all owner.(default: OWNERS)


## Example usage

The following will subscribe `issue_comment` and `pull_request_review` event and if the times is satisfied it will approve this pr.

**Recommend** Use docker image:
```yaml
name: lgtm workflow 

on:
  pull_request_review:
  issue_comment:

jobs:
  build-and-push:
    runs-on: ubuntu-16.04
    timeout-minutes: 3
    steps:
    - name: Checkout
      uses: actions/checkout@v2

    - name: lgtm action
      uses: docker://jonyhy/lgtm-action:v1
      with:
        token: ${{ secrets.GITHUB_TOKEN }}
```

### Coding style

[CODEFMT](https://github.com/golang/go/wiki/CodeReviewComments)

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/jonyhy96/lgtm-action/tags). 

## Authors

* **HAO YUN** - *Initial work* - [haoyun](https://github.com/jonyhy96)

See also the list of [contributors](CONTRIBUTORS) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

## Acknowledgments

* nothing

