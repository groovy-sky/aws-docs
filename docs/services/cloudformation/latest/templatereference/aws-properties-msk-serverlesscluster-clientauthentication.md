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

_Type_: [Sasl](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-msk-serverlesscluster-sasl.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::MSK::ServerlessCluster

Iam
