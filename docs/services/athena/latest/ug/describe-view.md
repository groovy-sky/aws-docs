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

See also [SHOW COLUMNS](https://docs.aws.amazon.com/athena/latest/ug/show-columns.html), [SHOW CREATE VIEW](https://docs.aws.amazon.com/athena/latest/ug/show-create-view.html), [SHOW VIEWS](https://docs.aws.amazon.com/athena/latest/ug/show-views.html), and [DROP VIEW](https://docs.aws.amazon.com/athena/latest/ug/drop-view.html).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DESCRIBE

DROP DATABASE
