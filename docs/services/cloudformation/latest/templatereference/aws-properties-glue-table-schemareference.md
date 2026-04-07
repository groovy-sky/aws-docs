This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Table SchemaReference

An object that references a schema stored in the AWS Glue Schema Registry.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SchemaId" : SchemaId,
  "SchemaVersionId" : String,
  "SchemaVersionNumber" : Integer
}

```

### YAML

```yaml

  SchemaId:
    SchemaId
  SchemaVersionId: String
  SchemaVersionNumber: Integer

```

## Properties

`SchemaId`

A structure that contains schema identity fields. Either this or the `SchemaVersionId` has to be
provided.

_Required_: No

_Type_: [SchemaId](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-glue-table-schemaid.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SchemaVersionId`

The unique ID assigned to a version of the schema. Either this or the `SchemaId` has to be provided.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SchemaVersionNumber`

The version number of the schema.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SchemaId

SerdeInfo
