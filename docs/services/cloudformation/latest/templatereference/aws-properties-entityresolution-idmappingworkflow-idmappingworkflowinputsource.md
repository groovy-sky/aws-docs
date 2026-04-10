This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EntityResolution::IdMappingWorkflow IdMappingWorkflowInputSource

An object containing `inputSourceARN`, `schemaName`, and
`type`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InputSourceARN" : String,
  "SchemaArn" : String,
  "Type" : String
}

```

### YAML

```yaml

  InputSourceARN: String
  SchemaArn: String
  Type: String

```

## Properties

`InputSourceARN`

An AWS Glue table Amazon Resource Name (ARN) or a matching workflow ARN for
the input source table.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws|aws-us-gov|aws-cn):entityresolution:[a-z]{2}-[a-z]{1,10}-[0-9]:[0-9]{12}:(idnamespace/[a-zA-Z_0-9-]{1,255})$|^arn:(aws|aws-us-gov|aws-cn):entityresolution:[a-z]{2}-[a-z]{1,10}-[0-9]:[0-9]{12}:(matchingworkflow/[a-zA-Z_0-9-]{1,255})$|^arn:(aws|aws-us-gov|aws-cn):glue:[a-z]{2}-[a-z]{1,10}-[0-9]:[0-9]{12}:(table/[a-zA-Z_0-9-]{1,255}/[a-zA-Z_0-9-]{1,255})$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SchemaArn`

The ARN (Amazon Resource Name) that AWS Entity Resolution generated for the
`SchemaMapping`.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws|aws-us-gov|aws-cn):entityresolution:.*:[0-9]+:(schemamapping/.*)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of ID namespace. There are two types: `SOURCE` and
`TARGET`.

The `SOURCE` contains configurations for `sourceId` data that will
be processed in an ID mapping workflow.

The `TARGET` contains a configuration of `targetId` which all
`sourceIds` will resolve to.

_Required_: No

_Type_: String

_Allowed values_: `SOURCE | TARGET`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IdMappingTechniques

IdMappingWorkflowOutputSource

All content copied from https://docs.aws.amazon.com/.
