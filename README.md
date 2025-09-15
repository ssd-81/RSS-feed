# RSS-feed
RSS feed accumulator using Golang, postGreSQL 

## setup
1. Requirements: golang, postgres 
2. Gator CLI
    run the command
    - `git clone https://github.com/ssd-81/RSS-feed.git`
    - `cd RSS-feed`
    - `go install .`
    - `go build -v -o app`
3. Ready to go


## Usage
`Note: the usage is based on the generation of binaries as per the above instructions`
---
**Commands**
- login: ./app login <username>
- register: ./app register <username>
- reset: ./app reset
- users: ./app users
- agg: ./app agg <time-interval(1s, 1m, 1h)>
- addfeed: ./app addfeed <name> <url>
- feeds: ./app feeds
- follow: ./app follow
- following: ./app following
- unfollow: ./app unfollow <url>
- browse: ./app browse <number of posts>


