#!/bin/bash

test ! -e .env && cp .docker/env/dev.env .env

fresh
