This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ServiceDiscovery::HttpNamespace

Creates an HTTP namespace. Service instances registered using an HTTP namespace can be
discovered using a `DiscoverInstances` request but can't be discovered using
DNS.

For the current quota on the number of namespaces that you can create using the same AWS account, see [AWS Cloud Map quotas](../../../cloud-map/latest/dg/cloud-map-limits.md) in the
_AWS Cloud Map Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ServiceDiscovery::HttpNamespace",
  "Properties" : {
      "Description" : String,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ServiceDiscovery::HttpNamespace
Properties:
  Description: String
  Name: String
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

_Required_: Yes

_Type_: String

_Pattern_: `^(?!arn:)[!-~]{1,1024}$`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags for the namespace. Each tag consists of a key and an optional value, both of
which you define. Tag keys can have a maximum character length of 128 characters, and tag
values can have a maximum length of 256 characters.

_Required_: No

_Type_: Array of [Tag](aws-properties-servicediscovery-httpnamespace-tag.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: Updates are not supported.

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `Id` for the namespace, such as
`ns-e4anhexample0004`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the namespace, such as
`arn:aws:service-discovery:us-east-1:123456789012:http-namespace/http-namespace-a1bzhi`.

`Id`

The ID of the namespace.

## Examples

### Create an HTTP namespace

The following example creates an HTTP namespace named
`example-namespace`.

#### JSON

```json

{
    "Resources": {
        "HttpNamespace": {
            "Type": "AWS::ServiceDiscovery::HttpNamespace",
            "Properties": {
                "Description": "AWS Cloud Map HTTP namespace for resources for example.com website",
                "Name": "example-namespace"
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
  HttpNamespace:
    Type: AWS::ServiceDiscovery::HttpNamespace
    Properties:
      Description: AWS Cloud Map HTTP namespace for resources for example.com website
      Name: example-namespace
```

## See also

- [CreateHttpNamespace](../../../cloud-map/latest/api/api-createhttpnamespace.md) in the _AWS Cloud Map API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Cloud Map

Tag

All content copied from https://docs.aws.amazon.com/.
