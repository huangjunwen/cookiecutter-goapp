#!/bin/bash

# Check docker installation.
command -v docker > /dev/null 2>&1 || { echo >&2 "docker is required but not installed. Aborting"; exit 1; }

# Check alembic installation.
command -v alembic > /dev/null 2>&1 || { echo >&2 "alembic is required but not installed. Aborting"; exit 1; }

# Check some dependencies.
pip show -q pymysql || { echo >&2 "pymysql is not installed. Aborting"; exit 1; }
