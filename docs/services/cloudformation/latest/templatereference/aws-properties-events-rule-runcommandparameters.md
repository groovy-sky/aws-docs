This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::Rule RunCommandParameters

This parameter contains the criteria (either InstanceIds or a tag) used to specify which
EC2 instances are to be sent the command.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RunCommandTargets" : [ RunCommandTarget, ... ]
}

```

### YAML

```yaml

  RunCommandTargets:
    - RunCommandTarget

```

## Properties

`RunCommandTargets`

Currently, we support including only one RunCommandTarget block, which specifies either an
array of InstanceIds or a tag.

_Required_: Yes

_Type_: Array of [RunCommandTarget](aws-properties-events-rule-runcommandtarget.md)

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RetryPolicy

RunCommandTarget

All content copied from https://docs.aws.amazon.com/.
