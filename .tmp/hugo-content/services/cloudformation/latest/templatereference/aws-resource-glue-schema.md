This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Schema

The `AWS::Glue::Schema` is an AWS Glue resource type that manages schemas in the AWS Glue Schema Registry.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Glue::Schema",
  "Properties" : {
      "CheckpointVersion" : SchemaVersion,
      "Compatibility" : String,
      "DataFormat" : String,
      "Description" : String,
      "Name" : String,
      "Registry" : Registry,
      "SchemaDefinition" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Glue::Schema
Properties:
  CheckpointVersion:
    SchemaVersion
  Compatibility: String
  DataFormat: String
  Description: String
  Name: String
  Registry:
    Registry
  SchemaDefinition: String
  Tags:
    - Tag

```

## Properties

`CheckpointVersion`

Specify the `VersionNumber` or the `IsLatest` for setting the checkpoint for the schema. This is only required for updating a checkpoint.

_Required_: No

_Type_: [SchemaVersion](aws-properties-glue-schema-schemaversion.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Compatibility`

The compatibility mode of the schema.

_Required_: Yes

_Type_: String

_Allowed values_: `NONE | DISABLED | BACKWARD | BACKWARD_ALL | FORWARD | FORWARD_ALL | FULL | FULL_ALL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataFormat`

The data format of the schema definition. Currently only `AVRO` is supported.

_Required_: Yes

_Type_: String

_Allowed values_: `AVRO | JSON | PROTOBUF`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

A description of the schema if specified when created.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

Name of the schema to be created of max length of 255, and may only contain letters, numbers, hyphen, underscore, dollar sign, or hash mark. No whitespace.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Registry`

The registry where a schema is stored.

_Required_: No

_Type_: [Registry](aws-properties-glue-schema-registry.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SchemaDefinition`

The schema definition using the `DataFormat` setting for `SchemaName`.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `170000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

AWS tags that contain a key value pair and may be searched by console, command line, or API.

_Required_: No

_Type_: Array of [Tag](aws-properties-glue-schema-tag.md)

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the schema.

`InitialSchemaVersionId`

Property description not available.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Registry

All content copied from https://docs.aws.amazon.com/.
