GIT_USER_ID = butlerhq
GIT_REPO_ID = airbyte-client-go
PACKAGE_NAME = airbyte

VERSION_TAG ?= v2.0.0

all:
	mkdir -p ${PACKAGE_NAME}
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen
	oapi-codegen -package ${PACKAGE_NAME} api/openapi.yaml > airbyte/airbyte.gen.go 

push-tag:
	git tag -a ${VERSION_TAG} -m "Tag api with version ${VERSION_TAG}"
	git push origin ${VERSION_TAG}
