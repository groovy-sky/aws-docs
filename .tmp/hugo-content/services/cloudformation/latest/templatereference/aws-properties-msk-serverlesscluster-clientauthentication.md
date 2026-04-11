This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::ServerlessCluster ClientAuthentication

Includes all client authentication information.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Sasl" : Sasl
}

```

### YAML

```yaml

  Sasl:
    Sasl

```

## Properties

`Sasl`

Details for client authentication using SASL. To turn on SASL, you must also turn on `EncryptionInTransit` by setting `inCluster` to true. You must set `clientBroker` to either `TLS` or `TLS_PLAINTEXT`. If you choose `TLS_PLAINTEXT`, then you must also set `unauthenticated` to true.

_Required_: Yes

_Type_: [Sasl](aws-properties-msk-serverlesscluster-sasl.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::MSK::ServerlessCluster

Iam

All content copied from https://docs.aws.amazon.com/.
