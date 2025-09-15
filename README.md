# WPress-Extractor Dockerized
A simple docker image that allows you to extract .wpress files created by the awesome All-in-one-Wp-Migration Wordpress plugin

## Dockerized Execution
First just build the image:
```bash
docker build -t wpress-extractor .
```

Then run the container with the following in mind:
- Pass the desired .wpress file as an argument to `--input`
- In order for the created files to be accessible, we need to run the container as our user
- By default output will be written to `/tmp/extractions`, you can override it with the `--output <directory>` argument but make sure it's also mounted inside the container
```bash
wpress_file="small-test.wpress"
docker run \
  --user $(id -u):$(id -g) \
  -v $(pwd)/${wpress_file}:/tmp/${wpress_file} \
  -v $(pwd)/extractions:/tmp/extractions \
  wpress-extractor \
  --input /tmp/${wpress_file}
```

For ease of use, you can also utilize our wrapper script
```bash
./wrapper-extract.sh small-test.wpress
```

### Credits and Disclaimer
The extractor source code was cloned from [https://github.com/yani-/wpress](https://github.com/yani-/wpress), directly in the repository in order to change the extract() function to accept an output directory argument. It's been 10 years since the original project was last updated (won't miss any upstream updates) and I am not very familiar with GO to fork the go package.
