# DEALLOCATE PREPARE

Removes the prepared statement with the specified name from the prepared statements in the
current workgroup.

## Synopsis

```

DEALLOCATE PREPARE statement_name
```

## Examples

The following example removes the `my_select1` prepared statement from the
current workgroup.

```

DEALLOCATE PREPARE my_select1
```

## Additional resources

[Use prepared statements](querying-with-prepared-statements-querying.md)

[PREPARE](sql-prepare.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EXECUTE

UNLOAD

All content copied from https://docs.aws.amazon.com/.
