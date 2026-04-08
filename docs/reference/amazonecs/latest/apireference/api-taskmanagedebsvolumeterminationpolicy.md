# TaskManagedEBSVolumeTerminationPolicy

The termination policy for the Amazon EBS volume when the task exits. For more
information, see [Amazon ECS volume termination policy](../../../../services/amazonecs/latest/developerguide/ebs-volumes.md#ebs-volume-types).

## Contents

**deleteOnTermination**

Indicates whether the volume should be deleted on when the task stops. If a value of
`true` is specified, Amazon ECS deletes the Amazon EBS volume on
your behalf when the task goes into the `STOPPED` state. If no value is
specified, the default value is `true` is used. When set to
`false`, Amazon ECS leaves the volume in your account.

Type: Boolean

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecs-2014-11-13/taskmanagedebsvolumeterminationpolicy.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecs-2014-11-13/taskmanagedebsvolumeterminationpolicy.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecs-2014-11-13/taskmanagedebsvolumeterminationpolicy.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TaskManagedEBSVolumeConfiguration

TaskOverride
