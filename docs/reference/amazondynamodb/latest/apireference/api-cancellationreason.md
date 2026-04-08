# CancellationReason

An ordered list of errors for each item in the request which caused the transaction to
get cancelled. The values of the list are ordered according to the ordering of the
`TransactWriteItems` request parameter. If no error occurred for the
associated item an error with a Null code and Null message will be present.

## Contents

###### Note

In the following list, the required parameters are described first.

**Code**

Status code for the result of the cancelled transaction.

Type: String

Required: No

**Item**

Item in the request which caused the transaction to get cancelled.

Type: String to [AttributeValue](api-attributevalue.md) object map

Key Length Constraints: Maximum length of 65535.

Required: No

**Message**

Cancellation reason message description.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/cancellationreason.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/cancellationreason.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/cancellationreason.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

BillingModeSummary

Capacity
