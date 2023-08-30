# Rest API Error Responses

RFC-7807

GET - HTTP Response Code: **404**

```javascript
    HTTP/1.1  404
    Content-Type: application/problem+json
 
    {
      "message": "The item does not exist"
    }
```

POST -  HTTP Response Code: **400**

```javascript
    HTTP/1.1  400
    Content-Type: application/problem+json
    
    {
        "message": "Validation errors in your request", /* skip or optional error message */
        "errors": [
            {
                "message": "Oops! The value is invalid",
                "code": 34,
                "field": "email"
            },
            {
                "message": "Oops! The format is not correct",
                "code": 35,
                "field": "phoneNumber"
            }
        ]
    }
```

PUT -  HTTP Response Code: **400**

```javascript
    HTTP/1.1  400
    Content-Type: application/problem+json
    
    {
        "message": "Validation errors in your request", /* skip or optional error message */
        "errors": {
            "email": [
              {
                "message": "Oops! The email is invalid",
                "code": 35
              }
            ],
            "phoneNumber": [
              {
                "message": "Oops! The phone number format is not correct",
                "code": 36
              }
            ]
        }
    }
```

VERB Unauthorized - HTTP Response Code: **401**

```javascript
    HTTP/1.1  401
    Content-Type: application/problem+json
 
    {
      "message": "Authentication credentials were missing or incorrect"
    }
```

VERB Forbidden - HTTP Response Code: **403**

```javascript
    HTTP/1.1  403
    Content-Type: application/problem+json
 
    {
      "message": "The request is understood, but it has been refused or access is not allowed"
    }
```

VERB Method Not Found - HTTP Response Code: **404**

```javascript
    HTTP/1.1  404
    Content-Type: application/problem+json
 
    {
      "message": "Not route"
    }
```

VERB Method Not Allowed - HTTP Response Code: **405**

```javascript
    HTTP/1.1  405
    Content-Type: application/problem+json
 
    {
      "message": "Method not allowed"
    }
```

VERB Conflict - HTTP Response Code: **409**

```javascript
    HTTP/1.1  409
    Content-Type: application/problem+json
 
    {
      "message": "Any message which should help the user to resolve the conflict"
    }
```

VERB Too Many Requests - HTTP Response Code: **429**

```javascript
    HTTP/1.1  429
    Content-Type: application/problem+json
 
    {
      "message": "The request cannot be served due to the rate limit having been exhausted for the resource"
    }
```

VERB Internal Server Error - HTTP Response Code: **500**

```javascript
    HTTP/1.1  500
    Content-Type: application/problem+json
 
    {
      "message": "Something is broken"
    }
```

VERB Service Unavailable - HTTP Response Code: **503**

```javascript
    HTTP/1.1  503
    Content-Type: application/problem+json
 
    {
      "message": "The server is up, but overloaded with requests. Try again later!"
    }
```
