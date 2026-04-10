This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RolesAnywhere::TrustAnchor

Creates a TrustAnchor.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::RolesAnywhere::TrustAnchor",
  "Properties" : {
      "Enabled" : Boolean,
      "Name" : String,
      "NotificationSettings" : [ NotificationSetting, ... ],
      "Source" : Source,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::RolesAnywhere::TrustAnchor
Properties:
  Enabled: Boolean
  Name: String
  NotificationSettings:
    - NotificationSetting
  Source:
    Source
  Tags:
    - Tag

```

## Properties

`Enabled`

Indicates whether the trust anchor is enabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the trust anchor.

_Required_: Yes

_Type_: String

_Pattern_: `[ a-zA-Z0-9-_]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotificationSettings`

A list of notification settings to be associated to the trust anchor.

_Required_: No

_Type_: Array of [NotificationSetting](aws-properties-rolesanywhere-trustanchor-notificationsetting.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

The trust anchor type and its related certificate data.

_Required_: Yes

_Type_: [Source](aws-properties-rolesanywhere-trustanchor-source.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to attach to the trust anchor.

_Required_: No

_Type_: Array of [Tag](aws-properties-rolesanywhere-trustanchor-tag.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns `TrustAnchorId`.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`TrustAnchorArn`

The ARN of the trust anchor.

`TrustAnchorId`

The unique identifier of the trust anchor.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

NotificationSetting

All content copied from https://docs.aws.amazon.com/.
