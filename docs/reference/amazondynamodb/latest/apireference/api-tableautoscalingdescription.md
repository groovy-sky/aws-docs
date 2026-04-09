# TableAutoScalingDescription

Represents the auto scaling configuration for a global table.

## Contents

###### Note

In the following list, the required parameters are described first.

**Replicas**

Represents replicas of the global table.

Type: Array of [ReplicaAutoScalingDescription](api-replicaautoscalingdescription.md) objects

Required: No

**TableName**

The name of the table.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

Pattern: `[a-zA-Z0-9_.-]+`

Required: No

**TableStatus**

The current state of the table:

- `CREATING` \- The table is being created.

- `UPDATING` \- The table is being updated.

- `DELETING` \- The table is being deleted.

- `ACTIVE` \- The table is ready for use.

Type: String

Valid Values: `CREATING | UPDATING | DELETING | ACTIVE | INACCESSIBLE_ENCRYPTION_CREDENTIALS | ARCHIVING | ARCHIVED | REPLICATION_NOT_AUTHORIZED`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/tableautoscalingdescription.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/tableautoscalingdescription.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/tableautoscalingdescription.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StreamSpecification

TableClassSummary

All content copied from https://docs.aws.amazon.com/.
