# ThrottlingReason

Represents the specific reason why a DynamoDB request was throttled and the
ARN of the impacted resource. This helps identify exactly what resource is being throttled,
what type of operation caused it, and why the throttling occurred.

## Contents

###### Note

In the following list, the required parameters are described first.

**reason**

The reason for throttling. The throttling reason follows a specific format:
`ResourceType+OperationType+LimitType`:

- Resource Type (What is being throttled): Table or Index

- Operation Type (What kind of operation): Read or Write

- Limit Type (Why the throttling occurred):

- `ProvisionedThroughputExceeded`: The request rate is
exceeding the [provisioned throughput capacity](../../../../services/dynamodb/latest/developerguide/provisioned-capacity-mode.md) (read or write capacity
units) configured for a table or a global secondary index (GSI) in
provisioned capacity mode.

- `AccountLimitExceeded`: The request rate has caused a table
or global secondary index (GSI) in on-demand mode to exceed the [per-table account-level service quotas](../../../../services/dynamodb/latest/developerguide/servicequotas.md#default-limits-throughput) for read/write
throughput in the current AWS Region.

- `KeyRangeThroughputExceeded`: The request rate directed at
a specific partition key value has exceeded the [internal partition-level throughput limits](../../../../services/dynamodb/latest/developerguide/bp-partition-key-design.md), indicating
uneven access patterns across the table's or GSI's key space.

- `MaxOnDemandThroughputExceeded`: The request rate has
exceeded the [configured maximum throughput limits](../../../../services/dynamodb/latest/developerguide/on-demand-capacity-mode-max-throughput.md) set for a table or
index in on-demand capacity mode.

Examples of complete throttling reasons:

- TableReadProvisionedThroughputExceeded

- IndexWriteAccountLimitExceeded

This helps identify exactly what resource is being throttled, what type of operation
caused it, and why the throttling occurred.

Type: String

Required: No

**resource**

The Amazon Resource Name (ARN) of the DynamoDB table or index that experienced the
throttling event.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/dynamodb-2012-08-10/ThrottlingReason)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/dynamodb-2012-08-10/ThrottlingReason)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/dynamodb-2012-08-10/ThrottlingReason)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

TimeToLiveDescription
