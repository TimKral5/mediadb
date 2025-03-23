# MediaDB Server

## Build Server
### Requirements

To build the server, BunJS must be installed.

### Install Dependencies

Before actually building the application, all dependencies must be
installed. This can be done with:

```bash
bun install
```

Make sure to run the command above in the `shared/` project as well.

### Build for Linux

The following command will compile a stand-alone executable from the
source code:

```bash
bun build \
  --compile \
  --outfile mdb \
  --target=bun-linux-x64 \
  ./index.ts
```

### Build for Windows

The following command will compile a stand-alone for Windows:

```bash
bun build \
  --compile \
  --outfile mdb \
  --target=bun-windows-x64 \
  ./index.ts
```
