This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::User UserPhoneConfig

Contains information about the phone configuration settings for a user.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AfterContactWorkTimeLimit" : Integer,
  "AutoAccept" : Boolean,
  "DeskPhoneNumber" : String,
  "PersistentConnection" : Boolean,
  "PhoneType" : String
}

```

### YAML

```yaml

  AfterContactWorkTimeLimit: Integer
  AutoAccept: Boolean
  DeskPhoneNumber: String
  PersistentConnection: Boolean
  PhoneType: String

```

## Properties

`AfterContactWorkTimeLimit`

The After Call Work (ACW) timeout setting, in seconds. This parameter has a minimum value of 0 and a maximum
value of 2,000,000 seconds (24 days). Enter 0 if you don't want to allocate a specific amount of ACW time. It
essentially means an indefinite amount of time. When the conversation ends, ACW starts; the agent must choose Close
contact to end ACW.

###### Note

When returned by a `SearchUsers` call, `AfterContactWorkTimeLimit` is returned in
milliseconds.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutoAccept`

The Auto accept setting.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeskPhoneNumber`

The phone number for the user's desk phone.

_Required_: No

_Type_: String

_Pattern_: `\+[1-9]\d{1,14}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PersistentConnection`

The persistent connection setting for the user.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PhoneType`

The phone type.

_Required_: No

_Type_: String

_Allowed values_: `SOFT_PHONE | DESK_PHONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

UserIdentityInfo

UserProficiency
