This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53Resolver::ResolverDNSSECConfig

The `AWS::Route53Resolver::ResolverDNSSECConfig` resource is a complex type that contains information about a configuration for DNSSEC validation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Route53Resolver::ResolverDNSSECConfig",
  "Properties" : {
      "ResourceId" : String
    }
}

```

### YAML

```yaml

Type: AWS::Route53Resolver::ResolverDNSSECConfig
Properties:
  ResourceId: String

```

## Properties

`ResourceId`

The ID of the virtual private cloud (VPC) that you're configuring the DNSSEC validation status for.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID. For example:

`{ "Ref": "rdsc-689d45d1ae623bf3" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The primary identifier of this `ResolverDNSSECConfig` resource. For example: `rdsc-689d45d1ae623bf3`.

`OwnerId`

The AWS account of the owner. For example: `111122223333`.

`ValidationStatus`

The current status of this `ResolverDNSSECConfig` resource. For example: `Enabled`.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Route53Resolver::ResolverConfig

AWS::Route53Resolver::ResolverEndpoint

All content copied from https://docs.aws.amazon.com/.
