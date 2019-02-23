git add .
echo "enter msg commit"
read $MSG 

git commit -m $MSG
git push 
