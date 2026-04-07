This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53RecoveryReadiness::RecoveryGroup

Creates a recovery group in Amazon Route 53 Application Recovery Controller. A recovery group represents your application.
It typically consists of two or more cells that are replicas of each other in terms of resources and functionality, so that you can fail over
from one to the other, for example, from one Region to another. You create recovery groups so you can use readiness checks to audit resources
in your application.

For more information, see
[Readiness checks, resource sets, \
and readiness scopes](https://docs.aws.amazon.com/r53recovery/latest/dg/recovery-readiness.recovery-groups.readiness-scope.html) in the Amazon Route 53 Application Recovery Controller Developer Guide.

Route 53 ARC Readiness supports us-east-1 and us-west-2 AWS Regions only.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Route53RecoveryReadiness::RecoveryGroup",
  "Properties" : {
      "Cells" : [ String, ... ],
      "RecoveryGroupName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Route53RecoveryReadiness::RecoveryGroup
Properties:
  Cells:
    - String
  RecoveryGroupName: String
  Tags:
    - Tag

```

## Properties

`Cells`

A list of the cell Amazon Resource Names (ARNs) in the recovery group.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `256 | 5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecoveryGroupName`

The name of the recovery group to create.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9_]+`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A collection of tags associated with a resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-route53recoveryreadiness-recoverygroup-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `RecoveryGroupName`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`RecoveryGroupArn`

The Amazon Resource Name (ARN) of the recovery group.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
