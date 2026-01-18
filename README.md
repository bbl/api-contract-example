# API Contract Example

An example project demonstrating API contract-first development using TypeSpec to define the contract, with code generation for both a Go server and TypeScript client.

## Overview

This project showcases a workflow where:
1. The API contract is defined in **TypeSpec** (`contract/main.tsp`)
2. TypeSpec compiles to an **OpenAPI 3.0** specification
3. **Go server code** is generated from the OpenAPI spec using `oapi-codegen`
4. **TypeScript client** is generated from the OpenAPI spec using `@hey-api/openapi-ts`

## Project Structure

```
.
├── contract/
│   └── main.tsp           # TypeSpec API definition
├── cmd/
│   └── main.go            # Go server implementation
├── generated/
│   ├── openapi.yaml       # Generated OpenAPI spec
│   ├── api/
│   │   └── api.gen.go     # Generated Go server code
│   └── ts-client/         # Generated TypeScript client
├── tspconfig.yaml         # TypeSpec configuration
├── oapi-codegen.yaml      # Go code generator configuration
├── openapi-ts.config.ts   # TypeScript client generator configuration
└── Makefile               # Build automation
```

## Prerequisites

- [Node.js](https://nodejs.org/) (for TypeSpec and TypeScript client generation)
- [Go](https://golang.org/) 1.24+
- npm or yarn

## Installation

Install the required dependencies:

```bash
make install-deps
```

This installs:
- TypeSpec compiler and HTTP/OpenAPI libraries
- `@hey-api/openapi-ts` for TypeScript client generation
- `oapi-codegen` for Go server code generation

## Usage

### 1. Define the API Contract

Edit `contract/main.tsp` to define your API. Example:

```typespec
import "@typespec/http";

using Http;

model Store {
  name: string;
  address: Address;
}

model Address {
  street: string;
  city: string;
}

@route("/stores")
interface Stores {
  list(@query filter: string): Store[];
  read(@path id: string): Store;
  @post create(@body store: Store): Store;
}
```

### 2. Generate Code

Generate both OpenAPI spec and Go server code:

```bash
make generate
```

Generate only the TypeScript client:

```bash
make generate-ts-client
```

Generate only the OpenAPI spec:

```bash
make openapi
```

### 3. Run the Go Server

```bash
go run ./cmd/main.go
```

The server starts on `http://localhost:8080`.

### 4. Use the TypeScript Client

Import and use the generated client in your TypeScript project:

```typescript
import { ApiClient } from './generated/ts-client';

// List stores
const stores = await ApiClient.storesList({ query: { filter: 'coffee' } });

// Create a store
const newStore = await ApiClient.storesCreate({
  body: {
    name: 'Coffee Shop',
    address: {
      street: '123 Main St',
      city: 'Seattle'
    }
  }
});

// Get a store by ID
const store = await ApiClient.storesRead({ path: { id: 'store-id' } });
```

## Configuration Files

### tspconfig.yaml
Configures TypeSpec to emit OpenAPI 3.0 spec to `generated/openapi.yaml`.

### oapi-codegen.yaml
Configures Go code generation with:
- Echo server handlers
- Strict server interface (request/response objects)
- Model types

### openapi-ts.config.ts
Configures TypeScript client generation with the `@hey-api/sdk` plugin.

## Clean Up

Remove all generated files:

```bash
make clean
```

## Tools Used

- [TypeSpec](https://typespec.io/) - API definition language
- [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) - Go code generator for OpenAPI
- [@hey-api/openapi-ts](https://github.com/hey-api/openapi-ts) - TypeScript client generator
- [Echo](https://echo.labstack.com/) - Go web framework
