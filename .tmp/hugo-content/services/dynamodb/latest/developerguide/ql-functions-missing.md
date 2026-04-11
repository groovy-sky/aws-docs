# Using the MISSING function with PartiQL for DynamoDB

Returns `TRUE` if the item does not contain the attribute specified. Only equality and inequality operators can be used with this function.

## Syntax

```nohighlight

 attributename IS | IS NOT  MISSING
```

## Arguments

`attributename`

(Required) The attribute name to look for.

## Return type

`bool`

## Examples

```sql

SELECT * FROM Music WHERE "Awards" is MISSING
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Begins\_with

Attribute\_type

All content copied from https://docs.aws.amazon.com/.
