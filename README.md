# Build Giphy Application with Go & React

This repository contains the code for a Giphy application that combines the power of a Go backend with a React frontend to display Giphy content. It allows you to search for Giphy images and view them in an interactive interface.

## Installation

Before you begin, ensure you have both Go and Node.js (with Yarn) installed on your machine.

### Server (Go)

1. Navigate to the `server` directory.

2. Create a `.env` file and add your Giphy API key:

   ```
   API_KEY=YOUR_API_KEY
   ```

3. Install the required Go dependencies:

   ```
   go get -d ./...
   ```

4. Run the Go server:
   ```
   go run main.go
   ```

The Go backend will start on http://localhost:8080.

### Client (React)

1. Navigate to the `client` directory.

2. Install project dependencies using Yarn:

   ```
   yarn install
   ```

3. Start the React development server:
   ```
   yarn dev
   ```

The React frontend will be accessible on http://localhost:5173/.

## Usage

1. Open your web browser and go to http://localhost:5173/ to access the Giphy application.

2. Enter your search query in the search field and click the "Search" button to display Giphy images related to your query.

3. Enjoy exploring Giphy images and viewing them in an interactive interface.

## Contributing

Contributions are welcome! If you'd like to enhance this project, please follow the standard GitHub Fork and Pull Request workflow.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

```

You can use this README.md as the introduction and installation guide for your Go and React-based Giphy application. Make sure to replace `"YOUR_API_KEY"` with your actual Giphy API key in the `.env` file on the server side.
```
