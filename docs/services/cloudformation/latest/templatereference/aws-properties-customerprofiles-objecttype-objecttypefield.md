This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::ObjectType ObjectTypeField

Represents a field in a ProfileObjectType.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContentType" : String,
  "Source" : String,
  "Target" : String
}

```

### YAML

```yaml

  ContentType: String
  Source: String
  Target: String

```

## Properties

`ContentType`

The content type of the field. Used for determining equality when searching.

_Required_: No

_Type_: String

_Allowed values_: `STRING | NUMBER | PHONE_NUMBER | EMAIL_ADDRESS | NAME`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

A field of a ProfileObject. For example: \_source.FirstName, where “\_source” is a
ProfileObjectType of a Zendesk user and “FirstName” is a field in that
ObjectType.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Target`

The location of the data in the standard ProfileObject model. For example:
\_profile.Address.PostalCode.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KeyMap

ObjectTypeKey

All content copied from https://docs.aws.amazon.com/.
