This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DLM::LifecyclePolicy ArchiveRetainRule

**\[Custom snapshot policies only\]** Specifies information about the archive storage tier retention period.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RetentionArchiveTier" : RetentionArchiveTier
}

```

### YAML

```yaml

  RetentionArchiveTier:
    RetentionArchiveTier

```

## Properties

`RetentionArchiveTier`

Information about retention period in the Amazon EBS Snapshots Archive. For more information, see
[Archive Amazon \
EBS snapshots](../../../ec2/latest/windowsguide/snapshot-archive.md).

_Required_: Yes

_Type_: [RetentionArchiveTier](aws-properties-dlm-lifecyclepolicy-retentionarchivetier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Action

ArchiveRule

All content copied from https://docs.aws.amazon.com/.
