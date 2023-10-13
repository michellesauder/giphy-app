# Giphy Viewer - Client-Side Application

This client-side React application serves as a Giphy viewer that allows you to search and view Giphy content based on your input. It interacts with the Go backend server to retrieve and display Giphy data.

## Features

- Search for Giphy content by entering a search query.
- Retrieve and display Giphy content in iframes using the provided URLs.
- Communicate with the Go backend server to fetch Giphy data.
- Error handling for failed data requests.

## Prerequisites

Before running this application, make sure you have the following prerequisites installed:

- Node.js: The project is built using React, and you need Node.js to run it.
- Yarn: Yarn is required for managing project dependencies.

## Installation and Setup

1. Clone this repository to your local machine:

   ```shell
   git clone git@github.com:michellesauder/giphy-app.git
   ```

2. Navigate to the project directory:

   ```shell
   cd client
   ```

3. Install project dependencies:

   ```shell
   yarn install
   ```

4. Start the application:

   ```shell
   yarn dev
   ```

   The application will run and can be accessed in your web browser at [http://localhost:5173/](http://localhost:5173/).

## Usage

1. Open the application in your web browser.

2. The search field allows you to enter a search query to find Giphy content related to your search.

3. Giphy content will be displayed in iframes below the search field.

4. If any errors occur while fetching Giphy data, you will see an error message in the console.

## Configuration

The Go backend server's endpoint is defined in the `ENDPOINT` constant at the beginning of the `App.tsx` file. You can change this URL to match the location of your Go server if needed.

## Contributing

Contributions to this project are welcome. To contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them.
4. Push the changes to your fork.
5. Submit a pull request to the main repository.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- The project utilizes the [SWR](https://swr.vercel.app/) library for handling data fetching.
- It is built with React, a JavaScript library for building user interfaces.

# React + TypeScript + Vite

This template provides a minimal setup to get React working in Vite with HMR and some ESLint rules.

Currently, two official plugins are available:

- [@vitejs/plugin-react](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react/README.md) uses [Babel](https://babeljs.io/) for Fast Refresh
- [@vitejs/plugin-react-swc](https://github.com/vitejs/vite-plugin-react-swc) uses [SWC](https://swc.rs/) for Fast Refresh

## Expanding the ESLint configuration

If you are developing a production application, we recommend updating the configuration to enable type aware lint rules:

- Configure the top-level `parserOptions` property like this:

```js
   parserOptions: {
    ecmaVersion: 'latest',
    sourceType: 'module',
    project: ['./tsconfig.json', './tsconfig.node.json'],
    tsconfigRootDir: __dirname,
   },
```

- Replace `plugin:@typescript-eslint/recommended` to `plugin:@typescript-eslint/recommended-type-checked` or `plugin:@typescript-eslint/strict-type-checked`
- Optionally add `plugin:@typescript-eslint/stylistic-type-checked`
- Install [eslint-plugin-react](https://github.com/jsx-eslint/eslint-plugin-react) and add `plugin:react/recommended` & `plugin:react/jsx-runtime` to the `extends` list
