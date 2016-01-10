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


## Simulate DHCP

* Install [dhcping][]

* Launch the server:

        $ sudo bin/shiva -d -backend boltdb -url "/tmp/shiva.db"
        2015/12/13 00:04:13 [DEBUG] [shiva] Init BoltDB storage : /tmp/shiva.db
        2015/12/13 00:04:13 [DEBUG] [shiva] Creating web service using DB<"/tmp/shiva.db">
        2015/12/13 00:04:13 [DEBUG] [shiva] Creates webservice with backend : DB<"/tmp/shiva.db">
        2015/12/13 00:04:13 [INFO] [shiva] Launch Shiva on 8080 using boltdb backend
        2015/12/13 00:06:45 [INFO] [shiva] DHCP Request from 00:20:18:56:29:8f
        2015/12/13 00:11:44 [DEBUG] [shiva] Check entry exists with key :  V)�
        2015/12/13 00:11:44 [INFO] [shiva] Find :
        2015/12/13 00:11:44 [INFO] [shiva] MAC Address unknown

* Simulate a DHCP request :

        $ sudo dhcping -s 127.0.0.1 -h 00:20:18:56:29:8f
        no answer


## Contributing

See [CONTRIBUTING](CONTRIBUTING.md).


## License

See [LICENSE](LICENSE) for the complete license.


## Changelog

A [changelog](ChangeLog.md) is available


## Contact

Nicolas Lamirault <nicolas.lamirault@gmail.com>


[badge-license]: https://img.shields.io/badge/license-Apache2-green.svg?style=flat

[Consul]: https://www.consul.io/
[Etcd]: https://github.com/coreos/etcd
[Zookeeper]: https://zookeeper.apache.org/
[BoltDB]: https://github.com/boltdb/bolt

[dhcping]: http://www.mavetju.org/unix/general.php

