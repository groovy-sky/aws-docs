This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMRServerless::Application CloudWatchLoggingConfiguration

The Amazon CloudWatch configuration for monitoring logs. You can configure your jobs
to send log information to CloudWatch.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "EncryptionKeyArn" : String,
  "LogGroupName" : String,
  "LogStreamNamePrefix" : String,
  "LogTypeMap" : [ LogTypeMapKeyValuePair, ... ]
}

```

### YAML

```yaml

  Enabled: Boolean
  EncryptionKeyArn: String
  LogGroupName: String
  LogStreamNamePrefix: String
  LogTypeMap:
    - LogTypeMapKeyValuePair

```

## Properties

`Enabled`

Enables CloudWatch logging.

_Required_: No

_Type_: Boolean

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`EncryptionKeyArn`

The AWS Key Management Service (KMS) key ARN to encrypt the logs that you store in CloudWatch Logs.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z0-9-]*):kms:[a-zA-Z0-9\-]*:(\d{12})?:key\/[a-zA-Z0-9-]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`LogGroupName`

The name of the log group in Amazon CloudWatch Logs where you want to publish your
logs.

_Required_: No

_Type_: String

_Pattern_: `^[\.\-_/#A-Za-z0-9]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`LogStreamNamePrefix`

Prefix for the CloudWatch log stream name.

_Required_: No

_Type_: String

_Pattern_: `^[^:*]*$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`LogTypeMap`

Property description not available.

_Required_: No

_Type_: Array of [LogTypeMapKeyValuePair](aws-properties-emrserverless-application-logtypemapkeyvaluepair.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AutoStopConfiguration

ConfigurationObject

All content copied from https://docs.aws.amazon.com/.
