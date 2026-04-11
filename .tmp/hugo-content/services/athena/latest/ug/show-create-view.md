# SHOW CREATE VIEW

Shows the SQL statement that created the specified Athena or Data Catalog view. The SQL returned
shows the create view syntax used in Athena. Calling `SHOW CREATE VIEW` on Data Catalog
views requires Lake Formation admin or view definer permissions.

## Synopsis

```sql

SHOW CREATE VIEW view_name
```

## Examples

```sql

SHOW CREATE VIEW orders_by_date
```

See also [CREATE VIEW and CREATE PROTECTED MULTI DIALECT VIEW](create-view.md) and [DROP VIEW](drop-view.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SHOW CREATE TABLE

SHOW DATABASES

All content copied from https://docs.aws.amazon.com/.
