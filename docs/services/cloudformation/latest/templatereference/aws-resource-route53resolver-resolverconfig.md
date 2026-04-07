This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53Resolver::ResolverConfig

A complex type that contains information about a Resolver configuration for a VPC.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Route53Resolver::ResolverConfig",
  "Properties" : {
      "AutodefinedReverseFlag" : String,
      "ResourceId" : String
    }
}

```

### YAML

```yaml

Type: AWS::Route53Resolver::ResolverConfig
Properties:
  AutodefinedReverseFlag: String
  ResourceId: String

```

## Properties

`AutodefinedReverseFlag`

Represents the desired status of `AutodefinedReverse`. The only supported value on creation is `DISABLE`.
Deletion of this resource will return `AutodefinedReverse` to its default value of `ENABLED`.

_Required_: Yes

_Type_: String

_Allowed values_: `DISABLE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceId`

The ID of the Amazon Virtual Private Cloud VPC or a Route 53 Profile that you're configuring Resolver for.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `ResolverConfiguration` ID.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AutodefinedReverse`

The status of whether or not the Route 53 Resolver will create autodefined rules for reverse DNS lookups. This is enabled by default.

`Id`

ID for the Route 53 Resolver configuration.

`OwnerId`

The owner account ID of the Amazon Virtual Private Cloud VPC.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::Route53Resolver::ResolverDNSSECConfig
