This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EntityResolution::MatchingWorkflow InputSource

An object containing `inputSourceARN`, `schemaName`, and
`applyNormalization`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApplyNormalization" : Boolean,
  "InputSourceARN" : String,
  "SchemaArn" : String
}

```

### YAML

```yaml

  ApplyNormalization: Boolean
  InputSourceARN: String
  SchemaArn: String

```

## Properties

`ApplyNormalization`

Normalizes the attributes defined in the schema in the input data. For example, if an
attribute has an `AttributeType` of `PHONE_NUMBER`, and the data in
the input table is in a format of 1234567890, AWS Entity Resolution will normalize this field
in the output to (123)-456-7890.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputSourceARN`

An object containing `inputSourceARN`, `schemaName`, and
`applyNormalization`.

_Required_: Yes

_Type_: String

_Pattern_: `arn:(aws|aws-us-gov|aws-cn):.*:.*:[0-9]+:.*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SchemaArn`

The name of the schema.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws|aws-us-gov|aws-cn):entityresolution:.*:[0-9]+:(schemamapping/.*)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

IncrementalRunConfig

IntermediateSourceConfiguration
