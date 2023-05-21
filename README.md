# mclogs

[![go report card](https://goreportcard.com/badge/github.com/joanlopez/mclogs)](https://goreportcard.com/report/github.com/joanlopez/mclogs)  [![docker download](https://img.shields.io/docker/pulls/joanlopez/mclogs.svg)](https://hub.docker.com/r/joanlopez/mclogs)

mclogs is a fake log generator, heavily inspired on [flog](https://github.com/mingrammer/flog) ❤️, with a wink to [McDonald's](https://mcdonalds.com/), and support for common log formats json and logfmt.

It is useful for testing some tasks which require log data like amazon kinesis log stream test.

> Thanks to [gofakeit](https://github.com/brianvoe/gofakeit) and [weightedrand](https://github.com/mroth/weightedrand/) 😘

## Installation

### Using go get

```bash
go get -u github.com/joanlopez/mclogs
```

### Using .tar.gz archive

Download gzip file from [Github Releases](https://github.com/joanlopez/mclogs/releases/latest) according to your OS. Then, copy the unzipped executable to under system path.

### Using [Docker](https://www.docker.com)

```
docker run -it --rm joanlopez/mclogs
```

## Usage

There are useful options. (`mclogs --help`)

```console
Options:
  -f, --format string      log format. available formats:
                           - json (default)
                           - logfmt
  -o, --output string      output filename. Path-like is allowed. (default "stdout")
  -n, --number integer     number of lines to generate.
  -b, --bytes integer      size of logs to generate (in bytes).
                           "bytes" will be ignored when "number" is set.
  -s, --sleep duration     fix creation time interval for each log (default unit "seconds"). It does not actually sleep.
                           examples: 10, 20ms, 5s, 1m
  -d, --delay duration     delay log generation speed (default unit "seconds").
                           examples: 10, 20ms, 5s, 1m
  -w, --overwrite          overwrite the existing log files.
  -l, --loop               loop output forever until killed.
```

```console
# Generate 100 lines of logs to stdout
$ mclogs

# Generate 200 lines of logs with a time interval of 10s for each log. It doesn't actually sleep while generating
$ mclogs -s 10s -n 200 

# Generate a log file with 1000 lines of logs, then overwrite existing log file
$ mclogs -n 1000 -o generated.log -w
```

## Supported Formats

- JSON
- Logfmt

## Supported Outputs

- Stdout
- File

## License

[MIT](LICENSE)
