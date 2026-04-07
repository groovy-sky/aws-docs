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

_Type_: [CsvOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-databrew-dataset-csvoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Excel`

Options that define how Excel input is to be interpreted by DataBrew.

_Required_: No

_Type_: [ExcelOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-databrew-dataset-exceloptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Json`

Options that define how JSON input is to be interpreted by DataBrew.

_Required_: No

_Type_: [JsonOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-databrew-dataset-jsonoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FilterValue

Input
