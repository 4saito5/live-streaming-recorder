# 使い方：
#  Dcokcer起動[Up]
#    ./docker.sh -u
#
#  Dcokcer終了[Down]
#    ./docker.sh -d

#!/bin/zsh


# shopt -s expand_aliases
# source ~/.zshrc
SCRIPT_DIR=$(cd $(dirname $0); pwd)
cd $SCRIPT_DIR

abort() { echo "$*" >&2; exit 1; }
unknown() { abort "unrecognized option '$1'"; }
required() { [ $# -gt 1 ] || abort "option '$1' requires an argument"; }

IS_LIMA=`which lima`
# echo ${IS_LIMA}
if ! echo "${IS_LIMA}" | grep -q "/lima"; then
  # echo "lima not found"
  DOCKER_COMPOSE="docker-compose"
else
  # echo "lima found"
  DOCKER_COMPOSE="lima nerdctl compose"
fi

ARG_ENV=''

up() {
    echo "up ..."
    # docker-compose -f ./prov/docker/docker-compose.yml up -d --build
    ${DOCKER_COMPOSE} -f ./prov/docker/docker-compose.yml up -d
}

down() {
    echo "down ..."
    ${DOCKER_COMPOSE} -f ./prov/docker/docker-compose.yml down
}

while [ $# -gt 0 ]; do
  case $1 in
    -u | --up ) up ;;
    -d | --down ) down ;;
    -?*) unknown "$@" ;;
    *) break
  esac
  shift
done
