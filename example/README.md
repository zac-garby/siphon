# Example

This is a simple example of usage of `db`. The schema is as follows:

```
me: user
map: <string:string>
nums: [float]

struct user {
    name: string
    email: string
    age: uint8
    friends: [user]
}
```

To run it, `$ cd` to the directory and run `$ siphon`.

## Stuff to try

Once the server is running (see above,) you can open a CLI session with it using `$ siphon-cli http://localhost:7913`:

```
? set nums
| [1, 2, 3, 4, 5]
OK

? nums
[
    1,
    2,
    3,
    4,
    5
]

? set me
| {
    "name": "foo",
    "email": "foo@example.com",
    "age": 100,
    "friends": []
}

? append nums
| 5
```
