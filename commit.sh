git add .

msg="commit at `date +%FT%H:%M:%S`"

git commit -m "$msg"

git pull origin -r

git push origin master