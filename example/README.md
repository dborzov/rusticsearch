rusticsearch-example
====================

Examples of configured [rusticsearch](https://github.com/dborzov/rusticsearch) servers:
searching the database of  Adventure Time characters and episodes.

We have an SQLite database file `database/example.db`. Let's open up SQLite shell and see the tables available:
```bash

  $ sqlite example.db
  > .tables
  episodes  characters appearances
  > .schema characters
  CREATE TABLE (
    name string
  )
```

So it goes.
