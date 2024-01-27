curl -s https://learn.01founders.co/assets/superhero/all.json | jq ' .[] | select(.id == 170 )|.name, .powerstats.power, .appearance.gender' 
#curl -s https://learn.01founders.co/assets/superhero/all.json | jq ' .[] | select(.id == 170 ).powerstats.power'
#curl -s https://learn.01founders.co/assets/superhero/all.json | jq ' .[] | select(.id == 170 ).appearance.gender'