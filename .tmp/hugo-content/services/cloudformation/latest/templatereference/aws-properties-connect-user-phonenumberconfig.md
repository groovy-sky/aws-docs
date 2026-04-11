This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::User PhoneNumberConfig

Configuration settings for phone type and phone number.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Channel" : String,
  "PhoneNumber" : String,
  "PhoneType" : String
}

```

### YAML

```yaml

  Channel: String
  PhoneNumber: String
  PhoneType: String

```

## Properties

`Channel`

The channel for this phone number configuration. **Only `VOICE` is supported for this data type.**

_Required_: Yes

_Type_: String

_Allowed values_: `VOICE | CHAT | TASK | EMAIL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PhoneNumber`

The phone number for the user's desk phone.

_Required_: No

_Type_: String

_Pattern_: `\+[1-9]\d{1,14}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PhoneType`

The phone type. Valid values: SOFT\_PHONE, DESK\_PHONE.

_Required_: Yes

_Type_: String

_Allowed values_: `SOFT_PHONE | DESK_PHONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PersistentConnectionConfig

Tag

All content copied from https://docs.aws.amazon.com/.
