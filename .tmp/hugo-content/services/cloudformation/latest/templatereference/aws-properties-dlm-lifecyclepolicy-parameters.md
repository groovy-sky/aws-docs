This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DLM::LifecyclePolicy Parameters

**\[Custom snapshot and AMI policies only\]** Specifies optional parameters for snapshot and AMI policies. The
set of valid parameters depends on the combination of policy type and target resource
type.

If you choose to exclude boot volumes and you specify tags that consequently exclude
all of the additional data volumes attached to an instance, then Amazon Data Lifecycle Manager will not create
any snapshots for the affected instance, and it will emit a `SnapshotsCreateFailed`
Amazon CloudWatch metric. For more information, see [Monitor your policies \
using Amazon CloudWatch](../../../ec2/latest/userguide/monitor-dlm-cw-metrics.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExcludeBootVolume" : Boolean,
  "ExcludeDataVolumeTags" : [ Tag, ... ],
  "NoReboot" : Boolean
}

```

### YAML

```yaml

  ExcludeBootVolume: Boolean
  ExcludeDataVolumeTags:
    - Tag
  NoReboot: Boolean

```

## Properties

`ExcludeBootVolume`

**\[Custom snapshot policies that target instances only\]** Indicates whether to exclude the root volume from multi-volume
snapshot sets. The default is `false`. If you specify `true`,
then the root volumes attached to targeted instances will be excluded from the multi-volume
snapshot sets created by the policy.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExcludeDataVolumeTags`

**\[Custom snapshot policies that target instances only\]** The tags used to identify data (non-root) volumes to exclude from
multi-volume snapshot sets.

If you create a snapshot lifecycle policy that targets instances and you specify tags for
this parameter, then data volumes with the specified tags that are attached to targeted
instances will be excluded from the multi-volume snapshot sets created by the policy.

_Required_: No

_Type_: Array of [Tag](aws-properties-dlm-lifecyclepolicy-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NoReboot`

**\[Custom AMI policies only\]** Indicates whether targeted instances are rebooted when the lifecycle policy
runs. `true` indicates that targeted instances are not rebooted when the policy
runs. `false` indicates that target instances are rebooted when the policy runs.
The default is `true` (instances are not rebooted).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FastRestoreRule

PolicyDetails

All content copied from https://docs.aws.amazon.com/.
