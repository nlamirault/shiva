# Shiva

[![License Apache 2][badge-license]](LICENSE)
[![GitHub version](https://badge.fury.io/gh/nlamirault%2Fshiva.svg)](https://badge.fury.io/gh/nlamirault%2Fshiva)

Master :
* [![Circle CI](https://circleci.com/gh/nlamirault/shiva/tree/master.svg?style=svg)](https://circleci.com/gh/nlamirault/shiva/tree/master)

Develop :
* [![Circle CI](https://circleci.com/gh/nlamirault/shiva/tree/develop.svg?style=svg)](https://circleci.com/gh/nlamirault/shiva/tree/develop)

*Shiva* provides a simple DHCP server.

Storage for IP addresses could be :
- Distributed : using [Redis][], [Consul][], [Etcd][] or [Zookeeper][] (*TODO*)
- Local : [BoltDB][]

A REST Api is provided to display informations.

## Installation

You can download the binaries :

* Architecture i386 [ [linux](https://bintray.com/artifact/download/nlamirault/oss/shiva-0.1.0_linux_386) / [darwin](https://bintray.com/artifact/download/nlamirault/oss/shiva-0.1.0_darwin_386) / [freebsd](https://bintray.com/artifact/download/nlamirault/oss/shiva-0.1.0_freebsd_386) / [netbsd](https://bintray.com/artifact/download/nlamirault/oss/shiva-0.1.0_netbsd_386) / [openbsd](https://bintray.com/artifact/download/nlamirault/oss/shiva-0.1.0_openbsd_386) / [windows](https://bintray.com/artifact/download/nlamirault/oss/shiva-0.1.0_windows_386.exe) ]
* Architecture amd64 [ [linux](https://bintray.com/artifact/download/nlamirault/oss/shiva-0.1.0_linux_amd64) / [darwin](https://bintray.com/artifact/download/nlamirault/oss/shiva-0.1.0_darwin_amd64) / [freebsd](https://bintray.com/artifact/download/nlamirault/oss/shiva-0.1.0_freebsd_amd64) / [netbsd](https://bintray.com/artifact/download/nlamirault/oss/shiva-0.1.0_netbsd_amd64) / [openbsd](https://bintray.com/artifact/download/nlamirault/oss/shiva-0.1.0_openbsd_amd64) / [windows](https://bintray.com/artifact/download/nlamirault/oss/shiva-0.1.0_windows_amd64.exe) ]
* Architecture arm [ [linux](https://bintray.com/artifact/download/nlamirault/oss/shiva-0.1.0_linux_arm) / [freebsd](https://bintray.com/artifact/download/nlamirault/oss/shiva-0.1.0_freebsd_arm) / [netbsd](https://bintray.com/artifact/download/nlamirault/oss/shiva-0.1.0_netbsd_arm) ]


## Usage

We will use [bat](https://github.com/astaxie/bat) to make HTTP request

* Install [dhcping][]

* Launch the server:

        $ sudo bin/shiva -d -backend boltdb -url "/tmp/shiva.db"
        [DEBUG] [shiva] Init BoltDB storage : /tmp/shiva.db
        [DEBUG] [shiva] Creating web service using DB<"/tmp/shiva.db">
        [DEBUG] [shiva] Creates webservice with backend : DB<"/tmp/shiva.db">
        [INFO] [shiva] Launch Shiva on 8080 using boltdb backend

* Simulate a DHCP request with unknown MAC address:

        $ sudo dhcping -s 127.0.0.1 -h 00:20:18:56:29:8f
        no answer

        [INFO] [shiva] DHCP Request from 00:20:18:56:29:8f
        [INFO] [shiva] Find :
        [INFO] [shiva] MAC Address unknown


* Add the MAC to the database :

        $ bat http://127.0.0.1:8089/api/v1/address mac=00:20:18:56:29:8f
        POST /api/v1/address HTTP/1.1
        Host: 127.0.0.1:8089
        Accept: application/json
        Accept-Encoding: gzip, deflate
        Content-Type: application/json
        User-Agent: bat/0.1.0


        {"mac":"00:20:18:56:29:8f"}


        HTTP/1.1 200 OK
        Content-Type : application/json; charset=utf-8
        Date : Thu, 25 Feb 2016 00:15:26 GMT
        Content-Length : 27


        {
            "mac": "00:20:18:56:29:8f"
        }

        $ bat http://127.0.0.1:8089/api/v1/address/00:20:18:56:29:8f
        GET /api/v1/address/00:20:18:56:29:8f HTTP/1.1
        Host: 127.0.0.1:8089
        Accept: application/json
        Accept-Encoding: gzip, deflate
        User-Agent: bat/0.1.0

        HTTP/1.1 200 OK
        Content-Type : application/json; charset=utf-8
        Date : Thu, 25 Feb 2016 00:17:16 GMT
        Content-Length : 27

        {
            "mac": "00:20:18:56:29:8f"
        }

* Perform another DHCP request :

        $ sudo dhcping -v -s 127.0.0.1 -h 00:20:18:56:29:8f
        Got answer from: 127.0.0.1

        [INFO] [shiva] DHCP Request from 00:20:18:56:29:8f
        [INFO] [shiva] Find : {"mac":"00:20:18:56:29:8f"}
        [INFO] [shiva] Assign IP 192.1.1.1 address to MAC Address 00:20:18:56:29:8f
        [INFO] [shiva] DHCP Release from 00:20:18:56:29:8f


## Development

* Initialize environment

        $ make init

* Build tool :

        $ make build

* Start backends :

        $ docker run -d -p 6379:6379 --name redis redis:3
        $ docker run -d -p 27017:27017 --name mongo mongo:3.1

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

[Redis]: http://redis.io/
[Consul]: https://www.consul.io/
[Etcd]: https://github.com/coreos/etcd
[Zookeeper]: https://zookeeper.apache.org/
[BoltDB]: https://github.com/boltdb/bolt

[dhcping]: http://www.mavetju.org/unix/general.php
