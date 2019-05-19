#!/bin/bash

git add .

cmit=""
if [ "$1" == '' ];
then cmit="commit at"
else cmit="commit problem $1 at"
fi

msg="$cmit `date +%FT%H:%M:%S`"

git commit -m "$msg"

git pull origin -r

git push origin master