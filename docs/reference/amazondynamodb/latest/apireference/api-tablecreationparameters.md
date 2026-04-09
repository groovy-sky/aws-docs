# TableCreationParameters

The parameters for the table created as part of the import operation.

## Contents

###### Note

In the following list, the required parameters are described first.

**AttributeDefinitions**

The attributes of the table created as part of the import operation.

Type: Array of [AttributeDefinition](api-attributedefinition.md) objects

Required: Yes

**KeySchema**

The primary key and option sort key of the table created as part of the import
operation.

Type: Array of [KeySchemaElement](api-keyschemaelement.md) objects

Array Members: Minimum number of 1 item.

Required: Yes

**TableName**

The name of the table created as part of the import operation.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

Pattern: `[a-zA-Z0-9_.-]+`

Required: Yes

**BillingMode**

The billing mode for provisioning the table created as part of the import operation.

Type: String

Valid Values: `PROVISIONED | PAY_PER_REQUEST`

Required: No

**GlobalSecondaryIndexes**

The Global Secondary Indexes (GSI) of the table to be created as part of the import
operation.

Type: Array of [GlobalSecondaryIndex](api-globalsecondaryindex.md) objects

Required: No

**OnDemandThroughput**

Sets the maximum number of read and write units for the specified on-demand table. If
you use this parameter, you must specify `MaxReadRequestUnits`,
`MaxWriteRequestUnits`, or both.

Type: [OnDemandThroughput](api-ondemandthroughput.md) object

Required: No

**ProvisionedThroughput**

Represents the provisioned throughput settings for the specified global secondary
index. You must use `ProvisionedThroughput` or
`OnDemandThroughput` based on your table’s capacity mode.

For current minimum and maximum provisioned throughput values, see [Service,\
Account, and Table Quotas](../../../../services/dynamodb/latest/developerguide/limits.md) in the _Amazon DynamoDB Developer_
_Guide_.

Type: [ProvisionedThroughput](api-provisionedthroughput.md) object

Required: No

**SSESpecification**

Represents the settings used to enable server-side encryption.

Type: [SSESpecification](api-ssespecification.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/tablecreationparameters.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/tablecreationparameters.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/tablecreationparameters.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TableClassSummary

TableDescription

All content copied from https://docs.aws.amazon.com/.
