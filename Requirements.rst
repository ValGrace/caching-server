Tasks 
 1. Build a CLI tool using Golang's Cobra library
 2. Build a caching server using Redis
 3. Build an actual server
 4. Create an API with GET requests

 Start -> Cache proxy server -> Response -> Actual server -> Response -> Cache
    Requirements        -> Golang
                        -> Redis
                        -> Cobra
                        -> HTTP
                        -> JSON
COMMANDS:
caching-proxy: Starts the caching proxy server

FLAGS:
 --port : Port for the server to listen on (default: 8080)
 --origin : URL of the server to which requests will be 
 --clear-cache : Clears the entire cache when set to true (default: false)
 
 HEADERS:
    --X-Cache:HIT or MISS ~ Indicate whether the response was served from cache or fetched from the origin server.