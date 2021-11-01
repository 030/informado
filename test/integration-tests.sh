#!/bin/bash -e

TOOL="${1:-./informado}"
DELIVERABLE="${DELIVERABLE:-informado}"

validate() {
  if [ -z "${TOOL}" ]; then
    echo "No deliverable defined. Assuming that 'go run main.go'
should be run."
    TOOL="go run main.go"
  fi
}

build() {
  source ./scripts/build.sh
}

cleanup() {
  echo "cleanup"
}

main() {
  validate
  build
}

trap cleanup EXIT
main
