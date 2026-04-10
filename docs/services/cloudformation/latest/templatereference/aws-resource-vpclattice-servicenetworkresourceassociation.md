This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VpcLattice::ServiceNetworkResourceAssociation

Associates the specified service network with the specified resource configuration. This
allows the resource configuration to receive connections through the service network, including
through a service network VPC endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::VpcLattice::ServiceNetworkResourceAssociation",
  "Properties" : {
      "PrivateDnsEnabled" : Boolean,
      "ResourceConfigurationId" : String,
      "ServiceNetworkId" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::VpcLattice::ServiceNetworkResourceAssociation
Properties:
  PrivateDnsEnabled: Boolean
  ResourceConfigurationId: String
  ServiceNetworkId: String
  Tags:
    - Tag

```

## Properties

`PrivateDnsEnabled`

Indicates if private DNS is enabled for the service network resource association.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceConfigurationId`

The ID of the resource configuration associated with the service network.

_Required_: No

_Type_: String

_Pattern_: `^rcfg-[0-9a-z]{17}$`

_Minimum_: `17`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ServiceNetworkId`

The ID of the service network associated with the resource configuration.

_Required_: No

_Type_: String

_Pattern_: `^((sn-[0-9a-z]{17})|(arn:[a-z0-9\-]+:vpc-lattice:[a-zA-Z0-9\-]+:\d{12}:servicenetwork/sn-[0-9a-z]{17}))$`

_Minimum_: `3`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A key-value pair to associate with a resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-vpclattice-servicenetworkresourceassociation-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the association.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the association.

`Id`

The ID of the association between the service network and resource configuration.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
