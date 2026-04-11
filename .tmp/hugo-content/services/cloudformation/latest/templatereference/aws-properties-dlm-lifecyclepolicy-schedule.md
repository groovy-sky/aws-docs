This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DLM::LifecyclePolicy Schedule

**\[Custom snapshot and AMI policies only\]** Specifies a schedule for a snapshot or AMI lifecycle policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ArchiveRule" : ArchiveRule,
  "CopyTags" : Boolean,
  "CreateRule" : CreateRule,
  "CrossRegionCopyRules" : [ CrossRegionCopyRule, ... ],
  "DeprecateRule" : DeprecateRule,
  "FastRestoreRule" : FastRestoreRule,
  "Name" : String,
  "RetainRule" : RetainRule,
  "ShareRules" : [ ShareRule, ... ],
  "TagsToAdd" : [ Tag, ... ],
  "VariableTags" : [ Tag, ... ]
}

```

### YAML

```yaml

  ArchiveRule:
    ArchiveRule
  CopyTags: Boolean
  CreateRule:
    CreateRule
  CrossRegionCopyRules:
    - CrossRegionCopyRule
  DeprecateRule:
    DeprecateRule
  FastRestoreRule:
    FastRestoreRule
  Name: String
  RetainRule:
    RetainRule
  ShareRules:
    - ShareRule
  TagsToAdd:
    - Tag
  VariableTags:
    - Tag

```

## Properties

`ArchiveRule`

**\[Custom snapshot policies that target volumes only\]** The snapshot archiving rule for the schedule. When you specify an archiving
rule, snapshots are automatically moved from the standard tier to the archive tier once the schedule's
retention threshold is met. Snapshots are then retained in the archive tier for the archive retention
period that you specify.

For more information about using snapshot archiving, see [Considerations for \
snapshot lifecycle policies](../../../ec2/latest/userguide/snapshot-ami-policy.md#dlm-archive).

_Required_: No

_Type_: [ArchiveRule](aws-properties-dlm-lifecyclepolicy-archiverule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CopyTags`

Copy all user-defined tags on a source volume to snapshots of the volume created by
this policy.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateRule`

The creation rule.

_Required_: No

_Type_: [CreateRule](aws-properties-dlm-lifecyclepolicy-createrule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CrossRegionCopyRules`

Specifies a rule for copying snapshots or AMIs across Regions.

###### Note

You can't specify cross-Region copy rules for policies that create snapshots on an
Outpost or in a Local Zone. If the policy creates snapshots in a Region, then snapshots
can be copied to up to three Regions or Outposts.

_Required_: No

_Type_: Array of [CrossRegionCopyRule](aws-properties-dlm-lifecyclepolicy-crossregioncopyrule.md)

_Minimum_: `0`

_Maximum_: `3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeprecateRule`

**\[Custom AMI policies only\]** The AMI deprecation rule for the schedule.

_Required_: No

_Type_: [DeprecateRule](aws-properties-dlm-lifecyclepolicy-deprecaterule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FastRestoreRule`

**\[Custom snapshot policies only\]** The rule for enabling fast snapshot restore.

_Required_: No

_Type_: [FastRestoreRule](aws-properties-dlm-lifecyclepolicy-fastrestorerule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the schedule.

_Required_: No

_Type_: String

_Pattern_: `[0-9A-Za-z _-]+`

_Minimum_: `0`

_Maximum_: `120`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetainRule`

The retention rule for snapshots or AMIs created by the policy.

_Required_: No

_Type_: [RetainRule](aws-properties-dlm-lifecyclepolicy-retainrule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareRules`

**\[Custom snapshot policies only\]** The rule for sharing snapshots with other AWS accounts.

_Required_: No

_Type_: Array of [ShareRule](aws-properties-dlm-lifecyclepolicy-sharerule.md)

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TagsToAdd`

The tags to apply to policy-created resources. These user-defined tags are in addition
to the AWS-added lifecycle tags.

_Required_: No

_Type_: Array of [Tag](aws-properties-dlm-lifecyclepolicy-tag.md)

_Minimum_: `0`

_Maximum_: `45`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VariableTags`

**\[AMI policies and snapshot policies that target instances only\]**
A collection of key/value pairs with values determined dynamically when the policy is
executed. Keys may be any valid Amazon EC2 tag key. Values must be in one of the two
following formats: `$(instance-id)` or `$(timestamp)`. Variable
tags are only valid for EBS Snapshot Management – Instance policies.

_Required_: No

_Type_: Array of [Tag](aws-properties-dlm-lifecyclepolicy-tag.md)

_Minimum_: `0`

_Maximum_: `45`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RetentionArchiveTier

Script

All content copied from https://docs.aws.amazon.com/.
