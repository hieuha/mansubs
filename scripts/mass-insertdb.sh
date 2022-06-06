while read -r domain;
do
   echo $domain | assetfinder -subs-only | mansubs -domain "$domain" -create;
done < domains.txt
