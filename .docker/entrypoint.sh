#!/bin/bash

if [ ! -f ".env" ]; then
  cp .env.example .env
fi

go test ./...

tail -f /dev/null