# SourceTableDetails

Contains the details of the table when the backup was created.

## Contents

###### Note

In the following list, the required parameters are described first.

**KeySchema**

Schema of the table.

Type: Array of [KeySchemaElement](api-keyschemaelement.md) objects

Array Members: Minimum number of 1 item.

Required: Yes

**ProvisionedThroughput**

Read IOPs and Write IOPS on the table when the backup was created.

Type: [ProvisionedThroughput](api-provisionedthroughput.md) object

Required: Yes

**TableCreationDateTime**

Time when the source table was created.

Type: Timestamp

Required: Yes

**TableId**

Unique identifier for the table for which the backup was created.

Type: String

Pattern: `[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`

Required: Yes

**TableName**

The name of the table for which the backup was created.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

Pattern: `[a-zA-Z0-9_.-]+`

Required: Yes

**BillingMode**

Controls how you are charged for read and write throughput and how you manage
capacity. This setting can be changed later.

- `PROVISIONED` \- Sets the read/write capacity mode to
`PROVISIONED`. We recommend using `PROVISIONED` for
predictable workloads.

- `PAY_PER_REQUEST` \- Sets the read/write capacity mode to
`PAY_PER_REQUEST`. We recommend using
`PAY_PER_REQUEST` for unpredictable workloads.

Type: String

Valid Values: `PROVISIONED | PAY_PER_REQUEST`

Required: No

**ItemCount**

Number of items in the table. Note that this is an approximate value.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**OnDemandThroughput**

Sets the maximum number of read and write units for the specified on-demand table. If
you use this parameter, you must specify `MaxReadRequestUnits`,
`MaxWriteRequestUnits`, or both.

Type: [OnDemandThroughput](api-ondemandthroughput.md) object

Required: No

**TableArn**

ARN of the table for which backup was created.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**TableSizeBytes**

Size of the table in bytes. Note that this is an approximate value.

Type: Long

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/sourcetabledetails.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/sourcetabledetails.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/sourcetabledetails.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3BucketSource

SourceTableFeatureDetails

All content copied from https://docs.aws.amazon.com/.
