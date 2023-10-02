# Go API

`app` package contains all application types and interfaces. Every other package that require them imports from
the `app`
package. This ensures there are no circular dependencies, flat hierarchical structure, and easily interchangeable
implementations.

## Notable packages

* pgx

## Notes

* Services generate IDs
* Use Google UUID v4 for now. Switch to UUID v7 when it's released
