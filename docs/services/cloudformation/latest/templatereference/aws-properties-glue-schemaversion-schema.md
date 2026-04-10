This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::SchemaVersion Schema

A wrapper structure to contain schema identity fields. Either `SchemaArn`, or `SchemaName` and `RegistryName` has to be provided.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RegistryName" : String,
  "SchemaArn" : String,
  "SchemaName" : String
}

```

### YAML

```yaml

  RegistryName: String
  SchemaArn: String
  SchemaName: String

```

## Properties

`RegistryName`

The name of the registry where the schema is stored. Either `SchemaArn`, or `SchemaName` and `RegistryName` has to be provided.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SchemaArn`

The Amazon Resource Name (ARN) of the schema. Either `SchemaArn`, or `SchemaName` and `RegistryName` has to be provided.

_Required_: No

_Type_: String

_Pattern_: `arn:(aws|aws-us-gov|aws-cn):glue:.*`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SchemaName`

The name of the schema. Either `SchemaArn`, or `SchemaName` and `RegistryName` has to be provided.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Glue::SchemaVersion

AWS::Glue::SchemaVersionMetadata

All content copied from https://docs.aws.amazon.com/.
