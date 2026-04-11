# Differences between a relational (SQL) database and DynamoDB when removing a table

In SQL, you use the `DROP TABLE` statement to remove a table. In Amazon DynamoDB,
you use the `DeleteTable` operation.

###### Topics

- [Removing a table with SQL](#SQLtoNoSQL.RemoveTable.SQL)

- [Removing a table in DynamoDB](#SQLtoNoSQL.RemoveTable.DynamoDB)

## Removing a table with SQL

When you no longer need a table and want to discard it permanently, you would use
the `DROP TABLE` statement in SQL.

```sql

DROP TABLE Music;
```

After a table is dropped, it cannot be recovered. (Some relational databases do
allow you to undo a `DROP TABLE` operation, but this is vendor-specific
functionality and it is not widely implemented.)

## Removing a table in DynamoDB

In DynamoDB, `DeleteTable` is a similar operation. In the following
example, the table is permanently deleted.

```nohighlight

{
    TableName: "Music"
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Deleting data from a table

Amazon DynamoDB learning resources and tools

All content copied from https://docs.aws.amazon.com/.
