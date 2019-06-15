#!/bin/bash

git add .

cmit=""
if [ "$1" == '' ];
then cmit="commit at"
else cmit="commit problem $1 at"
fi

msg="$cmit `date +%F%H:%M:%S`"

git commit -m "$msg"
