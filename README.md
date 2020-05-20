# HTTP Server

We have been asked to create a web server where users can track how many games have won.

>- `GET /players/{name}` should return a number indicating the number of wins
>- `POST /players/{name}` should record a win for that name, incrementing for every subsequent `POST`

We follow the TDD approach, getting working software as quickly as we can and then making small iterative improvements until we have the solution. By taking this approach we

>- keep the problem space small at any given time
>- don't go down rabbit holes
>- if we ever get stuck or lost, doing a revert wouldn't lose loads of work

## Red, Green, Refactor - Note on TDD

We emphasize the TDD process of writing a test, watching it fail (red), write a *minimal* amount of code to make it work (green) and then refactor.

This discipline of writing the minimal amount of code is important in terms of the safety TDD provides. You should be striving to get out of "**red**" as soon as you can.

Kent Beck describes it as:

> Make the test work quickly, committing whatever sins necessary in the process.

These sins can be committed because we will refactor afterwards backed by the safety of the tests.

## Chicken and egg

*How can we incrementally build our server?*

We can't `GET` a player without having stored something and it seems hard to know if `POST` has worked without the `GET` endpoint already existing.

This is where *mocking* shines.

>- `GET` will need a `PlayerStore` thing to get scores for a player. This should be an interface so when we test we can create a simple stub to test out code without needing to have implemented any actual storage code.
>- For `POST` we can *spy* on its calls to `PlayerStore` to make sure it stores players correctly. Our implementation of saving won't be coupled to retrieval.
>- For having some working software quickly we can make a very simple in-memory implementation and then later we can create an implementation backed by whatever storage mechanism we prefer.

## Build the application

Build and run the application with the following terminal commands:

```bash
go build && ./http-server
```

Then use `curl`, in a seperate terminal window, to communicate with the server.
The following sends a `win` that will be recorded for Pepper.

```bash
curl -X POST http://localhost:5000/player/Pepper
```

You can view the result with the command:
```bash
curl http://localhost:5000/player/Pepper
```

or launching your browser to `http://localhost:5000/players/Pepper`.

## References
You can find the main tutorial article [here](https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server)

***

## To do

Implement the following using TDD.

- [x] Pick a store => We'll use PostgreSQL.
- [ ] Make `PostgresPlayerStore` implement `PlayerStore`
- [ ] Plug it into integration test
- [ ] Plug it into `main()`
