This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::TLSInspectionConfiguration ServerCertificate

Any AWS Certificate Manager (ACM) Secure Sockets Layer/Transport Layer Security (SSL/TLS) server certificate that's associated with a [ServerCertificateConfiguration](../userguide/aws-properties-networkfirewall-tlsinspectionconfiguration-servercertificateconfiguration.md). Used in a [TLSInspectionConfiguration](aws-resource-networkfirewall-tlsinspectionconfiguration.md) for inspection of inbound traffic to your firewall. You must request or import a SSL/TLS certificate into ACM for each domain Network Firewall needs to decrypt and inspect. AWS Network Firewall uses the SSL/TLS certificates to decrypt specified inbound SSL/TLS traffic going to your firewall. For information about working with certificates in AWS Certificate Manager, see [Request a public certificate](../../../acm/latest/userguide/gs-acm-request-public.md) or [Importing certificates](../../../acm/latest/userguide/import-certificate.md) in the _AWS Certificate Manager User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ResourceArn" : String
}

```

### YAML

```yaml

  ResourceArn: String

```

## Properties

`ResourceArn`

The Amazon Resource Name (ARN) of the AWS Certificate Manager SSL/TLS server certificate that's used for inbound SSL/TLS inspection.

_Required_: No

_Type_: String

_Pattern_: `^(arn:aws.*)$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PortRange

ServerCertificateConfiguration

All content copied from https://docs.aws.amazon.com/.
