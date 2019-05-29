# The binary to build (just the basename).
BIN := otowari

# Where to push the docker image.
REGISTRY ?= otoware_kuop

# This version-strategy uses git tags to set the version string
VERSION := $(shell git describe --tags --always --dirty)
