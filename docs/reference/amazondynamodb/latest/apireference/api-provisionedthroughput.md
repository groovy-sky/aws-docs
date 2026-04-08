# ProvisionedThroughput

Represents the provisioned throughput settings for the specified global secondary
index. You must use `ProvisionedThroughput` or
`OnDemandThroughput` based on your table’s capacity mode.

For current minimum and maximum provisioned throughput values, see [Service,\
Account, and Table Quotas](../../../../services/dynamodb/latest/developerguide/limits.md) in the _Amazon DynamoDB Developer_
_Guide_.

## Contents

###### Note

In the following list, the required parameters are described first.

**ReadCapacityUnits**

The maximum number of strongly consistent reads consumed per second before DynamoDB
returns a `ThrottlingException`. For more information, see [Specifying\
Read and Write Requirements](../../../../services/dynamodb/latest/developerguide/provisionedthroughput.md) in the _Amazon DynamoDB Developer_
_Guide_.

If read/write capacity mode is `PAY_PER_REQUEST` the value is set to
0.

Type: Long

Valid Range: Minimum value of 1.

Required: Yes

**WriteCapacityUnits**

The maximum number of writes consumed per second before DynamoDB returns a
`ThrottlingException`. For more information, see [Specifying\
Read and Write Requirements](../../../../services/dynamodb/latest/developerguide/provisionedthroughput.md) in the _Amazon DynamoDB Developer_
_Guide_.

If read/write capacity mode is `PAY_PER_REQUEST` the value is set to
0.

Type: Long

Valid Range: Minimum value of 1.

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/provisionedthroughput.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/provisionedthroughput.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/provisionedthroughput.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Projection

ProvisionedThroughputDescription
