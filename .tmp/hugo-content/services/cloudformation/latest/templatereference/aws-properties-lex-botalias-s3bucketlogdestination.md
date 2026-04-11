This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::BotAlias S3BucketLogDestination

Specifies an Amazon S3 bucket for logging audio conversations

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KmsKeyArn" : String,
  "LogPrefix" : String,
  "S3BucketArn" : String
}

```

### YAML

```yaml

  KmsKeyArn: String
  LogPrefix: String
  S3BucketArn: String

```

## Properties

`KmsKeyArn`

The Amazon Resource Name (ARN) of an AWS Key Management Service
(KMS) key for encrypting audio log files stored in an Amazon S3 bucket.

_Required_: No

_Type_: String

_Pattern_: `^arn:[\w\-]+:kms:[\w\-]+:[\d]{12}:(?:key\/[\w\-]+|alias\/[a-zA-Z0-9:\/_\-]{1,256})$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogPrefix`

The S3 prefix to assign to audio log files.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3BucketArn`

The Amazon Resource Name (ARN) of an Amazon S3 bucket where audio
log files are stored.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:[\w\-]+:s3:::[a-z0-9][\.\-a-z0-9]{1,61}[a-z0-9]$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LambdaCodeHook

SentimentAnalysisSettings

All content copied from https://docs.aws.amazon.com/.
