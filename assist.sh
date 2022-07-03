#!/usr/bin/env bash 

NC=$'\e[0m' # No Color
BOLD=$'\033[1m'
UNDERLINE=$'\033[4m'
RED=$'\e[31m'
GREEN=$'\e[32m'
BLUE=$'\e[34m'
ORANGE=$'\x1B[33m'

function check_stack_exist(){
    stack=$1
    stack_count=$(pulumi stack  ls | grep -v NAME | grep $stack | wc -l)
    if [ $stack_count = 0 ];then
        echo -e "${RED}Stack $stack Does Not Exists. Exiting...${NC}\n" 
        return 1
    fi
}

function setup(){
    stack=$1
    pulumi stack init learning-cloud/ephemeral-iaac/$stack
    pulumi config set aws:region us-east-1 # any valid AWS region will work
    pulumi stack select learning-cloud/ephemeral-iaac/$stack
    echo -e "${GREEN}Stack $stack created successfully${NC}\n" 
}

function teardown(){
    stack=$1
    check_stack_exist $stack || exit 1
    pulumi destroy --yes
    pulumi stack select learning-cloud/ephemeral-iaac/$stack
    pulumi stack rm learning-cloud/ephemeral-iaac/$stack --yes
    echo -e "${GREEN}Stack $stack destroyed successfully${NC}\n" 
}

opt="$1"
stack="${2:-dev}"
choice=$( tr '[:upper:]' '[:lower:]' <<<"$opt" )

echo -e "\n${BLUE}Executing Action : $choice | Env: $stack ${NC}"

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