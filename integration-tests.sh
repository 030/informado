#!/bin/bash -e

TOOL="${1:-./informado}"
DELIVERABLE="${DELIVERABLE:-informado}"

validate(){
    if [ -z "${TOOL}" ]; then
        echo "No deliverable defined. Assuming that 'go run main.go' 
should be run."
        TOOL="go run main.go"
    fi
}

build(){
  echo "TRAVIS_TAG: '$TRAVIS_TAG' DELIVERABLE: '$DELIVERABLE'"
  go build -ldflags "-X informado/cmd.Version=${TRAVIS_TAG}" -o "${DELIVERABLE}"
  $SHA512_CMD "${TOOL}" > "${DELIVERABLE}.sha512.txt"
  chmod +x "${DELIVERABLE}"
}

cleanup(){
  echo "cleanup"
}

main(){
  validate
  build
}

trap cleanup EXIT
main