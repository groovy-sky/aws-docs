This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Timestream::ScheduledQuery S3Configuration

Details on S3 location for error reports that result from running a query.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketName" : String,
  "EncryptionOption" : String,
  "ObjectKeyPrefix" : String
}

```

### YAML

```yaml

  BucketName: String
  EncryptionOption: String
  ObjectKeyPrefix: String

```

## Properties

`BucketName`

Name of the S3 bucket under which error reports will be created.

_Required_: Yes

_Type_: String

_Pattern_: `[a-z0-9][\.\-a-z0-9]{1,61}[a-z0-9]`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EncryptionOption`

Encryption at rest options for the error reports. If no encryption option is specified,
Timestream will choose SSE\_S3 as default.

_Required_: No

_Type_: String

_Allowed values_: `SSE_S3 | SSE_KMS`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ObjectKeyPrefix`

Prefix for the error report key. Timestream by default adds the following prefix to the
error report path.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9|!\-_*'\(\)]([a-zA-Z0-9]|[!\-_*'\(\)\/.])+`

_Minimum_: `1`

_Maximum_: `896`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NotificationConfiguration

ScheduleConfiguration

All content copied from https://docs.aws.amazon.com/.
