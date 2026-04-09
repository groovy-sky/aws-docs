# GlobalTableDescription

Contains details about the global table.

## Contents

###### Note

In the following list, the required parameters are described first.

**CreationDateTime**

The creation time of the global table.

Type: Timestamp

Required: No

**GlobalTableArn**

The unique identifier of the global table.

Type: String

Required: No

**GlobalTableName**

The global table name.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

Pattern: `[a-zA-Z0-9_.-]+`

Required: No

**GlobalTableStatus**

The current state of the global table:

- `CREATING` \- The global table is being created.

- `UPDATING` \- The global table is being updated.

- `DELETING` \- The global table is being deleted.

- `ACTIVE` \- The global table is ready for use.

Type: String

Valid Values: `CREATING | ACTIVE | DELETING | UPDATING`

Required: No

**ReplicationGroup**

The Regions where the global table has replicas.

Type: Array of [ReplicaDescription](api-replicadescription.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/globaltabledescription.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/globaltabledescription.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/globaltabledescription.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GlobalTable

GlobalTableGlobalSecondaryIndexSettingsUpdate

All content copied from https://docs.aws.amazon.com/.
