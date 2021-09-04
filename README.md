# Weather Playlist API

Welcome to the Weather Playlist API ⛅📡🎧

Weather Playlist API is a micro-service able to accept RESTful requests receiving as parameter
either city name or lat long coordinates and returns a playlist suggestion according to the current temperature.

# Requires

- [Make](https://www.gnu.org/software/make/manual/make.html)
- [Docker](https://www.docker.com)
- [Docker Compose](https://docs.docker.com/compose)

# Get Started

To run the project localhost you can follow these commands:

```sh
# build docker image
make build

# start containers
make up

# run the tests
make test

# show the logs
make logs
```

Now, if you want to test from your browser open [localhost:3000](http://localhost:3000)

The result is:

```json
{
  title: "Welcome to the Weather Playlist API",
  samples: [
    "http://localhost:3000/playlist?city=campinas",
    "http://localhost:3000/playlist?city=salvador",
    "http://localhost:3000/playlist?city=fortaleza",
    "http://localhost:3000/playlist?city=florianópolis",
    "http://localhost:3000/playlist?lat=-3.1019&lon=-60.025",
    "http://localhost:3000/playlist?lat=13.8333&lon=-88.9167"
  ]
}
```

Try call passing our city name
[http://localhost:3000/playlist?city=florianópolis](http://localhost:3000/playlist?city=florianópolis)
or the latitude and longitude
[http://localhost:3000/playlist?lat=-3.1019&lon=-60.025](http://localhost:3000/playlist?lat=-3.1019&lon=-60.025)

---

Inspired by [ifood-backend-challenge](https://github.com/ifood/vemproifood-backend)
