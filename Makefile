# Copyright (c) 2024 Intel Corporation
# All rights reserved.
# SPDX-License-Identifier: BSD-3-Clause

SHELL := /bin/bash
APPNAME := attestation-app-gsc
VERSION := v0.1.0

PROXY_EXISTS := $(shell if [[ "${https_proxy}" || "${http_proxy}" || "${no_proxy}" ]]; then echo 1; else echo 0; fi)
DOCKER_PROXY_FLAGS := ""
ifeq ($(PROXY_EXISTS),1)
    DOCKER_PROXY_FLAGS = --build-arg http_proxy="${http_proxy}" --build-arg https_proxy="${https_proxy}" --build-arg no_proxy="${no_proxy}"
else
    DOCKER_PROXY_FLAGS =
endif
export DOCKER_BUILDKIT=1

makefile_path := $(realpath $(lastword $(MAKEFILE_LIST)))
makefile_dir := $(dir $(makefile_path))
OUTDIR := $(addprefix $(makefile_dir),out)
TMPDIR := $(addprefix $(makefile_dir),tmp)
GSC_GITURL := https://github.com/gramineproject/gsc.git
GSC_GITCOMMIT := v1.6

.PHONY: gramine

docker:
	cat Dockerfile | docker build \
		-t $(APPNAME):$(VERSION) \
		--progress=plain \
		${DOCKER_PROXY_FLAGS} \
		-f - .

setup-gsc:
	rm -rf gsc
	git clone $(GSC_GITURL) && cd gsc && git checkout $(GSC_GITCOMMIT)
	cp gramine.yaml gsc/config.yaml

gramine: setup-gsc docker
	cp gramine.manifest gsc/
	cd gsc && openssl genrsa -3 -out gramine-enclave-key.pem 3072
	cd gsc && ./gsc build --no-cache --rm ${DOCKER_PROXY_FLAGS} $(APPNAME):$(VERSION) gramine.manifest
	cd gsc && ./gsc sign-image $(APPNAME):$(VERSION) gramine-enclave-key.pem

all: gramine clean

clean:
	if pushd $(makefile_dir); then \
		rm -rf $(OUTDIR) $(TMPDIR); \
		rm -f docker.timestamp; \
	fi;
