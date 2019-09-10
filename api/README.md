## Content
- [API Specification](#api-specification)
  - [Scripts](#scripts)
    - [Server](#server)
    - [Compilation Scripts](#compilation-scripts)
    - [Utilities](#utilities)
  - [Collaboration guidelines & tips](#collaboration-guidelines--tips)
  - [Known Issues](#known-issues)

# API Specification
Created in **YAML** with the **OpenAPI 3.0** specification.

You can inspect the application's API by running `yarn live`

## Scripts

Run scripts with `yarn` or `npm`

### Server

| NAME    | DESCRIPTION                                                                 |
| ------- | --------------------------------------------------------------------------- |
| start   | Shortcut for `yarn live` (see bellow)                                       |
| swagger | Compiles & Runs swagger server                                              |
| live    | Runs a swagger server & watches live yaml file changes (doesn't hot reload) |

### Compilation Scripts

| NAME            | DESCRIPTION                                                   |
| --------------- | ------------------------------------------------------------- |
| compile         | Compiles yaml & swagger                                       |
| compile-yaml    | Bundles up `swagger.yaml` into `swagger.inc.yaml`             |
| compile-swagger | Browserifies a bundle .js file that is used to run swagger-ui |  |

### Utilities
| NAME       | DESCRIPTION                                           |
| ---------- | ----------------------------------------------------- |
| cc         | Compiles yaml and copies it to clipboard              |
| watch-yaml | Watches yaml file changes and re-compiles when needed |


## Collaboration guidelines & tips
- `./index.yaml` is the only file you're allowed to `$include` a file from a remote folder
- Use `$refer: '#/components/[...]'` instead of `$include: './[...]'` in any other file to refer to a component
- `yamlinc` is an npm package that will track any `$include: 'file'` statement and import the text that comes from the file.

## Known Issues
- `live` doesn't hot reload
- In order to get the new updated content with `live` you have to **hard reload**. Soft reload will not update old content.