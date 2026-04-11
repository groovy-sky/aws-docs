This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition EcsProperties

An object that contains the properties for the Amazon ECS resources of a job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TaskProperties" : [ EcsTaskProperties, ... ]
}

```

### YAML

```yaml

  TaskProperties:
    - EcsTaskProperties

```

## Properties

`TaskProperties`

An object that contains the properties for the Amazon ECS task definition of a job.

###### Note

This object is currently limited to one task element. However, the task element can run up to 10 containers.

_Required_: Yes

_Type_: Array of [EcsTaskProperties](aws-properties-batch-jobdefinition-ecstaskproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Device

EcsTaskProperties

All content copied from https://docs.aws.amazon.com/.
