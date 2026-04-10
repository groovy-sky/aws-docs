This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EntityResolution::IdNamespace IdNamespaceInputSource

An object containing `inputSourceARN` and `schemaName`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InputSourceARN" : String,
  "SchemaName" : String
}

```

### YAML

```yaml

  InputSourceARN: String
  SchemaName: String

```

## Properties

`InputSourceARN`

An AWS Glue table Amazon Resource Name (ARN) or a matching workflow ARN for
the input source table.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws|aws-us-gov|aws-cn):entityresolution:[a-z]{2}-[a-z]{1,10}-[0-9]:[0-9]{12}:(matchingworkflow/[a-zA-Z_0-9-]{1,255})$|^arn:(aws|aws-us-gov|aws-cn):glue:[a-z]{2}-[a-z]{1,10}-[0-9]:[0-9]{12}:(table/[a-zA-Z_0-9-]{1,255}/[a-zA-Z_0-9-]{1,255})$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SchemaName`

The name of the schema.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z_0-9-]*$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IdNamespaceIdMappingWorkflowProperties

NamespaceProviderProperties

All content copied from https://docs.aws.amazon.com/.
