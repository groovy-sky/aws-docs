This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket SseKmsEncryptedObjects

A container for filter information for the selection of S3 objects encrypted with AWS KMS.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Status" : String
}

```

### YAML

```yaml

  Status: String

```

## Properties

`Status`

Specifies whether Amazon S3 replicates objects created with server-side encryption using an AWS KMS key
stored in AWS Key Management Service.

_Required_: Yes

_Type_: String

_Allowed values_: `Disabled | Enabled`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SourceSelectionCriteria

StorageClassAnalysis

All content copied from https://docs.aws.amazon.com/.
