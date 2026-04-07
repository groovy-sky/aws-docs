# TaskVolumeConfiguration

Configuration settings for the task volume that was `configuredAtLaunch`
that weren't set during `RegisterTaskDef`.

## Contents

**name**

The name of the volume. This value must match the volume name from the
`Volume` object in the task definition.

Type: String

Required: Yes

**managedEBSVolume**

The configuration for the Amazon EBS volume that Amazon ECS creates and manages on
your behalf. These settings are used to create each Amazon EBS volume, with one volume
created for each task. The Amazon EBS volumes are visible in your account in the Amazon
EC2 console once they are created.

Type: [TaskManagedEBSVolumeConfiguration](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_TaskManagedEBSVolumeConfiguration.html) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ecs-2014-11-13/TaskVolumeConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ecs-2014-11-13/TaskVolumeConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ecs-2014-11-13/TaskVolumeConfiguration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TaskSet

TimeoutConfiguration
