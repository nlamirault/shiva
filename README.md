# Shiva

[![License Apache 2][badge-license]](LICENSE)
[![GitHub version](https://badge.fury.io/gh/nlamirault%2Fshiva.svg)](https://badge.fury.io/gh/nlamirault%2Fshiva)

Master :
* [![Circle CI](https://circleci.com/gh/nlamirault/shiva/tree/master.svg?style=svg)](https://circleci.com/gh/nlamirault/shiva/tree/master)

Develop :
* [![Circle CI](https://circleci.com/gh/nlamirault/shiva/tree/develop.svg?style=svg)](https://circleci.com/gh/nlamirault/shiva/tree/develop)

*Shiva* provides a simple DHCP server.

Storage for IP addresses could be :
- Distributed : using [Consul][], [Etcd][] or [Zookeeper][]
- Local : [BoltDB][]

A REST Api is provided to display informations.


## Usage


## Development

* Initialize environment

        $ make init

* Build tool :

        $ make build

* Launch unit tests :

        $ make test

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md).


## License

See [LICENSE](LICENSE) for the complete license.


## Changelog

A [changelog](ChangeLog.md) is available


## Contact

Nicolas Lamirault <nicolas.lamirault@gmail.com>


[badge-license]: https://img.shields.io/badge/license-Apache2-green.svg?style=flat
