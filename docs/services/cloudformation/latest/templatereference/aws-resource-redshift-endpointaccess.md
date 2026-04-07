This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Redshift::EndpointAccess

Creates a Redshift-managed VPC endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Redshift::EndpointAccess",
  "Properties" : {
      "ClusterIdentifier" : String,
      "EndpointName" : String,
      "ResourceOwner" : String,
      "SubnetGroupName" : String,
      "VpcSecurityGroupIds" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Redshift::EndpointAccess
Properties:
  ClusterIdentifier: String
  EndpointName: String
  ResourceOwner: String
  SubnetGroupName: String
  VpcSecurityGroupIds:
    - String

```

## Properties

`ClusterIdentifier`

The cluster identifier of the cluster associated with the endpoint.

_Required_: Yes

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EndpointName`

The name of the endpoint.

_Required_: Yes

_Type_: String

_Pattern_: `^(?=^[a-z][a-z0-9]*(-[a-z0-9]+)*$).{1,30}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceOwner`

The AWS account ID of the owner of the cluster.

_Required_: No

_Type_: String

_Pattern_: `^\d{12}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetGroupName`

The subnet group name where Amazon Redshift chooses to deploy the endpoint.

_Required_: Yes

_Type_: String

_Pattern_: `^(?=^[a-zA-Z0-9-]+$).{1,255}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VpcSecurityGroupIds`

The security group that defines the ports, protocols, and sources for inbound traffic that you are authorizing into your endpoint.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Address`

The DNS address of the endpoint.

`EndpointCreateTime`

The time (UTC) that the endpoint was created.

`EndpointStatus`

The status of the endpoint.

`Port`

The port number on which the cluster accepts incoming connections.

`VpcEndpoint.VpcEndpointId`

The connection endpoint ID for connecting an Amazon Redshift cluster through the proxy.

`VpcEndpoint.VpcId`

The VPC identifier that the endpoint is associated.

`VpcSecurityGroups`

The security groups associated with the endpoint.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

NetworkInterface
