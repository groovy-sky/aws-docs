This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis AnalysisSourceTemplate

The source template of an analysis.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Arn" : String,
  "DataSetReferences" : [ DataSetReference, ... ]
}

```

### YAML

```yaml

  Arn: String
  DataSetReferences:
    - DataSetReference

```

## Properties

`Arn`

The Amazon Resource Name (ARN) of the source template of an analysis.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataSetReferences`

The dataset references of the source template of an analysis.

_Required_: Yes

_Type_: Array of [DataSetReference](aws-properties-quicksight-analysis-datasetreference.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AnalysisSourceEntity

AnchorDateConfiguration

All content copied from https://docs.aws.amazon.com/.
