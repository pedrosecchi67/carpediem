#!/bin/bash

# set path back to original directory
ordir="$(pwd)"
cd "$(dirname $(realpath $0))"

# remove previous files
if [ -f PoetryFoundationData.csv ]; then
  rm PoetryFoundationData.csv
fi

if [ -f poetry-database.sqlite3 ]; then
  rm poetry-database.sqlite3
fi

# curl
curl -L -o ./poetry-foundation-poems.zip\
  https://www.kaggle.com/api/v1/datasets/download/tgdivy/poetry-foundation-poems

# unzip and remove
unzip poetry-foundation-poems.zip
rm poetry-foundation-poems.zip

# run cleaning and sqlite3 dbase
PYTHON="python3"
if [ ! -z `which "$PYTHON" | grep "no * in"` ]; then
  echo "Python not found! Database not built. Try specifying executable in curl_and_build.sh"
else  
  $PYTHON clean.py
fi

cd "$ordir"
