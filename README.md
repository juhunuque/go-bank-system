

## Migration

### Create migration file
`migrate create -ext sql -dir db/migration -seq init_schema`

### Execute migrations
`migrate -path db/migration -database "postgresql://test:test@localhost:5432/test?sslmode=disable" -verbose up`

## Tasks
- List  menu: `task --list`

## Mock
`mockgen -package mockdb -destination db/mock/store.go github.com/juhunuque/simplebank/db/sqlc Store`

## Resources
- https://github.com/techschool/simplebank

- DB Migrations: https://github.com/golang-migrate/migrate (CLI Version)
- TaskFile: https://taskfile.dev/#/
- TaskFile Guide: https://dev.to/stack-labs/introduction-to-taskfile-a-makefile-alternative-h92
- Data layer: https://github.com/kyleconroy/sqlc
- PG Driver lib: https://github.com/lib/pq (db/sql is a generic implementations and requires this lib to interact with PG)
- Assertion for tests: https://github.com/stretchr/testify
- Load config from env: https://github.com/spf13/viper
- Mock DB: https://github.com/golang/mock
- Manage UUID: https://github.com/google/uuid
- Handle JWT: https://github.com/dgrijalva/jwt-go
- Handle PASETO: https://github.com/o1egl/paseto

## Appendix

* go mod tidy

  ensures that the "go.mod" file matches the source code in the module. It adds any missing module requirements necessary to build the current module's packages and dependencies, if there are some not used dependencies go mod tidy will remove those from go.