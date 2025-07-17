#!/bin/bash

set -xe

root=$HOME/code/projects/ahsanul-qawaid

audio_program=mpg321

case $1 in
  "اَ") $audio_program ${root}/11/01.mp3 ;;
  "بَ") $audio_program ${root}/11/01.mp3 ;;
  "تََ) $audio_program ${root}/11/01.mp3 ;;
  *) echo "we got nothing for ${1} buddy" ;;
esac

exit 0
