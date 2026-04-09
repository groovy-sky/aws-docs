# DEALLOCATE PREPARE

Removes the prepared statement with the specified name from the list of
prepared statements in the current workgroup.

## Syntax

```sql

DEALLOCATE PREPARE statement_name
```

`statement_name` is the name of the prepared
statement to be removed.

## Example

The following example removes the `my_select1` prepared
statement from the current workgroup.

```sql

DEALLOCATE PREPARE my_select1
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EXECUTE

Use the Athena console

All content copied from https://docs.aws.amazon.com/.
