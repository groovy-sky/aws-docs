This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::User UserProficiency

###### Note

A predefined attribute must be created before using `UserProficiencies`
in the Cloudformation _User_ template. For more information, see
[Predefined\
attributes](../../../connect/latest/adminguide/predefined-attributes.md).

Proficiency of a user.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AttributeName" : String,
  "AttributeValue" : String,
  "Level" : Number
}

```

### YAML

```yaml

  AttributeName: String
  AttributeValue: String
  Level: Number

```

## Properties

`AttributeName`

The name of user’s proficiency. You must use a predefined attribute name that is
present in the Amazon Connect instance.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AttributeValue`

The value of user’s proficiency. You must use a predefined attribute value that is
present in the Amazon Connect instance.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Level`

The level of the proficiency. The valid values are 1, 2, 3, 4 and 5.

_Required_: Yes

_Type_: Number

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UserPhoneConfig

VoiceEnhancementConfig

All content copied from https://docs.aws.amazon.com/.
