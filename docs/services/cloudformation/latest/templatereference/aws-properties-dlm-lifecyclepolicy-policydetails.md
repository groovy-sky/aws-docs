This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DLM::LifecyclePolicy PolicyDetails

Specifies the configuration of a lifecycle policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Actions" : [ Action, ... ],
  "CopyTags" : Boolean,
  "CreateInterval" : Integer,
  "CrossRegionCopyTargets" : [ CrossRegionCopyTarget, ... ],
  "EventSource" : EventSource,
  "Exclusions" : Exclusions,
  "ExtendDeletion" : Boolean,
  "Parameters" : Parameters,
  "PolicyLanguage" : String,
  "PolicyType" : String,
  "ResourceLocations" : [ String, ... ],
  "ResourceType" : String,
  "ResourceTypes" : [ String, ... ],
  "RetainInterval" : Integer,
  "Schedules" : [ Schedule, ... ],
  "TargetTags" : [ Tag, ... ]
}

```

### YAML

```yaml

  Actions:
    - Action
  CopyTags: Boolean
  CreateInterval: Integer
  CrossRegionCopyTargets:
    - CrossRegionCopyTarget
  EventSource:
    EventSource
  Exclusions:
    Exclusions
  ExtendDeletion: Boolean
  Parameters:
    Parameters
  PolicyLanguage: String
  PolicyType: String
  ResourceLocations:
    - String
  ResourceType: String
  ResourceTypes:
    - String
  RetainInterval: Integer
  Schedules:
    - Schedule
  TargetTags:
    - Tag

```

## Properties

`Actions`

**\[Event-based policies only\]** The actions to be performed when the event-based policy is activated. You can specify
only one action per policy.

_Required_: No

_Type_: Array of [Action](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dlm-lifecyclepolicy-action.html)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CopyTags`

**\[Default policies only\]** Indicates whether the policy should copy tags from the source resource
to the snapshot or AMI. If you do not specify a value, the default is `false`.

Default: false

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateInterval`

**\[Default policies only\]** Specifies how often the policy should run and create snapshots or AMIs.
The creation frequency can range from 1 to 7 days. If you do not specify a value, the
default is 1.

Default: 1

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CrossRegionCopyTargets`

**\[Default policies only\]** Specifies destination Regions for snapshot or AMI copies. You can specify
up to 3 destination Regions. If you do not want to create cross-Region copies, omit this
parameter.

_Required_: No

_Type_: Array of [CrossRegionCopyTarget](aws-properties-dlm-lifecyclepolicy-crossregioncopytarget.md)

_Minimum_: `0`

_Maximum_: `3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventSource`

**\[Event-based policies only\]** The event that activates the event-based policy.

_Required_: No

_Type_: [EventSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dlm-lifecyclepolicy-eventsource.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Exclusions`

**\[Default policies only\]** Specifies exclusion parameters for volumes or instances for which you
do not want to create snapshots or AMIs. The policy will not create snapshots or AMIs
for target resources that match any of the specified exclusion parameters.

_Required_: No

_Type_: [Exclusions](aws-properties-dlm-lifecyclepolicy-exclusions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExtendDeletion`

**\[Default policies only\]** Defines the snapshot or AMI retention behavior for the policy if the
source volume or instance is deleted, or if the policy enters the error, disabled, or
deleted state.

By default ( **ExtendDeletion=false**):

- If a source resource is deleted, Amazon Data Lifecycle Manager will continue to delete previously
created snapshots or AMIs, up to but not including the last one, based on the
specified retention period. If you want Amazon Data Lifecycle Manager to delete all snapshots or AMIs,
including the last one, specify `true`.

- If a policy enters the error, disabled, or deleted state, Amazon Data Lifecycle Manager stops deleting
snapshots and AMIs. If you want Amazon Data Lifecycle Manager to continue deleting snapshots or AMIs,
including the last one, if the policy enters one of these states, specify
`true`.

If you enable extended deletion ( **ExtendDeletion=true**),
you override both default behaviors simultaneously.

If you do not specify a value, the default is `false`.

Default: false

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Parameters`

**\[Custom snapshot and AMI policies only\]** A set of optional parameters for snapshot and AMI lifecycle policies.

###### Note

If you are modifying a policy that was created or previously modified using the Amazon
Data Lifecycle Manager console, then you must include this parameter and specify either
the default values or the new values that you require. You can't omit this parameter or
set its values to null.

_Required_: No

_Type_: [Parameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dlm-lifecyclepolicy-parameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyLanguage`

The type of policy to create. Specify one of the following:

- `SIMPLIFIED` To create a default policy.

- `STANDARD` To create a custom policy.

_Required_: No

_Type_: String

_Allowed values_: `SIMPLIFIED | STANDARD`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyType`

The type of policy. Specify `EBS_SNAPSHOT_MANAGEMENT`
to create a lifecycle policy that manages the lifecycle of Amazon EBS snapshots. Specify `IMAGE_MANAGEMENT`
to create a lifecycle policy that manages the lifecycle of EBS-backed AMIs. Specify `EVENT_BASED_POLICY `
to create an event-based policy that performs specific actions when a defined event occurs in your AWS account.

The default is `EBS_SNAPSHOT_MANAGEMENT`.

_Required_: No

_Type_: String

_Allowed values_: `EBS_SNAPSHOT_MANAGEMENT | IMAGE_MANAGEMENT | EVENT_BASED_POLICY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceLocations`

**\[Custom snapshot and AMI policies only\]** The location of the resources to backup.

- If the source resources are located in a Region, specify `CLOUD`. In this case,
the policy targets all resources of the specified type with matching target tags across all
Availability Zones in the Region.

- **\[Custom snapshot policies only\]** If the source resources are located in a Local Zone, specify `LOCAL_ZONE`.
In this case, the policy targets all resources of the specified type with matching target
tags across all Local Zones in the Region.

- If the source resources are located on an Outpost in your account, specify `OUTPOST`.
In this case, the policy targets all resources of the specified type with matching target
tags across all of the Outposts in your account.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceType`

**\[Default policies only\]** Specify the type of default policy to create.

- To create a default policy for EBS snapshots, that creates snapshots of all volumes in the
Region that do not have recent backups, specify `VOLUME`.

- To create a default policy for EBS-backed AMIs, that creates EBS-backed
AMIs from all instances in the Region that do not have recent backups, specify
`INSTANCE`.

_Required_: No

_Type_: String

_Allowed values_: `VOLUME | INSTANCE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceTypes`

**\[Custom snapshot policies only\]** The target resource type for snapshot and AMI lifecycle policies. Use `VOLUME ` to
create snapshots of individual volumes or use `INSTANCE` to create multi-volume
snapshots from the volumes for an instance.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetainInterval`

**\[Default policies only\]** Specifies how long the policy should retain snapshots or AMIs before
deleting them. The retention period can range from 2 to 14 days, but it must be greater
than the creation frequency to ensure that the policy retains at least 1 snapshot or
AMI at any given time. If you do not specify a value, the default is 7.

Default: 7

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Schedules`

**\[Custom snapshot and AMI policies only\]** The schedules of policy-defined actions for snapshot and AMI lifecycle policies. A policy
can have up to four schedules—one mandatory schedule and up to three optional schedules.

_Required_: No

_Type_: Array of [Schedule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dlm-lifecyclepolicy-schedule.html)

_Minimum_: `1`

_Maximum_: `4`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetTags`

**\[Custom snapshot and AMI policies only\]** The single tag that identifies targeted resources for this policy.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dlm-lifecyclepolicy-tag.html)

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [PolicyDetails](https://docs.aws.amazon.com/dlm/latest/APIReference/API_PolicyDetails.html) in the
_Amazon Data Lifecycle Manager API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Parameters

RetainRule
