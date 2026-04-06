# DROP VIEW

Drops (deletes) an existing Athena or AWS Glue Data Catalog view. The optional `IF EXISTS`
clause causes the error to be suppressed if the view does not exist.

For Data Catalog views, drops the view only if Athena view syntax (dialect) is present in the
Data Catalog view. For example, if a user calls `DROP VIEW` from Athena, the view is
dropped only if the Athena dialect exists in the view. Otherwise, the operation fails.
Dropping Data Catalog views requires Lake Formation admin or view definer permissions.

For more information, see [Work with views](https://docs.aws.amazon.com/athena/latest/ug/views.html) and [Use Data Catalog views in Athena](https://docs.aws.amazon.com/athena/latest/ug/views-glue.html).

## Synopsis

```sql

DROP VIEW [ IF EXISTS ] view_name
```

## Examples

```sql

DROP VIEW orders_by_date
```

```

DROP VIEW IF EXISTS orders_by_date
```

See also [CREATE VIEW and CREATE PROTECTED MULTI DIALECT VIEW](create-view.md), [SHOW COLUMNS](https://docs.aws.amazon.com/athena/latest/ug/show-columns.html), [SHOW CREATE VIEW](https://docs.aws.amazon.com/athena/latest/ug/show-create-view.html), [SHOW VIEWS](https://docs.aws.amazon.com/athena/latest/ug/show-views.html), and [DESCRIBE VIEW](describe-view.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DROP TABLE

MSCK REPAIR TABLE
