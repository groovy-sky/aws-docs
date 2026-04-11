# ReplaceRootVolumeTask

Information about a root volume replacement task.

## Contents

**completeTime**

The time the task completed.

Type: String

Required: No

**deleteReplacedRootVolume**

Indicates whether the original root volume is to be deleted after the root volume
replacement task completes.

Type: Boolean

Required: No

**imageId**

The ID of the AMI used to create the replacement root volume.

Type: String

Required: No

**instanceId**

The ID of the instance for which the root volume replacement task was created.

Type: String

Required: No

**replaceRootVolumeTaskId**

The ID of the root volume replacement task.

Type: String

Required: No

**snapshotId**

The ID of the snapshot used to create the replacement root volume.

Type: String

Required: No

**startTime**

The time the task was started.

Type: String

Required: No

**TagSet.N**

The tags assigned to the task.

Type: Array of [Tag](api-tag.md) objects

Required: No

**taskState**

The state of the task. The task can be in one of the following states:

- `pending` \- the replacement volume is being created.

- `in-progress` \- the original volume is being detached and the
replacement volume is being attached.

- `succeeded` \- the replacement volume has been successfully attached
to the instance and the instance is available.

- `failing` \- the replacement task is in the process of failing.

- `failed` \- the replacement task has failed but the original root
volume is still attached.

- `failing-detached` \- the replacement task is in the process of failing.
The instance might have no root volume attached.

- `failed-detached` \- the replacement task has failed and the instance
has no root volume attached.

Type: String

Valid Values: `pending | in-progress | failing | succeeded | failed | failed-detached`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/replacerootvolumetask.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/replacerootvolumetask.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/replacerootvolumetask.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

RemovePrefixListEntry

RequestFilterPortRange
