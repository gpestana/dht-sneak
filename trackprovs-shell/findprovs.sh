#!/bin/bash

# parses cid from input
cid=$1

# gets provider's addr
prov=$(ipfs dht findprovs ${cid} -n 1)
prov_info=$(ipfs id ${prov})
addrs=$(ipfs id -f="<addrs>")

# filter ipv4 addresses
ipv4addrs=""
for a in ${addrs}
do
	if [ ${a:3:1} = "4" ] 
	then 
		ipv4addrs="${ipv4addrs}, ${a}"
	fi
done

out=$( jq -n \
				--arg p "$prov" \
				--arg pa "[${ipv4addrs:2}]" \
				--arg cid "$cid" \
				--arg ts "$(date +%s)" \
				--arg lc "{}" \
				'{prov: $p, prov_addr: $pa, content_id: $cid, timestamp: $ts, location: $lc}' )

# note: not a valid JSON
echo "${out},"

