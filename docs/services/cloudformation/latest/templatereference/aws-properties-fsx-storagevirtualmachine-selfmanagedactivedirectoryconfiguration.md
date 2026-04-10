This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::StorageVirtualMachine SelfManagedActiveDirectoryConfiguration

The configuration that Amazon FSx uses to join the ONTAP storage virtual machine
(SVM) to your self-managed (including on-premises) Microsoft Active Directory directory.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DnsIps" : [ String, ... ],
  "DomainJoinServiceAccountSecret" : String,
  "DomainName" : String,
  "FileSystemAdministratorsGroup" : String,
  "OrganizationalUnitDistinguishedName" : String,
  "Password" : String,
  "UserName" : String
}

```

### YAML

```yaml

  DnsIps:
    - String
  DomainJoinServiceAccountSecret: String
  DomainName: String
  FileSystemAdministratorsGroup: String
  OrganizationalUnitDistinguishedName: String
  Password: String
  UserName: String

```

## Properties

`DnsIps`

A list of IP addresses of DNS servers or domain controllers in the
self-managed AD directory. FSx for ONTAP SVMs support up to three IP addresses, FSx for Windows supports up to two.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainJoinServiceAccountSecret`

The Amazon Resource Name (ARN) of the AWS Secrets Manager secret containing the self-managed Active Directory domain join service account credentials.
When provided, Amazon FSx uses the credentials stored in this secret to join the file system to your self-managed Active Directory domain.

The secret must contain two key-value pairs:

- `CUSTOMER_MANAGED_ACTIVE_DIRECTORY_USERNAME` \- The username for the service account

- `CUSTOMER_MANAGED_ACTIVE_DIRECTORY_PASSWORD` \- The password for the service account

For more information, see
[Using Amazon FSx for Windows with your self-managed Microsoft Active Directory](../../../fsx/latest/windowsguide/self-manage-prereqs.md) or
[Using Amazon FSx for ONTAP with your self-managed Microsoft Active Directory](../../../fsx/latest/ontapguide/self-manage-prereqs.md).

_Required_: No

_Type_: String

_Pattern_: `^arn:[^:]{1,63}:secretsmanager:[a-z0-9-]+:[0-9]{12}:secret:[a-zA-Z0-9/_+=.@-]+-[a-zA-Z0-9]{6}$`

_Minimum_: `64`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainName`

The fully qualified domain name of the self-managed AD directory, such as
`corp.example.com`.

_Required_: No

_Type_: String

_Pattern_: `^[^\u0000\u0085\u2028\u2029\r\n]{1,255}$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FileSystemAdministratorsGroup`

(Optional) The name of the domain group whose members are granted administrative
privileges for the file system. Administrative privileges include taking ownership of
files and folders, setting audit controls (audit ACLs) on files and folders, and
administering the file system remotely by using the FSx Remote PowerShell.
The group that you specify must already exist in your domain. If you don't provide one,
your AD domain's Domain Admins group is used.

_Required_: No

_Type_: String

_Pattern_: `^[^\u0000\u0085\u2028\u2029\r\n]{1,256}$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OrganizationalUnitDistinguishedName`

(Optional) The fully qualified distinguished name of the organizational unit within
your self-managed AD directory. Amazon
FSx only accepts OU as the direct parent of the file system. An example is
`OU=FSx,DC=yourdomain,DC=corp,DC=com`. To learn more, see [RFC 2253](https://tools.ietf.org/html/rfc2253). If none is provided, the
FSx file system is created in the default location of your self-managed AD directory.

###### Important

Only Organizational Unit (OU) objects can be the direct parent of the file system
that you're creating.

_Required_: No

_Type_: String

_Pattern_: `^[^\u0000\u0085\u2028\u2029\r\n]{1,2000}$`

_Minimum_: `1`

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Password`

The password for the service account on your self-managed AD domain that Amazon FSx
will use to join to your AD domain.

_Required_: No

_Type_: String

_Pattern_: `^.{1,256}$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserName`

The user name for the service account on your self-managed AD domain that Amazon FSx
will use to join to your AD domain. This account must have the permission to join
computers to the domain in the organizational unit provided in
`OrganizationalUnitDistinguishedName`, or in the default location of your
AD domain.

_Required_: No

_Type_: String

_Pattern_: `^[^\u0000\u0085\u2028\u2029\r\n]{1,256}$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ActiveDirectoryConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
