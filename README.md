# URL Shortener

## How to run
1. Make sure docker is installed and running and run the following command

```bash 
docker compose up
```

2. Make a POST request to localhost:3000/shorten with the body using the body containing JSON structure

```json
{"url": "http://www.google.com"}

```
- Note that the URL has to start with http:// or https:// due to "validation"

- The expected response is 

```json
{"code": "1381a8b29f7"}
```

3. Take the code from the response and make a GET request to localhost:3000/\<code\> the expected response is 

```json
{"url": "http://www.google.com"}
```
