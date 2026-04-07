This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::ObjectType KeyMap

A unique key map that can be used to map data to the profile.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "ObjectTypeKeyList" : [ ObjectTypeKey, ... ]
}

```

### YAML

```yaml

  Name: String
  ObjectTypeKeyList:
    - ObjectTypeKey

```

## Properties

`Name`

Name of the key.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ObjectTypeKeyList`

A list of ObjectTypeKey.

_Required_: No

_Type_: Array of [ObjectTypeKey](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-customerprofiles-objecttype-objecttypekey.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FieldMap

ObjectTypeField
