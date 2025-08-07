# Gator 

Gator is a cli scraper for rss feeds. Allowing you to fetch rss feeds and view them in your terminal. 

## Pre-requisites

- Install [go](https://go.dev/doc/install) with a minimum version of 1.24.4
- Install [postgres](https://www.postgresql.org/download/)

- Use `go install github.com/nacen-dev/gator`

## Running the project

`gator <command>`

### List of commands

- login: gator login <username>
- register: gator register username
- reset: gator reset
- users: gator users
- agg: gator agg (<time_duration> e.g. 1m - 1 minute, 30s - 30 seconds, 1h - 1hour)
- addfeed: gator addfeed <feed_name> <feed_url>
- feeds: gator feeds
- follow: gator follow <feed_url>
- following: gator following
- unfollow: gator unfollow <feed_url>
- browse: gator browse (<optional_limit> the default is 2 if it is not provided)

