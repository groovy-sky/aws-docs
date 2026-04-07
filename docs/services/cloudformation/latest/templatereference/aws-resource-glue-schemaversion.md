This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::SchemaVersion

The `AWS::Glue::SchemaVersion` is an AWS Glue resource type that manages schema versions of schemas in the AWS Glue Schema Registry.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Glue::SchemaVersion",
  "Properties" : {
      "Schema" : Schema,
      "SchemaDefinition" : String
    }
}

```

### YAML

```yaml

Type: AWS::Glue::SchemaVersion
Properties:
  Schema:
    Schema
  SchemaDefinition: String

```

## Properties

`Schema`

The schema that includes the schema version.

_Required_: Yes

_Type_: [Schema](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-glue-schemaversion-schema.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SchemaDefinition`

The schema definition for the schema version.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `170000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`VersionId`

Property description not available.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Schema
