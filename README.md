# Gator 

Gator is a cli scraper for rss feeds. Allowing you to fetch rss feeds and view them in your terminal. 

## Installation

- Make sure you have the latest [Go toolchain](https://golang.org/dl/)
- A local [Postgres](https://www.postgresql.org/download/) database. 
- 
- You can then install `gator` with:

```bash
go install github.com/nacen-dev/gator
```

## Config

Create a `.gatorconfig.json` file in your home directory with the following structure:

```json
{
  "db_url": "postgres://username:@localhost:5432/database?sslmode=disable"
}
```
Replace the values with your database connection string.

# Usage

Create a new user:

```bash
gator register <name>
```

Add a feed:

```bash
gator addfeed <url>
```

Start the aggregator:

```bash
gator agg 30s
```

View the posts:

```bash
gator browse [limit]
```

There are a few other commands you'll need as well:

- `gator login <name>` - Log in as a user that already exists
- `gator users` - List all users
- `gator feeds` - List all feeds
- `gator follow <url>` - Follow a feed that already exists in the database
- `gator unfollow <url>` - Unfollow a feed that already exists in the database
- `gator agg (<time_duration>` e.g. (1m - 1 minute, 30s - 30 seconds, 1h - 1hour)
- `gator register username` - Register a user
- `gator reset` - Reset all data which removes all saved post, feed, and user data
- `gator addfeed <feed_name> <feed_url>` - Add a feed to follow for the user
- `gator following` - List the feeds followed by the user
- `gator browse <optional_limit>` - List posts from the feed by default 2 posts are retrieved


