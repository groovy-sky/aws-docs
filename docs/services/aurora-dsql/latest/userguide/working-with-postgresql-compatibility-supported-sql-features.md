# Supported SQL for Aurora DSQL

Aurora DSQL supports a wide range of core PostgreSQL SQL features. In the following sections, you
can learn about general PostgreSQL expression support. This list is not exhaustive.

## `SELECT` command

Aurora DSQL supports the following clauses of the `SELECT` command.

Primary clauseSupported clauses

`FROM`

`GROUP BY`

`ALL`, `DISTINCT`

`ORDER BY`

`ASC`, `DESC`, `NULLS`

`LIMIT`

`DISTINCT`

`HAVING`

`USING`

`WITH` (common table expressions)

`INNER JOIN`

`ON`

`OUTER JOIN`

`LEFT`, `RIGHT`, `FULL`,
`ON`

`CROSS JOIN`

`ON`

`UNION`

`ALL`

`INTERSECT`

`ALL`

`EXCEPT`

`ALL`

`OVER`

`RANK ()`, `PARTITION BY`

`FOR UPDATE`

## Data Definition Language (DDL)

Aurora DSQL supports the following PostgreSQL DDL commands.

CommandPrimary ClauseSupported Clauses

`CREATE`

`TABLE`

For information about the supported syntax of the `CREATE TABLE`
command, see [CREATE TABLE](create-table-syntax-support.md).

`ALTER`

`TABLE`

For information about the supported syntax of the `ALTER TABLE`
command, see [ALTER TABLE](alter-table-syntax-support.md).

`DROP`

`TABLE`

`CREATE`

`[UNIQUE] INDEX ASYNC`

You can use this command with the following parameters: `ON`,
`NULLS FIRST`, `NULLS LAST`.

For information about the supported syntax of the `CREATE INDEX
                      ASYNC` command, see [Asynchronous indexes in Aurora DSQL](working-with-create-index-async.md).

`DROP`

`INDEX`

`CREATE`

`VIEW`

For more information about the supported syntax of the `CREATE
                      VIEW` command, see [CREATE VIEW](create-view.md).

`ALTER``VIEW`

For information about the supported syntax of the `ALTER VIEW`
command, see [ALTER VIEW](alter-view-syntax-support.md).

`DROP``VIEW`For information about the supported syntax of the `DROP VIEW`
command, see [DROP VIEW](drop-view-overview.md).

`CREATE`

`SEQUENCE`

For information about the supported syntax of the `CREATE
                      SEQUENCE` command, see [CREATE SEQUENCE](create-sequence-syntax-support.md).

`ALTER`

`SEQUENCE`

For information about the supported syntax of the `ALTER
                      SEQUENCE` command, see [ALTER SEQUENCE](alter-sequence-syntax-support.md).

`DROP`

`SEQUENCE`

For information about the supported syntax of the `DROP
                      SEQUENCE` command, see [DROP SEQUENCE](drop-sequence-syntax-support.md).

`CREATE`

`ROLE`, `WITH`

`CREATE`

`FUNCTION`

`LANGUAGE SQL`

`CREATE`

`DOMAIN`

## Data Manipulation Language (DML)

Aurora DSQL supports the following PostgreSQL DML commands.

CommandPrimary clauseSupported clauses

`INSERT`

`INTO`

`VALUES`

`SELECT`

`UPDATE`

`SET`

`WHERE (SELECT)`

`FROM, WITH`

`DELETE``FROM``USING`, `WHERE`

## Data Control Language (DCL)

Aurora DSQL supports the following PostgreSQL DCL commands.

CommandSupported clauses

`GRANT`

`ON`, `TO`

`REVOKE`

`ON`, `FROM`, `CASCADE`,
`RESTRICT`

## Transaction Control Language (TCL)

Aurora DSQL supports the following PostgreSQL TCL commands.

CommandSupported clausesAlias

`COMMIT`

\[ `WORK` \| `TRANSACTION`\]

\[ `AND NO CHAIN`\]

`END`

`BEGIN`

\[ `WORK` \| `TRANSACTION`\]

\[ `ISOLATION LEVEL REPEATABLE READ`\]

\[ `READ WRITE` \| `READ ONLY`\]

`START TRANSACTION`

\[ `ISOLATION LEVEL REPEATABLE READ`\]

\[ `READ WRITE` \| `READ ONLY`\]

`ROLLBACK`

\[ `WORK` \| `TRANSACTION`\]

\[ `AND NO CHAIN`\]

`ABORT`

## Utility commands

Aurora DSQL supports the following PostgreSQL utility commands:

- `EXPLAIN`

- `ANALYZE` (relation name only)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Supported data types

Supported subsets of SQL commands

All content copied from https://docs.aws.amazon.com/.
