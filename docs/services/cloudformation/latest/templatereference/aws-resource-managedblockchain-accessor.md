This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ManagedBlockchain::Accessor

Creates a new accessor for use with Amazon Managed Blockchain service that supports token based access.
The accessor contains information required for token based access.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ManagedBlockchain::Accessor",
  "Properties" : {
      "AccessorType" : String,
      "NetworkType" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ManagedBlockchain::Accessor
Properties:
  AccessorType: String
  NetworkType: String
  Tags:
    - Tag

```

## Properties

`AccessorType`

The type of the accessor.

###### Note

Currently, accessor type is restricted to `BILLING_TOKEN`.

_Required_: Yes

_Type_: String

_Allowed values_: `BILLING_TOKEN`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetworkType`

The blockchain network that the `Accessor` token is created for.

###### Note

We recommend using the appropriate `networkType`
value for the blockchain network that you are creating the `Accessor`
token for. You cannot use the value `ETHEREUM_MAINNET_AND_GOERLI` to
specify a `networkType` for your Accessor token.

The default value of `ETHEREUM_MAINNET_AND_GOERLI` is only applied:

- when the `CreateAccessor` action does not set a `networkType`.

- to all existing `Accessor` tokens that were created before the `networkType` property was introduced.

_Required_: No

_Type_: String

_Allowed values_: `ETHEREUM_GOERLI | ETHEREUM_MAINNET | ETHEREUM_MAINNET_AND_GOERLI | POLYGON_MAINNET | POLYGON_MUMBAI`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags assigned to the Accessor.

For more information about tags, see [Tagging Resources](../../../managed-blockchain/latest/ethereum-dev/tagging-resources.md) in the _Amazon Managed Blockchain Ethereum Developer Guide_, or [Tagging Resources](../../../managed-blockchain/latest/hyperledger-fabric-dev/tagging-resources.md) in the _Amazon Managed Blockchain Hyperledger Fabric Developer Guide_.

_Required_: No

_Type_: Array of [Tag](aws-properties-managedblockchain-accessor-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Accessor ID.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the accessor. For more information about
ARNs and their format, see [Amazon Resource \
Names (ARNs)](../../../../general/latest/gr/aws-arns-and-namespaces.md) in the _AWS General Reference_.

`BillingToken`

The billing token is a property of the accessor. Use this token to make Ethereum API calls to your
Ethereum node. The billing token is used to track your accessor object for billing Ethereum API
requests made to your Ethereum nodes.

`CreationDate`

The creation date and time of the accessor.

`Id`

The unique identifier of the accessor.

`Status`

The current status of the accessor.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Managed Blockchain

Tag

All content copied from https://docs.aws.amazon.com/.
