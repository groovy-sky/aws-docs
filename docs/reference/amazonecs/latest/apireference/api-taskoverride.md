# TaskOverride

The overrides that are associated with a task.

## Contents

**containerOverrides**

One or more container overrides that are sent to a task.

Type: Array of [ContainerOverride](api-containeroverride.md) objects

Required: No

**cpu**

The CPU override for the task.

Type: String

Required: No

**ephemeralStorage**

The ephemeral storage setting override for the task.

###### Note

This parameter is only supported for tasks hosted on Fargate that use the
following platform versions:

- Linux platform version `1.4.0` or later.

- Windows platform version `1.0.0` or later.

Type: [EphemeralStorage](api-ephemeralstorage.md) object

Required: No

**executionRoleArn**

The Amazon Resource Name (ARN) of the task execution role override for the task. For
more information, see [Amazon ECS task\
execution IAM role](../../../../services/amazonecs/latest/developerguide/task-execution-iam-role.md) in the _Amazon Elastic Container Service_
_Developer Guide_.

Type: String

Required: No

**inferenceAcceleratorOverrides**

The Elastic Inference accelerator override for the task.

Type: Array of [InferenceAcceleratorOverride](api-inferenceacceleratoroverride.md) objects

Required: No

**memory**

The memory override for the task.

Type: String

Required: No

**taskRoleArn**

The Amazon Resource Name (ARN) of the role that containers in this task can assume.
All containers in this task are granted the permissions that are specified in this role.
For more information, see [IAM Role for Tasks](../../../../services/amazonecs/latest/developerguide/task-iam-roles.md)
in the _Amazon Elastic Container Service Developer Guide_.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecs-2014-11-13/taskoverride.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecs-2014-11-13/taskoverride.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecs-2014-11-13/taskoverride.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TaskManagedEBSVolumeTerminationPolicy

TaskSet
