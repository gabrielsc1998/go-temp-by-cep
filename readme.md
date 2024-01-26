## Temp by Cep - Go

This is a simple web server that returns the temperature of a city, passing zipcode (CEP) in celsius, kelvin and fahrenheit.

### Environment variables

- To use this app, create a `.env` file in the root of the project, following the `.env.example` file.
- OBS: for the `WEATHER_API_KEY` variable, you need to create an account on [OpenWeather](https://openweathermap.org/) and get your API key.

### Usage - local (without docker)

```bash
# run the app
go run cmd/main.go

# run app with air
air
```

### Usage - local (with dockerfile)

```bash
# build the image
docker build -t temp-by-cep .
# run the image
docker run temp-by-cep
```

### Test - GCP

- Invalid cep: https://go-temp-by-cep-d6o5asnjiq-uc.a.run.app/temp/city/123abd
- Not found cep: https://go-temp-by-cep-d6o5asnjiq-uc.a.run.app/temp/city/11222000
- Valid cep: https://go-temp-by-cep-d6o5asnjiq-uc.a.run.app/temp/city/01031970