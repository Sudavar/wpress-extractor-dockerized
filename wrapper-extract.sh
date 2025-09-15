#!/bin/bash

set -e

if [ -z "${1}" ]; then
  echo "Please provide an wpress archive as argument"
  exit 1
fi
archive="${1}"
# create full path to archive variable
archive_full_path=$(readlink -f "${1}")

extraction_path=$(pwd)"/extractions"
mkdir -p ${extraction_path}

docker run --user $(id -u):$(id -g) -v ${archive_full_path}:/tmp/${archive} -v ${extraction_path}:${extraction_path} wpress-extractor --input /tmp/${archive} --output ${extraction_path}
