This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::LocationSMB

The `AWS::DataSync::LocationSMB` resource specifies a Server Message Block
(SMB) location that AWS DataSync can use as a transfer source or
destination.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataSync::LocationSMB",
  "Properties" : {
      "AgentArns" : [ String, ... ],
      "AuthenticationType" : String,
      "CmkSecretConfig" : CmkSecretConfig,
      "CustomSecretConfig" : CustomSecretConfig,
      "DnsIpAddresses" : [ String, ... ],
      "Domain" : String,
      "KerberosKeytab" : String,
      "KerberosKrb5Conf" : String,
      "KerberosPrincipal" : String,
      "MountOptions" : MountOptions,
      "Password" : String,
      "ServerHostname" : String,
      "Subdirectory" : String,
      "Tags" : [ Tag, ... ],
      "User" : String
    }
}

```

### YAML

```yaml

Type: AWS::DataSync::LocationSMB
Properties:
  AgentArns:
    - String
  AuthenticationType: String
  CmkSecretConfig:
    CmkSecretConfig
  CustomSecretConfig:
    CustomSecretConfig
  DnsIpAddresses:
    - String
  Domain: String
  KerberosKeytab: String
  KerberosKrb5Conf: String
  KerberosPrincipal: String
  MountOptions:
    MountOptions
  Password: String
  ServerHostname: String
  Subdirectory: String
  Tags:
    - Tag
  User: String

```

## Properties

`AgentArns`

Specifies the DataSync agent (or agents) that can connect to your SMB file
server. You specify an agent by using its Amazon Resource Name (ARN).

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `128 | 8`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthenticationType`

Property description not available.

_Required_: No

_Type_: String

_Allowed values_: `NTLM | KERBEROS`

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

_Type_: [CmkSecretConfig](aws-properties-datasync-locationsmb-cmksecretconfig.md)

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

_Type_: [CustomSecretConfig](aws-properties-datasync-locationsmb-customsecretconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DnsIpAddresses`

Property description not available.

_Required_: No

_Type_: Array of String

_Minimum_: `7`

_Maximum_: `15 | 2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Domain`

Specifies the Windows domain name that your SMB file server belongs to. This parameter
applies only if `AuthenticationType` is set to `NTLM`.

If you have multiple domains in your environment, configuring this parameter makes sure
that DataSync connects to the right file server.

_Required_: No

_Type_: String

_Pattern_: `^([A-Za-z0-9]+[A-Za-z0-9-.]*)*[A-Za-z0-9-]*[A-Za-z0-9]$`

_Maximum_: `253`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KerberosKeytab`

Property description not available.

_Required_: No

_Type_: String

_Maximum_: `87384`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KerberosKrb5Conf`

Property description not available.

_Required_: No

_Type_: String

_Maximum_: `174764`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KerberosPrincipal`

Property description not available.

_Required_: No

_Type_: String

_Pattern_: `^.+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MountOptions`

Specifies the version of the SMB protocol that DataSync uses to access your SMB
file server.

_Required_: No

_Type_: [MountOptions](aws-properties-datasync-locationsmb-mountoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Password`

Specifies the password of the user who can mount your SMB file server and has permission
to access the files and folders involved in your transfer. This parameter applies only if
`AuthenticationType` is set to `NTLM`.

_Required_: No

_Type_: String

_Pattern_: `^.{0,104}$`

_Maximum_: `104`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerHostname`

Specifies the domain name or IP address (IPv4 or IPv6) of the SMB file server that your DataSync agent connects to.

###### Note

If you're using Kerberos authentication, you must specify a domain name.

_Required_: No

_Type_: String

_Pattern_: `^(([a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9\-]*[A-Za-z0-9])$`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Subdirectory`

Specifies the name of the share exported by your SMB file server where DataSync
will read or write data. You can include a subdirectory in the share path (for example,
`/path/to/subdirectory`). Make sure that other SMB clients in your network can
also mount this path.

To copy all data in the subdirectory, DataSync must be able to mount the SMB
share and access all of its data. For more information, see [Providing DataSync access to SMB file servers](../../../datasync/latest/userguide/create-smb-location.md#configuring-smb-permissions).

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-\+\./\(\)\$\p{Zs}]+$`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Specifies labels that help you categorize, filter, and search for your AWS
resources. We recommend creating at least a name tag for your location.

_Required_: No

_Type_: Array of [Tag](aws-properties-datasync-locationsmb-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`User`

Specifies the user that can mount and access the files, folders, and file metadata in your
SMB file server. This parameter applies only if `AuthenticationType` is set to
`NTLM`.

For information about choosing a user with the right level of access for your transfer,
see [Providing DataSync access to SMB file servers](../../../datasync/latest/userguide/create-smb-location.md#configuring-smb-permissions).

_Required_: No

_Type_: String

_Pattern_: `^[^\x5B\x5D\\/:;|=,+*?]{1,104}$`

_Maximum_: `104`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref`
function, `Ref` returns the location resource Amazon Resource Name (ARN). For
example:

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

The Amazon Resource Name (ARN) of the specified SMB location.

`LocationUri`

The URI of the specified SMB location.

## Examples

### Creating an SMB location

The following example specifies an SMB location for DataSync. In
this example, the SMB location uses the domain `EXAMPLE` with SMB
version 3. The server hostname is `MyServer@example.com`, and the SMB
location is in the `/share` subdirectory. This example also specifies
the user ID `user-1` for NTLM authentication.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "Creates an SMB location for DataSync",
    "Resources": {
        "LocationSMB": {
            "Type": "AWS::DataSync::LocationSMB",
            "Properties": {
                "AgentArns": [
                    "arn:aws:datasync:us-east-2:111222333444:agent/agent-0b0addbeef44b3nfs,",
                    "arn:aws:datasync:us-east-2:111222333444:agent/agent-2345noo35nnee1123ovo3"
                ],
                "Domain": "EXAMPLE",
                "AuthenticationType": "NTLM",
                "MountOptions": {
                    "Version": "SMB3"
                },
                "Password": "Password",
                "ServerHostname": "MyServer.example.com",
                "Subdirectory": "/share",
                "User": "user-1"
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: Creates an SMB location for DataSync
Resources:
  LocationSMB:
    Type: AWS::DataSync::LocationSMB
    Properties:
      AgentArns:
        - arn:aws:datasync:us-east-2:111222333444:agent/agent-0b0addbeef44b3nfs,
        - arn:aws:datasync:us-east-2:111222333444:agent/agent-2345noo35nnee1123ovo3
      Domain: EXAMPLE
      AuthenticationType: NTLM
      MountOptions:
        Version: SMB3
      Password: Password
      ServerHostname: MyServer.example.com
      Subdirectory: /share
      User: user-1
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

CmkSecretConfig

All content copied from https://docs.aws.amazon.com/.
