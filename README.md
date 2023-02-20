# song-merger

## How to run
```bash
go run main.go
```

The command above starts a server listening to `127.0.0.1:8000`, to which you can send requests to all the routes.

### Current Routes
- `/songs`
  - Receives a JSON with song parameters
  - Sample JSON below:

```json
[
    {
        "name": "ride",
        "artist": "twenty-one-pilots",
        "musicalTone": 4
    },
    {
        "name": "dust-in-the-wind",
        "artist": "Kansas",
        "musicalTone": 3
    }
]
```

## To-dos
- [X] Put text into an output file
- [X] Get song names from a JSON
- [ ] Get song parameters from a JSON
  - [X] Key
  - [X] Capo
  - [ ] Toggle tabs
- [X] Get song title to put in merged HTML
- [X] Receive a list of JSON objects in the payload and treat each one as a separate song
- [X] Put all scores into a single HTML document
- [X] Convert HTML to PDF (maybe?)
