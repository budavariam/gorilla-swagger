#!/bin/bash

git submodule add https://github.com/swagger-api/swagger-ui.git external/swagger-ui
mkdir -p swaggerui
cp -r external/swagger-ui/dist/ ./swaggerui

DEFAULT_VALUE="https://petstore.swagger.io/v2/swagger.json"
LOCAL_GENERATED_FILE="./swagger.json"
sed -i "s_${DEFAULT_VALUE}_${LOCAL_GENERATED_FILE}_" ./swaggerui/swagger-initializer.js