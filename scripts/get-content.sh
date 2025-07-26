#!/bin/bash

#set -xe

PAGE_MIN=17
PAGE_MAX=37 # loop <= PAGE_MAX

root_dir=$HOME/code/projects/ahsanul-qawaid/
audio_dir=${root_dir}audio/
website=https://myelm.co.uk/aq-new/

# page_number="$1"
# if [ "$page_number" == "" ];
# then
#   echo "parameter 1 empty: expected page number" && exit 1
# elif [ "$page_number" -lt "$PAGE_MIN" ];
# then
#   echo "parameter 1 less than ${PAGE_MIN}: expected number greater then or equal to ${PAGE_MIN}"  && exit 2
# elif [ "$page_number" -gt "$PAGE_MAX" ];
# then
#   echo "parameter 1 greater than ${PAGE_MAX}: expected number less than or equal to ${PAGE_MIN}"  && exit 3
# fi

for (( page_number=$PAGE_MIN; page_number<=$PAGE_MAX; ++page_number ));
do
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
    idx=""
  
    if [ "$i" -lt "10" ];
    then
      idx="0${i}"
    else
      idx="${i}"
    fi
  
    file="_${idx}"
    curl "${website}${page_path}${file}.mp3" -o "${page_dir}${idx}.mp3" &
  done
  
  wait
  
  # clean up extra fetch content
  for file in $(ls $page_dir);
  do
    filetype=$(file -b ${page_dir}${file} | sed 's/,.*//g')
    if [ "$filetype" == "HTML document" ];
    then
      rm ${page_dir}${file} &
    fi
  done
  
  wait
done

exit 0
