# List or search columns for a specified table or view

You can list all columns for a table, all columns for a view, or search for a column
by name in a specified database and table.

To list the columns, use a `SELECT *` query. In the `FROM`
clause, specify `information_schema.columns`. In the `WHERE`
clause, use `table_schema='database_name'` to
specify the database and `table_name =
                'table_name'` to specify the table or view that has
the columns that you want to list.

###### Example– Listing all columns for a specified table

The following example query lists all columns for the table
`rdspostgresqldb1_public_account`.

```sql

SELECT *
FROM   information_schema.columns
WHERE  table_schema = 'rdspostgresql'
       AND table_name = 'rdspostgresqldb1_public_account'
```

The following table shows sample results.

table\_catalogtable\_schematable\_namecolumn\_nameordinal\_positioncolumn\_defaultis\_nullabledata\_typecommentextra\_info1awsdatacatalogrdspostgresqlrdspostgresqldb1\_public\_accountpassword1YESvarchar2awsdatacatalogrdspostgresqlrdspostgresqldb1\_public\_accountuser\_id2YESinteger3awsdatacatalogrdspostgresqlrdspostgresqldb1\_public\_accountcreated\_on3YEStimestamp4awsdatacatalogrdspostgresqlrdspostgresqldb1\_public\_accountlast\_login4YEStimestamp5awsdatacatalogrdspostgresqlrdspostgresqldb1\_public\_accountemail5YESvarchar6awsdatacatalogrdspostgresqlrdspostgresqldb1\_public\_accountusername6YESvarchar

###### Example– Listing the columns for a specified view

The following example query lists all the columns in the `default`
database for the view `arrayview`.

```sql

SELECT *
FROM   information_schema.columns
WHERE  table_schema = 'default'
       AND table_name = 'arrayview'
```

The following table shows sample results.

table\_catalogtable\_schematable\_namecolumn\_nameordinal\_positioncolumn\_defaultis\_nullabledata\_typecommentextra\_info1awsdatacatalogdefaultarrayviewsearchdate1YESvarchar2awsdatacatalogdefaultarrayviewsid2YESvarchar3awsdatacatalogdefaultarrayviewbtid3YESvarchar4awsdatacatalogdefaultarrayviewp4YESvarchar5awsdatacatalogdefaultarrayviewinfantprice5YESvarchar6awsdatacatalogdefaultarrayviewsump6YESvarchar7awsdatacatalogdefaultarrayviewjourneymaparray7YESarray(varchar)

###### Example– Searching for a column by name in a specified database and table

The following example query searches for metadata for the `sid` column
in the `arrayview` view of the `default` database.

```sql

SELECT *
FROM   information_schema.columns
WHERE  table_schema = 'default'
       AND table_name = 'arrayview'
       AND column_name='sid'
```

The following table shows a sample result.

table\_catalogtable\_schematable\_namecolumn\_nameordinal\_positioncolumn\_defaultis\_nullabledata\_typecommentextra\_info1awsdatacatalogdefaultarrayviewsid2YESvarchar

[Document Conventions](../../../../general/latest/gr/docconventions.md)

List partitions

List columns in common

All content copied from https://docs.aws.amazon.com/.
