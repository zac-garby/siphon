# Example

This is a simple example of usage of `db`. The schema is as follows:

```
me: user
map: <string:string>

struct user {
    name: string
    email: string
    age: uint8
    friends: [user]
}
```

To run it, `$ cd` to the directory and run `$ db`.

## Stuff to try

### Set some data

Set the field `me` to contain some actual data. Send a `POST` request to `localhost:7913/set?selector=me` with a post body looking something like this:

```json
{
    "name": "your name",
    "email": "your email",
    "age": 16,
    "friends": [
        {
            "name": "something",
            "email": "some@thi.ng",
            "age": 17,
            "friends": []
        },
        {
            "name": "foo",
            "email": "bar",
            "age": 16,
            "friends": []
        }
    ]
}
```

### Query the database

Make a `GET` request to `localhost:7913/json?selector=me` to retrieve the data you just stored. You can also do something a bit more complicated, for example instead of using the selector `me`, you could use `me.friends[age=16].name`, which will return the names all friends who are 16 years old.
