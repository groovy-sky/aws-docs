This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::NetworkInsightsAccessScope

Describes a Network Access Scope. A Network Access Scope defines outbound (egress) and inbound (ingress)
traffic patterns, including sources, destinations, paths, and traffic types.

Network Access Analyzer identifies unintended network access to your resources on
AWS. When you start an analysis on a Network Access Scope, Network
Access Analyzer produces findings. For more information, see the [Network Access Analyzer\
User Guide](../../../vpc/latest/network-access-analyzer.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::NetworkInsightsAccessScope",
  "Properties" : {
      "ExcludePaths" : [ AccessScopePathRequest, ... ],
      "MatchPaths" : [ AccessScopePathRequest, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EC2::NetworkInsightsAccessScope
Properties:
  ExcludePaths:
    - AccessScopePathRequest
  MatchPaths:
    - AccessScopePathRequest
  Tags:
    - Tag

```

## Properties

`ExcludePaths`

The paths to exclude.

_Required_: No

_Type_: Array of [AccessScopePathRequest](aws-properties-ec2-networkinsightsaccessscope-accessscopepathrequest.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MatchPaths`

The paths to match.

_Required_: No

_Type_: Array of [AccessScopePathRequest](aws-properties-ec2-networkinsightsaccessscope-accessscopepathrequest.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-networkinsightsaccessscope-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the network insights scope.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedDate`

The creation date.

`NetworkInsightsAccessScopeArn`

The ARN of the Network Access Scope.

`NetworkInsightsAccessScopeId`

The ID of the Network Access Scope.

`UpdatedDate`

The last updated date.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PortRange

AccessScopePathRequest

All content copied from https://docs.aws.amazon.com/.
