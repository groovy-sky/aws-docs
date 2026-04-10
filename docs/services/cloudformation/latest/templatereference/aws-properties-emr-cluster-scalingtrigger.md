This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::Cluster ScalingTrigger

`ScalingTrigger` is a subproperty of the `ScalingRule` property type. `ScalingTrigger` determines the conditions that trigger an automatic scaling activity.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudWatchAlarmDefinition" : CloudWatchAlarmDefinition
}

```

### YAML

```yaml

  CloudWatchAlarmDefinition:
    CloudWatchAlarmDefinition

```

## Properties

`CloudWatchAlarmDefinition`

The definition of a CloudWatch metric alarm. When the defined alarm conditions are met
along with other trigger parameters, scaling activity begins.

_Required_: Yes

_Type_: [CloudWatchAlarmDefinition](aws-properties-emr-cluster-cloudwatchalarmdefinition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScalingRule

ScriptBootstrapActionConfig

All content copied from https://docs.aws.amazon.com/.
