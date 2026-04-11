This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::Cluster VpcConnectivityClientAuthentication

Includes all client authentication information for VpcConnectivity.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Sasl" : VpcConnectivitySasl,
  "Tls" : VpcConnectivityTls
}

```

### YAML

```yaml

  Sasl:
    VpcConnectivitySasl
  Tls:
    VpcConnectivityTls

```

## Properties

`Sasl`

Details for VpcConnectivity ClientAuthentication using SASL.

_Required_: No

_Type_: [VpcConnectivitySasl](aws-properties-msk-cluster-vpcconnectivitysasl.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tls`

Details for VpcConnectivity ClientAuthentication using TLS.

_Required_: No

_Type_: [VpcConnectivityTls](aws-properties-msk-cluster-vpcconnectivitytls.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VpcConnectivity

VpcConnectivityIam

All content copied from https://docs.aws.amazon.com/.
