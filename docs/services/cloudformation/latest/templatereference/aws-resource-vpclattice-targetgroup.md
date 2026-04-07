This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VpcLattice::TargetGroup

Creates a target group. A target group is a collection of targets, or compute resources,
that run your application or service. A target group can only be used by a single service.

For more information, see [Target groups](https://docs.aws.amazon.com/vpc-lattice/latest/ug/target-groups.html) in the
_Amazon VPC Lattice User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::VpcLattice::TargetGroup",
  "Properties" : {
      "Config" : TargetGroupConfig,
      "Name" : String,
      "Tags" : [ Tag, ... ],
      "Targets" : [ Target, ... ],
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::VpcLattice::TargetGroup
Properties:
  Config:
    TargetGroupConfig
  Name: String
  Tags:
    - Tag
  Targets:
    - Target
  Type: String

```

## Properties

`Config`

The target group configuration.

_Required_: No

_Type_: [TargetGroupConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-vpclattice-targetgroup-targetgroupconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the target group. The name must be unique within the account. The valid
characters are a-z, 0-9, and hyphens (-). You can't use a hyphen as the first or last
character, or immediately after another hyphen.

If you don't specify a name, CloudFormation generates one. However, if
you specify a name, and later want to replace the resource, you must specify a new
name.

_Required_: No

_Type_: String

_Pattern_: `^(?!tg-)(?![-])(?!.*[-]$)(?!.*[-]{2})[a-z0-9-]+$`

_Minimum_: `3`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags for the target group.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-vpclattice-targetgroup-tag.html)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Targets`

Describes a target.

_Required_: No

_Type_: Array of [Target](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-vpclattice-targetgroup-target.html)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of target group.

_Required_: Yes

_Type_: String

_Allowed values_: `IP | LAMBDA | INSTANCE | ALB`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the target group.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the target group.

`CreatedAt`

The date and time that the target group was created, specified in ISO-8601 format.

`Id`

The ID of the target group.

`LastUpdatedAt`

The date and time that the target group was last updated, specified in ISO-8601
format.

`Status`

The operation's status. You can retry the operation if the status is
`CREATE_FAILED`. However, if you retry it while the status is
`CREATE_IN_PROGRESS`, there is no change in the status.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

HealthCheckConfig
