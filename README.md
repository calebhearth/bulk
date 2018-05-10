# Bulk

`bulk` is a utility to insert very large sets of records into postgres, using as
few `INSERT` statements as possible. Its usefulness lies in that it avoids the
limit Postgresql has of a maximum of 65535 bind variables per statement.


While it is trivial for a database driver to insert multiple records, most
`sql.Driver` implementations do not handle this issue, which is present in very
few use cases (such as importing datasets from one database to another, or bulk
inserting from a CSV in code which performs transformations on data).

This SQL that a driver may write will overflow the bind variable limit. It is
possible, but unwieldy, to build this on your own in code.

```sql
INSERT
  INTO people (id, name)
  VALUES
    ($1::bigint, $2),
    ($3::bigint, $4),
    -- 16,382 more people row values, which just exceed 65535 bind variables
;
```

`bulk` abstracts the math to determine how many rows can be inserted at once,
and performs the insert in as few statements as possible. It returns a
`sql.Result` that includues the final `LastInsertId` and sum of `RowsAffected`.

```go
package main

import (
  "database/sql"
  "database/sql/driver"

  "github.com/calebthompson/bulk"
)

func InsertPeople(people []models.Person) (sql.Result, error) {
  rows := [][]driver.Value{}
  for _, a := range people {
    rows = append(rows, []driver.Value{
      a.ID,
      a.Name,
    })
  }
  return bulk.NewInsert(
    dal.db.DB(),
    `
      INSERT
        INTO people (id, name)
        VALUES <values>
    `,
    []string{"bigint", ""},
  ).Exec(rows)
}
```

In the case above statement, if `people` contains 1,000,000 entries and there
are 2 columns being inserted, 1,000,000 \* 2 / 65535 ~= 30.5 statements would be
needed, so 30 "full" statements would be executed and one "half full" statement
would be executed. The performance of 31 statements that take longer to run is
generally significantly faster vs 1,000,000 statements that are very fast to
run.
