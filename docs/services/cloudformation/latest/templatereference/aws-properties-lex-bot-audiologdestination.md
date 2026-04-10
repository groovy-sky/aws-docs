This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot AudioLogDestination

The location of audio log files collected when conversation logging
is enabled for a bot.

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

Specifies the Amazon S3 bucket where the audio files are
stored.

_Required_: Yes

_Type_: [S3BucketLogDestination](aws-properties-lex-bot-s3bucketlogdestination.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AudioAndDTMFInputSpecification

AudioLogSetting

All content copied from https://docs.aws.amazon.com/.
