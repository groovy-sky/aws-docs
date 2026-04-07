# Evolve Iceberg table schema

Iceberg schema updates are metadata-only changes. No data files are changed when you
perform a schema update.

The Iceberg format supports the following schema evolution changes:

- Add – Adds a new column to a table or to
a nested `struct`.

- Drop – Removes an existing column from a
table or nested `struct`.

- Rename – Renames an existing column or
field in a nested `struct`.

- Reorder – Changes the order of
columns.

- Type promotion – Widens the type of a
column, `struct` field, `map` key, `map` value,
or `list` element. Currently, the following cases are supported for
Iceberg tables:

- integer to big integer

- float to double

- increasing the precision of a decimal type

You can use the DDL statements in this section to modify Iceberg table schema.

###### Topics

- [ALTER TABLE ADD COLUMNS](https://docs.aws.amazon.com/athena/latest/ug/querying-iceberg-alter-table-add-columns.html)

- [ALTER TABLE DROP COLUMN](https://docs.aws.amazon.com/athena/latest/ug/querying-iceberg-alter-table-drop-column.html)

- [ALTER TABLE CHANGE COLUMN](https://docs.aws.amazon.com/athena/latest/ug/querying-iceberg-alter-table-change-column.html)

- [SHOW COLUMNS](https://docs.aws.amazon.com/athena/latest/ug/querying-iceberg-show-columns.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SHOW TBLPROPERTIES

ALTER TABLE ADD COLUMNS
