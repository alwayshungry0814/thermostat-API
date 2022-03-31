# thermostat-API 

The only external package is the gorilla/mux package for routing. 

```go get -u github.com/gorilla/mux```

The following requests are all that's needed to test the api. There is no database implementation yet. 

To get the current temperature:

```
curl --location --request GET 'localhost:8000/temperature'
```

To get current temperature setting:

```
curl --location --request GET 'localhost:8000/temperatureSetting'
```

To get thermostat state:

```
curl --location --request GET 'localhost:8000/thermostatState'
```

To change thermostat temperature:
```
curl --location --request PUT 'localhost:8000/temperatureSetting' \
--header 'Content-Type: application/json' \
--data-raw '
    {
        "coolTemp": 73,
        "heatTemp": 68
    }
'
```
