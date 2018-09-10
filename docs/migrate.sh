#!/usr/bin/env bash


if [[ "$1" =~ ^(--help|help|-help)$ ]]; then
  echo "migrate a database using docker flyway container. requires env vars PG_HOST, PG_PORT, PG_PASSWORD, PG_USER"
  exit 0
fi

if [[ -z "$PG_HOST" ]]; then
  echo "PG_HOST required"
  exit 1
fi
if [[ -z "$PG_PORT" ]]; then
  echo "PG_PORT required"
  exit 1
fi

if [[ -z "$PG_PASSWORD" ]]; then
  echo "$PG_PASSWORD required"
  exit 1
fi

if [[ -z "$PG_USER" ]]; then
  echo "PG_USER required"
  exit 1
fi


docker run --rm -v $(pwd)/sql:/flyway/sql boxfuse/flyway migrate -url=jdbc:postgresql://${PG_HOST}:5432/${PG_USER} -user=${PG_USER} -password=${PG_PASSWORD} -baselineOnMigrate=true