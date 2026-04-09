# Using expressions in DynamoDB

In Amazon DynamoDB, you can use _expressions_ to specify which attributes to
read from an item, write data when a condition is met, specify how to update an item, define
queries and filter the results of a query.

This table describes the basic expression grammar and the available kinds of
expressions.

Expression typeDescriptionProjection expressionA projection expression identifies the attributes that you want to
retrieve from an item when you use operations such as GetItem, Query, or
Scan.Condition expressionA condition expression determines which items should be modified when you
use the PutItem, UpdateItem, and DeleteItem operations.Update expressionAn update expression specifies how UpdateItem will modify the attributes
of an item— for example, setting a scalar value or removing elements from a
list or a map.Key condition expressionA key condition expression determines which items a query will read from
a table or index.Filter expressionA filter expression determines which items among the Query results should
be returned to you. All the other results are discarded.

For information about expression syntax and more detailed information about each type of
expression, see the following sections.

###### Topics

- [Referring to item attributes when using expressions in DynamoDB](expressions-attributes.md)

- [Expression attribute names (aliases) in DynamoDB](expressions-expressionattributenames.md)

- [Using expression attribute values in DynamoDB](expressions-expressionattributevalues.md)

- [Using projection expressions in DynamoDB](expressions-projectionexpressions.md)

- [Using update expressions in DynamoDB](expressions-updateexpressions.md)

- [Condition and filter expressions, operators, and functions in DynamoDB](expressions-operatorsandfunctions.md)

- [DynamoDB condition expression CLI example](expressions-conditionexpressions.md)

###### Note

For backward compatibility, DynamoDB also supports conditional parameters that do not use
expressions. For more information, see [Legacy DynamoDB conditional parameters](legacyconditionalparameters.md).

New applications should use expressions rather than the legacy parameters.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Item sizes and formats

Specifying item attributes

All content copied from https://docs.aws.amazon.com/.
