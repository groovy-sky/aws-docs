This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe BatchContainerOverrides

The overrides that are sent to a container.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Command" : [ String, ... ],
  "Environment" : [ BatchEnvironmentVariable, ... ],
  "InstanceType" : String,
  "ResourceRequirements" : [ BatchResourceRequirement, ... ]
}

```

### YAML

```yaml

  Command:
    - String
  Environment:
    - BatchEnvironmentVariable
  InstanceType: String
  ResourceRequirements:
    - BatchResourceRequirement

```

## Properties

`Command`

The command to send to the container that overrides the default command from the Docker
image or the task definition.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Environment`

The environment variables to send to the container. You can add new environment
variables, which are added to the container at launch, or you can override the existing
environment variables from the Docker image or the task definition.

###### Note

Environment variables cannot start with " `
                                AWS Batch
                            `". This
naming convention is reserved for variables that AWS Batch sets.

_Required_: No

_Type_: Array of [BatchEnvironmentVariable](aws-properties-pipes-pipe-batchenvironmentvariable.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceType`

The instance type to use for a multi-node parallel job.

###### Note

This parameter isn't applicable to single-node container jobs or jobs that run on
Fargate resources, and shouldn't be provided.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceRequirements`

The type and amount of resources to assign to a container. This overrides the settings
in the job definition. The supported resources include `GPU`,
`MEMORY`, and `VCPU`.

_Required_: No

_Type_: Array of [BatchResourceRequirement](aws-properties-pipes-pipe-batchresourcerequirement.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BatchArrayProperties

BatchEnvironmentVariable

All content copied from https://docs.aws.amazon.com/.
