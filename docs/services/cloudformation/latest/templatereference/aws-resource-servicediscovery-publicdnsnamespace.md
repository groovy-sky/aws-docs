This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ServiceDiscovery::PublicDnsNamespace

Creates a public namespace based on DNS, which is visible on the internet. The namespace
defines your service naming scheme. For example, if you name your namespace
`example.com` and name your service `backend`, the resulting DNS name for
the service is `backend.example.com`. You can discover instances that were registered
with a public DNS namespace by using either a `DiscoverInstances` request or using
DNS. For the current quota on the number of namespaces that you can create using the same AWS account, see [AWS Cloud Map quotas](https://docs.aws.amazon.com/cloud-map/latest/dg/cloud-map-limits.html) in the
_AWS Cloud Map Developer Guide_.

###### Important

The `CreatePublicDnsNamespace` API operation is not supported in the AWS GovCloud (US) Regions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ServiceDiscovery::PublicDnsNamespace",
  "Properties" : {
      "Description" : String,
      "Name" : String,
      "Properties" : Properties,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ServiceDiscovery::PublicDnsNamespace
Properties:
  Description: String
  Name: String
  Properties:
    Properties
  Tags:
    - Tag

```

## Properties

`Description`

A description for the namespace.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name that you want to assign to this namespace.

###### Note

Do not include sensitive information in the name. The name is publicly available using DNS
queries.

_Required_: Yes

_Type_: String

_Pattern_: `^([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?$`

_Maximum_: `253`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Properties`

Properties for the
public DNS namespace.

_Required_: No

_Type_: [Properties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-servicediscovery-publicdnsnamespace-properties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags for the namespace. Each tag consists of a key and an optional value, both of
which you define. Tag keys can have a maximum character length of 128 characters, and tag
values can have a maximum length of 256 characters.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-servicediscovery-publicdnsnamespace-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: Updates are not supported.

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the value of `Id` for the namespace, such as
`ns-e4anhexample0004`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the public namespace.

`HostedZoneId`

The ID for the Route 53 hosted zone that AWS Cloud Map creates when you create a
namespace.

`Id`

The ID of the public namespace.

## Examples

### Create a public DNS namespace

The following example creates a public DNS namespace named
`example.com`.

#### JSON

```json

{
    "Resources": {
        "PublicDnsNamespace": {
            "Type": "AWS::ServiceDiscovery::PublicDnsNamespace",
            "Properties": {
                "Description": "AWS Cloud Map public DNS namespace for example.com website",
                "Name": "example-public-http.com",
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
  PublicDnsNamespace:
    Type: AWS::ServiceDiscovery::PublicDnsNamespace
    Properties:
      Description: AWS Cloud Map public DNS namespace for example.com website
      Name: example-public-http.com
      Properties:
        DnsProperties:
          SOA:
            TTL: 100
```

## See also

- [CreatePublicDnsNamespace](https://docs.aws.amazon.com/cloud-map/latest/api/API_CreatePublicDnsNamespace.html) in the _AWS Cloud Map API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Properties
