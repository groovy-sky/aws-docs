This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::StorageVirtualMachine

Creates a storage virtual machine (SVM) for an Amazon FSx for ONTAP file system.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::FSx::StorageVirtualMachine",
  "Properties" : {
      "ActiveDirectoryConfiguration" : ActiveDirectoryConfiguration,
      "FileSystemId" : String,
      "Name" : String,
      "RootVolumeSecurityStyle" : String,
      "SvmAdminPassword" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::FSx::StorageVirtualMachine
Properties:
  ActiveDirectoryConfiguration:
    ActiveDirectoryConfiguration
  FileSystemId: String
  Name: String
  RootVolumeSecurityStyle: String
  SvmAdminPassword: String
  Tags:
    - Tag

```

## Properties

`ActiveDirectoryConfiguration`

Describes the Microsoft Active Directory configuration to which the SVM is joined, if applicable.

_Required_: No

_Type_: [ActiveDirectoryConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-fsx-storagevirtualmachine-activedirectoryconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FileSystemId`

Specifies the FSx for ONTAP file system on which to create the SVM.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the SVM.

_Required_: Yes

_Type_: String

_Pattern_: `^[^\u0000\u0085\u2028\u2029\r\n]{1,47}$`

_Minimum_: `1`

_Maximum_: `47`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RootVolumeSecurityStyle`

The security style of the root volume of the SVM. Specify one of the following values:

- `UNIX` if the file system is managed by a UNIX
administrator, the majority of users are NFS clients, and an application
accessing the data uses a UNIX user as the service account.

- `NTFS` if the file system is managed by a Microsoft Windows
administrator, the majority of users are SMB clients, and an application
accessing the data uses a Microsoft Windows user as the service account.

- `MIXED` This is an advanced setting. For more information, see
[Volume security style](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/volume-security-style.html)
in the Amazon FSx for NetApp ONTAP User Guide.

_Required_: No

_Type_: String

_Allowed values_: `UNIX | NTFS | MIXED`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SvmAdminPassword`

Specifies the password to use when logging on to the SVM using a secure shell (SSH) connection to the SVM's management endpoint.
Doing so enables you to manage the SVM using the NetApp ONTAP CLI or REST API. If you do not specify a password, you can still
use the file system's `fsxadmin` user to manage the SVM. For more information, see
[Managing SVMs using the NetApp ONTAP CLI](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/managing-resources-ontap-apps.html#vsadmin-ontap-cli) in the _FSx for ONTAP User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of `Tag` values, with a maximum of 50 elements.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-fsx-storagevirtualmachine-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource ID, such as `svm-01234567890123456`.
For example:

`{"Ref": "svm_logical_id"}` returns

`svm-01234567890123456`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ResourceARN`

Returns the storage virtual machine's Amazon Resource Name (ARN).

Example: `arn:aws:fsx:us-east-2:111111111111:storage-virtual-machine/fs-0123456789abcdef1/svm-01234567890123456`

`StorageVirtualMachineId`

Returns the storgage virtual machine's system generated ID.

Example: `svm-0123456789abcedf1`

`UUID`

Returns the storage virtual machine's system generated unique identifier (UUID).

Example: `abcd0123-cd45-ef67-11aa-1111aaaa23bc`

## Examples

### Create an Amazon FSx for NetApp ONTAP Storage Virtual Machine

The following examples create an Amazon FSx for NetApp ONTAP storage
virtual machine (SVN) that's joined to a self-managed Active Directory
domain.

#### JSON

```json

{
  "OntapStorageVirtualMachineWithAllConfigs": {
    "Type": "AWS::FSx::StorageVirtualMachine",
    "Properties": {
      "ActiveDirectoryConfiguration": {
        "NetBiosName": "svm1",
        "SelfManagedActiveDirectoryConfiguration": {
          "DnsIps": [
            "10.0.10.67"
          ],
          "DomainName": "CFN-CUSTOMER-AD.SIMBA.LOCAL",
          "FileSystemAdministratorsGroup": "Domain Admins",
          "OrganizationalUnitDistinguishedName": "OU=cfn-customer-ad,DC=cfn-customer-ad,DC=simba,DC=local",
          "Password": {
            "Fn::Join": [
              ":",
              [
                "{{resolve:secretsmanager",
                {
                  "Fn::ImportValue": "CustomerADCredentialName"
              },
              "SecretString}}"
            ]
          ]
        },
        "UserName": "Admin"
      }
    },
    "FileSystemId": {
      "Ref": "OntapMultiAzFileSystemWithAllConfigs"
    },
    "Name": "svm1",
    "RootVolumeSecurityStyle": "UNIX",
    "SvmAdminPassword": {
      "Password": {
        "Fn::Join": [
        ":",
        [
          "{{resolve:secretsmanager",
            {
              "Fn::ImportValue": "CustomerADCredentialName"
              },
              "SecretString}}"
            ]
          ]
        }
      },
      "Tags": [
        {
          "Key": "Name",
          "Value": "OntapSvm"
        }
      ]
    }
  }
}
```

#### YAML

```yaml

OntapStorageVirtualMachineWithAllConfigs:
  Type: "AWS::FSx::StorageVirtualMachine"
  Properties:
    ActiveDirectoryConfiguration:
      NetBiosName: "svm1"
      SelfManagedActiveDirectoryConfiguration:
        DnsIps: ["10.0.10.67"]
        DomainName: "CFN-CUSTOMER-AD.SIMBA.LOCAL"
        FileSystemAdministratorsGroup: "Domain Admins"
        OrganizationalUnitDistinguishedName: "OU=cfn-customer-ad,DC=cfn-customer-ad,DC=simba,DC=local"
        Password:
          !Join
          - ':'
          - - '{{resolve:secretsmanager'
            - !ImportValue CustomerADCredentialName
            - 'SecretString}}'
        UserName: "Admin"
    FileSystemId: !Ref OntapMultiAzFileSystemWithAllConfigs
    Name: "svm1"
    RootVolumeSecurityStyle: "UNIX"
    SvmAdminPassword:
      Password:
        !Join
        - ':'
        - - '{{resolve:secretsmanager'
          - !ImportValue CustomerADCredentialName
          - 'SecretString}}'
    Tags:
      - Key: "Name"
        Value: "OntapSvm"
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

ActiveDirectoryConfiguration
