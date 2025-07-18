#!/bin/bash

#set -xe

root_dir=$HOME/code/projects/ahsanul-qawaid/audio/basic-letters/

letters="alif
ba
ta
tha
jeem
hha
kha
daal
thaal
ra
za
seen
sheen
saad
dhwaad
twa
zwa
ain
ghain
fa
qaaf
kaaf
laam
meem
noon
waw
ha
hamza
ya
"

n_letter=29
i=0
for l in $letters;
do
  i=$(( i + 1))
  if [ "$i" -lt "10" ];
  then
    mv "${root_dir}${l}.mp3" "${root_dir}0${i}.mp3"
  else
    mv "${root_dir}${l}.mp3" "${root_dir}${i}.mp3"
  fi
done
