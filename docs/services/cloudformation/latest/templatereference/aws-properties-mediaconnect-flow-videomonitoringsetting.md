This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::Flow VideoMonitoringSetting

Specifies the configuration for video stream metrics monitoring.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BlackFrames" : BlackFrames,
  "FrozenFrames" : FrozenFrames
}

```

### YAML

```yaml

  BlackFrames:
    BlackFrames
  FrozenFrames:
    FrozenFrames

```

## Properties

`BlackFrames`

Detects video frames that are black.

_Required_: No

_Type_: [BlackFrames](aws-properties-mediaconnect-flow-blackframes.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FrozenFrames`

Detects video frames that have not changed.

_Required_: No

_Type_: [FrozenFrames](aws-properties-mediaconnect-flow-frozenframes.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

VpcInterface

All content copied from https://docs.aws.amazon.com/.
