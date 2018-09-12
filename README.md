# Siphon DB

Siphon is a NoSQL database with a straightforward and powerful query syntax. Querying is done using _selectors_ - here are a few:

```ruby
# Get a list of the friends of user no. 5
users[5].friends

# Get all the posts belonging to the user called "foo" with more than 10 likes
users[name="foo"].posts[likes>10]

# Returns the content of the first post posted by an admin
posts[user~/^admin-/][0].content
```

Data can be modified in a similar way.

## Features

 - Selector syntax for querying and adding new data
 - Access the database via the HTTP API
 - Define database structure with a schema
 - Responses available as JSON
 - Persistant storage on disk
 - Support for graphQL queries

## TODO

 - Improve selector syntax
    - Allow every kind of literal in comparisons
    - Add boolean operators, e.g. `users[name="foo" | (age>10 & age<20)]`
    - Add full expression support, e.g. `num_pairs[a + b > 5]`
    - Add global variables, e.g. `sessions[expires >= $TIMESTAMP]`

## Schema

A schema defines the structure and types of the database. It might look something like this:

```go
users: [user]
posts: [post]

struct user {
    id: int
    name: string
    pass: string
    posts: [post]
}

struct post {
    id: int
    title: string
    content: string
    likes: uint
}
```

> Note: it wouldn't be advised to store posts in two places (top-level `posts` field and `user.posts`, but rather you should store them in a hashmap, mapping IDs to posts.)

## Modifying data

Given the schema defined above, you could add a new user by sending a POST request to `/append?selector=users` with the given JSON data:

```json
{
    "id": 56,
    "name": "foo",
    "pass": "...",
    "posts": [
        {
            "id": 132,
            "title": "Hello, world!",
            "content": "This is a post ... the end",
            "likes": 34912
        },
        {
            "id": 158,
            "title": "Another post",
            "content": "...",
            "likes": 10
        }
    ]
}
```

> Note: Number literals are polymorphic - since the required type is known (e.g. `likes` is `uint`), the JSON number values are cast to the correct type. The same happens with strings/regexps.

A number of routes are supported for modifying data. Here's a full list: (`data` represents the POSTed data.)

Route        | Description
-------------|---------------------------------------------
`/set`       | Sets the value to `data`
`/unset`     | Deletes the selected keys/indexes
`/append`    | Appends `data` to the current value
`/prepend`   | Prepends `data` to the current value
`/key`       | Sets key `data.key` to `data.value` (also works for struct fields and list indices)
`/empty`     | Empties a list or hashmap
