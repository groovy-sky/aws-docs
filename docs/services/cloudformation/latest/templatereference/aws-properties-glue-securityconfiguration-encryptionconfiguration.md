This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::SecurityConfiguration EncryptionConfiguration

Specifies an encryption configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudWatchEncryption" : CloudWatchEncryption,
  "JobBookmarksEncryption" : JobBookmarksEncryption,
  "S3Encryptions" : [ S3Encryption, ... ]
}

```

### YAML

```yaml

  CloudWatchEncryption:
    CloudWatchEncryption
  JobBookmarksEncryption:
    JobBookmarksEncryption
  S3Encryptions:
    - S3Encryption

```

## Properties

`CloudWatchEncryption`

The encryption configuration for Amazon CloudWatch.

_Required_: No

_Type_: [CloudWatchEncryption](aws-properties-glue-securityconfiguration-cloudwatchencryption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JobBookmarksEncryption`

The encryption configuration for job bookmarks.

_Required_: No

_Type_: [JobBookmarksEncryption](aws-properties-glue-securityconfiguration-jobbookmarksencryption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Encryptions`

The encyption configuration for Amazon Simple Storage Service (Amazon S3) data.

_Required_: No

_Type_: Array of [S3Encryption](aws-properties-glue-securityconfiguration-s3encryption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CloudWatchEncryption

JobBookmarksEncryption

All content copied from https://docs.aws.amazon.com/.
