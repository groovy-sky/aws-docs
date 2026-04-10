This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::Cluster VpcConnectivity

VPC connection control settings for brokers.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClientAuthentication" : VpcConnectivityClientAuthentication
}

```

### YAML

```yaml

  ClientAuthentication:
    VpcConnectivityClientAuthentication

```

## Properties

`ClientAuthentication`

VPC connection control settings for brokers.

_Required_: No

_Type_: [VpcConnectivityClientAuthentication](aws-properties-msk-cluster-vpcconnectivityclientauthentication.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Unauthenticated

VpcConnectivityClientAuthentication

All content copied from https://docs.aws.amazon.com/.
