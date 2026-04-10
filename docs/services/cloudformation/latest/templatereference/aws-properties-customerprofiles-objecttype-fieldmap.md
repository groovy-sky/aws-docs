This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::ObjectType FieldMap

A map of the name and ObjectType field.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "ObjectTypeField" : ObjectTypeField
}

```

### YAML

```yaml

  Name: String
  ObjectTypeField:
    ObjectTypeField

```

## Properties

`Name`

Name of the field.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ObjectTypeField`

Represents a field in a ProfileObjectType.

_Required_: No

_Type_: [ObjectTypeField](aws-properties-customerprofiles-objecttype-objecttypefield.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CustomerProfiles::ObjectType

KeyMap

All content copied from https://docs.aws.amazon.com/.
