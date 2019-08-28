#!/bin/bash

test ! -e dev.env && cp .docker/env/dev.env dev.env

fresh
