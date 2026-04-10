This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::DomainConfiguration ClientCertificateConfig

An object that speciﬁes the client certificate conﬁguration for a domain.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClientCertificateCallbackArn" : String
}

```

### YAML

```yaml

  ClientCertificateCallbackArn: String

```

## Properties

`ClientCertificateCallbackArn`

The ARN of the Lambda function that IoT invokes after mutual TLS authentication during the connection.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `170`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AuthorizerConfig

ServerCertificateConfig

All content copied from https://docs.aws.amazon.com/.
