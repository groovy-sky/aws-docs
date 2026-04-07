This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket Tiering

The S3 Intelligent-Tiering storage class is designed to optimize storage costs by automatically
moving data to the most cost-effective storage access tier, without additional operational
overhead.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccessTier" : String,
  "Days" : Integer
}

```

### YAML

```yaml

  AccessTier: String
  Days: Integer

```

## Properties

`AccessTier`

S3 Intelligent-Tiering access tier. See [Storage class for\
automatically optimizing frequently and infrequently accessed objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html#sc-dynamic-data-access) for a list of access
tiers in the S3 Intelligent-Tiering storage class.

_Required_: Yes

_Type_: String

_Allowed values_: `ARCHIVE_ACCESS | DEEP_ARCHIVE_ACCESS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Days`

The number of consecutive days of no access after which an object will be eligible to be
transitioned to the corresponding tier. The minimum number of days specified for Archive Access tier
must be at least 90 days and Deep Archive Access tier must be at least 180 days. The maximum can be up to
2 years (730 days).

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TargetObjectKeyFormat

TopicConfiguration
