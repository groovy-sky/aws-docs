---
title: "ReplicaGlobalSecondaryIndexAutoScalingDescription"
---

# ReplicaGlobalSecondaryIndexAutoScalingDescription

Represents the auto scaling configuration for a replica global secondary index.

## Contents

###### Note

In the following list, the required parameters are described first.

**IndexName**

The name of the global secondary index.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

Pattern: `[a-zA-Z0-9_.-]+`

Required: No

**IndexStatus**

The current state of the replica global secondary index:

- `CREATING` \- The index is being created.

- `UPDATING` \- The table/index configuration is being updated. The
table/index remains available for data operations when
`UPDATING`

- `DELETING` \- The index is being deleted.

- `ACTIVE` \- The index is ready for use.

Type: String

Valid Values: `CREATING | UPDATING | DELETING | ACTIVE`

Required: No

**ProvisionedReadCapacityAutoScalingSettings**

Represents the auto scaling settings for a global table or global secondary
index.

Type: [AutoScalingSettingsDescription](api-autoscalingsettingsdescription.md) object

Required: No

**ProvisionedWriteCapacityAutoScalingSettings**

Represents the auto scaling settings for a global table or global secondary
index.

Type: [AutoScalingSettingsDescription](api-autoscalingsettingsdescription.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/dynamodb-2012-08-10/ReplicaGlobalSecondaryIndexAutoScalingDescription)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/dynamodb-2012-08-10/ReplicaGlobalSecondaryIndexAutoScalingDescription)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/dynamodb-2012-08-10/ReplicaGlobalSecondaryIndexAutoScalingDescription)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReplicaGlobalSecondaryIndex

ReplicaGlobalSecondaryIndexAutoScalingUpdate

All content copied from https://docs.aws.amazon.com/.
