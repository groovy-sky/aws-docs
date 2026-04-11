This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::Cluster VpcConnectivitySasl

Details for client authentication using SASL for VpcConnectivity.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Iam" : VpcConnectivityIam,
  "Scram" : VpcConnectivityScram
}

```

### YAML

```yaml

  Iam:
    VpcConnectivityIam
  Scram:
    VpcConnectivityScram

```

## Properties

`Iam`

Details for ClientAuthentication using IAM for VpcConnectivity.

_Required_: No

_Type_: [VpcConnectivityIam](aws-properties-msk-cluster-vpcconnectivityiam.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scram`

Details for SASL/SCRAM client authentication for VpcConnectivity.

_Required_: No

_Type_: [VpcConnectivityScram](aws-properties-msk-cluster-vpcconnectivityscram.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VpcConnectivityIam

VpcConnectivityScram

All content copied from https://docs.aws.amazon.com/.
