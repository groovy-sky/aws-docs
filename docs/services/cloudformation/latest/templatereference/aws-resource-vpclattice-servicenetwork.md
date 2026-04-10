This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VpcLattice::ServiceNetwork

Creates a service network. A service network is a logical boundary for a collection of
services. You can associate services and VPCs with a service network.

For more information, see [Service networks](../../../vpc-lattice/latest/ug/service-networks.md) in the
_Amazon VPC Lattice User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::VpcLattice::ServiceNetwork",
  "Properties" : {
      "AuthType" : String,
      "Name" : String,
      "SharingConfig" : SharingConfig,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::VpcLattice::ServiceNetwork
Properties:
  AuthType: String
  Name: String
  SharingConfig:
    SharingConfig
  Tags:
    - Tag

```

## Properties

`AuthType`

The type of IAM policy.

- `NONE`: The resource does not use an IAM policy. This is the default.

- `AWS_IAM`: The resource uses an IAM policy. When this type is used, auth is enabled and an auth policy is required.

_Required_: No

_Type_: String

_Allowed values_: `NONE | AWS_IAM`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the service network. The name must be unique to the account. The valid
characters are a-z, 0-9, and hyphens (-). You can't use a hyphen as the first or last
character, or immediately after another hyphen.

If you don't specify a name, CloudFormation generates one. However, if
you specify a name, and later want to replace the resource, you must specify a new
name.

_Required_: No

_Type_: String

_Pattern_: `^(?!servicenetwork-)(?![-])(?!.*[-]$)(?!.*[-]{2})[a-z0-9-]+$`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SharingConfig`

Specify if the service network should be enabled for sharing.

_Required_: No

_Type_: [SharingConfig](aws-properties-vpclattice-servicenetwork-sharingconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags for the service network.

_Required_: No

_Type_: Array of [Tag](aws-properties-vpclattice-servicenetwork-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the service network.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the service network.

`CreatedAt`

The date and time that the service network was created, specified in ISO-8601
format.

`Id`

The ID of the service network.

`LastUpdatedAt`

The date and time of the last update, specified in ISO-8601 format.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

SharingConfig

All content copied from https://docs.aws.amazon.com/.
