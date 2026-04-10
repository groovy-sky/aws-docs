This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppRunner::VpcConnector

The `AWS::AppRunner::VpcConnector` resource is an AWS App Runner resource type that specifies an App Runner VPC connector.

App Runner requires this resource when you want to associate your App Runner service to a custom Amazon Virtual Private Cloud (Amazon VPC).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppRunner::VpcConnector",
  "Properties" : {
      "SecurityGroups" : [ String, ... ],
      "Subnets" : [ String, ... ],
      "Tags" : [ Tag, ... ],
      "VpcConnectorName" : String
    }
}

```

### YAML

```yaml

Type: AWS::AppRunner::VpcConnector
Properties:
  SecurityGroups:
    - String
  Subnets:
    - String
  Tags:
    - Tag
  VpcConnectorName: String

```

## Properties

`SecurityGroups`

A list of IDs of security groups that App Runner should use for access to AWS resources under the specified subnets. If not specified, App Runner uses the
default security group of the Amazon VPC. The default security group allows all outbound traffic.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Subnets`

A list of IDs of subnets that App Runner should use when it associates your service with a custom Amazon VPC. Specify IDs of subnets of a single
Amazon VPC. App Runner determines the Amazon VPC from the subnets you specify.

###### Note

App Runner only supports subnets of IP address type _IPv4_ and _dual stack_ (IPv4 and IPv6).

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A list of metadata items that you can associate with your VPC connector resource. A tag is a key-value pair.

###### Note

A `VpcConnector` is immutable, so you cannot update its tags. To change the
tags, replace the resource. To replace a `VpcConnector`, you must provide a new
combination of security groups.

_Required_: No

_Type_: Array of [Tag](aws-properties-apprunner-vpcconnector-tag.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VpcConnectorName`

A name for the VPC connector.

If you don't specify a name, CloudFormation generates a name for your VPC connector.

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z0-9][A-Za-z0-9-\\_]{3,39}$`

_Minimum_: `4`

_Maximum_: `40`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`VpcConnectorArn`

The Amazon Resource Name (ARN) of this VPC connector.

`VpcConnectorRevision`

The revision of this VPC connector. It's unique among all the active connectors ( `"Status": "ACTIVE"`) that share the same
`Name`.

###### Note

At this time, App Runner supports only one revision per name.

## Examples

### VPC connector

This example illustrates creating a VPC connector with two subnets and two security groups.

#### JSON

```json

{
  "Type" : "AWS::AppRunner::VpcConnector",
  "Properties" : {
    "VpcConnectorName": "my-vpc-connector",
    "Subnets": ["subnet-123", "subnet-456"],
    "SecurityGroups": ["sg-123", "sg-456"]
  }
}
```

#### YAML

```yaml

Type: AWS::AppRunner::VpcConnector
Properties:
  VpcConnectorName: my-vpc-connector
  Subnets:
    - subnet-123
    - subnet-456
  SecurityGroups:
    - sg-123
    - sg-456
```

## See also

- [Enabling Amazon VPC access for your service](../../../apprunner/latest/dg/network-vpc.md) in the
_AWS App Runner Developer Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
