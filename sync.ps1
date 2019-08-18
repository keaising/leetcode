git add -A

$cmit=""
if ($null -eq $args[0]) {
    $cmit="commit at "
} else {
    $cmit="commit problem $($args[0]) at "
} 

$msg=$cmit + (Get-Date -Format "yyyy-MM-ddTHH:mm:ss")

Write-Output $msg

git commit -m $msg

git pull origin master -r

git push origin master
