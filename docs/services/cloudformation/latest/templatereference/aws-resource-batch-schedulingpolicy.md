This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::SchedulingPolicy

The `AWS::Batch::SchedulingPolicy` resource specifies the parameters for an
AWS Batch scheduling policy. For more information, see [Scheduling\
Policies](https://docs.aws.amazon.com/batch/latest/userguide/scheduling_policies.html) in the _AWS Batch User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Batch::SchedulingPolicy",
  "Properties" : {
      "FairsharePolicy" : FairsharePolicy,
      "Name" : String,
      "QuotaSharePolicy" : QuotaSharePolicy,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::Batch::SchedulingPolicy
Properties:
  FairsharePolicy:
    FairsharePolicy
  Name: String
  QuotaSharePolicy:
    QuotaSharePolicy
  Tags:
    Key: Value

```

## Properties

`FairsharePolicy`

The fair-share scheduling policy details. Only one of fairsharePolicy or quotaSharePolicy can be set.
Once set, this policy type cannot be removed or changed to a quotaSharePolicy.

_Required_: No

_Type_: [FairsharePolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-batch-schedulingpolicy-fairsharepolicy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the fair-share scheduling policy. It can be up to 128 letters long. It can contain
uppercase and lowercase letters, numbers, hyphens (-), and underscores (\_).

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`QuotaSharePolicy`

The quota share scheduling policy details.

_Required_: No

_Type_: [QuotaSharePolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-batch-schedulingpolicy-quotasharepolicy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags that you apply to the scheduling policy to help you categorize and organize your
resources. Each tag consists of a key and an optional value. For more information, see [Tagging AWS\
Resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in _AWS General Reference_.

These tags can be updated or removed using the [TagResource](https://docs.aws.amazon.com/batch/latest/APIReference/API_TagResource.html) and [UntagResource](https://docs.aws.amazon.com/batch/latest/APIReference/API_UntagResource.html) API operations.

_Required_: No

_Type_: Object of String

_Pattern_: `.*`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the scheduling policy ARN, such as `arn:aws:batch:us-east-1:111122223333:scheduling-policy/HighPriority`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

Returns the scheduling policy ARN, such as `arn:aws:batch:us-east-1:111122223333:scheduling-policy/HighPriority`.

## See also

- [Scheduling\
Policy Parameters](https://docs.aws.amazon.com/batch/latest/userguide/scheduling_policy_parameters.html) in the _AWS Batch User Guide_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

QuotaShareResourceSharingConfiguration

FairsharePolicy
