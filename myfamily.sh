curl -s https://acad.learn2earn.ng/assets/superhero/all.json \
  | jq -r --arg id "$HERO_ID" '.[] | select(.id == ($id | tonumber)) | .connections.relatives | gsub("\\n"; "\n")'
