#!/bin/bash -ex
tinygo flash --target=badger2040-w --size=short --stack-size=8kb .
tinygo monitor
