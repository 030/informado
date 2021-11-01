#!/bin/bash -e
TRAVIS_TAG="${TRAVIS_TAG:-local}"
SHA512_CMD="${SHA512_CMD:-sha512sum}"
export DELIVERABLE="${DELIVERABLE:-informado}"

echo "GITHUB_TAG: '$GITHUB_TAG' DELIVERABLE: '$DELIVERABLE'"
cd cmd/informado
go build -ldflags "-X informado/cmd.Version=${GITHUB_TAG}" -o "${DELIVERABLE}"
$SHA512_CMD "${DELIVERABLE}" >"${DELIVERABLE}.sha512.txt"
chmod +x "${DELIVERABLE}"
cd ../..
