# URL Shortener

## How to run
1. Make sure docker is installed and running and run the following command

```bash 
docker compose up
```

2. Make a POST request to localhost:3000/shorten with the body using the following JSON structure

```json
{"url": "https://www.google.com"}

```
- Note that the URL has to start with http:// or https:// due to "validation"

3. Take the code from the response and make a GET request to localhost:3000/<code>.


