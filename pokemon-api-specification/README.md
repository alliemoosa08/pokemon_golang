# Pokemon-API Specification

## Getting Started

### Prerequisites

You'll need to install a few things first:

- [Node](https://nodejs.org/en/) 18 or equivalent LTS

### Installation

In the root of this repository, run:

```bash
npm install
```

### Open API 3.0 Development Environment

API documentation is in the [OpenAPI 3.0](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.0.md) standard (formerly known as Swagger).

[swagger-ui-watcher](https://github.com/moon0326/swagger-ui-watcher) is a useful tool for developing the API specification.

Start up with:

```bash
npm start
```

Compile and build the specification with:

```bash
npm run build
```

This will create a fully dereferenced, standalone specification file in the /dist/api.json

To validate the schema run:

```bash
npm test
```

## Testing

```bash
npm run test
```

Followed by

## Technologies

- [Open API 3.0 ](https://github.com/OAI/OpenAPI-Specification)