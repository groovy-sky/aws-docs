---
title: "WriteRequest"
---

# WriteRequest

Represents an operation to perform - either `DeleteItem` or
`PutItem`. You can only request one of these operations, not both, in a
single `WriteRequest`. If you do need to perform both of these operations,
you need to provide two separate `WriteRequest` objects.

## Contents

###### Note

In the following list, the required parameters are described first.

**DeleteRequest**

A request to perform a `DeleteItem` operation.

Type: [DeleteRequest](api-deleterequest.md) object

Required: No

**PutRequest**

A request to perform a `PutItem` operation.

Type: [PutRequest](api-putrequest.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/dynamodb-2012-08-10/WriteRequest)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/dynamodb-2012-08-10/WriteRequest)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/dynamodb-2012-08-10/WriteRequest)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WarmThroughput

DynamoDB Accelerator

All content copied from https://docs.aws.amazon.com/.
