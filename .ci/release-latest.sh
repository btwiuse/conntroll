#!/bin/bash

ROOT_DIR=$(dirname $(dirname $(realpath -m ${0})))

# unecessary
refresh(){
##git push --delete origin latest || true
##git tag -d latest || true
##git tag latest HEAD
##git push origin latest
  github release delete latest
}

github(){
  hub -C $ROOT_DIR ${@}
}

upload(){
  github release | grep latest || github release create -m 'latest' latest
  github release edit -m 'latest' latest -a "$1"
}

loop_unix(){
  cd ${ROOT_DIR}/bin
  ls -1d {android,darwin,linux,freebsd,openbsd,netbsd}/*/ | while read dir; do
    pushd $dir
    compressed="${OLDPWD}/${dir////-}k0s.tar.gz"
    tar cz * > "$compressed"
    upload $(realpath -m $compressed)
    popd
  done
}

loop_windows(){
  cd ${ROOT_DIR}/bin
  ls -1d windows/*/ | while read dir; do
    pushd $dir
    compressed="${OLDPWD}/${dir////-}k0s.zip"
    zip - * > "$compressed"
    upload $(realpath -m $compressed)
    popd
  done
}

main(){
  refresh
  loop_unix
  loop_windows
}

main