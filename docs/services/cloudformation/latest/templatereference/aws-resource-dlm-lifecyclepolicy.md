This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DLM::LifecyclePolicy

Specifies a lifecycle policy, which is used to automate operations on Amazon EBS resources.

The properties are required when you add a lifecycle policy and optional when you update a lifecycle policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DLM::LifecyclePolicy",
  "Properties" : {
      "CopyTags" : Boolean,
      "CreateInterval" : Integer,
      "CrossRegionCopyTargets" : [ CrossRegionCopyTarget, ... ],
      "DefaultPolicy" : String,
      "Description" : String,
      "Exclusions" : Exclusions,
      "ExecutionRoleArn" : String,
      "ExtendDeletion" : Boolean,
      "PolicyDetails" : PolicyDetails,
      "RetainInterval" : Integer,
      "State" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DLM::LifecyclePolicy
Properties:
  CopyTags: Boolean
  CreateInterval: Integer
  CrossRegionCopyTargets:
    - CrossRegionCopyTarget
  DefaultPolicy: String
  Description: String
  Exclusions:
    Exclusions
  ExecutionRoleArn: String
  ExtendDeletion: Boolean
  PolicyDetails:
    PolicyDetails
  RetainInterval: Integer
  State: String
  Tags:
    - Tag

```

## Properties

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

_Type_: Array of [CrossRegionCopyTarget](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dlm-lifecyclepolicy-crossregioncopytarget.html)

_Minimum_: `0`

_Maximum_: `3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultPolicy`

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

`Description`

A description of the lifecycle policy. The characters ^\[0-9A-Za-z \_-\]+$ are
supported.

_Required_: Conditional

_Type_: String

_Pattern_: `[0-9A-Za-z _-]+`

_Minimum_: `0`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Exclusions`

**\[Default policies only\]** Specifies exclusion parameters for volumes or instances for which you
do not want to create snapshots or AMIs. The policy will not create snapshots or AMIs
for target resources that match any of the specified exclusion parameters.

_Required_: No

_Type_: [Exclusions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dlm-lifecyclepolicy-exclusions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExecutionRoleArn`

The Amazon Resource Name (ARN) of the IAM role used to run the operations specified by
the lifecycle policy.

_Required_: Conditional

_Type_: String

_Pattern_: `arn:aws(-[a-z]{1,3}){0,2}:iam::\d+:role/.*`

_Minimum_: `0`

_Maximum_: `2048`

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

`PolicyDetails`

The configuration details of the lifecycle policy.

###### Important

If you create a default policy, you can specify the request parameters either in
the request body, or in the PolicyDetails request structure, but not both.

_Required_: Conditional

_Type_: [PolicyDetails](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dlm-lifecyclepolicy-policydetails.html)

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

`State`

The activation state of the lifecycle policy.

_Required_: Conditional

_Type_: String

_Allowed values_: `ENABLED | DISABLED | ERROR`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to apply to the lifecycle policy during creation.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dlm-lifecyclepolicy-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the lifecycle policy.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the lifecycle policy.

## Examples

### Creating a Lifecycle Policy

The following example demonstrates how to create a basic snapshot lifecycle policy with a cross-Region copy rule.

#### JSON

```json

{
    "Description": "Basic LifecyclePolicy",
    "Resources": {
        "BasicLifecyclePolicy": {
            "Type": "AWS::DLM::LifecyclePolicy",
            "Properties": {
                "Description": "Lifecycle Policy using CloudFormation",
                "State": "ENABLED",
                "ExecutionRoleArn": "arn:aws:iam::123456789012:role/service-role/AWSDataLifecycleManagerDefaultRole",
                "PolicyDetails": {
                    "ResourceTypes": [
                        "VOLUME"
                    ],
                    "TargetTags": [{
                        "Key": "costcenter",
                        "Value": "115"
                    }],
                    "Schedules": [{
                        "Name": "Daily Snapshots",
                        "TagsToAdd": [{
                            "Key": "type",
                            "Value": "DailySnapshot"
                        }],
                        "CreateRule": {
                            "Interval": 12,
                            "IntervalUnit": "HOURS",
                            "Times": [
                                "13:00"
                            ]
                        },
                        "RetainRule": {
                            "Count": 1
                        },
                        "CopyTags": true,
                        "CrossRegionCopyRules": [{
                            "Encrypted": false,
                            "Target": "us-east-1"
                        }]
                   }]
               }
           }
        }
    }
}
```

#### YAML

```yaml

Description: Basic LifecyclePolicy
Resources:
  BasicLifecyclePolicy:
    Type: AWS::DLM::LifecyclePolicy
    Properties:
      Description: Lifecycle Policy using CloudFormation
      State: ENABLED
      ExecutionRoleArn: arn:aws:iam::123456789012:role/service-role/AWSDataLifecycleManagerDefaultRole
      PolicyDetails:
        ResourceTypes:
        - VOLUME
        TargetTags:
        - Key: costcenter
          Value: '115'
        Schedules:
        - Name: Daily Snapshots
          TagsToAdd:
          - Key: type
            Value: DailySnapshot
          CreateRule:
            Interval: 12
            IntervalUnit: HOURS
            Times:
            - '13:00'
          RetainRule:
            Count: 1
          CopyTags: true
          CrossRegionCopyRules:
          - Encrypted: false
            Target: us-east-1
```

## See also

- [CreateLifecyclePolicy](https://docs.aws.amazon.com/dlm/latest/APIReference/API_CreateLifecyclePolicy.html) in the
_Amazon Data Lifecycle Manager API Reference_

- [Automating the Amazon EBS Snapshot Lifecycle](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/snapshot-lifecycle.html)
in the _Amazon Elastic Compute Cloud User Guide_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon Data Lifecycle Manager

Action
