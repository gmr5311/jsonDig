# jsonDig

A Go package for searching through a given json string for a specific key and expecting it's value to be what you expected.

### Importing

  ```Go
  import "github.com/gmr5311/jsondig"
  ```

### Documentation

### Usage

In order to search through a given json string you will need to run the following function.

```Go
  jsondig.FindInJSON(jsonString, key, value)
```

where `jsonString` is the json string that you are searching through, `key` is the string value of the key that you are looking for, and `value` is the value that you expect the key to have when it is found. If the key is found and it has the value that you expect then the function will return true. Otherwise it will return false.
