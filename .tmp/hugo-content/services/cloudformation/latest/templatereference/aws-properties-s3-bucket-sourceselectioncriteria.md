This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket SourceSelectionCriteria

A container that describes additional filters for identifying the source objects that you
want to replicate. You can choose to enable or disable the replication of these
objects.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ReplicaModifications" : ReplicaModifications,
  "SseKmsEncryptedObjects" : SseKmsEncryptedObjects
}

```

### YAML

```yaml

  ReplicaModifications:
    ReplicaModifications
  SseKmsEncryptedObjects:
    SseKmsEncryptedObjects

```

## Properties

`ReplicaModifications`

A filter that you can specify for selection for modifications on replicas.

_Required_: No

_Type_: [ReplicaModifications](aws-properties-s3-bucket-replicamodifications.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SseKmsEncryptedObjects`

A container for filter information for the selection of Amazon S3 objects encrypted with
AWS KMS.

_Required_: No

_Type_: [SseKmsEncryptedObjects](aws-properties-s3-bucket-ssekmsencryptedobjects.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ServerSideEncryptionRule

SseKmsEncryptedObjects

All content copied from https://docs.aws.amazon.com/.
