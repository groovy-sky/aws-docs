This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::QuotaShare

Creates an AWS Batch quota share. Each quota share operates as a virtual queue with a configured compute capacity, resource sharing strategy, and borrow limits.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Batch::QuotaShare",
  "Properties" : {
      "CapacityLimits" : [ QuotaShareCapacityLimit, ... ],
      "JobQueue" : String,
      "PreemptionConfiguration" : QuotaSharePreemptionConfiguration,
      "QuotaShareName" : String,
      "ResourceSharingConfiguration" : QuotaShareResourceSharingConfiguration,
      "State" : String,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::Batch::QuotaShare
Properties:
  CapacityLimits:
    - QuotaShareCapacityLimit
  JobQueue: String
  PreemptionConfiguration:
    QuotaSharePreemptionConfiguration
  QuotaShareName: String
  ResourceSharingConfiguration:
    QuotaShareResourceSharingConfiguration
  State: String
  Tags:
    Key: Value

```

## Properties

`CapacityLimits`

A list that specifies the quantity and type of compute capacity allocated to the quota share.

_Required_: Yes

_Type_: Array of [QuotaShareCapacityLimit](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-batch-quotashare-quotasharecapacitylimit.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JobQueue`

Property description not available.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PreemptionConfiguration`

Specifies the preemption behavior for jobs in a quota share.

_Required_: Yes

_Type_: [QuotaSharePreemptionConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-batch-quotashare-quotasharepreemptionconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QuotaShareName`

The name of the quota share.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceSharingConfiguration`

Specifies whether a quota share reserves, lends, or both lends and borrows idle compute capacity.

_Required_: Yes

_Type_: [QuotaShareResourceSharingConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-batch-quotashare-quotashareresourcesharingconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`State`

The state of the quota share.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Object of String

_Pattern_: `.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`QuotaShareArn`

The Amazon Resource Name (ARN) of the quota share.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ServiceEnvironmentOrder

QuotaShareCapacityLimit
