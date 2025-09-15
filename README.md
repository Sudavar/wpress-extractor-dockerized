# WPress-Extractor Dockerized
A simple docker image that allows you to extract .wpress files created by the awesome All-in-one-Wp-Migration Wordpress plugin

## Dockerized Execution
First just build the image:
```bash
docker build -t wpress-extractor .
```

Then run the container, passing the desired .wpress file as an argument:
```bash
wpress_file="small-test.wpress"
docker run -v $(pwd)/${wpress_file}:/tmp/${wpress_file} -v /tmp/extractions:/tmp/extractions wpress-extractor:latest /tmp/${wpress_file}
```

or utilize our wrapper script:
```bash
./extract.sh small-test.wpress
```

### Credits and Disclaimer
The extractor source code: [https://github.com/yani-/wpress](https://github.com/yani-/wpress).

The All-in-one-Wp-Migration plugin: [https://wordpress.org/plugins/all-in-one-wp-migration/](https://wordpress.org/plugins/all-in-one-wp-migration/).
