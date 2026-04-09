# Using the CONTAINS function with PartiQL for DynamoDB

Returns `TRUE` if the attribute specified by the path is one of the
following:

- A String that contains a particular substring.

- A Set that contains a particular element within the set.

For more information, see the DynamoDB [contains](expressions-operatorsandfunctions.md#Expressions.OperatorsAndFunctions.Functions) function.

## Syntax

```nohighlight

contains( path, substring )
```

## Arguments

`path`

(Required) The attribute name or document path to use.

`substring`

(Required) The attribute substring or set member to check for. For
more information, see the DynamoDB [contains](expressions-operatorsandfunctions.md#Expressions.OperatorsAndFunctions.Functions) function.

## Return type

`bool`

## Examples

```sql

SELECT * FROM "Orders" WHERE "OrderID"=1 AND contains("Address", 'Kirkland')
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Attribute\_type

Size

All content copied from https://docs.aws.amazon.com/.
