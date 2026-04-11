This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppStream::Stack UserSetting

Specifies an action and whether the action is enabled or disabled for users during their streaming sessions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : String,
  "MaximumLength" : Integer,
  "Permission" : String
}

```

### YAML

```yaml

  Action: String
  MaximumLength: Integer
  Permission: String

```

## Properties

`Action`

The action that is enabled or disabled.

_Required_: Yes

_Type_: String

_Allowed values_: `CLIPBOARD_COPY_FROM_LOCAL_DEVICE | CLIPBOARD_COPY_TO_LOCAL_DEVICE | FILE_UPLOAD | FILE_DOWNLOAD | PRINTING_TO_LOCAL_DEVICE | DOMAIN_PASSWORD_SIGNIN | DOMAIN_SMART_CARD_SIGNIN | AUTO_TIME_ZONE_REDIRECTION`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumLength`

Specifies the number of characters that can be copied by end users from the local device to the remote session, and to the local device from the remote session.

This can be specified only for the `CLIPBOARD_COPY_FROM_LOCAL_DEVICE` and `CLIPBOARD_COPY_TO_LOCAL_DEVICE` actions.

This defaults to 20,971,520 (20 MB) when unspecified and the permission is `ENABLED`. This can't be specified when the permission is `DISABLED`.

The value can be between 1 and 20,971,520 (20 MB).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Permission`

Indicates whether the action is enabled or disabled.

_Required_: Yes

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::AppStream::StackFleetAssociation

All content copied from https://docs.aws.amazon.com/.
