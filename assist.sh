#!/usr/bin/env bash 

NC=$'\e[0m' # No Color
BOLD=$'\033[1m'
UNDERLINE=$'\033[4m'
RED=$'\e[31m'
GREEN=$'\e[32m'
BLUE=$'\e[34m'
ORANGE=$'\x1B[33m'

function setup(){
    stack=$1
    pulumi stack init learning-cloud/$stack
    pulumi config set aws:region us-east-1 # any valid AWS region will work
    pulumi stack select learning-cloud/ephemeral-iaac/$stack
}

function teardown(){
    pulumi destroy --yes
    pulumi stack rm learning-cloud/$stack --yes
}

opt="$1"
stack="${2:-dev}"
choice=$( tr '[:upper:]' '[:lower:]' <<<"$opt" )

case ${choice} in
    "setup")
        setup $stack
    ;;
    "teardown")
        teardown $stack
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