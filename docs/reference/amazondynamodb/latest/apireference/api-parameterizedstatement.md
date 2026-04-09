# ParameterizedStatement

Represents a PartiQL statement that uses parameters.

## Contents

###### Note

In the following list, the required parameters are described first.

**Statement**

A PartiQL statement that uses parameters.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 8192.

Required: Yes

**Parameters**

The parameter values.

Type: Array of [AttributeValue](api-attributevalue.md) objects

Array Members: Minimum number of 1 item.

Required: No

**ReturnValuesOnConditionCheckFailure**

An optional parameter that returns the item attributes for a PartiQL
`ParameterizedStatement` operation that failed a condition check.

There is no additional cost associated with requesting a return value aside from the
small network and processing overhead of receiving a larger response. No read capacity
units are consumed.

Type: String

Valid Values: `ALL_OLD | NONE`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/parameterizedstatement.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/parameterizedstatement.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/parameterizedstatement.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OnDemandThroughputOverride

PointInTimeRecoveryDescription

All content copied from https://docs.aws.amazon.com/.
