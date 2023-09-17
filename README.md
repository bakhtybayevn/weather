# GO Weather App
This Go application serves as a simple weather information retrieval and display system using the OpenWeatherMap API. It allows you to fetch weather data for a specified city and provides additional features such as displaying temperature in either Celsius or Fahrenheit.
# Features
1. Retrieve weather information for a specified city.
2. Display temperature in Celsius or Fahrenheit (default is Celsius).
3. Show additional weather details including pressure, humidity, and description.
4. Easily configurable via a JSON file for API key management.
# Prerequisites
Before you start using this application, ensure you have the following:
1. Go installed on your local machine. You can download it from the official website: [Go Downloads](https://go.dev/dl/).
2. An OpenWeatherMap API key. You can obtain one by signing up at [OpenWeatherMap](https://openweathermap.org/api).
# Configuration
1. Clone the repository to your local machine:
git clone https://github.com/bakhtybayevn/weather.git
2. Create a configuration file named .apiConfig in the root directory of the project. The file should contain your OpenWeatherMap API key in the following format:
{
  "APIKey": "YOUR_OPENWEATHERMAP_API_KEY_HERE"
}
3. Install required Go dependencies using:
go get
# Usage
1. Start the Go application:
go run main.go
2. Access the application through a web browser or API requests.
   1. To retrieve weather information for a city, use the following URL format:
   http://localhost:8580/weather/CITY_NAME
   Replace CITY_NAME with the name of the city you want to retrieve weather information for.
   2. To specify the temperature unit (Celsius or Fahrenheit), add the unit query parameter to the URL. For example, to get weather information in Fahrenheit for New York:
   http://localhost:8580/weather/NewYork?unit=fahrenheit
# API Response
The API response will be in JSON format and include the following information:
1. city: The name of the city.
2. temperature: The temperature in the specified unit (Celsius or Fahrenheit).
3. pressure: The atmospheric pressure.
4. humidity: The relative humidity.
5. description: A description of the weather conditions.
# Contributing
Feel free to contribute to this project by opening issues, suggesting improvements, or submitting pull requests.
