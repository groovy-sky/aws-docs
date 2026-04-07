# BundleTask

Describes a bundle task.

## Contents

**bundleId**

The ID of the bundle task.

Type: String

Required: No

**error**

If the task fails, a description of the error.

Type: [BundleTaskError](api-bundletaskerror.md) object

Required: No

**instanceId**

The ID of the instance associated with this bundle task.

Type: String

Required: No

**progress**

The level of task completion, as a percent (for example, 20%).

Type: String

Required: No

**startTime**

The time this task started.

Type: Timestamp

Required: No

**state**

The state of the task.

Type: String

Valid Values: `pending | waiting-for-shutdown | bundling | storing | cancelling | complete | failed`

Required: No

**storage**

The Amazon S3 storage locations.

Type: [Storage](api-storage.md) object

Required: No

**updateTime**

The time of the most recent update for the task.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/bundletask.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/bundletask.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/bundletask.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

BlockPublicAccessStates

BundleTaskError
