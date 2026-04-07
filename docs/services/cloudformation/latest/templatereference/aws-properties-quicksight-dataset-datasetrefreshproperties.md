This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet DataSetRefreshProperties

The refresh properties of a dataset.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FailureConfiguration" : RefreshFailureConfiguration,
  "RefreshConfiguration" : RefreshConfiguration
}

```

### YAML

```yaml

  FailureConfiguration:
    RefreshFailureConfiguration
  RefreshConfiguration:
    RefreshConfiguration

```

## Properties

`FailureConfiguration`

The failure configuration for a dataset.

_Required_: No

_Type_: [RefreshFailureConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-dataset-refreshfailureconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RefreshConfiguration`

The refresh configuration for a dataset.

_Required_: No

_Type_: [RefreshConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-dataset-refreshconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DatasetParameter

DataSetStringComparisonFilterCondition
