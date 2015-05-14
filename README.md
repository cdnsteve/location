# location: Geolocation REST API in Go

A very lightweight Geolocation REST API using Google App Engine. It relies on the Google App Engine generated headers to collect IP information. 
See example at: [http://gold-cycling-94509.appspot.com/](http://gold-cycling-94509.appspot.com/)

## Usage

`curl http://gold-cycling-94509.appspot.com/`

Returns JSON with IP address location information.

```
{
    "Country": "CA",
    "Region": "on",
    "City": "kingston",
    "LatLong": [
        "XX.xxxxx,-XX.xxxxx"
        ]
}
```

## License

 MIT
