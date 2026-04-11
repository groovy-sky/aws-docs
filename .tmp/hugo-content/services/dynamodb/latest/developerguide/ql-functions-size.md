# Using the SIZE function with PartiQL for amazon DynamoDB

Returns a number representing an attribute's size in bytes. The following are valid data types
for use with size. For more information, see the DynamoDB [size](expressions-operatorsandfunctions.md#Expressions.OperatorsAndFunctions.Functions) function.

## Syntax

```nohighlight

size( path)
```

## Arguments

`path`

(Required) The attribute name or document path.

For supported types, see DynamoDB [size](expressions-operatorsandfunctions.md#Expressions.OperatorsAndFunctions.Functions) function.

## Return type

`int`

## Examples

```sql

 SELECT * FROM "Orders" WHERE "OrderID"=1 AND size("Image") >300
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Contains

Operators

All content copied from https://docs.aws.amazon.com/.
