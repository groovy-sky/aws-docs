This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSM::ResourceDataSync S3Destination

Information about the target S3 bucket for the resource data sync.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketName" : String,
  "BucketPrefix" : String,
  "BucketRegion" : String,
  "KMSKeyArn" : String,
  "SyncFormat" : String
}

```

### YAML

```yaml

  BucketName: String
  BucketPrefix: String
  BucketRegion: String
  KMSKeyArn: String
  SyncFormat: String

```

## Properties

`BucketName`

The name of the S3 bucket where the aggregated data is stored.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BucketPrefix`

An Amazon S3 prefix for the bucket.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BucketRegion`

The AWS Region with the S3 bucket targeted by the resource data sync.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KMSKeyArn`

The ARN of an encryption key for a destination in Amazon S3. Must belong to the same
Region as the destination S3 bucket.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SyncFormat`

A supported sync format. The following format is currently supported: JsonSerDe

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AwsOrganizationsSource

SyncSource

All content copied from https://docs.aws.amazon.com/.
