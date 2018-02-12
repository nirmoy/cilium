#!/usr/bin/env bash

function log_msg {
  MSG="$1"
  echo "***************************************"
  echo "*"
  echo "*"
  echo "* ${MSG}"
  echo "*"
  echo "*"
  echo "***************************************"
}

function retry_function {
  set +e
  FUNC="$1"
  COUNTER=0
  MAX_TRIES=10
  echo "beginning trying up to ${MAX_TRIES} times function \"${FUNC}\""
  while [ $COUNTER -lt $MAX_TRIES ]; do
    log_msg "on attempt ${COUNTER} of function \"${FUNC}\""
    ${FUNC}
    if [[ "$?" == "0" ]] ; then
      echo "running of \"${FUNC}\" successful"
      echo 
      echo
      echo
      set -e
      return 0
    fi
    sleep 1
    let COUNTER=COUNTER+1
  done

  log_msg "running function \"${FUNC}\" ${MAX_TRIES} times did not succeed"
  set -e
  return 1
}

export GOPATH=/home/vagrant/go
echo 'export PATH=/home/vagrant/go/bin:/usr/local/clang/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games' >> /home/vagrant/.bashrc
export PATH=/home/vagrant/go/bin:/usr/local/clang/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games
