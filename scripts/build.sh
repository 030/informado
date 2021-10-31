#!/bin/bash -e
TRAVIS_TAG="${TRAVIS_TAG:-local}"
SHA512_CMD="${SHA512_CMD:-sha512sum}"
DELIVERABLE="${DELIVERABLE:-informado}"

echo "TRAVIS_TAG: '$TRAVIS_TAG' DELIVERABLE: '$DELIVERABLE'"
cd cmd/informado
go build -ldflags "-X informado/cmd.Version=${TRAVIS_TAG}" -o "${DELIVERABLE}"
$SHA512_CMD "${DELIVERABLE}" >"${DELIVERABLE}.sha512.txt"
chmod +x "${DELIVERABLE}"
