This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DLM::LifecyclePolicy EncryptionConfiguration

**\[Event-based policies only\]** Specifies the encryption settings for cross-Region snapshot copies created by
event-based policies.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CmkArn" : String,
  "Encrypted" : Boolean
}

```

### YAML

```yaml

  CmkArn: String
  Encrypted: Boolean

```

## Properties

`CmkArn`

The Amazon Resource Name (ARN) of the AWS KMS key to use for EBS encryption. If
this parameter is not specified, the default KMS key for the account is used.

_Required_: No

_Type_: String

_Pattern_: `arn:aws(-[a-z]{1,4}){0,2}:kms:([a-z]+-){2,3}\d:\d+:key/.*`

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Encrypted`

To encrypt a copy of an unencrypted snapshot when encryption by default is not enabled, enable
encryption using this parameter. Copies of encrypted snapshots are encrypted, even if this
parameter is false or when encryption by default is not enabled.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeprecateRule

EventParameters

All content copied from https://docs.aws.amazon.com/.
