This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Dataset Variable

An instance of a variable to be passed to the `containerAction` execution. Each
variable must have a name and a value given by one of `stringValue`,
`datasetContentVersionValue`, or `outputFileUriValue`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DatasetContentVersionValue" : DatasetContentVersionValue,
  "DoubleValue" : Number,
  "OutputFileUriValue" : OutputFileUriValue,
  "StringValue" : String,
  "VariableName" : String
}

```

### YAML

```yaml

  DatasetContentVersionValue:
    DatasetContentVersionValue
  DoubleValue: Number
  OutputFileUriValue:
    OutputFileUriValue
  StringValue:
    String
  VariableName: String

```

## Properties

`DatasetContentVersionValue`

The value of the variable as a structure that specifies a dataset content version.

_Required_: No

_Type_: [DatasetContentVersionValue](aws-properties-iotanalytics-dataset-datasetcontentversionvalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DoubleValue`

The value of the variable as a double (numeric).

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputFileUriValue`

The value of the variable as a structure that specifies an output file URI.

_Required_: No

_Type_: [OutputFileUriValue](aws-properties-iotanalytics-dataset-outputfileurivalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StringValue`

The value of the variable as a string.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VariableName`

The name of the variable.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TriggeringDataset

VersioningConfiguration

All content copied from https://docs.aws.amazon.com/.
