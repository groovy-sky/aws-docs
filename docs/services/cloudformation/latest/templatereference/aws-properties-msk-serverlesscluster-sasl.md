This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::ServerlessCluster Sasl

Details for client authentication using SASL. To turn on SASL, you must also turn on `EncryptionInTransit` by setting `inCluster` to true. You must set `clientBroker` to either `TLS` or `TLS_PLAINTEXT`. If you choose `TLS_PLAINTEXT`, then you must also set `unauthenticated` to true.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Iam" : Iam
}

```

### YAML

```yaml

  Iam:
    Iam

```

## Properties

`Iam`

Details for ClientAuthentication using IAM.

_Required_: Yes

_Type_: [Iam](aws-properties-msk-serverlesscluster-iam.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Iam

VpcConfig

All content copied from https://docs.aws.amazon.com/.
