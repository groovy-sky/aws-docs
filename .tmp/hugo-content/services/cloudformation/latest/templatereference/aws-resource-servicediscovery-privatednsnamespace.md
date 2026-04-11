This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ServiceDiscovery::PrivateDnsNamespace

Creates a private namespace based on DNS, which is visible only inside a specified Amazon
VPC. The namespace defines your service naming scheme. For example, if you name your namespace
`example.com` and name your service `backend`, the resulting DNS name for
the service is `backend.example.com`. Service instances that are registered using a
private DNS namespace can be discovered using either a `DiscoverInstances` request or
using DNS. For the current quota on the number of namespaces that you can create using the same
AWS account, see [AWS Cloud Map quotas](../../../cloud-map/latest/dg/cloud-map-limits.md) in the
_AWS Cloud Map Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ServiceDiscovery::PrivateDnsNamespace",
  "Properties" : {
      "Description" : String,
      "Name" : String,
      "Properties" : Properties,
      "Tags" : [ Tag, ... ],
      "Vpc" : String
    }
}

```

### YAML

```yaml

Type: AWS::ServiceDiscovery::PrivateDnsNamespace
Properties:
  Description: String
  Name: String
  Properties:
    Properties
  Tags:
    - Tag
  Vpc: String

```

## Properties

`Description`

A description for the namespace.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name that you want to assign to this namespace. When you create a private DNS namespace,
AWS Cloud Map automatically creates an Amazon Route 53 private hosted zone that has the same name as the
namespace.

_Required_: Yes

_Type_: String

_Pattern_: `^(?!arn:)[!-~]{1,253}$`

_Maximum_: `253`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Properties`

Properties for the
private DNS namespace.

_Required_: No

_Type_: [Properties](aws-properties-servicediscovery-privatednsnamespace-properties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags for the namespace. Each tag consists of a key and an optional value, both of
which you define. Tag keys can have a maximum character length of 128 characters, and tag
values can have a maximum length of 256 characters.

_Required_: No

_Type_: Array of [Tag](aws-properties-servicediscovery-privatednsnamespace-tag.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: Updates are not supported.

`Vpc`

The ID of the Amazon VPC that you want to associate the namespace with.

_Required_: Yes

_Type_: String

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the value of `Id` for the namespace, such as
`ns-e4anhexample0004`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the private namespace.

`HostedZoneId`

The ID for the Route 53 hosted zone that AWS Cloud Map creates when you create a
namespace.

`Id`

The ID of the private namespace.

## Examples

### Create a private DNS namespace

The following example creates a private DNS namespace named
`private-example.com`.

#### JSON

```json

{
    "Resources": {
        "PrivateDnsNamespace": {
            "Type": "AWS::ServiceDiscovery::PrivateDnsNamespace",
            "Properties": {
                "Description": "AWS Cloud Map private DNS namespace for resources for example.com website",
                "Vpc": "vpc-12345678",
                "Name": "private-example.com",
                "Properties": {
                    "DnsProperties": {
                        "SOA": {
                            "TTL": 100
                        }
                    }
                }
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
  PrivateDnsNamespace:
    Type: AWS::ServiceDiscovery::PrivateDnsNamespace
    Properties:
      Description: AWS Cloud Map private DNS namespace for resources for example.com website
      Vpc: vpc-12345678
      Name: private-example.com
      Properties:
        DnsProperties:
          SOA:
            TTL: 100
```

## See also

- [CreatePrivateDnsNamespace](../../../cloud-map/latest/api/api-createprivatednsnamespace.md) in the _AWS Cloud Map API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ServiceDiscovery::Instance

PrivateDnsPropertiesMutable

All content copied from https://docs.aws.amazon.com/.
