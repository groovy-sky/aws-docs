This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaTailor::Channel TimeShiftConfiguration

The configuration for time-shifted viewing.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxTimeDelaySeconds" : Number
}

```

### YAML

```yaml

  MaxTimeDelaySeconds: Number

```

## Properties

`MaxTimeDelaySeconds`

The maximum time delay for time-shifted viewing. The minimum allowed maximum time delay is 0 seconds,
and the maximum allowed maximum time delay is 21600 seconds (6 hours).

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::MediaTailor::ChannelPolicy

All content copied from https://docs.aws.amazon.com/.
