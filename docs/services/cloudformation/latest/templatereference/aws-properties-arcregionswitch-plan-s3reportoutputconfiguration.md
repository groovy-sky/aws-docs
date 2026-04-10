This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ARCRegionSwitch::Plan S3ReportOutputConfiguration

Configuration for delivering generated reports to an Amazon S3 bucket.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketOwner" : String,
  "BucketPath" : String
}

```

### YAML

```yaml

  BucketOwner: String
  BucketPath: String

```

## Properties

`BucketOwner`

The AWS account ID that owns the S3 bucket. Required to ensure the bucket is still owned by the same expected owner at generation time.

_Required_: No

_Type_: String

_Pattern_: `^\d{12}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BucketPath`

The S3 bucket name and optional prefix where reports are stored. Format: bucket-name or bucket-name/prefix.

_Required_: No

_Type_: String

_Pattern_: `^(?:s3://)?[a-z0-9][a-z0-9-]{1,61}[a-z0-9](?:/[^/ ][^/]*)*/?$`

_Minimum_: `3`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Route53ResourceRecordSet

Service

All content copied from https://docs.aws.amazon.com/.
