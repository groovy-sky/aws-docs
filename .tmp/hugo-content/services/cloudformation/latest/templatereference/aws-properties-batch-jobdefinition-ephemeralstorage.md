This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition EphemeralStorage

The amount of ephemeral storage to allocate for the task. This parameter is used to expand
the total amount of ephemeral storage available, beyond the default amount, for tasks hosted on
AWS Fargate.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SizeInGiB" : Integer
}

```

### YAML

```yaml

  SizeInGiB: Integer

```

## Properties

`SizeInGiB`

The total amount, in GiB, of ephemeral storage to set for the task. The minimum supported
value is `21` GiB and the maximum supported value is `200` GiB.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Environment

EvaluateOnExit

All content copied from https://docs.aws.amazon.com/.
