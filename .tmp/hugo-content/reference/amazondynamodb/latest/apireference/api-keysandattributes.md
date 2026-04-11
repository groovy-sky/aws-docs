# KeysAndAttributes

Represents a set of primary keys and, for each key, the attributes to retrieve from
the table.

For each primary key, you must provide _all_ of the key attributes.
For example, with a simple primary key, you only need to provide the partition key. For
a composite primary key, you must provide _both_ the partition key
and the sort key.

## Contents

###### Note

In the following list, the required parameters are described first.

**Keys**

The primary key attribute values that define the items and the attributes associated
with the items.

Type: Array of string to [AttributeValue](api-attributevalue.md) object maps

Array Members: Minimum number of 1 item. Maximum number of 100 items.

Key Length Constraints: Maximum length of 65535.

Required: Yes

**AttributesToGet**

This is a legacy parameter. Use `ProjectionExpression` instead. For more
information, see [Legacy\
Conditional Parameters](../../../../services/dynamodb/latest/developerguide/legacyconditionalparameters.md) in the _Amazon DynamoDB Developer_
_Guide_.

Type: Array of strings

Array Members: Minimum number of 1 item.

Length Constraints: Maximum length of 65535.

Required: No

**ConsistentRead**

The consistency of a read operation. If set to `true`, then a strongly
consistent read is used; otherwise, an eventually consistent read is used.

Type: Boolean

Required: No

**ExpressionAttributeNames**

One or more substitution tokens for attribute names in an expression. The following
are some use cases for using `ExpressionAttributeNames`:

- To access an attribute whose name conflicts with a DynamoDB reserved
word.

- To create a placeholder for repeating occurrences of an attribute name in an
expression.

- To prevent special characters in an attribute name from being misinterpreted
in an expression.

Use the **#** character in an expression to dereference
an attribute name. For example, consider the following attribute name:

- `Percentile`

The name of this attribute conflicts with a reserved word, so it cannot be used
directly in an expression. (For the complete list of reserved words, see [Reserved Words](../../../../services/dynamodb/latest/developerguide/reservedwords.md) in the _Amazon DynamoDB Developer_
_Guide_). To work around this, you could specify the following for
`ExpressionAttributeNames`:

- `{"#P":"Percentile"}`

You could then use this substitution in an expression, as in this example:

- `#P = :val`

###### Note

Tokens that begin with the **:** character are
_expression attribute values_, which are placeholders for the
actual value at runtime.

For more information on expression attribute names, see [Accessing Item Attributes](../../../../services/dynamodb/latest/developerguide/expressions-accessingitemattributes.md) in the _Amazon DynamoDB Developer_
_Guide_.

Type: String to string map

Value Length Constraints: Maximum length of 65535.

Required: No

**ProjectionExpression**

A string that identifies one or more attributes to retrieve from the table. These
attributes can include scalars, sets, or elements of a JSON document. The attributes in
the `ProjectionExpression` must be separated by commas.

If no attribute names are specified, then all attributes will be returned. If any of
the requested attributes are not found, they will not appear in the result.

For more information, see [Accessing Item Attributes](../../../../services/dynamodb/latest/developerguide/expressions-accessingitemattributes.md) in the _Amazon DynamoDB Developer_
_Guide_.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/keysandattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/keysandattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/keysandattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ItemResponse

KeySchemaElement

All content copied from https://docs.aws.amazon.com/.
