This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::StorageVirtualMachine ActiveDirectoryConfiguration

Describes the self-managed Microsoft Active Directory to which you want to join the SVM.
Joining an Active Directory provides user authentication and access control for SMB clients,
including Microsoft Windows and macOS clients accessing the file system.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NetBiosName" : String,
  "SelfManagedActiveDirectoryConfiguration" : SelfManagedActiveDirectoryConfiguration
}

```

### YAML

```yaml

  NetBiosName: String
  SelfManagedActiveDirectoryConfiguration:
    SelfManagedActiveDirectoryConfiguration

```

## Properties

`NetBiosName`

The NetBIOS name of the Active Directory computer object that will be created for your SVM.

_Required_: No

_Type_: String

_Pattern_: `^[^\u0000\u0085\u2028\u2029\r\n]{1,255}$`

_Minimum_: `1`

_Maximum_: `15`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelfManagedActiveDirectoryConfiguration`

The configuration that Amazon FSx uses to join the ONTAP storage virtual machine
(SVM) to your self-managed (including on-premises) Microsoft Active Directory directory.

_Required_: No

_Type_: [SelfManagedActiveDirectoryConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-fsx-storagevirtualmachine-selfmanagedactivedirectoryconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::FSx::StorageVirtualMachine

SelfManagedActiveDirectoryConfiguration
