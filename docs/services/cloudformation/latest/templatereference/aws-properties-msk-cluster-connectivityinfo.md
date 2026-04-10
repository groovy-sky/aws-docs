This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::Cluster ConnectivityInfo

Broker access controls.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NetworkType" : String,
  "PublicAccess" : PublicAccess,
  "VpcConnectivity" : VpcConnectivity
}

```

### YAML

```yaml

  NetworkType: String
  PublicAccess:
    PublicAccess
  VpcConnectivity:
    VpcConnectivity

```

## Properties

`NetworkType`

Property description not available.

_Required_: No

_Type_: String

_Allowed values_: `IPV4 | DUAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PublicAccess`

Access control settings for the cluster's brokers.

_Required_: No

_Type_: [PublicAccess](aws-properties-msk-cluster-publicaccess.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcConnectivity`

VPC connection control settings for brokers.

_Required_: No

_Type_: [VpcConnectivity](aws-properties-msk-cluster-vpcconnectivity.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConfigurationInfo

EBSStorageInfo

All content copied from https://docs.aws.amazon.com/.
