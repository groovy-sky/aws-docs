# TaskEphemeralStorage

The amount of ephemeral storage to allocate for the task.

## Contents

**kmsKeyId**

Specify an AWS Key Management Service key ID to encrypt the ephemeral storage for the
task.

Type: String

Required: No

**sizeInGiB**

The total amount, in GiB, of the ephemeral storage to set for the task. The minimum
supported value is `20` GiB and the maximum supported value is
`200` GiB.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecs-2014-11-13/taskephemeralstorage.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecs-2014-11-13/taskephemeralstorage.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecs-2014-11-13/taskephemeralstorage.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TaskDefinitionPlacementConstraint

TaskManagedEBSVolumeConfiguration
