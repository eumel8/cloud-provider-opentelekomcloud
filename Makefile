# Copyright 2020 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

GOOS ?= $(shell go env GOOS)
SOURCES := $(shell find . -type f  -name '*.go')
LDFLAGS := ""

# Images management
REGISTRY_REGION?="eu-de"
ACCESS_KEY?=""
REGISTRY_LOGIN_KEY?=""

# Set you version by env or using latest tags from git
VERSION?=$(shell git describe --tags)

opentelekomcloud-controller-manager: $(SOURCES)
	CGO_ENABLED=0 GOOS=$(GOOS) go build \
		-ldflags $(LDFLAGS) \
		-o opentelekomcloud-controller-manager \
		cmd/cloud-controller-manager/cloud-controller-manager.go

clean:
	rm -rf opentelekomcloud-controller-manager

verify:
	hack/verify.sh

test:
	go test ./...
