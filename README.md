# Weather Fetcher

This Go project fetches and displays the current weather information for the location associated with your IP address. It utilizes two external APIs:
- A geo-location API to determine your city and country based on your IP address.
- A weather API to fetch the current weather for the detected location.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [APIs Used](#apis-used)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Installation

### Prerequisites

Ensure that you have [Go](https://golang.org/dl/) installed on your machine. You can verify the installation by running:

```bash
go version
```

### Clone the Repository

```bash
git clone https://github.com/scotty-c/weather.git
cd weather
```

### Build the Project

You can build the project using the Go build tool:

```bash
go build -o weather
```

### Run the Application

After building the project, you can run the application using:

```bash
./weather
```

This will print the current weather information for your location.

## Usage

By default, the application uses the following APIs:
- Geo-location API: `https://api.seeip.org/geoip`
- Weather API: `https://v3.wttr.in/`

The application will fetch your city and country based on your IP address and then fetch the current weather for that location.

### Example Output

```plaintext
Test City: Clear skies
```

This output shows the weather information for the detected location.

## Project Structure

```
weather/
│
├── main.go          # The main application logic
├── main_test.go     # Unit tests for the application
├── README.md        # Project documentation
```

## APIs Used

1. **Geo-location API**:
   - URL: `https://api.seeip.org/geoip`
   - This API returns the city and country based on the user's IP address.

2. **Weather API**:
   - URL: `https://v3.wttr.in/<city>+<country>?format=%c+%t`
   - This API returns the current weather for the specified city and country.

## Testing

Unit tests are provided in the `main_test.go` file. The tests use mock servers to simulate the external APIs.

### Running Tests

You can run the tests using the following command:

```bash
make test
```

This command will execute all the tests and display the results.

## Contributing

Contributions are welcome! If you would like to contribute to the project, please fork the repository, create a new branch for your feature or bugfix, and submit a pull request. Make sure to include tests for any new functionality.

### Steps to Contribute:
1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -am 'Add new feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Create a new Pull Request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.