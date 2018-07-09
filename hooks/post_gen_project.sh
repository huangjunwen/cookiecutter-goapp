#!/bin/bash

# Creates an initial revision for the database.
alembic -c database/migrations/alembic.ini revision -m "init"

# Replace import path.
baseimport=${PWD#$GOPATH/src/}
find . -type f -name "*.go" -exec sed -i.tmpbak "s#%%baseimport%%#$baseimport#g" {} \;
find . -type f -name "*.go.tmpbak" -exec rm {} \;
