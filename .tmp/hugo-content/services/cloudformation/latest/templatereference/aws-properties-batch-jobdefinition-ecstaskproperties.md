This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition EcsTaskProperties

The properties for a task definition that describes the container and volume definitions of
an Amazon ECS task. You can specify which Docker images to use, the required resources, and other
configurations related to launching the task definition through an Amazon ECS service or task.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Containers" : [ TaskContainerProperties, ... ],
  "EnableExecuteCommand" : Boolean,
  "EphemeralStorage" : EphemeralStorage,
  "ExecutionRoleArn" : String,
  "IpcMode" : String,
  "NetworkConfiguration" : NetworkConfiguration,
  "PidMode" : String,
  "PlatformVersion" : String,
  "RuntimePlatform" : RuntimePlatform,
  "TaskRoleArn" : String,
  "Volumes" : [ Volume, ... ]
}

```

### YAML

```yaml

  Containers:
    - TaskContainerProperties
  EnableExecuteCommand: Boolean
  EphemeralStorage:
    EphemeralStorage
  ExecutionRoleArn: String
  IpcMode: String
  NetworkConfiguration:
    NetworkConfiguration
  PidMode: String
  PlatformVersion: String
  RuntimePlatform:
    RuntimePlatform
  TaskRoleArn: String
  Volumes:
    - Volume

```

## Properties

`Containers`

This object is a list of containers.

_Required_: No

_Type_: Array of [TaskContainerProperties](aws-properties-batch-jobdefinition-taskcontainerproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableExecuteCommand`

Determines whether execute command functionality is turned on for this task. If `true`, execute
command functionality is turned on all the containers in the task.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EphemeralStorage`

The amount of ephemeral storage to allocate for the task. This parameter is used to expand
the total amount of ephemeral storage available, beyond the default amount, for tasks hosted on
AWS Fargate.

_Required_: No

_Type_: [EphemeralStorage](aws-properties-batch-jobdefinition-ephemeralstorage.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExecutionRoleArn`

The Amazon Resource Name (ARN) of the execution role that AWS Batch can assume. For jobs that run on Fargate
resources, you must provide an execution role. For more information, see [AWS Batch execution IAM role](../../../batch/latest/userguide/execution-iam-role.md)
in the _AWS Batch User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IpcMode`

The IPC resource namespace to use for the containers in the task. The valid values are
`host`, `task`, or `none`.

If `host` is specified, all containers within the tasks that specified the
`host` IPC mode on the same container instance share the same IPC resources with the
host Amazon EC2 instance.

If `task` is specified, all containers within the specified `task`
share the same IPC resources.

If `none` is specified, the IPC resources within the containers of a task are
private, and are not shared with other containers in a task or on the container instance.

If no value is specified, then the IPC resource namespace sharing depends on the Docker
daemon setting on the container instance. For more information, see [IPC settings](https://docs.docker.com/engine/reference/run) in
the Docker run reference.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkConfiguration`

The network configuration for jobs that are running on Fargate resources. Jobs that are
running on Amazon EC2 resources must not specify this parameter.

_Required_: No

_Type_: [NetworkConfiguration](aws-properties-batch-jobdefinition-networkconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PidMode`

The process namespace to use for the containers in the task. The valid values are
`host` or `task`. For example, monitoring sidecars might need
`pidMode` to access information about other containers running in the same
task.

If `host` is specified, all containers within the tasks that specified the
`host` PID mode on the same container instance share the process namespace with the
host Amazon EC2 instance.

If `task` is specified, all containers within the specified task share the same
process namespace.

If no value is specified, the default is a private namespace for each container. For more
information, see [PID settings](https://docs.docker.com/engine/reference/run) in the Docker run reference.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PlatformVersion`

The Fargate platform version where the jobs are running. A platform version is specified
only for jobs that are running on Fargate resources. If one isn't specified, the
`LATEST` platform version is used by default. This uses a recent, approved version of
the Fargate platform for compute resources. For more information, see [AWS Fargate\
platform versions](../../../amazonecs/latest/developerguide/platform-versions.md) in the _Amazon Elastic Container Service Developer Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuntimePlatform`

An object that represents the compute environment architecture for AWS Batch jobs on
Fargate.

_Required_: No

_Type_: [RuntimePlatform](aws-properties-batch-jobdefinition-runtimeplatform.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TaskRoleArn`

The Amazon Resource Name (ARN) that's associated with the Amazon ECS task.

###### Note

This is object is comparable to [ContainerProperties:jobRoleArn](../../../../reference/batch/latest/apireference/api-containerproperties.md).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Volumes`

A list of volumes that are associated with the job.

_Required_: No

_Type_: Array of [Volume](aws-properties-batch-jobdefinition-volume.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EcsProperties

EFSAuthorizationConfig

All content copied from https://docs.aws.amazon.com/.
