# UpdateGlobalSecondaryIndexAction

Represents the new provisioned throughput settings to be applied to a global secondary
index.

## Contents

###### Note

In the following list, the required parameters are described first.

**IndexName**

The name of the global secondary index to be updated.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

Pattern: `[a-zA-Z0-9_.-]+`

Required: Yes

**OnDemandThroughput**

Updates the maximum number of read and write units for the specified global secondary
index. If you use this parameter, you must specify `MaxReadRequestUnits`,
`MaxWriteRequestUnits`, or both.

Type: [OnDemandThroughput](api-ondemandthroughput.md) object

Required: No

**ProvisionedThroughput**

Represents the provisioned throughput settings for the specified global secondary
index.

For current minimum and maximum provisioned throughput values, see [Service,\
Account, and Table Quotas](../../../../services/dynamodb/latest/developerguide/limits.md) in the _Amazon DynamoDB Developer_
_Guide_.

Type: [ProvisionedThroughput](api-provisionedthroughput.md) object

Required: No

**WarmThroughput**

Represents the warm throughput value of the new provisioned throughput settings to be
applied to a global secondary index.

Type: [WarmThroughput](api-warmthroughput.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/updateglobalsecondaryindexaction.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/updateglobalsecondaryindexaction.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/updateglobalsecondaryindexaction.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Update

UpdateKinesisStreamingConfiguration

All content copied from https://docs.aws.amazon.com/.
