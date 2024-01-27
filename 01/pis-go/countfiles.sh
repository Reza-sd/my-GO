#b= find . -type f | wc -l
#a=$(( b + 103 ))
#echo $a
#find -type f | wc -l
find -type f,d | wc -l
