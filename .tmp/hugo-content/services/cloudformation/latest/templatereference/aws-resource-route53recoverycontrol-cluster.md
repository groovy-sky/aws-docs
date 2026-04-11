This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53RecoveryControl::Cluster

Creates a cluster in Amazon Route 53 Application Recovery Controller. A cluster is a set of redundant Regional endpoints that you can run Route 53 ARC API calls against to update or get the state
of one or more routing controls.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Route53RecoveryControl::Cluster",
  "Properties" : {
      "Name" : String,
      "NetworkType" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Route53RecoveryControl::Cluster
Properties:
  Name: String
  NetworkType: String
  Tags:
    - Tag

```

## Properties

`Name`

Name of the cluster. You can use any non-white space character in the name except the following: & > < ' (single quote) " (double quote) ; (semicolon).

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetworkType`

The network-type can either be IPV4 or DUALSTACK.

_Required_: No

_Type_: String

_Allowed values_: `IPV4 | DUALSTACK`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags associated with the cluster.

_Required_: No

_Type_: Array of [Tag](aws-properties-route53recoverycontrol-cluster-tag.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `ClusterArn` object.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ClusterArn`

The Amazon Resource Name (ARN) of the cluster.

`ClusterEndpoints`

An array of endpoints for the cluster. You specify one of these endpoints when you want to set or retrieve a routing control state
in the cluster.

`Status`

The deployment status of the cluster. Status can be one of the following: PENDING, DEPLOYED, PENDING\_DELETION.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Route 53 Recovery Control

ClusterEndpoint

All content copied from https://docs.aws.amazon.com/.
