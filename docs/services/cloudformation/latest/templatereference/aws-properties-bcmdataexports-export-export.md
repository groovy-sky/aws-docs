This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BCMDataExports::Export Export

The details that are available for an export.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataQuery" : DataQuery,
  "Description" : String,
  "DestinationConfigurations" : DestinationConfigurations,
  "ExportArn" : String,
  "Name" : String,
  "RefreshCadence" : RefreshCadence
}

```

### YAML

```yaml

  DataQuery:
    DataQuery
  Description: String
  DestinationConfigurations:
    DestinationConfigurations
  ExportArn: String
  Name: String
  RefreshCadence:
    RefreshCadence

```

## Properties

`DataQuery`

The data query for this specific data export.

_Required_: Yes

_Type_: [DataQuery](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bcmdataexports-export-dataquery.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description for this specific data export.

_Required_: No

_Type_: String

_Pattern_: `^[\S\s]*$`

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DestinationConfigurations`

The destination configuration for this specific data export.

_Required_: Yes

_Type_: [DestinationConfigurations](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bcmdataexports-export-destinationconfigurations.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExportArn`

The Amazon Resource Name (ARN) for this export.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:(bcm-data-exports):[-a-z0-9]*:[0-9]{12}:[-a-zA-Z0-9/:_]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of this specific data export.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9A-Za-z\-_]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RefreshCadence`

The cadence for AWS to update the export in your S3 bucket.

_Required_: Yes

_Type_: [RefreshCadence](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bcmdataexports-export-refreshcadence.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DestinationConfigurations

RefreshCadence
