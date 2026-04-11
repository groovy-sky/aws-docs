This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DLM::LifecyclePolicy CrossRegionCopyRule

**\[Custom snapshot and AMI policies only\]** Specifies a cross-Region copy rule for a snapshot and AMI policies.

###### Note

To specify a cross-Region copy action for event-based polices, use
[CrossRegionCopyAction](../../../../reference/dlm/latest/apireference/api-crossregioncopyaction.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CmkArn" : String,
  "CopyTags" : Boolean,
  "DeprecateRule" : CrossRegionCopyDeprecateRule,
  "Encrypted" : Boolean,
  "RetainRule" : CrossRegionCopyRetainRule,
  "Target" : String,
  "TargetRegion" : String
}

```

### YAML

```yaml

  CmkArn: String
  CopyTags: Boolean
  DeprecateRule:
    CrossRegionCopyDeprecateRule
  Encrypted: Boolean
  RetainRule:
    CrossRegionCopyRetainRule
  Target: String
  TargetRegion: String

```

## Properties

`CmkArn`

The Amazon Resource Name (ARN) of the AWS KMS key to use for EBS encryption. If this
parameter is not specified, the default KMS key for the account is used.

_Required_: No

_Type_: String

_Pattern_: `arn:aws(-[a-z]{1,4}){0,2}:kms:([a-z]+-){2,3}\d:\d+:key/.*`

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CopyTags`

Indicates whether to copy all user-defined tags from the source snapshot or AMI to the
cross-Region copy.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeprecateRule`

**\[Custom AMI policies only\]** The AMI deprecation rule for cross-Region AMI copies created by the rule.

_Required_: No

_Type_: [CrossRegionCopyDeprecateRule](aws-properties-dlm-lifecyclepolicy-crossregioncopydeprecaterule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Encrypted`

To encrypt a copy of an unencrypted snapshot if encryption by default is not enabled,
enable encryption using this parameter. Copies of encrypted snapshots are encrypted,
even if this parameter is false or if encryption by default is not enabled.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetainRule`

The retention rule that indicates how long the cross-Region snapshot or AMI copies are
to be retained in the destination Region.

_Required_: No

_Type_: [CrossRegionCopyRetainRule](aws-properties-dlm-lifecyclepolicy-crossregioncopyretainrule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Target`

###### Note

Use this parameter for snapshot policies only. For AMI policies, use
**TargetRegion** instead.

**\[Custom snapshot policies only\]** The target Region or the Amazon Resource Name (ARN) of the target Outpost for the
snapshot copies.

_Required_: No

_Type_: String

_Pattern_: `^[\w:\-\/\*]+$`

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetRegion`

###### Note

Use this parameter for AMI policies only. For snapshot policies, use
**Target** instead. For snapshot policies
created before the **Target** parameter
was introduced, this parameter indicates the target Region for snapshot
copies.

**\[Custom AMI policies only\]** The target Region or the Amazon Resource Name (ARN) of the target Outpost for the
snapshot copies.

_Required_: No

_Type_: String

_Pattern_: `([a-z]+-){2,3}\d`

_Minimum_: `0`

_Maximum_: `16`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CrossRegionCopyRetainRule

CrossRegionCopyTarget

All content copied from https://docs.aws.amazon.com/.
