This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EntityResolution::SchemaMapping

Creates a schema mapping, which defines the schema of the input customer records table.
The `SchemaMapping` also provides AWS Entity Resolution with some metadata about the
table, such as the attribute types of the columns and which columns to match on.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EntityResolution::SchemaMapping",
  "Properties" : {
      "Description" : String,
      "MappedInputFields" : [ SchemaInputAttribute, ... ],
      "SchemaName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EntityResolution::SchemaMapping
Properties:
  Description: String
  MappedInputFields:
    - SchemaInputAttribute
  SchemaName: String
  Tags:
    - Tag

```

## Properties

`Description`

A description of the schema.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MappedInputFields`

A list of `MappedInputFields`. Each `MappedInputField` corresponds
to a column the source data table, and contains column name plus additional information
that AWS Entity Resolution uses for matching.

_Required_: Yes

_Type_: Array of [SchemaInputAttribute](aws-properties-entityresolution-schemamapping-schemainputattribute.md)

_Minimum_: `2`

_Maximum_: `35`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SchemaName`

The name of the schema. There can't be multiple `SchemaMappings` with the
same name.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z_0-9-]*$`

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags used to organize, track, or control access for this resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-entityresolution-schemamapping-tag.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EntityResolution::PolicyStatement

SchemaInputAttribute

All content copied from https://docs.aws.amazon.com/.
