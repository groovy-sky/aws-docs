This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Dataset FormatOptions

Represents a set of options that define the structure of either comma-separated value (CSV),
Excel, or JSON input.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Csv" : CsvOptions,
  "Excel" : ExcelOptions,
  "Json" : JsonOptions
}

```

### YAML

```yaml

  Csv:
    CsvOptions
  Excel:
    ExcelOptions
  Json:
    JsonOptions

```

## Properties

`Csv`

Options that define how CSV input is to be interpreted by DataBrew.

_Required_: No

_Type_: [CsvOptions](aws-properties-databrew-dataset-csvoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Excel`

Options that define how Excel input is to be interpreted by DataBrew.

_Required_: No

_Type_: [ExcelOptions](aws-properties-databrew-dataset-exceloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Json`

Options that define how JSON input is to be interpreted by DataBrew.

_Required_: No

_Type_: [JsonOptions](aws-properties-databrew-dataset-jsonoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FilterValue

Input

All content copied from https://docs.aws.amazon.com/.
