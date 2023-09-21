# lazy-platform-auth

GORM uuid

For postgresql, here is what I did:

`go get github.com/google/uuid`

Use `uuid.UUID` (from "github.com/google/uuid"), as type,
e.g
```
ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
```
Add uuid-ossp extension for postgres database,
e.g
```
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```
Then, when you call DB's Create() method, the uuid is generated automatically.