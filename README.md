# Svelte + TypeScript Frontend and Go Backend Setup

This guide will walk you through setting up a simple **Svelte** frontend with **TypeScript** and a **Go** backend.

Project for a realtime quiz platform played on a large screen, similar to [Kahoot](https://kahoot.com). Backend in Go with Fiber HTTP server, frontend in vanilla Svelte.

## Step 1: Prerequisites

Before starting, ensure you have the following installed:

- [Node.js](https://nodejs.org/): For the Svelte frontend (includes npm).
- [Go](https://golang.org/): For the backend API.
- [npm](https://www.npmjs.com/): Node.js package manager.
- [Git](https://git-scm.com/): To manage source code (optional but recommended).


## Step 2. Run the Go API

Start the Go server by running the following command:

```go
go run main.go
```

## Step 3: Setting up the Svelte Frontend

Install the TypeScript support for Svelte:
```js
npm install --save-dev typescript svelte-preprocess @tsconfig/svelte
```

And then run the following command to start the Svelte development server:
```js
npm run dev
```

## Step 4: Build and Deploy

### 4.1 Build the Svelte Project

To create a production-ready build of your Svelte frontend, run:

```js
npm run build
```

### 4.2 Deploy the Go Backend

To deploy the Go backend, you can build it using:

```go
go build -o go-backend
```

### 4.3 Serve the Svelte Build

You can use any static file server to serve your Svelte frontend, for example:

```bash
npm install -g serve
serve -s public
```

## Conclusion
You now have a fully functional application with a Go backend and a Svelte frontend with TypeScript. This is a basic setup and can be expanded with more complex routes, database integration, authentication, and so on.