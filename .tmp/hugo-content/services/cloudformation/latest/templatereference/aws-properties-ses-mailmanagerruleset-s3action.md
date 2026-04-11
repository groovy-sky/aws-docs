This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerRuleSet S3Action

Writes the MIME content of the email to an S3 bucket.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ActionFailurePolicy" : String,
  "RoleArn" : String,
  "S3Bucket" : String,
  "S3Prefix" : String,
  "S3SseKmsKeyId" : String
}

```

### YAML

```yaml

  ActionFailurePolicy: String
  RoleArn: String
  S3Bucket: String
  S3Prefix: String
  S3SseKmsKeyId: String

```

## Properties

`ActionFailurePolicy`

A policy that states what to do in the case of failure. The action will fail if there
are configuration errors. For example, the specified the bucket has been deleted.

_Required_: No

_Type_: String

_Allowed values_: `CONTINUE | DROP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of the IAM Role to use while writing to S3. This role must have access to
the s3:PutObject, kms:Encrypt, and kms:GenerateDataKey APIs for the given bucket.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9:_/+=,@.#-]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Bucket`

The bucket name of the S3 bucket to write to.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9.-]+$`

_Minimum_: `1`

_Maximum_: `62`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Prefix`

The S3 prefix to use for the write to the s3 bucket.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9!_.*'()/-]+$`

_Minimum_: `1`

_Maximum_: `62`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3SseKmsKeyId`

The KMS Key ID to use to encrypt the message in S3.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9-:/]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RuleVerdictToEvaluate

SendAction

All content copied from https://docs.aws.amazon.com/.
