#!/bin/bash

DOWNLOAD_URL='https://codeload.github.com/BBC-News/wraith/zip/master'
CWD=`pwd`

## Helpers
is_osx() {
  [ `/usr/bin/sw_vers -productVersion 2>/dev/null` ]
}

is_brew_installed() {
  [ `which brew` ]
}

ESC_SEQ="\x1b["
COL_RESET=$ESC_SEQ"39;49;00m"
COL_RED=$ESC_SEQ"31;01m"
COL_GREEN=$ESC_SEQ"32;01m"
COL_YELLOW=$ESC_SEQ"33;01m"

red() {
  echo -e "${COL_RED}${@}${COL_RESET}"
}

yellow() {
  echo -e "${COL_YELLOW}${@}${COL_RESET}"
}

green() {
  echo -e "${COL_GREEN}${@}${COL_RESET}"
}

colour() {
  case $1 in
    "reset") echo "\033[0m$@" ;;
    "blue") echo "\033[1;34m$@" ;;
  esac
}




download() {
  green "Downloading Wraith"
  curl --progress-bar $DOWNLOAD_URL --output $CWD/wraith.zip
}

extract() {
  echo ""
  green "Extracting to current directory"
  unzip -qq $CWD/wraith.zip
  mv wraith-master wraith
  rm -f $CWD/wraith.zip
}

install_brew_packages() {
  if [ -z `which convert` ]; then
    brew install imagemagick
  fi

  if [ -z `which phantomjs` ]; then
    brew install phantomjs
  fi
}

download
extract

if is_osx; then
  if is_brew_installed; then
    install_brew_packages
  else
    red "Homebrew isn't installed, please install it. http://brew.sh"
    exit
  fi
else
  if [ -z `which convert` ]; then
    red "Missing ImageMagick"
  fi

  if [ -z `which phantomjs` ]; then
    red "Missing phantomjs"
  fi
fi

green "You can now run Wraith, awesome!"

# vi: set ft=bash