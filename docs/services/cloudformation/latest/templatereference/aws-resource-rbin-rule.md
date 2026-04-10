This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Rbin::Rule

Creates a Recycle Bin retention rule. You can create two types of retention rules:

- **Tag-level retention rules** \- These retention rules use
resource tags to identify the resources to protect. For each retention rule, you specify one or
more tag key and value pairs. Resources (of the specified type) that have at least one of these
tag key and value pairs are automatically retained in the Recycle Bin upon deletion. Use this
type of retention rule to protect specific resources in your account based on their tags.

- **Region-level retention rules** \- These retention rules,
by default, apply to all of the resources (of the specified type) in the Region, even if the
resources are not tagged. However, you can specify exclusion tags to exclude resources that have
specific tags. Use this type of retention rule to protect all resources of a specific type in a
Region.

For more information, see [Create Recycle Bin retention rules](../../../ebs/latest/userguide/recycle-bin.md) in the _Amazon EBS User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Rbin::Rule",
  "Properties" : {
      "Description" : String,
      "ExcludeResourceTags" : [ ResourceTag, ... ],
      "LockConfiguration" : UnlockDelay,
      "ResourceTags" : [ ResourceTag, ... ],
      "ResourceType" : String,
      "RetentionPeriod" : RetentionPeriod,
      "Status" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Rbin::Rule
Properties:
  Description: String
  ExcludeResourceTags:
    - ResourceTag
  LockConfiguration:
    UnlockDelay
  ResourceTags:
    - ResourceTag
  ResourceType: String
  RetentionPeriod:
    RetentionPeriod
  Status: String
  Tags:
    - Tag

```

## Properties

`Description`

The retention rule description.

_Required_: No

_Type_: String

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExcludeResourceTags`

\[Region-level retention rules only\] Specifies the exclusion tags to use to identify resources that are to be excluded,
or ignored, by a Region-level retention rule. Resources that have any of these tags are not retained by the retention rule
upon deletion.

You can't specify exclusion tags for tag-level retention rules.

_Required_: No

_Type_: Array of [ResourceTag](aws-properties-rbin-rule-resourcetag.md)

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LockConfiguration`

Information about the retention rule lock configuration.

_Required_: No

_Type_: [UnlockDelay](aws-properties-rbin-rule-unlockdelay.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceTags`

\[Tag-level retention rules only\] Specifies the resource tags to use to identify resources that are to be retained by a
tag-level retention rule. For tag-level retention rules, only deleted resources, of the specified resource type, that
have one or more of the specified tag key and value pairs are retained. If a resource is deleted, but it does not have
any of the specified tag key and value pairs, it is immediately deleted without being retained by the retention rule.

You can add the same tag key and value pair to a maximum or five retention rules.

To create a Region-level retention rule, omit this parameter. A Region-level retention rule
does not have any resource tags specified. It retains all deleted resources of the specified
resource type in the Region in which the rule is created, even if the resources are not tagged.

_Required_: No

_Type_: Array of [ResourceTag](aws-properties-rbin-rule-resourcetag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceType`

The resource type to be retained by the retention rule. Currently, only EBS volumes, EBS snapshots, and EBS-backed AMIs
are supported.

- To retain EBS volumes, specify `EBS_VOLUME`.

- To retain EBS snapshots, specify `EBS_SNAPSHOT`

- To retain EBS-backed AMIs, specify `EC2_IMAGE`.

_Required_: Yes

_Type_: String

_Allowed values_: `EBS_SNAPSHOT | EC2_IMAGE | EBS_VOLUME`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RetentionPeriod`

Information about the retention period for which the retention rule is to
retain resources.

_Required_: Yes

_Type_: [RetentionPeriod](aws-properties-rbin-rule-retentionperiod.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The state of the retention rule. Only retention rules that are in the `available`
state retain resources.

_Required_: No

_Type_: String

_Pattern_: `pending|available`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Information about the tags to assign to the retention rule.

_Required_: No

_Type_: Array of [Tag](aws-properties-rbin-rule-tag.md)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the retention rule.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the retention rule.

`Identifier`

The unique ID of the retention rule.

`LockState`

\[Region-level retention rules only\] The lock state for the retention rule.

- `locked` \- The retention rule is locked and can't be modified or deleted.

- `pending_unlock` \- The retention rule has been unlocked but it is still within
the unlock delay period. The retention rule can be modified or deleted only after the unlock
delay period has expired.

- `unlocked` \- The retention rule is unlocked and it can be modified or deleted by
any user with the required permissions.

- `null` \- The retention rule has never been locked. Once a retention rule has
been locked, it can transition between the `locked` and `unlocked` states
only; it can never transition back to `null`.

## Examples

### Create a retention rule

This example creates a retention rule that retains EBS snapshots for 7 days after deletion.

#### JSON

```json

"MyRule" : {
  "Type": "AWS::Rbin::Rule",
  "Properties": {
    "Description": "my new retention rule",
    "LockConfiguration": {
      "UnlockDelayUnit": "DAYS",
      "UnlockDelayValue": 15
    },
    "ResourceTags": [{
      "Key": "environment",
      "Value": "testing"
    }],
    "ResourceType": "EBS_SNAPSHOT",
    "RetentionPeriod": {
      "RetentionPeriodUnit": "DAYS",
      "RetentionPeriodValue": 7
    },
    "Status": "available",
    "Tags": [{
      "Key": "purpose",
      "Value": "testing"
    }]
  }
}
```

#### YAML

```yaml

MyRule:
  Type: AWS::Rbin::Rule
  Properties:
    Description: my new retention rule
    LockConfiguration:
      UnlockDelayUnit: DAYS
      UnlockDelayValue: 15
    ResourceTags:
    - Key: environment
      Value: testing
    ResourceType: EBS_SNAPSHOT
    RetentionPeriod:
      RetentionPeriodUnit: DAYS
      RetentionPeriodValue: 7
    Status: available
    Tags:
    - Key: purpose
      Value: testing
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Recycle Bin

ResourceTag

All content copied from https://docs.aws.amazon.com/.
