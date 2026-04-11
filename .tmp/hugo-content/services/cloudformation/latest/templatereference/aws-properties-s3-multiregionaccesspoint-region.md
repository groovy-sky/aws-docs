This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::MultiRegionAccessPoint Region

A bucket associated with a specific Region when creating Multi-Region Access
Points.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Bucket" : String,
  "BucketAccountId" : String
}

```

### YAML

```yaml

  Bucket: String
  BucketAccountId: String

```

## Properties

`Bucket`

The name of the associated bucket for the Region.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z0-9][a-z0-9//.//-]*[a-z0-9]$`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BucketAccountId`

The AWS account ID that owns the Amazon S3 bucket that's associated with this Multi-Region Access Point.

_Required_: No

_Type_: String

_Pattern_: `^[0-9]{12}$`

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PublicAccessBlockConfiguration

AWS::S3::MultiRegionAccessPointPolicy

All content copied from https://docs.aws.amazon.com/.
