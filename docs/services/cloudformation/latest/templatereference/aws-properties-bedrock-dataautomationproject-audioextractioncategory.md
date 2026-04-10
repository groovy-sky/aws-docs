This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataAutomationProject AudioExtractionCategory

Settings for generating data from audio.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "State" : String,
  "TypeConfiguration" : AudioExtractionCategoryTypeConfiguration,
  "Types" : [ String, ... ]
}

```

### YAML

```yaml

  State: String
  TypeConfiguration:
    AudioExtractionCategoryTypeConfiguration
  Types:
    - String

```

## Properties

`State`

Whether generating categorical data from audio is enabled.

_Required_: Yes

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TypeConfiguration`

This element contains information about extractions from different types. Used to enable speaker and channel labeling for transcripts.

_Required_: No

_Type_: [AudioExtractionCategoryTypeConfiguration](aws-properties-bedrock-dataautomationproject-audioextractioncategorytypeconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Types`

The types of data to generate.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Bedrock::DataAutomationProject

AudioExtractionCategoryTypeConfiguration

All content copied from https://docs.aws.amazon.com/.
