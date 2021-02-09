# Euterpe

``service.teapot/`` is an example service that aims to demonstrate as much functionality as possible.

Following the following conventions on API design:
- https://cloud.google.com/apis/design/standard_fields
- https://cloud.google.com/apis/design/naming_convention

## Scripts

### Generating from protos

```shell
make protos
```

## Creating migration

```shell
migrate create -ext sql -dir ./service.<servicename>/migrations <migrationname>
```