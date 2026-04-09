# ReplicaGlobalSecondaryIndexAutoScalingUpdate

Represents the auto scaling settings of a global secondary index for a replica that
will be modified.

## Contents

###### Note

In the following list, the required parameters are described first.

**IndexName**

The name of the global secondary index.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

Pattern: `[a-zA-Z0-9_.-]+`

Required: No

**ProvisionedReadCapacityAutoScalingUpdate**

Represents the auto scaling settings to be modified for a global table or global
secondary index.

Type: [AutoScalingSettingsUpdate](api-autoscalingsettingsupdate.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/replicaglobalsecondaryindexautoscalingupdate.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/replicaglobalsecondaryindexautoscalingupdate.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/replicaglobalsecondaryindexautoscalingupdate.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReplicaGlobalSecondaryIndexAutoScalingDescription

ReplicaGlobalSecondaryIndexDescription

All content copied from https://docs.aws.amazon.com/.
