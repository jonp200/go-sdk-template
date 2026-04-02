# go-sdk-template

A GitHub template repository for Go library/SDK projects.

## Using This Template

Click **Use this template** on GitHub to create a new repository from this template.

On the first push to `main`, the [bootstrap workflow](.github/workflows/bootstrap.yml)
automatically:

1. Installs the **latest stable Go** and updates the `go` directive in `go.mod`.
2. Runs `go mod tidy` to synchronise dependencies.
3. Replaces `.gitignore` with the **latest canonical [Go `.gitignore`](https://github.com/github/gitignore/blob/main/Go.gitignore)**
   from the official `github/gitignore` repository.
4. Auto-commits any changes back to the default branch (only when something actually changed).

## Getting Started

Import the SDK in your Go project:

```go
import sdk "github.com/jonp200/go-sdk-template"
```

> **After using the template**, replace `github.com/jonp200/go-sdk-template` with your own
> module path in `go.mod` and any import statements.

### Example

```go
package main

import (
    "fmt"
    sdk "github.com/jonp200/go-sdk-template"
)

func main() {
    fmt.Println(sdk.Hello("World")) // Hello, World!
}
```

## Running Tests

```bash
go test ./...
```

## CI

A [CI workflow](.github/workflows/ci.yml) runs `go test ./...` on every push and pull request.

## Bootstrap Workflow

The [bootstrap workflow](.github/workflows/bootstrap.yml) runs once on the **first push to
`main`** (and can also be triggered manually via `workflow_dispatch`).

| Step | What it does |
|------|-------------|
| Check last commit | Skips the entire job if the previous commit was authored by `github-actions[bot]` or its message contains the `[bootstrap]` marker — preventing infinite loops. |
| Install latest Go | Uses `actions/setup-go` with `go-version: stable` to resolve the current stable release. |
| Update `go.mod` | Runs `go mod edit -go <version>` and `go mod tidy`. |
| Fetch `.gitignore` | Downloads `Go.gitignore` from `https://raw.githubusercontent.com/github/gitignore/main/Go.gitignore`. |
| Auto-commit | Stages `go.mod`, `go.sum`, and `.gitignore`; commits with `chore: update Go version and .gitignore [bootstrap]`; pushes only when the diff is non-empty. |

The workflow uses `contents: write` permission so it can push directly to the default branch,
and `contents: read` everywhere else to follow the principle of least privilege.
