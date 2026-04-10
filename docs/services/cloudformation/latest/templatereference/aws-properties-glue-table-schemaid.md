This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Table SchemaId

A structure that contains schema identity fields. Either this or the `SchemaVersionId` has to be
provided.

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

The name of the schema registry that contains the schema.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SchemaArn`

The Amazon Resource Name (ARN) of the schema. One of `SchemaArn` or `SchemaName` has to be
provided.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SchemaName`

The name of the schema. One of `SchemaArn` or `SchemaName` has to be provided.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Order

SchemaReference

All content copied from https://docs.aws.amazon.com/.
