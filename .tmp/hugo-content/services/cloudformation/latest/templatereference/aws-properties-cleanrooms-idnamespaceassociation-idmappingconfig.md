This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::IdNamespaceAssociation IdMappingConfig

The configuration settings for the ID mapping table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowUseAsDimensionColumn" : Boolean
}

```

### YAML

```yaml

  AllowUseAsDimensionColumn: Boolean

```

## Properties

`AllowUseAsDimensionColumn`

An indicator as to whether you can use your column as a dimension column in the ID mapping table ( `TRUE`) or not ( `FALSE`).

Default is `FALSE`.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CleanRooms::IdNamespaceAssociation

IdNamespaceAssociationInputReferenceConfig

All content copied from https://docs.aws.amazon.com/.
