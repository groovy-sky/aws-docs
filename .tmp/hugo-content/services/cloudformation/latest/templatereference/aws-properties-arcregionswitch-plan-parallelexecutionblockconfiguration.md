This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ARCRegionSwitch::Plan ParallelExecutionBlockConfiguration

Configuration for steps that should be executed in parallel during a Region switch.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Steps" : [ Step, ... ]
}

```

### YAML

```yaml

  Steps:
    - Step

```

## Properties

`Steps`

The steps for a parallel execution block.

_Required_: Yes

_Type_: Array of [Step](aws-properties-arcregionswitch-plan-step.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LambdaUngraceful

RdsCreateCrossRegionReplicaConfiguration

All content copied from https://docs.aws.amazon.com/.
