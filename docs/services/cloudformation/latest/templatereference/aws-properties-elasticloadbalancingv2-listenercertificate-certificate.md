This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancingV2::ListenerCertificate Certificate

Specifies an SSL server certificate for the certificate list of a secure
listener.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CertificateArn" : String
}

```

### YAML

```yaml

  CertificateArn: String

```

## Properties

`CertificateArn`

The Amazon Resource Name (ARN) of the certificate.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ElasticLoadBalancingV2::ListenerCertificate

AWS::ElasticLoadBalancingV2::ListenerRule
