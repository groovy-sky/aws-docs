This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::TLSInspectionConfiguration CheckCertificateRevocationStatus

When enabled, Network Firewall checks if the server certificate presented by the server in the SSL/TLS connection has a revoked or unkown status. If the certificate has an unknown or revoked status, you must specify the actions that Network Firewall takes on outbound traffic. To check the certificate revocation status, you must also specify a `CertificateAuthorityArn` in [ServerCertificateConfiguration](aws-resource-networkfirewall-servercertificateconfiguration.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RevokedStatusAction" : String,
  "UnknownStatusAction" : String
}

```

### YAML

```yaml

  RevokedStatusAction: String
  UnknownStatusAction: String

```

## Properties

`RevokedStatusAction`

Configures how Network Firewall processes traffic when it determines that the certificate presented by the server in the SSL/TLS connection has a revoked status.

- **PASS** \- Allow the connection to continue, and pass subsequent packets to the stateful engine for inspection.

- **DROP** \- Network Firewall closes the connection and drops subsequent packets for that connection.

- **REJECT** \- Network Firewall sends a TCP reject packet back to your client. The service closes the connection and drops subsequent packets for that connection. `REJECT` is available only for TCP traffic.

_Required_: No

_Type_: String

_Allowed values_: `PASS | DROP | REJECT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UnknownStatusAction`

Configures how Network Firewall processes traffic when it determines that the certificate presented by the server in the SSL/TLS connection has an unknown status, or a status that cannot be determined for any other reason, including when the service is unable to connect to the OCSP and CRL endpoints for the certificate.

- **PASS** \- Allow the connection to continue, and pass subsequent packets to the stateful engine for inspection.

- **DROP** \- Network Firewall closes the connection and drops subsequent packets for that connection.

- **REJECT** \- Network Firewall sends a TCP reject packet back to your client. The service closes the connection and drops subsequent packets for that connection. `REJECT` is available only for TCP traffic.

_Required_: No

_Type_: String

_Allowed values_: `PASS | DROP | REJECT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Address

PortRange

All content copied from https://docs.aws.amazon.com/.
