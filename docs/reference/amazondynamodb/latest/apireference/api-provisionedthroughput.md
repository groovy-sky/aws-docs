# ProvisionedThroughput

Represents the provisioned throughput settings for the specified global secondary
index. You must use `ProvisionedThroughput` or
`OnDemandThroughput` based on your table’s capacity mode.

For current minimum and maximum provisioned throughput values, see [Service,\
Account, and Table Quotas](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html) in the _Amazon DynamoDB Developer_
_Guide_.

## Contents

###### Note

In the following list, the required parameters are described first.

**ReadCapacityUnits**

The maximum number of strongly consistent reads consumed per second before DynamoDB
returns a `ThrottlingException`. For more information, see [Specifying\
Read and Write Requirements](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughput.html) in the _Amazon DynamoDB Developer_
_Guide_.

If read/write capacity mode is `PAY_PER_REQUEST` the value is set to
0.

Type: Long

Valid Range: Minimum value of 1.

Required: Yes

**WriteCapacityUnits**

The maximum number of writes consumed per second before DynamoDB returns a
`ThrottlingException`. For more information, see [Specifying\
Read and Write Requirements](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughput.html) in the _Amazon DynamoDB Developer_
_Guide_.

If read/write capacity mode is `PAY_PER_REQUEST` the value is set to
0.

Type: Long

Valid Range: Minimum value of 1.

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/dynamodb-2012-08-10/ProvisionedThroughput)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/dynamodb-2012-08-10/ProvisionedThroughput)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/dynamodb-2012-08-10/ProvisionedThroughput)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Projection

ProvisionedThroughputDescription
