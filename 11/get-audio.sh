#!/bin/bash

set -xe

for (( i=0; i<=45; ++i )) 
do
  if [ $i -lt "10" ];
  then
    wget "https://myelm.co.uk/aq-new/_011/_0${i}.mp3" &
  else
    wget "https://myelm.co.uk/aq-new/_011/_${i}.mp3" &
  fi
done

wait
exit 0
