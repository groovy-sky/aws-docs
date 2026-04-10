This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::AnalysisTemplate ColumnClassificationDetails

Contains classification information for data columns, including mappings that specify how columns should be handled during synthetic data generation and privacy analysis.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColumnMapping" : [ SyntheticDataColumnProperties, ... ]
}

```

### YAML

```yaml

  ColumnMapping:
    - SyntheticDataColumnProperties

```

## Properties

`ColumnMapping`

A mapping that defines the classification of data columns for synthetic data generation and specifies how each column should be handled during the privacy-preserving data synthesis process.

_Required_: Yes

_Type_: Array of [SyntheticDataColumnProperties](aws-properties-cleanrooms-analysistemplate-syntheticdatacolumnproperties.md)

_Minimum_: `5`

_Maximum_: `1000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AnalysisTemplateArtifacts

ErrorMessageConfiguration

All content copied from https://docs.aws.amazon.com/.
