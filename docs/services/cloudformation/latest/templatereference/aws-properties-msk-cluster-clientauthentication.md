This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::Cluster ClientAuthentication

The `ClientAuthentication` property type specifies Property description not available. for an [AWS::MSK::Cluster](aws-resource-msk-cluster.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Sasl" : Sasl,
  "Tls" : Tls,
  "Unauthenticated" : Unauthenticated
}

```

### YAML

```yaml

  Sasl:
    Sasl
  Tls:
    Tls
  Unauthenticated:
    Unauthenticated

```

## Properties

`Sasl`

Details for client authentication using SASL. To turn on SASL, you must also turn on `EncryptionInTransit` by setting `inCluster` to true. You must set `clientBroker` to either `TLS` or `TLS_PLAINTEXT`. If you choose `TLS_PLAINTEXT`, then you must also set `unauthenticated` to true.

_Required_: No

_Type_: [Sasl](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-msk-cluster-sasl.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tls`

Details for ClientAuthentication using TLS. To turn on TLS access control, you must also turn on `EncryptionInTransit` by setting `inCluster` to true and `clientBroker` to `TLS`.

_Required_: No

_Type_: [Tls](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-msk-cluster-tls.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Unauthenticated`

Details for ClientAuthentication using no authentication.

_Required_: No

_Type_: [Unauthenticated](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-msk-cluster-unauthenticated.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

BrokerNodeGroupInfo

CloudWatchLogs
