This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMRServerless::Application ManagedPersistenceMonitoringConfiguration

The managed log persistence configuration for a job run.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "EncryptionKeyArn" : String
}

```

### YAML

```yaml

  Enabled: Boolean
  EncryptionKeyArn: String

```

## Properties

`Enabled`

Enables managed logging and defaults to true. If set to false, managed logging will be
turned off.

_Required_: No

_Type_: Boolean

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`EncryptionKeyArn`

The KMS key ARN to encrypt the logs stored in managed log persistence.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z0-9-]*):kms:[a-zA-Z0-9\-]*:(\d{12})?:key\/[a-zA-Z0-9-]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LogTypeMapKeyValuePair

MaximumAllowedResources

All content copied from https://docs.aws.amazon.com/.
