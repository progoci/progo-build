#!/bin/bash

dep ensure

pwd
test ! -e .env && cp .docker/env/dev.env .env

watcher -run github.com/progoci/progo/services/container/cmd  -watch github.com/progoci/progo/services/container
