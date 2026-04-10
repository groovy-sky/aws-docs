This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::ClientVpnEndpoint CertificateAuthenticationRequest

Information about the client certificate to be used for authentication.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClientRootCertificateChainArn" : String
}

```

### YAML

```yaml

  ClientRootCertificateChainArn: String

```

## Properties

`ClientRootCertificateChainArn`

The ARN of the client certificate. The certificate must be signed by a certificate
authority (CA) and it must be provisioned in AWS Certificate Manager (ACM).

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EC2::ClientVpnEndpoint

ClientAuthenticationRequest

All content copied from https://docs.aws.amazon.com/.
