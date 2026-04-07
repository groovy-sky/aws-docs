# List partitions for a specific table

You can use `SHOW PARTITIONS table_name` to list
the partitions for a specified table, as in the following example.

```sql

SHOW PARTITIONS cloudtrail_logs_test2
```

You can also use a `$partitions` metadata query to list the partition
numbers and partition values for a specific table.

###### Example– Querying the partitions for a table using the $partitions syntax

The following example query lists the partitions for the table
`cloudtrail_logs_test2` using the `$partitions`
syntax.

```sql

SELECT * FROM default."cloudtrail_logs_test2$partitions" ORDER BY partition_number
```

The following table shows sample results.

table\_catalogtable\_schematable\_nameYearMonthDay1awsdatacatalogdefaultcloudtrail\_logs\_test2202008102awsdatacatalogdefaultcloudtrail\_logs\_test2202008113awsdatacatalogdefaultcloudtrail\_logs\_test220200812

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

List tables

List columns for one table or view
