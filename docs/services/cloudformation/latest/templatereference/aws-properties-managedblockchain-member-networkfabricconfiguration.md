This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ManagedBlockchain::Member NetworkFabricConfiguration

Hyperledger Fabric configuration properties for the network.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Edition" : String
}

```

### YAML

```yaml

  Edition: String

```

## Properties

`Edition`

The edition of Amazon Managed Blockchain that the network uses. Valid values are
`standard` and `starter`. For more information, see [Amazon Managed Blockchain Pricing](https://aws.amazon.com/managed-blockchain/pricing)

_Required_: Yes

_Type_: String

_Allowed values_: `STARTER | STANDARD`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NetworkConfiguration

NetworkFrameworkConfiguration

All content copied from https://docs.aws.amazon.com/.
