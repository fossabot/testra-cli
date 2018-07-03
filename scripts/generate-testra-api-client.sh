#!/bin/bash

wget -O swagger.yaml https://raw.githubusercontent.com/testra-tech/testra-api/master/swagger.yaml

swagger generate client -f swagger.yaml -A testra -t ../api

rm -f swagger.yaml