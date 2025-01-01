# Score Board (Sportradar interview task)

This repo was created for interview process. The app tracks football games and provides two modes:
* **interactive mode** - provides very simple user interface for interactions like starting a game, updating game events, finishing a game and showing summary
* **demo mode** - doesn't provide any user interface, just creates some pre-defined events for games and logs them when they are processed

## Requirements

To run the app you will need:
* Go (1.22.2 or higher)
* Docker
* make

## Running

To run the app in interactive mode, use:
```sh
make run
```
or
```sh
make run-ui
```

This will run the app with user interface allowing user interactions with scoreboard.

---

Tu run the app in demo mode, with some pre-defined events which should be sent to score board, use:
```sh
make run-demo
```

--- 

Check all available commands:
```sh
make help
```

## Tests

Run tests:
```sh
make test
```

## Linters

For linters enabled in golangci please check `.golangci.yaml` file.
To lint the code in cloned repository run:
```sh
make golangci
```

## Q&A

**Why event based solution?**

I believe good solution should show places when it can grow. Simplest possible implementation could do the same, but it would look "finished". With event based implementation it is easy to imagine how the solution may grow - we can add more events for other iteractions with stored matches or add some middleware for additional actions for events, like sending the event throuh Pub/Sub or storing them in a file.

**Why are there no tests of ui package?**

The package was added just to allow interactions for easier checking the task. The main part of the code responsible for all the logic is scoreboard package and this package has tests implemented.

**Why there are only unit tests?**

There just wasn't too much to test, so other types of tests would be redundant.

**Why such simple user interface?**

I specialise in backend solutions - please don't expect me to do some fancy UI if I don't really need to ;)

**Why there is no docker compose?**

For such simple app without other components it's just not needed.
