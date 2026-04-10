This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition TaskContainerDependency

A list of containers that this task depends on.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Condition" : String,
  "ContainerName" : String
}

```

### YAML

```yaml

  Condition: String
  ContainerName: String

```

## Properties

`Condition`

The dependency condition of the container. The following are the available conditions and
their behavior:

- `START` \- This condition emulates the behavior of links and volumes today. It
validates that a dependent container is started before permitting other containers to start.

- `COMPLETE` \- This condition validates that a dependent container runs to
completion (exits) before permitting other containers to start. This can be useful for
nonessential containers that run a script and then exit. This condition can't be set on an
essential container.

- `SUCCESS` \- This condition is the same as `COMPLETE`, but it also
requires that the container exits with a zero status. This condition can't be set on an
essential container.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContainerName`

A unique identifier for the container.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Secret

TaskContainerProperties

All content copied from https://docs.aws.amazon.com/.
