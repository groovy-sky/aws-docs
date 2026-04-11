This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe EcsContainerOverride

The overrides that are sent to a container. An empty container override can be passed
in. An example of an empty container override is `{"containerOverrides": [ ] }`.
If a non-empty container override is specified, the `name` parameter must be
included.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Command" : [ String, ... ],
  "Cpu" : Integer,
  "Environment" : [ EcsEnvironmentVariable, ... ],
  "EnvironmentFiles" : [ EcsEnvironmentFile, ... ],
  "Memory" : Integer,
  "MemoryReservation" : Integer,
  "Name" : String,
  "ResourceRequirements" : [ EcsResourceRequirement, ... ]
}

```

### YAML

```yaml

  Command:
    - String
  Cpu: Integer
  Environment:
    - EcsEnvironmentVariable
  EnvironmentFiles:
    - EcsEnvironmentFile
  Memory: Integer
  MemoryReservation: Integer
  Name: String
  ResourceRequirements:
    - EcsResourceRequirement

```

## Properties

`Command`

The command to send to the container that overrides the default command from the Docker
image or the task definition. You must also specify a container name.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Cpu`

The number of `cpu` units reserved for the container, instead of the default
value from the task definition. You must also specify a container name.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Environment`

The environment variables to send to the container. You can add new environment
variables, which are added to the container at launch, or you can override the existing
environment variables from the Docker image or the task definition. You must also specify a
container name.

_Required_: No

_Type_: Array of [EcsEnvironmentVariable](aws-properties-pipes-pipe-ecsenvironmentvariable.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnvironmentFiles`

A list of files containing the environment variables to pass to a container, instead of
the value from the container definition.

_Required_: No

_Type_: Array of [EcsEnvironmentFile](aws-properties-pipes-pipe-ecsenvironmentfile.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Memory`

The hard limit (in MiB) of memory to present to the container, instead of the default
value from the task definition. If your container attempts to exceed the memory specified
here, the container is killed. You must also specify a container name.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MemoryReservation`

The soft limit (in MiB) of memory to reserve for the container, instead of the default
value from the task definition. You must also specify a container name.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the container that receives the override. This parameter is required if any
override is specified.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceRequirements`

The type and amount of a resource to assign to a container, instead of the default value
from the task definition. The only supported resource is a GPU.

_Required_: No

_Type_: Array of [EcsResourceRequirement](aws-properties-pipes-pipe-ecsresourcerequirement.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DimensionMapping

EcsEnvironmentFile

All content copied from https://docs.aws.amazon.com/.
