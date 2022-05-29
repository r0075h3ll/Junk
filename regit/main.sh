githubInfo=$1
target=$2
lang=$3
query=$4

result=`curl -s -u $githubInfo https://api.github.com/search/code?q=\"$target\"+$query+in:file+language:$lang&s=indexed`
resultCount=`echo $result | jq '.total_count'`

if [[ $? -ne 0 ]]; then
	printf "\n[-] jq not installed\n"
	exit
fi


for (( index=0; index<=$resultCount; index++ ))
do
	url=`echo $result | jq -r --argjson ind $index '.items[$ind].html_url' | sed -e 's/blob/raw/g'`

	if [[ $url == "null" ]]; then
		printf "\n[!] Links will be saved in file: link_to_files.txt"
		printf "\n[-] You hit the rate limit, try again after a couple of minutes\n"
		break
	else
		echo $url >> link_to_files.txt
	fi
done


if [ -f "link_to_files.txt" ]; then
	sort -u link_to_files.txt -o link_to_files.txt
fi