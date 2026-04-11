This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::Flow FrozenFrames

Configures settings for the `FrozenFrames` metric.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "State" : String,
  "ThresholdSeconds" : Integer
}

```

### YAML

```yaml

  State: String
  ThresholdSeconds: Integer

```

## Properties

`State`

Indicates whether the `FrozenFrames` metric is enabled or disabled.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThresholdSeconds`

Specifies the number of consecutive seconds of a static image that triggers an event or alert.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Fmtp

GatewayBridgeSource

All content copied from https://docs.aws.amazon.com/.
