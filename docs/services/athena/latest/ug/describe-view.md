# DESCRIBE VIEW

Shows the list of columns for the specified Athena or AWS Glue Data Catalog view. Useful for
examining the attributes of a complex view.

For Data Catalog views, the output of the statement is controlled by Lake Formation access control and
shows only the columns that the caller has access to.

## Synopsis

```sql

DESCRIBE [db_name.]view_name
```

## Example

```sql

DESCRIBE orders
```

See also [SHOW COLUMNS](show-columns.md), [SHOW CREATE VIEW](show-create-view.md), [SHOW VIEWS](show-views.md), and [DROP VIEW](drop-view.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DESCRIBE

DROP DATABASE

All content copied from https://docs.aws.amazon.com/.
