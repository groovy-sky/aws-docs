This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::LocationFSxONTAP Protocol

Specifies the data transfer protocol that AWS DataSync uses to access your
Amazon FSx file system.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NFS" : NFS,
  "SMB" : SMB
}

```

### YAML

```yaml

  NFS:
    NFS
  SMB:
    SMB

```

## Properties

`NFS`

Specifies the Network File System (NFS) protocol configuration that DataSync uses to
access your FSx for ONTAP file system's storage virtual machine (SVM).

_Required_: No

_Type_: [NFS](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datasync-locationfsxontap-nfs.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SMB`

Specifies the Server Message Block (SMB) protocol configuration that DataSync uses to
access your FSx for ONTAP file system's SVM.

_Required_: No

_Type_: [SMB](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datasync-locationfsxontap-smb.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

NfsMountOptions

SMB
