# song-merger

## How to run
```bash
go run main.go
```

The command above starts a server listening to `localhost:8000`, to which you can send requests to all the routes.

### Current Routes
- `/songs`
  - Receives a JSON with song parameters
  - Sample JSON below:

```json
{
    "name": "tempo-perdido",
    "artist": "legiao-urbana",
    "tone": 0
}
```

## To-dos
- [X] Put text into an output file
- [X] Get song names from a JSON
- [ ] Get song parameters from a JSON
  - [X] Key
  - [X] Capo
  - [ ] Toggle tabs
- [ ] Get song title to put in merged HTML
- [ ] Receive a list of JSON objects in the payload and treat each one as a separate song
- [ ] Put all scores into a single HTML document
- [ ] Convert HTML to PDF (maybe?)
