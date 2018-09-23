# todo

This example demonstrates what an app using siphon might look like. It is a simple TODO app where you can manage a list of TODOs stored in a siphon database.

## Usage

First, start a siphon server on port 7913: `$ siphon`. Then, navigate to this directory and `$ go run main.go`. It should give you a prompt in which you can enter commands. The available commands are `ls`, `add`, `done`, `rm`, `clear`, and `quit`.

 - `ls` lists all TODOs
 - `add` adds a new TODO
 - `done` toggles a TODO between completed and not completed
 - `rm` removes a TODO
 - `clear` clears all saved TODOs
 - `quit` quits the client