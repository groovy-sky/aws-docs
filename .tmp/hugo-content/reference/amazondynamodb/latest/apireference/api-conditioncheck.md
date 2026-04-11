# ConditionCheck

Represents a request to perform a check that an item exists or to check the condition
of specific attributes of the item.

## Contents

###### Note

In the following list, the required parameters are described first.

**ConditionExpression**

A condition that must be satisfied in order for a conditional update to succeed. For
more information, see [Condition expressions](../../../../services/dynamodb/latest/developerguide/expressions-conditionexpressions.md) in the _Amazon DynamoDB Developer_
_Guide_.

Type: String

Required: Yes

**Key**

The primary key of the item to be checked. Each element consists of an attribute name
and a value for that attribute.

Type: String to [AttributeValue](api-attributevalue.md) object map

Key Length Constraints: Maximum length of 65535.

Required: Yes

**TableName**

Name of the table for the check item request. You can also provide the Amazon Resource Name (ARN) of
the table in this parameter.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: Yes

**ExpressionAttributeNames**

One or more substitution tokens for attribute names in an expression. For more
information, see [Expression attribute names](../../../../services/dynamodb/latest/developerguide/expressions-expressionattributenames.md) in the _Amazon DynamoDB Developer_
_Guide_.

Type: String to string map

Value Length Constraints: Maximum length of 65535.

Required: No

**ExpressionAttributeValues**

One or more values that can be substituted in an expression. For more information, see
[Condition expressions](../../../../services/dynamodb/latest/developerguide/expressions-conditionexpressions.md) in the _Amazon DynamoDB Developer_
_Guide_.

Type: String to [AttributeValue](api-attributevalue.md) object map

Required: No

**ReturnValuesOnConditionCheckFailure**

Use `ReturnValuesOnConditionCheckFailure` to get the item attributes if the
`ConditionCheck` condition fails. For
`ReturnValuesOnConditionCheckFailure`, the valid values are: NONE and
ALL\_OLD.

Type: String

Valid Values: `ALL_OLD | NONE`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/conditioncheck.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/conditioncheck.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/conditioncheck.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Condition

ConsumedCapacity

All content copied from https://docs.aws.amazon.com/.
