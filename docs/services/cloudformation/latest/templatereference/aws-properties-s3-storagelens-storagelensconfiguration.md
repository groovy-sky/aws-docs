This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::StorageLens StorageLensConfiguration

This is the property of the Amazon S3 Storage Lens configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccountLevel" : AccountLevel,
  "AwsOrg" : AwsOrg,
  "DataExport" : DataExport,
  "Exclude" : BucketsAndRegions,
  "ExpandedPrefixesDataExport" : StorageLensExpandedPrefixesDataExport,
  "Id" : String,
  "Include" : BucketsAndRegions,
  "IsEnabled" : Boolean,
  "PrefixDelimiter" : String,
  "StorageLensArn" : String
}

```

### YAML

```yaml

  AccountLevel:
    AccountLevel
  AwsOrg:
    AwsOrg
  DataExport:
    DataExport
  Exclude:
    BucketsAndRegions
  ExpandedPrefixesDataExport:
    StorageLensExpandedPrefixesDataExport
  Id: String
  Include:
    BucketsAndRegions
  IsEnabled: Boolean
  PrefixDelimiter: String
  StorageLensArn: String

```

## Properties

`AccountLevel`

This property contains the details of the account-level metrics for Amazon S3 Storage Lens
configuration.

_Required_: Yes

_Type_: [AccountLevel](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-storagelens-accountlevel.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AwsOrg`

This property contains the details of the AWS Organization for the S3
Storage Lens configuration.

_Required_: No

_Type_: [AwsOrg](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-storagelens-awsorg.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataExport`

This property contains the details of this S3 Storage Lens configuration's metrics
export.

_Required_: No

_Type_: [DataExport](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-storagelens-dataexport.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Exclude`

This property contains the details of the bucket and or Regions excluded for Amazon S3
Storage Lens configuration.

_Required_: No

_Type_: [BucketsAndRegions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-storagelens-bucketsandregions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExpandedPrefixesDataExport`

This property configures your S3 Storage Lens expanded prefixes metrics
report.

_Required_: No

_Type_: [StorageLensExpandedPrefixesDataExport](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-storagelens-storagelensexpandedprefixesdataexport.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

This property contains the details of the ID of the S3 Storage Lens configuration.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9\-_.]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Include`

This property contains the details of the bucket and or Regions included for Amazon S3
Storage Lens configuration.

_Required_: No

_Type_: [BucketsAndRegions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-storagelens-bucketsandregions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsEnabled`

This property contains the details of whether the Amazon S3 Storage Lens configuration is
enabled.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrefixDelimiter`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StorageLensArn`

This property contains the details of the ARN of the S3 Storage Lens configuration. This
property is read-only.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SSEKMS

StorageLensExpandedPrefixesDataExport
