
# Access Method `gitHub` and `github` - Github Commit Access

## Synopsis

```yaml
type: gitHub/v1
```

Provided blobs use the following media type for: `application/x-tgz`

The artifact content is provided as gnu-zipped tar archive

### Description

This method implements the access of the content of a git commit stored in a
GitHub repository.

Supported specification version is `v1`

### Specification Versions

#### Version `v1`

The type specific specification fields are:

- **`repoUrl`**  *string*

  Repository URL with or without scheme.

- **`ref`** (optional) *string*

  Original ref used to get the commit from. mutually exclusive with `commit`.

- **`commit`** (optional) *string*

  The sha/id of the git commit. mutually exclusive with `ref`.

### Go Bindings

The go binding can be found [here](method.go)
