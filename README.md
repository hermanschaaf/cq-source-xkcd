# CloudQuery XKCD Source Plugin

[![test](https://github.com/cloudquery/cq-source-xkcd/actions/workflows/test.yaml/badge.svg)](https://github.com/cloudquery/cq-source-xkcd/actions/workflows/test.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/cloudquery/cq-source-xkcd)](https://goreportcard.com/report/github.com/cloudquery/cq-source-xkcd)

This CloudQuery source plugin fetches data from the [XKCD API](https://xkcd.com/json.html), allowing you to load the XKCD comic data into any CloudQuery-supported destination (e.g. PostgreSQL, Elasticsearch, Snowflake, etc.).

## Links

 - [CloudQuery Quickstart Guide](https://www.cloudquery.io/docs/quickstart)
 - [Supported Tables](docs/tables/README.md)

## Configuration

The following source configuration file will sync all comics to a local SQLite database. See [the CloudQuery Quickstart](https://www.cloudquery.io/docs/quickstart) for more information on how to configure the source and destination.

```yaml
kind: source
spec:
  name: xkcd
  path: hermanschaaf/xkcd
  version: v1.0.0
  destinations: ["sqlite"]
  spec:
---
kind: destination
spec:
  name: sqlite
  path: cloudquery/sqlite
  version: v1.2.3
  spec:
    connection_string: ./db.sqlite
```
