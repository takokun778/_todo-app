#!/bin/bash

# typescript
# axios
# task
rm -rf ./generated/typescript/axios/task

docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:latest generate -i /local/generated/openapiv3/task.openapi.yaml -g typescript-axios -o /local/generated/typescript/axios/task --additional-properties=supportsES6=true

# user
rm -rf ./generated/typescript/axios/user

docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:latest generate -i /local/generated/openapiv3/user.openapi.yaml -g typescript-axios -o /local/generated/typescript/axios/user --additional-properties=supportsES6=true
