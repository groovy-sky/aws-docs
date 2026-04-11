This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::Standard StandardsControl

Provides details about an individual security control. For a list of Security Hub CSPM controls, see [Security Hub CSPM controls reference](../../../securityhub/latest/userguide/securityhub-controls-reference.md) in the _AWS Security Hub CSPM User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Reason" : String,
  "StandardsControlArn" : String
}

```

### YAML

```yaml

  Reason: String
  StandardsControlArn: String

```

## Properties

`Reason`

A user-defined reason for changing a control's enablement status in a
specified standard. If you are disabling a control, then this property is required.

_Required_: No

_Type_: String

_Pattern_: `.*\S.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StandardsControlArn`

The Amazon Resource Name (ARN) of the control.

_Required_: Yes

_Type_: String

_Pattern_: `arn:aws\S*:securityhub:\S*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SecurityHub::Standard

Next

All content copied from https://docs.aws.amazon.com/.
