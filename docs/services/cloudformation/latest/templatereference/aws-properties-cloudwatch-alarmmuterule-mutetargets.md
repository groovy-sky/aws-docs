This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudWatch::AlarmMuteRule MuteTargets

Specifies which alarms this rule applies to.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AlarmNames" : [ String, ... ]
}

```

### YAML

```yaml

  AlarmNames:
    - String

```

## Properties

`AlarmNames`

Specifies which alarms this rule applies to.

_Required_: Yes

_Type_: Array of String

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::CloudWatch::AlarmMuteRule

Rule
