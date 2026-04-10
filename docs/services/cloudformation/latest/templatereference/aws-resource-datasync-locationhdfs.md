This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::LocationHDFS

The `AWS::DataSync::LocationHDFS` resource specifies an endpoint for a
Hadoop Distributed File System (HDFS).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataSync::LocationHDFS",
  "Properties" : {
      "AgentArns" : [ String, ... ],
      "AuthenticationType" : String,
      "BlockSize" : Integer,
      "CmkSecretConfig" : CmkSecretConfig,
      "CustomSecretConfig" : CustomSecretConfig,
      "KerberosKeytab" : String,
      "KerberosKrb5Conf" : String,
      "KerberosPrincipal" : String,
      "KmsKeyProviderUri" : String,
      "NameNodes" : [ NameNode, ... ],
      "QopConfiguration" : QopConfiguration,
      "ReplicationFactor" : Integer,
      "SimpleUser" : String,
      "Subdirectory" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DataSync::LocationHDFS
Properties:
  AgentArns:
    - String
  AuthenticationType: String
  BlockSize: Integer
  CmkSecretConfig:
    CmkSecretConfig
  CustomSecretConfig:
    CustomSecretConfig
  KerberosKeytab: String
  KerberosKrb5Conf: String
  KerberosPrincipal: String
  KmsKeyProviderUri: String
  NameNodes:
    - NameNode
  QopConfiguration:
    QopConfiguration
  ReplicationFactor: Integer
  SimpleUser: String
  Subdirectory: String
  Tags:
    - Tag

```

## Properties

`AgentArns`

The Amazon Resource Names (ARNs) of the DataSync agents that can connect to your
HDFS cluster.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `128 | 4`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthenticationType`

Property description not available.

_Required_: Yes

_Type_: String

_Allowed values_: `SIMPLE | KERBEROS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BlockSize`

The size of data blocks to write into the HDFS cluster. The block size must be a multiple
of 512 bytes. The default block size is 128 mebibytes (MiB).

_Required_: No

_Type_: Integer

_Minimum_: `1048576`

_Maximum_: `1073741824`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CmkSecretConfig`

Specifies configuration information for a DataSync-managed secret, such as an
authentication token, secret key, password, or Kerberos keytab that DataSync uses
to access a specific storage location, with a customer-managed AWS KMS key.

###### Note

You can use either `CmkSecretConfig` or `CustomSecretConfig` to
provide credentials for a `CreateLocation` request. Do not provide both
parameters for the same request.

_Required_: No

_Type_: [CmkSecretConfig](aws-properties-datasync-locationhdfs-cmksecretconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomSecretConfig`

Specifies configuration information for a customer-managed Secrets Manager secret where
a storage location credentials is stored in Secrets Manager as plain text (for authentication
token, secret key, or password) or as binary (for Kerberos keytab). This configuration includes
the secret ARN, and the ARN for an IAM role that provides access to the secret.

###### Note

You can use either `CmkSecretConfig` or `CustomSecretConfig` to
provide credentials for a `CreateLocation` request. Do not provide both
parameters for the same request.

_Required_: No

_Type_: [CustomSecretConfig](aws-properties-datasync-locationhdfs-customsecretconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KerberosKeytab`

The Kerberos key table (keytab) that contains mappings between the defined Kerberos
principal and the encrypted keys. Provide the base64-encoded file text. If
`KERBEROS` is specified for `AuthType`, this value is
required.

_Required_: No

_Type_: String

_Maximum_: `87384`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KerberosKrb5Conf`

The `krb5.conf` file that contains the Kerberos configuration information.
You can load the `krb5.conf` by providing a string of the file's contents or
an Amazon S3 presigned URL of the file. If `KERBEROS` is specified for
`AuthType`, this value is required.

_Required_: No

_Type_: String

_Maximum_: `174764`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KerberosPrincipal`

The Kerberos principal with access to the files and folders on the HDFS cluster.

###### Note

If `KERBEROS` is specified for `AuthenticationType`, this
parameter is required.

_Required_: No

_Type_: String

_Pattern_: `^.+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyProviderUri`

The URI of the HDFS cluster's Key Management Server (KMS).

_Required_: No

_Type_: String

_Pattern_: `^kms:\/\/http[s]?@(([a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9\-]*[A-Za-z0-9])(;(([a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9\-]*[A-Za-z0-9]))*:[0-9]{1,5}\/kms$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NameNodes`

The NameNode that manages the HDFS namespace. The NameNode performs operations such as
opening, closing, and renaming files and directories. The NameNode contains the information to
map blocks of data to the DataNodes. You can use only one NameNode.

_Required_: Yes

_Type_: Array of [NameNode](aws-properties-datasync-locationhdfs-namenode.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QopConfiguration`

The Quality of Protection (QOP) configuration specifies the Remote Procedure Call (RPC)
and data transfer protection settings configured on the Hadoop Distributed File System (HDFS)
cluster. If `QopConfiguration` isn't specified, `RpcProtection` and
`DataTransferProtection` default to `PRIVACY`. If you set
`RpcProtection` or `DataTransferProtection`, the other parameter
assumes the same value.

_Required_: No

_Type_: [QopConfiguration](aws-properties-datasync-locationhdfs-qopconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplicationFactor`

The number of DataNodes to replicate the data to when writing to the HDFS cluster. By
default, data is replicated to three DataNodes.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SimpleUser`

The user name used to identify the client on the host operating system.

###### Note

If `SIMPLE` is specified for `AuthenticationType`, this parameter
is required.

_Required_: No

_Type_: String

_Pattern_: `^[_.A-Za-z0-9][-_.A-Za-z0-9]*$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Subdirectory`

A subdirectory in the HDFS cluster. This subdirectory is used to read data from or write
data to the HDFS cluster. If the subdirectory isn't specified, it will default to
`/`.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-\+\./\(\)\$\p{Zs}]+$`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The key-value pair that represents the tag that you want to add to the location. The value
can be an empty string. We recommend using tags to name your resources.

_Required_: No

_Type_: Array of [Tag](aws-properties-datasync-locationhdfs-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref`
function, `Ref` returns the location resource ARN. For example:

`arn:aws:datasync:us-east-2:111222333444:location/loc-07db7abfc326c50s3`

For more information about using the `Ref` function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified
attribute of this type. The following are the available attributes and sample return
values.

For more information about using the `Fn::GetAtt` intrinsic function, see
[Fn::GetAtt](../userguide/intrinsic-function-reference-getatt.md).

`CmkSecretConfig.SecretArn`

Property description not available.

`LocationArn`

The Amazon Resource Name (ARN) of the HDFS cluster location to describe.

`LocationUri`

The URI of the HDFS cluster location.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

CmkSecretConfig

All content copied from https://docs.aws.amazon.com/.
