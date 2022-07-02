#!/usr/bin/env bash 

NC=$'\e[0m' # No Color
BOLD=$'\033[1m'
UNDERLINE=$'\033[4m'
RED=$'\e[31m'
GREEN=$'\e[32m'
BLUE=$'\e[34m'
ORANGE=$'\x1B[33m'

function setup(){
    pulumi stack init dev
    pulumi config set aws:region us-east-1 # any valid AWS region will work
}

function teardown(){
    pulumi destroy --yes
    pulumi stack rm dev --yes
}

opt="$1"
choice=$( tr '[:upper:]' '[:lower:]' <<<"$opt" )

case ${choice} in
    "setup")
        setup
    ;;
    "teardown")
        teardown
    ;;
    *)
    echo "${RED}Usage: assist.sh < setup | teardown >  ${NC}"
cat <<-EOF
Commands:
---------
  setup       -> Build & Configure
  teardown    -> Teardown & Destroy
EOF
    ;;
esac