This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Transfer::Connector SftpConfig

A structure that contains the parameters for an SFTP connector object.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxConcurrentConnections" : Integer,
  "TrustedHostKeys" : [ String, ... ],
  "UserSecretId" : String
}

```

### YAML

```yaml

  MaxConcurrentConnections: Integer
  TrustedHostKeys:
    - String
  UserSecretId: String

```

## Properties

`MaxConcurrentConnections`

Specify the number of concurrent connections that your connector creates to the remote
server. The default value is `1`. The maximum values is
`5`.

###### Note

If you are using the AWS Management Console, the default value is
`5`.

This parameter specifies the number of active connections that your connector can
establish with the remote server at the same time. Increasing this value can enhance
connector performance when transferring large file batches by enabling parallel
operations.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrustedHostKeys`

The public portion of the host key, or keys, that are used to identify the external
server to which you are connecting. You can use the `ssh-keyscan` command
against the SFTP server to retrieve the necessary key.

###### Note

`TrustedHostKeys` is optional for `CreateConnector`. If not
provided, you can use `TestConnection` to retrieve the server host key
during the initial connection attempt, and subsequently update the connector with
the observed host key.

When creating connectors with egress config (VPC\_LATTICE type connectors), since host
name is not something we can verify, the only accepted trusted host key format is
`key-type key-body` without the host name. For example: `ssh-rsa
                AAAAB3Nza...<long-string-for-public-key>`

The three standard SSH public key format elements are `<key type>`,
`<body base64>`, and an optional `<comment>`, with
spaces between each element. Specify only the `<key type>` and
`<body base64>`: do not enter the `<comment>` portion of
the key.

For the trusted host key, AWS Transfer Family accepts RSA and ECDSA keys.

- For RSA keys, the `<key type>` string is
`ssh-rsa`.

- For ECDSA keys, the `<key type>` string is either
`ecdsa-sha2-nistp256`, `ecdsa-sha2-nistp384`, or
`ecdsa-sha2-nistp521`, depending on the size of the key you
generated.

Run this command to retrieve the SFTP server host key, where your SFTP server name is
`ftp.host.com`.

`ssh-keyscan ftp.host.com`

This prints the public host key to standard output.

`ftp.host.com ssh-rsa
            AAAAB3Nza...<long-string-for-public-key>`

Copy and paste this string into the `TrustedHostKeys` field for the
`create-connector` command or into the **Trusted host**
**keys** field in the console.

For VPC Lattice type connectors (VPC\_LATTICE), remove the hostname from the key and
use only the `key-type key-body` format. In this example, it should be:
`ssh-rsa AAAAB3Nza...<long-string-for-public-key>`

_Required_: No

_Type_: Array of String

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserSecretId`

The identifier for the secret (in AWS Secrets Manager) that contains the SFTP user's
private key, password, or both. The identifier must be the Amazon Resource Name (ARN) of
the secret.

###### Note

- Required when creating an SFTP connector

- Optional when updating an existing SFTP connector

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConnectorVpcLatticeEgressConfig

Tag

All content copied from https://docs.aws.amazon.com/.
