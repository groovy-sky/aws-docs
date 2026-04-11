This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::Cluster EncryptionAtRest

The data-volume encryption details. You can't update encryption at rest settings for existing clusters.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataVolumeKMSKeyId" : String
}

```

### YAML

```yaml

  DataVolumeKMSKeyId: String

```

## Properties

`DataVolumeKMSKeyId`

The ARN of the Amazon KMS key for encrypting data at rest. If you don't specify a KMS key, MSK creates one for you and uses it.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EBSStorageInfo

EncryptionInfo

All content copied from https://docs.aws.amazon.com/.
