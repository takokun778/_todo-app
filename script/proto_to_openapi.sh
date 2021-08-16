#!/bin/bash

TMPDIR=generated/tmp

OUTDIR=generated/openapiv3

rm -rf ${OUTDIR}

mkdir ${TMPDIR} 

# protoc を利用して proto ファイルから openapi v2 の json コードを生成
protoc ./proto/proto/user/**/*.proto --openapiv2_out ./${TMPDIR} --openapiv2_opt logtostderr=true

protoc ./proto/proto/task/**/*.proto --openapiv2_out ./${TMPDIR} --openapiv2_opt logtostderr=true

mkdir ${OUTDIR}

# api を利用して version2 の json ファイルから v3 のyaml ファイルを生成
curl -X POST "https://converter.swagger.io/api/convert" \
     -H "Content-Type: application/json" \
     -H "Accept: application/yaml" \
     -d @./${TMPDIR}/proto/proto/task/v1/task.swagger.json \
     -o ./${OUTDIR}/task.openapi.yaml

curl -X POST "https://converter.swagger.io/api/convert" \
     -H "Content-Type: application/json" \
     -H "Accept: application/yaml" \
     -d @./${TMPDIR}/proto/proto/user/v1/user.swagger.json \
     -o ./${OUTDIR}/user.openapi.yaml

rm -rf ${TMPDIR}
