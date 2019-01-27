#!/bin/bash

wget -O swagger.yaml https://raw.githubusercontent.com/testra/testra-api/master/swagger.yaml

# yarn global add yamljs (or) npm i -g yamljs
yaml2json swagger.yaml -p -i4 > swagger.json

swagger generate client -f swagger.json -A testra -t api

rm -f swagger.yaml swagger.json