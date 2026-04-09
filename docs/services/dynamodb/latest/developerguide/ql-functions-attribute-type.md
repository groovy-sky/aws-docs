# Using the ATTRIBUTE\_TYPE function with PartiQL for DynamoDB

Returns `TRUE` if the attribute at the specified path is of a particular
data type.

## Syntax

```nohighlight

attribute_type( attributename, type )
```

## Arguments

`attributename`

(Required) The attribute name to use.

`type`

(Required) The attribute type to check for. For a list of valid
values, see DynamoDB [attribute\_type](expressions-operatorsandfunctions.md#Expressions.OperatorsAndFunctions.Functions).

## Return type

`bool`

## Examples

```sql

SELECT * FROM "Music" WHERE attribute_type("Artist", 'S')
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Missing

Contains

All content copied from https://docs.aws.amazon.com/.
