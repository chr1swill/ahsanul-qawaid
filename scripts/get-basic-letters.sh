#!/bin/bash

#set -xe

root_dir=$HOME/code/projects/ahsanul-qawaid/
audio_dir=${root_dir}audio/basic-letters/
website=https://myelm.co.uk/aq-new/

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

mkdir -p $audio_dir

for l in $letters;
do
  #echo "${website}_001/${l}.mp3"
  curl "${website}_001/${l}.mp3" -o "${audio_dir}${l}.mp3" &
  #echo -e "url:${website}_001/${l}.mp3\n\toutput:${audio_dir}${l}.mp3"
done

wait
exit 0
