# Task

Details on a task in a cluster.

## Contents

**attachments**

The Elastic Network Adapter that's associated with the task if the task uses the
`awsvpc` network mode.

Type: Array of [Attachment](api-attachment.md) objects

Required: No

**attributes**

The attributes of the task

Type: Array of [Attribute](api-attribute.md) objects

Required: No

**availabilityZone**

The Availability Zone for the task.

Type: String

Required: No

**capacityProviderName**

The capacity provider that's associated with the task.

Type: String

Required: No

**clusterArn**

The ARN of the cluster that hosts the task.

Type: String

Required: No

**connectivity**

The connectivity status of a task.

Type: String

Valid Values: `CONNECTED | DISCONNECTED`

Required: No

**connectivityAt**

The Unix timestamp for the time when the task last went into `CONNECTED`
status.

Type: Timestamp

Required: No

**containerInstanceArn**

The ARN of the container instances that host the task.

Type: String

Required: No

**containers**

The containers that's associated with the task.

Type: Array of [Container](api-container.md) objects

Required: No

**cpu**

The number of CPU units used by the task as expressed in a task definition. It can be
expressed as an integer using CPU units (for example, `1024`). It can also be
expressed as a string using vCPUs (for example, `1 vCPU` or `1
				vcpu`). String values are converted to an integer that indicates the CPU units
when the task definition is registered.

If you're using the EC2 launch type or the external launch type, this field is
optional. Supported values are between `128` CPU units ( `0.125`
vCPUs) and `196608` CPU units ( `192` vCPUs). If you do not specify
a value, the parameter is ignored.

This field is required for Fargate. For information about the valid values, see [Task\
size](../../../../services/amazonecs/latest/developerguide/task-definition-parameters.md#task_size) in the _Amazon Elastic Container Service Developer_
_Guide_.

Type: String

Required: No

**createdAt**

The Unix timestamp for the time when the task was created. More specifically, it's for
the time when the task entered the `PENDING` state.

Type: Timestamp

Required: No

**desiredStatus**

The desired status of the task. For more information, see [Task\
Lifecycle](../../../../services/amazonecs/latest/developerguide/task-lifecycle.md).

Type: String

Required: No

**enableExecuteCommand**

Determines whether execute command functionality is turned on for this task. If
`true`, execute command functionality is turned on all the containers in
the task.

Type: Boolean

Required: No

**ephemeralStorage**

The ephemeral storage settings for the task.

Type: [EphemeralStorage](api-ephemeralstorage.md) object

Required: No

**executionStoppedAt**

The Unix timestamp for the time when the task execution stopped.

Type: Timestamp

Required: No

**fargateEphemeralStorage**

The Fargate ephemeral storage settings for the task.

Type: [TaskEphemeralStorage](api-taskephemeralstorage.md) object

Required: No

**group**

The name of the task group that's associated with the task.

Type: String

Required: No

**healthStatus**

The health status for the task. It's determined by the health of the essential
containers in the task. If all essential containers in the task are reporting as
`HEALTHY`, the task status also reports as `HEALTHY`. If any
essential containers in the task are reporting as `UNHEALTHY` or
`UNKNOWN`, the task status also reports as `UNHEALTHY` or
`UNKNOWN`.

###### Note

The Amazon ECS container agent doesn't monitor or report on Docker health checks
that are embedded in a container image and not specified in the container
definition. For example, this includes those specified in a parent image or from the
image's Dockerfile. Health check parameters that are specified in a container
definition override any Docker health checks that are found in the container
image.

Type: String

Valid Values: `HEALTHY | UNHEALTHY | UNKNOWN`

Required: No

**inferenceAccelerators**

The Elastic Inference accelerator that's associated with the task.

Type: Array of [InferenceAccelerator](api-inferenceaccelerator.md) objects

Required: No

**lastStatus**

The last known status for the task. For more information, see [Task\
Lifecycle](../../../../services/amazonecs/latest/developerguide/task-lifecycle.md).

Type: String

Required: No

**launchType**

The infrastructure where your task runs on. For more information, see [Amazon\
ECS launch types](../../../../services/amazonecs/latest/developerguide/launch-types.md) in the _Amazon Elastic Container Service Developer_
_Guide_.

Type: String

Valid Values: `EC2 | FARGATE | EXTERNAL | MANAGED_INSTANCES`

Required: No

**memory**

The amount of memory (in MiB) that the task uses as expressed in a task definition. It
can be expressed as an integer using MiB (for example, `1024`). If it's
expressed as a string using GB (for example, `1GB` or `1 GB`),
it's converted to an integer indicating the MiB when the task definition is
registered.

If you use the EC2 launch type, this field is optional.

If you use the Fargate launch type, this field is required. You must use one of the
following values. The value that you choose determines the range of supported values for
the `cpu` parameter.

- 512 (0.5 GB), 1024 (1 GB), 2048 (2 GB) - Available `cpu` values:
256 (.25 vCPU)

- 1024 (1 GB), 2048 (2 GB), 3072 (3 GB), 4096 (4 GB) - Available
`cpu` values: 512 (.5 vCPU)

- 2048 (2 GB), 3072 (3 GB), 4096 (4 GB), 5120 (5 GB), 6144 (6 GB), 7168 (7 GB),
8192 (8 GB) - Available `cpu` values: 1024 (1 vCPU)

- Between 4096 (4 GB) and 16384 (16 GB) in increments of 1024 (1 GB) - Available
`cpu` values: 2048 (2 vCPU)

- Between 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB) - Available
`cpu` values: 4096 (4 vCPU)

- Between 16 GB and 60 GB in 4 GB increments - Available `cpu`
values: 8192 (8 vCPU)

This option requires Linux platform `1.4.0` or later.

- Between 32GB and 120 GB in 8 GB increments - Available `cpu`
values: 16384 (16 vCPU)

This option requires Linux platform `1.4.0` or later.

Type: String

Required: No

**overrides**

One or more container overrides.

Type: [TaskOverride](api-taskoverride.md) object

Required: No

**platformFamily**

The operating system that your tasks are running on. A platform family is specified
only for tasks that use the Fargate launch type.

All tasks that run as part of this service must use the same
`platformFamily` value as the service (for example,
`LINUX.`).

Type: String

Required: No

**platformVersion**

The platform version where your task runs on. A platform version is only specified for
tasks that use the Fargate launch type. If you didn't specify one, the
`LATEST` platform version is used. For more information, see [AWS Fargate Platform Versions](../../../../services/amazonecs/latest/developerguide/platform-versions.md) in the _Amazon Elastic_
_Container Service Developer Guide_.

Type: String

Required: No

**pullStartedAt**

The Unix timestamp for the time when the container image pull began.

Type: Timestamp

Required: No

**pullStoppedAt**

The Unix timestamp for the time when the container image pull completed.

Type: Timestamp

Required: No

**startedAt**

The Unix timestamp for the time when the task started. More specifically, it's for the
time when the task transitioned from the `PENDING` state to the
`RUNNING` state.

Type: Timestamp

Required: No

**startedBy**

The tag specified when a task is started. If an Amazon ECS service started the task,
the `startedBy` parameter contains the deployment ID of that service.

Type: String

Required: No

**stopCode**

The stop code indicating why a task was stopped. The `stoppedReason` might
contain additional details.

For more information about stop code, see [Stopped tasks\
error codes](../../../../services/amazonecs/latest/developerguide/stopped-task-error-codes.md) in the _Amazon ECS Developer Guide_.

Type: String

Valid Values: `TaskFailedToStart | EssentialContainerExited | UserInitiated | ServiceSchedulerInitiated | SpotInterruption | TerminationNotice`

Required: No

**stoppedAt**

The Unix timestamp for the time when the task was stopped. More specifically, it's for
the time when the task transitioned from the `RUNNING` state to the
`STOPPED` state.

Type: Timestamp

Required: No

**stoppedReason**

The reason that the task was stopped.

Type: String

Required: No

**stoppingAt**

The Unix timestamp for the time when the task stops. More specifically, it's for the
time when the task transitions from the `RUNNING` state to
`STOPPING`.

Type: Timestamp

Required: No

**tags**

The metadata that you apply to the task to help you categorize and organize the task.
Each tag consists of a key and an optional value. You define both the key and
value.

The following basic restrictions apply to tags:

- Maximum number of tags per resource - 50

- For each resource, each tag key must be unique, and each tag key can have only
one value.

- Maximum key length - 128 Unicode characters in UTF-8

- Maximum value length - 256 Unicode characters in UTF-8

- If your tagging schema is used across multiple services and resources,
remember that other services may have restrictions on allowed characters.
Generally allowed characters are: letters, numbers, and spaces representable in
UTF-8, and the following characters: + - = . \_ : / @.

- Tag keys and values are case-sensitive.

- Do not use `aws:`, `AWS:`, or any upper or lowercase
combination of such as a prefix for either keys or values as it is reserved for
AWS use. You cannot edit or delete tag keys or values with
this prefix. Tags with this prefix do not count against your tags per resource
limit.

Type: Array of [Tag](api-tag.md) objects

Array Members: Minimum number of 0 items. Maximum number of 50 items.

Required: No

**taskArn**

The Amazon Resource Name (ARN) of the task.

Type: String

Required: No

**taskDefinitionArn**

The ARN of the task definition that creates the task.

Type: String

Required: No

**version**

The version counter for the task. Every time a task experiences a change that starts a
CloudWatch event, the version counter is incremented. If you replicate your Amazon ECS
task state with CloudWatch Events, you can compare the version of a task reported by the
Amazon ECS API actions with the version reported in CloudWatch Events for the task
(inside the `detail` object) to verify that the version in your event stream
is current.

Type: Long

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecs-2014-11-13/task.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecs-2014-11-13/task.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecs-2014-11-13/task.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Tag

TaskDefinition
