# REST api in Go 

# Run 
    # Run on port 8000
        go build & ./golang-rest-api



# APIs
    curl -X GET \
        http://localhost:8000/api/books \
        -H 'Accept: */*' \
        -H 'Cache-Control: no-cache' \
        -H 'Connection: keep-alive' \
        -H 'Host: localhost:8000' \
        -H 'Postman-Token: 07ee3447-56f2-4e69-b93d-f2fabc1ead0a,df254efc-ebcf-4976-bb47-bfb61d5dc1bc' \
        -H 'User-Agent: PostmanRuntime/7.13.0' \
        -H 'accept-encoding: gzip, deflate' \
        -H 'cache-control: no-cache'


    curl -X POST \
        http://localhost:8000/api/books \
        -H 'Accept: */*' \
        -H 'Cache-Control: no-cache' \
        -H 'Connection: keep-alive' \
        -H 'Content-Type: application/json' \
        -H 'Host: localhost:8000' \
        -H 'Postman-Token: 97dda288-5509-4816-8054-a1da398711b8,7f1321e7-7394-4db6-94a1-f06e00b63364' \
        -H 'User-Agent: PostmanRuntime/7.13.0' \
        -H 'accept-encoding: gzip, deflate' \
        -H 'cache-control: no-cache' \
        -H 'content-length: 139' \
        -d '{
            "isbn": "4242425",
            "title": "title Book 3",
            "author": {
                "firstname": "Ranjeet2",
                "lastname": "Kumar2"
            }
        }'


    curl -X PUT \
        http://localhost:8000/api/books/1 \
        -H 'Accept: */*' \
        -H 'Cache-Control: no-cache' \
        -H 'Connection: keep-alive' \
        -H 'Content-Type: application/json' \
        -H 'Host: localhost:8000' \
        -H 'Postman-Token: 7f6de2fb-bbe7-49ee-8f80-4b9c5089dde9,07b4df95-f4b2-4fd5-be8e-9ed1a6882c48' \
        -H 'User-Agent: PostmanRuntime/7.13.0' \
        -H 'accept-encoding: gzip, deflate' \
        -H 'cache-control: no-cache' \
        -H 'content-length: 139' \
        -d '{
            "isbn": "4242425",
            "title": "title Book 3",
            "author": {
                "firstname": "Ranjeet2",
                "lastname": "Kumar2"
            }
        }'
