# Contributing Guide
## Development Environment

Information about how to setup and run the project can be obtained
from the respective READMEs:

- [MediaDB Repository/Docker](../README.md)
- [MediaDB Server](../server/README.md)
- [MediaDB Client](../client/README.md)

## Development Workflow

```bash
# Clone the repository
git clone https://github.com/TimKral5/mediadb.git

# Create a feature branch
git checkout -b feature/<issue_id>--<issue_title>

# Or create a fix if it's a bugfix
git checkout -b fix/<issue_id>--<issue_title>

# After applying changes, run the tests
MDB_ENDPOINT=<endpoint> bun test

# Push after applying the changes
git push --set-upstream origin <branch-name>
```

The develpment workflow begins with an issue that requires action.
From that issue, a designated branch is created, on which the changes
are implemented.

When the changes are done, the tests need to be run. If they complete
without fail, the changes can be pushed to the repository.

Now, a pull request can be created. It will be reviewed by one of the
project's maintainers.
