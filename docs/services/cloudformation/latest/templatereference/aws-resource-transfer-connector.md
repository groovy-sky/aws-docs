This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Transfer::Connector

Creates the connector, which captures the parameters for a connection for the AS2 or
SFTP protocol. For AS2, the connector is required for sending files to an externally
hosted AS2 server. For SFTP, the connector is required when sending files to an SFTP
server or receiving files from an SFTP server. For more details about connectors, see
[Configure AS2\
connectors](https://docs.aws.amazon.com/transfer/latest/userguide/configure-as2-connector.html) and [Create SFTP\
connectors](https://docs.aws.amazon.com/transfer/latest/userguide/configure-sftp-connector.html).

###### Note

You must specify exactly one configuration object: either for AS2
( `As2Config`) or SFTP ( `SftpConfig`).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Transfer::Connector",
  "Properties" : {
      "AccessRole" : String,
      "As2Config" : As2Config,
      "EgressConfig" : ConnectorEgressConfig,
      "EgressType" : String,
      "LoggingRole" : String,
      "SecurityPolicyName" : String,
      "SftpConfig" : SftpConfig,
      "Tags" : [ Tag, ... ],
      "Url" : String
    }
}

```

### YAML

```yaml

Type: AWS::Transfer::Connector
Properties:
  AccessRole: String
  As2Config:
    As2Config
  EgressConfig:
    ConnectorEgressConfig
  EgressType: String
  LoggingRole: String
  SecurityPolicyName: String
  SftpConfig:
    SftpConfig
  Tags:
    - Tag
  Url: String

```

## Properties

`AccessRole`

Connectors are used to send files using either the AS2 or SFTP protocol. For the access role,
provide the Amazon Resource Name (ARN) of the AWS Identity and Access Management role to use.

**For AS2 connectors**

With AS2, you can send files by calling `StartFileTransfer` and specifying the
file paths in the request parameter, `SendFilePaths`. We use the file’s parent
directory (for example, for `--send-file-paths /bucket/dir/file.txt`, parent
directory is `/bucket/dir/`) to temporarily store a processed AS2 message file,
store the MDN when we receive them from the partner, and write a final JSON file containing
relevant metadata of the transmission. So, the `AccessRole` needs to provide read
and write access to the parent directory of the file location used in the
`StartFileTransfer` request. Additionally, you need to provide read and write
access to the parent directory of the files that you intend to send with
`StartFileTransfer`.

If you are using Basic authentication for your AS2 connector, the access role requires the
`secretsmanager:GetSecretValue` permission for the secret. If the secret is encrypted using
a customer-managed key instead of the AWS managed key in Secrets Manager, then the role also
needs the `kms:Decrypt` permission for that key.

**For SFTP connectors**

Make sure that the access role provides
read and write access to the parent directory of the file location
that's used in the `StartFileTransfer` request.
Additionally, make sure that the role provides
`secretsmanager:GetSecretValue` permission to AWS Secrets Manager.

_Required_: Yes

_Type_: String

_Pattern_: `arn:.*role/.*`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`As2Config`

A structure that contains the parameters for an AS2 connector object.

_Required_: No

_Type_: [As2Config](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-transfer-connector-as2config.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EgressConfig`

Current egress configuration of the connector, showing how traffic is routed to the
SFTP server. Contains VPC Lattice settings when using VPC\_LATTICE egress type.

When using the VPC\_LATTICE egress type, AWS Transfer Family uses a managed Service
Network to simplify the resource sharing process.

_Required_: No

_Type_: [ConnectorEgressConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-transfer-connector-connectoregressconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EgressType`

Type of egress configuration for the connector. SERVICE\_MANAGED uses Transfer Family
managed NAT gateways, while VPC\_LATTICE routes traffic through customer VPCs using VPC
Lattice.

_Required_: No

_Type_: String

_Allowed values_: `SERVICE_MANAGED | VPC_LATTICE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoggingRole`

The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that allows a connector to turn
on CloudWatch logging for Amazon S3 events. When set, you can view connector
activity in your CloudWatch logs.

_Required_: No

_Type_: String

_Pattern_: `arn:.*role/.*`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityPolicyName`

The text name of the security policy for the specified connector.

_Required_: No

_Type_: String

_Pattern_: `TransferSFTPConnectorSecurityPolicy-[A-Za-z0-9-]+`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SftpConfig`

A structure that contains the parameters for an SFTP connector object.

_Required_: No

_Type_: [SftpConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-transfer-connector-sftpconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Key-value pairs that can be used to group and search for connectors.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-transfer-connector-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Url`

The URL of the partner's AS2 or SFTP endpoint.

When creating AS2 connectors or service-managed SFTP connectors
(connectors without egress configuration), you must provide a URL to
specify the remote server endpoint. For VPC Lattice type connectors,
the URL must be null.

_Required_: No

_Type_: String

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ConnectorId`

The service-assigned ID of the connector that is created.

`ErrorMessage`

Error message providing details when the connector is in ERRORED status. Contains
information to help troubleshoot connector creation or operation failures.

`ServiceManagedEgressIpAddresses`

The list of egress IP addresses of this connector. These IP addresses are assigned
automatically when you create the connector.

`Status`

Current status of the connector. PENDING indicates creation/update in progress, ACTIVE
means ready for operations, and ERRORED indicates a failure requiring attention.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

As2Config
