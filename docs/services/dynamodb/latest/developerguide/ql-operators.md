# PartiQL arithmetic, comparison, and logical operators for DynamoDB

PartiQL in Amazon DynamoDB supports the following [SQL standard\
operators](https://www.w3schools.com/sql/sql_operators.asp).

###### Note

Any SQL operators that are not included in this list are not currently supported in
DynamoDB.

## Arithmetic operators

OperatorDescription`+`Add`-`Subtract

## Comparison operators

OperatorDescription`=`Equal to`<>`Not Equal to`!=`Not Equal to`>`Greater than`<`Less than`>=`Greater than or equal to`<=`Less than or equal to

## Logical operators

OperatorDescription`AND``TRUE` if all the conditions separated by `AND`
are `TRUE``BETWEEN`

`TRUE` if the operand is within the range of comparisons.

This operator is inclusive of the lower and upper bound of the operands on which you apply it.

`IN`

`TRUE` if the operand is equal to one of a list of
expressions (at max 50 hash attribute values or at max 100 non-key
attribute values).

Results are returned in pages of up to 10 items. If the
`IN` list contains more values, you must use the
`NextToken` returned in the response to retrieve subsequent
pages.

`IS``TRUE` if the operand is a given, PartiQL data type, including
`NULL` or `MISSING``NOT`Reverses the value of a given Boolean expression`OR``TRUE` if any of the conditions separated by
`OR` are `TRUE`

For more information about using logical operators, see [Making comparisons](expressions-operatorsandfunctions.md#Expressions.OperatorsAndFunctions.Comparators) and [Logical evaluations](expressions-operatorsandfunctions.md#Expressions.OperatorsAndFunctions.LogicalEvaluations).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Size

Transactions
