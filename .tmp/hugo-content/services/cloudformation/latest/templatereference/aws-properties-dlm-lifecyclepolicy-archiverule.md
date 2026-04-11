This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DLM::LifecyclePolicy ArchiveRule

**\[Custom snapshot policies only\]** Specifies a snapshot archiving rule for a schedule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RetainRule" : ArchiveRetainRule
}

```

### YAML

```yaml

  RetainRule:
    ArchiveRetainRule

```

## Properties

`RetainRule`

Information about the retention period for the snapshot archiving rule.

_Required_: Yes

_Type_: [ArchiveRetainRule](aws-properties-dlm-lifecyclepolicy-archiveretainrule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ArchiveRetainRule

CreateRule

All content copied from https://docs.aws.amazon.com/.
