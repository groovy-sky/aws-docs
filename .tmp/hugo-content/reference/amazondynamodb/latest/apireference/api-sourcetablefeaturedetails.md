# SourceTableFeatureDetails

Contains the details of the features enabled on the table when the backup was created.
For example, LSIs, GSIs, streams, TTL.

## Contents

###### Note

In the following list, the required parameters are described first.

**GlobalSecondaryIndexes**

Represents the GSI properties for the table when the backup was created. It includes
the IndexName, KeySchema, Projection, and ProvisionedThroughput for the GSIs on the
table at the time of backup.

Type: Array of [GlobalSecondaryIndexInfo](api-globalsecondaryindexinfo.md) objects

Required: No

**LocalSecondaryIndexes**

Represents the LSI properties for the table when the backup was created. It includes
the IndexName, KeySchema and Projection for the LSIs on the table at the time of backup.

Type: Array of [LocalSecondaryIndexInfo](api-localsecondaryindexinfo.md) objects

Required: No

**SSEDescription**

The description of the server-side encryption status on the table when the backup was
created.

Type: [SSEDescription](api-ssedescription.md) object

Required: No

**StreamDescription**

Stream settings on the table when the backup was created.

Type: [StreamSpecification](api-streamspecification.md) object

Required: No

**TimeToLiveDescription**

Time to Live settings on the table when the backup was created.

Type: [TimeToLiveDescription](api-timetolivedescription.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/sourcetablefeaturedetails.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/sourcetablefeaturedetails.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/sourcetablefeaturedetails.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SourceTableDetails

SSEDescription

All content copied from https://docs.aws.amazon.com/.
