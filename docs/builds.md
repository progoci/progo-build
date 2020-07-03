Builds
======

Structure
----------

* A build consists of services
* A service consist of a docker image and steps
* A step consists of a plugin and a set of options and commands

Plugins
------

At the moment, only the plugin `Command` is available. It allows to run bash
command lines.

Under the hood
--------------

Every time that a build is created, the following things occur.

* A docker network is created with ID in the form `buildID.progo`. This network
  is useful to prevent containers from different builds to communicate with each
  other.
* At the moment, only one service per build can be created. However, the code
  was written in a way to support multiple services per build in the future.
  A service is simply a docker container with the given image and steps.
