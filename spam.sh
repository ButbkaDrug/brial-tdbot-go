#!/bin/bash
for (( ; ; ))
do
    curl -H "Accept: text/plain" https://icanhazdadjoke.com/ | kilogram send message -d 6379992146
    sleep 10
done
