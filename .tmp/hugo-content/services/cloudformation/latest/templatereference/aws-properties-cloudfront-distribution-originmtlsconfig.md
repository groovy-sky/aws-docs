This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::Distribution OriginMtlsConfig

Configures mutual TLS authentication between CloudFront and your origin server.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClientCertificateArn" : String
}

```

### YAML

```yaml

  ClientCertificateArn: String

```

## Properties

`ClientCertificateArn`

The Amazon Resource Name (ARN) of the client certificate stored in AWS Certificate Manager (ACM) that CloudFront uses to authenticate with your origin using Mutual TLS.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OriginGroups

OriginShield

All content copied from https://docs.aws.amazon.com/.
