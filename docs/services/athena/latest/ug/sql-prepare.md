# PREPARE

Creates a SQL statement with the name `statement_name` to be run at a later
time. The statement can include parameters represented by question marks. To supply values
for the parameters and run the prepared statement, use [EXECUTE](https://docs.aws.amazon.com/athena/latest/ug/sql-execute.html).

## Synopsis

```sql

PREPARE statement_name FROM statement
```

The following table describes the parameters.

ParameterDescription`statement_name`The name of the statement to be prepared. The name must be unique
within the workgroup.`statement`A `SELECT`, `CTAS`, or `INSERT INTO`
query.

###### Note

The maximum number of prepared statements in a workgroup is 1000.

## Examples

The following example prepares a select query without parameters.

```sql

PREPARE my_select1 FROM
SELECT * FROM nation
```

The following example prepares a select query that includes parameters. The values for
`productid` and `quantity` will be supplied by the
`USING` clause of an `EXECUTE` statement:

```sql

PREPARE my_select2 FROM
SELECT order FROM orders WHERE productid = ? and quantity < ?
```

The following example prepares an insert query.

```sql

PREPARE my_insert FROM
INSERT INTO cities_usa (city, state)
SELECT city, state
FROM cities_world
WHERE country = ?
```

## Additional resources

[Use prepared statements](https://docs.aws.amazon.com/athena/latest/ug/querying-with-prepared-statements-querying.html)

[EXECUTE](https://docs.aws.amazon.com/athena/latest/ug/sql-execute.html)

[DEALLOCATE PREPARE](https://docs.aws.amazon.com/athena/latest/ug/sql-deallocate-prepare.html)

[INSERT INTO](insert-into.md)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Understand EXPLAIN results

EXECUTE
