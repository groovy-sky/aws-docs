This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53RecoveryReadiness::Cell

Creates a cell in recovery group in Amazon Route 53 Application Recovery Controller. A cell in Route 53 ARC represents replicas or independent units of
failover in your application. It groups within it all the AWS resources that are necessary for your application to run independently.
Typically, you would have define one set of resources in a primary cell and another set in a standby cell in your recovery group.

After you set up the cells for your application, you can create readiness checks in Route 53 ARC to continually audit readiness for AWS
resource quotas, capacity, network routing policies, and other predefined rules.

You can set up notifications about changes that would affect your ability to fail over to a replica and recover. However, you should make decisions about whether to
fail away from or to a replica based on your monitoring and health check systems. You should consider readiness checks as a complementary service to those systems.

Route 53 ARC Readiness supports us-east-1 and us-west-2 AWS Regions only.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Route53RecoveryReadiness::Cell",
  "Properties" : {
      "CellName" : String,
      "Cells" : [ String, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Route53RecoveryReadiness::Cell
Properties:
  CellName: String
  Cells:
    - String
  Tags:
    - Tag

```

## Properties

`CellName`

The name of the cell to create.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9_]+`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Cells`

A list of cell Amazon Resource Names (ARNs) contained within this cell, for use in nested cells. For example, Availability Zones within specific AWS Regions.

_Required_: No

_Type_: Array of String

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A collection of tags associated with a resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-route53recoveryreadiness-cell-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `CellName`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CellArn`

The ARN of the cell.

`ParentReadinessScopes`

The readiness scope for the cell, which can be the Amazon Resource Name (ARN) of a cell or the ARN of a recovery group.
Although this is a list, it can currently have only one element.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon Route 53 Recovery Readiness

Tag
