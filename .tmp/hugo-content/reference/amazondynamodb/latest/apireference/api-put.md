# Put

Represents a request to perform a `PutItem` operation.

## Contents

###### Note

In the following list, the required parameters are described first.

**Item**

A map of attribute name to attribute values, representing the primary key of the item
to be written by `PutItem`. All of the table's primary key attributes must be
specified, and their data types must match those of the table's key schema. If any
attributes are present in the item that are part of an index key schema for the table,
their types must match the index key schema.

Type: String to [AttributeValue](api-attributevalue.md) object map

Key Length Constraints: Maximum length of 65535.

Required: Yes

**TableName**

Name of the table in which to write the item. You can also provide the Amazon Resource Name (ARN) of
the table in this parameter.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: Yes

**ConditionExpression**

A condition that must be satisfied in order for a conditional update to
succeed.

Type: String

Required: No

**ExpressionAttributeNames**

One or more substitution tokens for attribute names in an expression.

Type: String to string map

Value Length Constraints: Maximum length of 65535.

Required: No

**ExpressionAttributeValues**

One or more values that can be substituted in an expression.

Type: String to [AttributeValue](api-attributevalue.md) object map

Required: No

**ReturnValuesOnConditionCheckFailure**

Use `ReturnValuesOnConditionCheckFailure` to get the item attributes if the
`Put` condition fails. For
`ReturnValuesOnConditionCheckFailure`, the valid values are: NONE and
ALL\_OLD.

Type: String

Valid Values: `ALL_OLD | NONE`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/put.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/put.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/put.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ProvisionedThroughputOverride

PutRequest

All content copied from https://docs.aws.amazon.com/.
