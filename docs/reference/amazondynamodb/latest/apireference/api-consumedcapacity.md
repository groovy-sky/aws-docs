# ConsumedCapacity

The capacity units consumed by an operation. The data returned includes the total
provisioned throughput consumed, along with statistics for the table and any indexes
involved in the operation. `ConsumedCapacity` is only returned if the request
asked for it. For more information, see [Provisioned capacity mode](../../../../services/dynamodb/latest/developerguide/provisioned-capacity-mode.md) in the _Amazon DynamoDB Developer_
_Guide_.

## Contents

###### Note

In the following list, the required parameters are described first.

**CapacityUnits**

The total number of capacity units consumed by the operation.

Type: Double

Required: No

**GlobalSecondaryIndexes**

The amount of throughput consumed on each global index affected by the
operation.

Type: String to [Capacity](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_Capacity.html) object map

Key Length Constraints: Minimum length of 3. Maximum length of 255.

Key Pattern: `[a-zA-Z0-9_.-]+`

Required: No

**LocalSecondaryIndexes**

The amount of throughput consumed on each local index affected by the
operation.

Type: String to [Capacity](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_Capacity.html) object map

Key Length Constraints: Minimum length of 3. Maximum length of 255.

Key Pattern: `[a-zA-Z0-9_.-]+`

Required: No

**ReadCapacityUnits**

The total number of read capacity units consumed by the operation.

Type: Double

Required: No

**Table**

The amount of throughput consumed on the table affected by the operation.

Type: [Capacity](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_Capacity.html) object

Required: No

**TableName**

The name of the table that was affected by the operation. If you had specified the
Amazon Resource Name (ARN) of a table in the input, you'll see the table ARN in the response.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**WriteCapacityUnits**

The total number of write capacity units consumed by the operation.

Type: Double

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/dynamodb-2012-08-10/ConsumedCapacity)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/dynamodb-2012-08-10/ConsumedCapacity)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/dynamodb-2012-08-10/ConsumedCapacity)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ConditionCheck

ContinuousBackupsDescription
