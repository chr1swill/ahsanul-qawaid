#!/bin/bash

#set -xe

PAGE_MIN=1
PAGE_MAX=56

root_dir=$HOME/code/projects/ahsanul-qawaid/
audio_dir=${root_dir}audio/
website=https://myelm.co.uk/aq-new/

page_number="$1"
if [ "$page_number" == "" ];
then
  echo "parameter 1 empty: expected page number" && exit 1
elif [ "$page_number" -lt "$PAGE_MIN" ];
then
  echo "parameter 1 less than ${PAGE_MIN}: expected number greater then or equal to ${PAGE_MIN}"  && exit 2
elif [ "$page_number" -gt "$PAGE_MAX" ];
then
  echo "parameter 1 greater than ${PAGE_MAX}: expected number less than or equal to ${PAGE_MIN}"  && exit 3
fi

page_dir="${audio_dir}${page_number}/"
if [ ! -d "${page_dir}" ];
then
  mkdir -p "$page_dir"
fi

page_path=""
if [ "$page_number" -lt "10" ];
then
  page_path="_00${page_number}/"
else
  page_path="_0${page_number}/"
fi

# TODO: max iters 100 is just a arbitrary number 
# probably should fact check 
# to not waste time with unnessary requests
for (( i=1; i<100; ++i ));
do
  file=""

  if [ "$i" -lt "10" ];
  then
    file="_0${i}"
  else
    file="_${i}"
  fi

  curl "${website}${page_path}${file}.mp3" -o "${page_dir}${i}.mp3" &
done

wait

exit 0
