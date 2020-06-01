#!/bin/bash
cd "$(dirname $0)"
if ! test -e jaqy.jar
then curl -sL https://github.com/Teradata/jaqy/releases/download/v1.1.0/jaqy-1.1.0.jar >jaqy.jar
fi
if ! test -e postgresql.jar
then curl -sL https://jdbc.postgresql.org/download/postgresql-42.2.12.jar >postgresql.jar
fi
echo "select * from messages;" | java -jar jaqy.jar --rcfile connect.rc

