#!/bin/bash

# Creates an initial revision for the database.
alembic -c database/migrations/alembic.ini revision -m "init"
