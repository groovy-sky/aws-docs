# Get

Specifies an item and related attribute values to retrieve in a
`TransactGetItem` object.

## Contents

###### Note

In the following list, the required parameters are described first.

**Key**

A map of attribute names to `AttributeValue` objects that specifies the
primary key of the item to retrieve.

Type: String to [AttributeValue](api-attributevalue.md) object map

Key Length Constraints: Maximum length of 65535.

Required: Yes

**TableName**

The name of the table from which to retrieve the specified item. You can also provide
the Amazon Resource Name (ARN) of the table in this parameter.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: Yes

**ExpressionAttributeNames**

One or more substitution tokens for attribute names in the ProjectionExpression
parameter.

Type: String to string map

Value Length Constraints: Maximum length of 65535.

Required: No

**ProjectionExpression**

A string that identifies one or more attributes of the specified item to retrieve from
the table. The attributes in the expression must be separated by commas. If no attribute
names are specified, then all attributes of the specified item are returned. If any of
the requested attributes are not found, they do not appear in the result.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/get.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/get.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/get.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FailureException

GlobalSecondaryIndex

All content copied from https://docs.aws.amazon.com/.
