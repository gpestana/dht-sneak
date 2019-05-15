# BitTorrent privacy

## The BitTorrent Anonymity Marketplace

The BitTorrent Anonymity Marketplace consists of a protocol in which peers participate in k swarms (whether they are interested in the swarm's content or not), obscuring its own intent while improving the availability of content in the network (i.e. caching of unwanted content is useful to the network). The mechanism presented in the paper intends to improve plausible deniability of content requesters/caching in BitTorrent, while providing the incentives for peers to collaborate in order to increase individual and network privacy, as well as performance. Three categories are taken into consideration: privacy, performance and incentives. From an incentive perspective, participating in the privacy protocol should also help with download. The reasons for not using Bittorrent-over-Tor (as another solution for the same problem) is lack of incentives in Tor and performance degradation. The protocol consists of a peer joining k swarms per interest content (i.e. downloading and seeding k-1 'cover' files and 1 file it is interested on) and cross share files between swarms for performance and privacy purposes. Each peer uniformly behaves as relay, seeder and leecher of content that if both of its interest or not. This bakes a plausible deniability layer in BitTorrent protocol.  

## Bitblender

Aims at routing traffic through a path of relays (without encryption, as in onion routing), which makes user requests indistinguishable from the set of k peers (k being the size of the relaying path)

## OneSwarm

Uses social networks (friend to friend routing) to make sure private information is only shared with trusted peers


