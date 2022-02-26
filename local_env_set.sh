while read line; do
    export $line
done < ./.env
# don't forget place a new line at the end of .env file!!