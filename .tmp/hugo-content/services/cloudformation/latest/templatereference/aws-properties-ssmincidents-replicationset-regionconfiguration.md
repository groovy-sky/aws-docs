This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSMIncidents::ReplicationSet RegionConfiguration

The `RegionConfiguration` property specifies the Region and AWS Key Management Service key to add
to the replication set.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SseKmsKeyId" : String
}

```

### YAML

```yaml

  SseKmsKeyId: String

```

## Properties

`SseKmsKeyId`

The AWS Key Management Service key ID to use to encrypt your replication set.

_Required_: Yes

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SSMIncidents::ReplicationSet

ReplicationRegion

All content copied from https://docs.aws.amazon.com/.
