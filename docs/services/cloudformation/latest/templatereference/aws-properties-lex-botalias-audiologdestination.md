This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::BotAlias AudioLogDestination

Specifies the S3 bucket location where audio logs are stored.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3Bucket" : S3BucketLogDestination
}

```

### YAML

```yaml

  S3Bucket:
    S3BucketLogDestination

```

## Properties

`S3Bucket`

The S3 bucket location where audio logs are stored.

_Required_: Yes

_Type_: [S3BucketLogDestination](aws-properties-lex-botalias-s3bucketlogdestination.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Lex::BotAlias

AudioLogSetting

All content copied from https://docs.aws.amazon.com/.
