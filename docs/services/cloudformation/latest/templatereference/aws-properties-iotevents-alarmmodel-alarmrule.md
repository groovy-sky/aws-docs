This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTEvents::AlarmModel AlarmRule

Defines when your alarm is invoked.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SimpleRule" : SimpleRule
}

```

### YAML

```yaml

  SimpleRule:
    SimpleRule

```

## Properties

`SimpleRule`

A rule that compares an input property value to a threshold value with a comparison operator.

_Required_: No

_Type_: [SimpleRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotevents-alarmmodel-simplerule.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AlarmEventActions

AssetPropertyTimestamp
