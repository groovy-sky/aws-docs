---
title: "BatchStatementRequest"
---

# BatchStatementRequest

A PartiQL batch statement request.

## Contents

###### Note

In the following list, the required parameters are described first.

**Statement**

A valid PartiQL statement.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 8192.

Required: Yes

**ConsistentRead**

The read consistency of the PartiQL batch request.

Type: Boolean

Required: No

**Parameters**

The parameters associated with a PartiQL statement in the batch request.

Type: Array of [AttributeValue](api-attributevalue.md) objects

Array Members: Minimum number of 1 item.

Required: No

**ReturnValuesOnConditionCheckFailure**

An optional parameter that returns the item attributes for a PartiQL batch request
operation that failed a condition check.

There is no additional cost associated with requesting a return value aside from the
small network and processing overhead of receiving a larger response. No read capacity
units are consumed.

Type: String

Valid Values: `ALL_OLD | NONE`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/dynamodb-2012-08-10/BatchStatementRequest)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/dynamodb-2012-08-10/BatchStatementRequest)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/dynamodb-2012-08-10/BatchStatementRequest)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BatchStatementError

BatchStatementResponse

All content copied from https://docs.aws.amazon.com/.
