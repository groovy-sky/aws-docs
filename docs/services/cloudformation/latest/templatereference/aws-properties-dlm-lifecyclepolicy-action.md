This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DLM::LifecyclePolicy Action

**\[Event-based policies only\]** Specifies an action for an event-based policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CrossRegionCopy" : [ CrossRegionCopyAction, ... ],
  "Name" : String
}

```

### YAML

```yaml

  CrossRegionCopy:
    - CrossRegionCopyAction
  Name: String

```

## Properties

`CrossRegionCopy`

The rule for copying shared snapshots across Regions.

_Required_: Yes

_Type_: Array of [CrossRegionCopyAction](aws-properties-dlm-lifecyclepolicy-crossregioncopyaction.md)

_Minimum_: `0`

_Maximum_: `3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A descriptive name for the action.

_Required_: Yes

_Type_: String

_Pattern_: `[0-9A-Za-z _-]+`

_Minimum_: `0`

_Maximum_: `120`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::DLM::LifecyclePolicy

ArchiveRetainRule

All content copied from https://docs.aws.amazon.com/.
