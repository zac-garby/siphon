# db

db is a NoSQL database with a straightforward and powerful query syntax. Querying is done using _selectors_ - here are a few:

```ruby
# Get a list of the friends of user no. 5
users[5].friends

# Get all the posts belonging to the user called "foo" with more than 10 likes
users[name="foo"].posts[likes>10]

# Returns the content of the first post posted by an admin
posts[user~/^admin-/][0].content

# Returns all posts sorted by number of likes
posts.sorted(self.likes)
```

Data can be modified in a similar way.

## Features

 - Selector syntax for querying and adding new data
 - Access the database via the HTTP API
 - Schema support to enforce types and database structure
 - Data available as JSON or Go binary data