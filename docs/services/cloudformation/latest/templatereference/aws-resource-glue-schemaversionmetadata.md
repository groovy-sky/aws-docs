This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::SchemaVersionMetadata

The `AWS::Glue::SchemaVersionMetadata` is an AWS Glue resource type that defines the metadata key-value pairs for a schema version in AWS Glue Schema Registry.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Glue::SchemaVersionMetadata",
  "Properties" : {
      "Key" : String,
      "SchemaVersionId" : String,
      "Value" : String
    }
}

```

### YAML

```yaml

Type: AWS::Glue::SchemaVersionMetadata
Properties:
  Key: String
  SchemaVersionId: String
  Value: String

```

## Properties

`Key`

A metadata key in a key-value pair for metadata.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SchemaVersionId`

The version number of the schema.

_Required_: Yes

_Type_: String

_Pattern_: `[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Value`

A metadata key's corresponding value.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Schema

AWS::Glue::SecurityConfiguration
