This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::CapacityManagerDataExport

Creates a new data export configuration for EC2 Capacity Manager. This allows you to automatically export capacity usage data to an S3 bucket on a scheduled basis.
The exported data includes metrics for On-Demand, Spot, and Capacity Reservations usage across your organization.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::CapacityManagerDataExport",
  "Properties" : {
      "OutputFormat" : String,
      "S3BucketName" : String,
      "S3BucketPrefix" : String,
      "Schedule" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EC2::CapacityManagerDataExport
Properties:
  OutputFormat: String
  S3BucketName: String
  S3BucketPrefix: String
  Schedule: String
  Tags:
    - Tag

```

## Properties

`OutputFormat`

The file format of the exported data.

_Required_: Yes

_Type_: String

_Allowed values_: `csv | parquet`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3BucketName`

The name of the S3 bucket where export files are delivered.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3BucketPrefix`

The S3 key prefix used for organizing export files within the bucket.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Schedule`

The frequency at which data exports are generated.

_Required_: Yes

_Type_: String

_Allowed values_: `hourly`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags associated with the data export configuration.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-capacitymanagerdataexport-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the export ID, such as `cmde-0f3f217ee9a5baead`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CapacityManagerDataExportId`

The unique identifier for the data export configuration.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon EC2

Tag

All content copied from https://docs.aws.amazon.com/.
