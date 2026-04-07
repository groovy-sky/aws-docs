This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::ProjectProfile ResourceTagParameter

The resource tag parameter of the project profile.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IsValueEditable" : Boolean,
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  IsValueEditable: Boolean
  Key: String
  Value: String

```

## Properties

`IsValueEditable`

Specifies whether the value of the resource tag parameter of the project profile is
editable at the project level.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Key`

The key of the resource tag parameter of the project profile.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w \.:/=+@-]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value of the resource tag parameter key of the project profile.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w \.:/=+@-]*$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Region

AWS::DataZone::SubscriptionTarget
