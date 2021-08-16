#!/bin/bash

# typescript
# fetch
# task
rm -rf ./generated/typescript/fetch/task

docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:latest generate -i /local/generated/openapiv3/task.openapi.yaml -g typescript-fetch -o /local/generated/typescript/fetch/task --additional-properties=supportsES6=true,typescriptThreePlus=true

# user
rm -rf ./generated/typescript/fetch/user

docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:latest generate -i /local/generated/openapiv3/user.openapi.yaml -g typescript-fetch -o /local/generated/typescript/fetch/user --additional-properties=supportsES6=true,typescriptThreePlus=true
