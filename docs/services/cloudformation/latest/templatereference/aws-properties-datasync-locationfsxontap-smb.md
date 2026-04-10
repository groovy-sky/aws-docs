This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::LocationFSxONTAP SMB

Specifies the Server Message Block (SMB) protocol configuration that AWS DataSync uses to access a storage virtual machine (SVM) on your Amazon FSx for
NetApp ONTAP file system. For more information, see [Accessing FSx for ONTAP file systems](../../../datasync/latest/userguide/create-ontap-location.md#create-ontap-location-access).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CmkSecretConfig" : CmkSecretConfig,
  "CustomSecretConfig" : CustomSecretConfig,
  "Domain" : String,
  "ManagedSecretConfig" : ManagedSecretConfig,
  "MountOptions" : SmbMountOptions,
  "Password" : String,
  "User" : String
}

```

### YAML

```yaml

  CmkSecretConfig:
    CmkSecretConfig
  CustomSecretConfig:
    CustomSecretConfig
  Domain: String
  ManagedSecretConfig:
    ManagedSecretConfig
  MountOptions:
    SmbMountOptions
  Password: String
  User: String

```

## Properties

`CmkSecretConfig`

Property description not available.

_Required_: No

_Type_: [CmkSecretConfig](aws-properties-datasync-locationfsxontap-cmksecretconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomSecretConfig`

Property description not available.

_Required_: No

_Type_: [CustomSecretConfig](aws-properties-datasync-locationfsxontap-customsecretconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Domain`

Specifies the name of the Windows domain that your storage virtual machine (SVM) belongs
to.

If you have multiple domains in your environment, configuring this setting makes sure that
DataSync connects to the right SVM.

If you have multiple Active Directory domains in your environment, configuring this
parameter makes sure that DataSync connects to the right SVM.

_Required_: No

_Type_: String

_Pattern_: `^([A-Za-z0-9]+[A-Za-z0-9-.]*)*[A-Za-z0-9-]*[A-Za-z0-9]$`

_Maximum_: `253`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManagedSecretConfig`

Property description not available.

_Required_: No

_Type_: [ManagedSecretConfig](aws-properties-datasync-locationfsxontap-managedsecretconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MountOptions`

Specifies how DataSync can access a location using the SMB protocol.

_Required_: Yes

_Type_: [SmbMountOptions](aws-properties-datasync-locationfsxontap-smbmountoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Password`

Specifies the password of a user who has permission to access your SVM.

_Required_: No

_Type_: String

_Pattern_: `^.{0,104}$`

_Maximum_: `104`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`User`

Specifies a user name that can mount the location and access the files, folders, and
metadata that you need in the SVM.

If you provide a user in your Active Directory, note the following:

- If you're using AWS Directory Service for Microsoft Active Directory, the user must be a member of
the AWS Delegated FSx Administrators group.

- If you're using a self-managed Active Directory, the user must be a member of
either the Domain Admins group or a custom group that you specified for file
system administration when you created your file system.

Make sure that the user has the permissions it needs to copy the data you want:

- `SE_TCB_NAME`: Required to set object ownership and file metadata.
With this privilege, you also can copy NTFS discretionary access lists
(DACLs).

- `SE_SECURITY_NAME`: May be needed to copy NTFS system access
control lists (SACLs). This operation specifically requires the Windows
privilege, which is granted to members of the Domain Admins group. If you
configure your task to copy SACLs, make sure that the user has the required
privileges. For information about copying SACLs, see [Ownership and permissions-related options](../../../datasync/latest/userguide/create-task.md#configure-ownership-and-permissions).

_Required_: Yes

_Type_: String

_Pattern_: `^[^\x5B\x5D\\/:;|=,+*?]{1,104}$`

_Maximum_: `104`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Protocol

SmbMountOptions

All content copied from https://docs.aws.amazon.com/.
