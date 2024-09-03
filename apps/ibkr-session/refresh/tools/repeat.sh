#!/bin/sh
DELAY_SEC=60
while true
do
    sleep ${DELAY_SEC}
    java -jar /usr/bin/app-standalone.jar
done