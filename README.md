# The-Joke-Teller
This api returns a random awful joke or pun that I have collected recently.

## Usage
* Start the server by running `go run main.go`.
* Make a GET request to `localhost:3000/api/jokes`.
> `curl http://localhost:3000/api/joke`
### The request will produce the following response
```response
{
    "id": 6,
    "joke": "Why did the coffee file a police report? It got mugged."
}
```
