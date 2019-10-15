#!/bin/bash

curl -s https://raw.githubusercontent.com/mukkhit/superhero-api/master/api/all.json | jq -M 'map(select(.id == 70)) | .[].name'