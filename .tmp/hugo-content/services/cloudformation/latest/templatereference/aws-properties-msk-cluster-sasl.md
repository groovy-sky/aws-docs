This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::Cluster Sasl

Details for client authentication using SASL. To turn on SASL, you must also turn on `EncryptionInTransit` by setting `inCluster` to true. You must set `clientBroker` to either `TLS` or `TLS_PLAINTEXT`. If you choose `TLS_PLAINTEXT`, then you must also set `unauthenticated` to true.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Iam" : Iam,
  "Scram" : Scram
}

```

### YAML

```yaml

  Iam:
    Iam
  Scram:
    Scram

```

## Properties

`Iam`

Details for ClientAuthentication using IAM.

_Required_: No

_Type_: [Iam](aws-properties-msk-cluster-iam.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scram`

Details for SASL/SCRAM client authentication.

_Required_: No

_Type_: [Scram](aws-properties-msk-cluster-scram.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3

Scram

All content copied from https://docs.aws.amazon.com/.
