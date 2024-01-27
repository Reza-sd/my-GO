#find -name "*.sh" | cut -b 3- | rev | cut -c 4- | rev | sort -r
#find -name "*.sh"
find  -name "*.sh" -type f -printf "%f\n" |  rev | cut -c 4- | rev | sort -r