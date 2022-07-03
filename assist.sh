#!/usr/bin/env bash 

NC=$'\e[0m' # No Color
BOLD=$'\033[1m'
UNDERLINE=$'\033[4m'
RED=$'\e[31m'
GREEN=$'\e[32m'
BLUE=$'\e[34m'
ORANGE=$'\x1B[33m'

organization="learning-cloud"
project="ephemeral-iaac"

function check_stack_exist(){
    stack=$1
    stack_count=$(pulumi stack  ls | grep -v NAME | grep $stack | wc -l)
    if [ $stack_count = 0 ];then
        echo -e "${RED}Stack $stack Does Not Exists. Exiting...${NC}\n" 
        return 1
    fi
}

function setup(){
    org_proj_stack=$1
    check_stack_exist $org_proj_stack || exit 1
    pulumi stack init $org_proj_stack
    pulumi config set aws:region us-east-1 # any valid AWS region will work
    pulumi stack select $org_proj_stack
    echo -e "${GREEN}Stack $org_proj_stack created successfully${NC}\n" 
}

function teardown(){
    org_proj_stack=$1
    check_stack_exist $org_proj_stack || exit 1
    pulumi destroy --yes
    pulumi stack select $org_proj_stack
    pulumi stack rm $org_proj_stack --yes
    echo -e "${GREEN}Stack $org_proj_stack destroyed successfully${NC}\n" 
}

opt="$1"
stack="${2:-dev}"
org_proj_stack="$organization/$project/$stack"

choice=$( tr '[:upper:]' '[:lower:]' <<<"$opt" )

echo -e "\n${BLUE}Executing Action : $choice | Env: $org_proj_stack ${NC}"

case ${choice} in
    "setup")
        setup $org_proj_stack
    ;;
    "teardown")
        teardown $org_proj_stack
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