# go_template_chainlink

# This is a basic implementation for a Chainlink External Adapter Go Template
----

## Please follow the steps in the original Chainlink repo to run a node
- [Chainlink](https://github.com/smartcontractkit/chainlink#install)
- Follow steps in this [video](https://www.youtube.com/watch?v=ZB3GLtQvgME&list=PLVP9aGDn-X0Shwzuvw12srE-O6WKsGvY_&index=7) by @patrickAlphaC
  

### Run
- go get -u github.com/gin-gonic/gin
- go run main.go

### Call Api Server
- ```curl -X POST -H "content-type:application/json" "http://localhost:8080/" --data '{ "id": "0", "data": { "q": "New York", "units": "metric" }}'```
  
#### Input
```json
{
    "id": "0",
    "Data":
    {
        "q": "New York"
    }
}
```

#### Output
```json
{
    "data": {
        "coord": {
            "lon": -74.006,
            "lat": 40.7143
        },
        "weather": [
            {
                "id": 800,
                "main": "Clear",
                "description": "clear sky",
                "icon": "01n"
            }
        ],
        "base": "stations",
        "main": {
            "temp": 269.17,
            "feels_like": 265.21,
            "temp_min": 265.56,
            "temp_max": 271.45,
            "pressure": 1037,
            "humidity": 52
        },
        "visibility": 10000,
        "wind": {
            "speed": 2.68,
            "deg": 231
        },
        "clouds": {
            "all": 0
        },
        "dt": 1641687756,
        "sys": {
            "type": 2,
            "id": 2008776,
            "country": "US",
            "sunrise": 1641644389,
            "sunset": 1641678315
        },
        "timezone": -18000,
        "id": 5128581,
        "name": "New York",
        "cod": 200
    },
    "jobRunID": "0",
    "result": 269.17,
    "statusCode": 200
}   
```