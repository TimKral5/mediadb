# MediaDB Client Library

This is a client library for the **MediaDB Server**. It may be in
projects that make use of the API.

## Usage

The project can be directly used by importing components from
`index.ts`.

## Bundling the Library

Bundling the library may be required for usage in the browser. It can
be done with following command:

```bash
bun build \
  --target browser \
  --outfile mediadb-client.js \
  ./index.ts
```

## Running the Development Environment

A dev-environment can be spun up with the following command:

```bash
bun ./test/index.html
```
