#!/bin/bash

cmit=""
if [ "$1" == '' ];
then cmit="commit at"
else cmit="commit problem $1 at"
fi

msg="$cmit `date +%FT%H:%M:%S`"

git commit -m "$msg"
