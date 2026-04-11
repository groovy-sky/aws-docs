This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::TLSInspectionConfiguration ServerCertificateConfiguration

Configures the AWS Certificate Manager certificates and scope that Network Firewall uses to decrypt and re-encrypt traffic using a [TLSInspectionConfiguration](aws-resource-networkfirewall-tlsinspectionconfiguration.md). You can configure `ServerCertificates` for inbound SSL/TLS inspection, a `CertificateAuthorityArn` for outbound SSL/TLS inspection, or both. For information about working with certificates for TLS inspection, see [Using SSL/TLS server certficiates with TLS inspection configurations](../../../network-firewall/latest/developerguide/tls-inspection-certificate-requirements.md) in the _AWS Network Firewall Developer Guide_.

###### Note

If a server certificate that's associated with your [TLSInspectionConfiguration](aws-resource-networkfirewall-tlsinspectionconfiguration.md) is revoked, deleted, or expired it can result in client-side TLS errors.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CertificateAuthorityArn" : String,
  "CheckCertificateRevocationStatus" : CheckCertificateRevocationStatus,
  "Scopes" : [ ServerCertificateScope, ... ],
  "ServerCertificates" : [ ServerCertificate, ... ]
}

```

### YAML

```yaml

  CertificateAuthorityArn: String
  CheckCertificateRevocationStatus:
    CheckCertificateRevocationStatus
  Scopes:
    - ServerCertificateScope
  ServerCertificates:
    - ServerCertificate

```

## Properties

`CertificateAuthorityArn`

The Amazon Resource Name (ARN) of the imported certificate authority (CA) certificate within AWS Certificate Manager (ACM) to use for outbound SSL/TLS inspection.

The following limitations apply:

- You can use CA certificates that you imported into ACM, but you can't generate CA certificates with ACM.

- You can't use certificates issued by AWS Private Certificate Authority.

For more information about configuring certificates for outbound inspection, see [Using SSL/TLS certificates with TLS inspection configurations](../../../network-firewall/latest/developerguide/tls-inspection-certificate-requirements.md) in the _AWS Network Firewall Developer Guide_.

For information about working with certificates in ACM, see [Importing certificates](../../../acm/latest/userguide/import-certificate.md) in the _AWS Certificate Manager User Guide_.

_Required_: No

_Type_: String

_Pattern_: `^(arn:aws.*)$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CheckCertificateRevocationStatus`

When enabled, Network Firewall checks if the server certificate presented by the server in the SSL/TLS connection has a revoked or unkown status. If the certificate has an unknown or revoked status, you must specify the actions that Network Firewall takes on outbound traffic. To check the certificate revocation status, you must also specify a `CertificateAuthorityArn` in [ServerCertificateConfiguration](aws-resource-networkfirewall-servercertificateconfiguration.md).

_Required_: No

_Type_: [CheckCertificateRevocationStatus](aws-properties-networkfirewall-tlsinspectionconfiguration-checkcertificaterevocationstatus.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scopes`

A list of scopes.

_Required_: No

_Type_: Array of [ServerCertificateScope](aws-properties-networkfirewall-tlsinspectionconfiguration-servercertificatescope.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerCertificates`

The list of server certificates to use for inbound SSL/TLS inspection.

_Required_: No

_Type_: Array of [ServerCertificate](aws-properties-networkfirewall-tlsinspectionconfiguration-servercertificate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ServerCertificate

ServerCertificateScope

All content copied from https://docs.aws.amazon.com/.
