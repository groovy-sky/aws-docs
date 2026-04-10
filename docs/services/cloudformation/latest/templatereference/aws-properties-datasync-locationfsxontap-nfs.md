This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::LocationFSxONTAP NFS

Specifies the Network File System (NFS) protocol configuration that AWS DataSync uses to access a storage virtual machine (SVM) on your Amazon FSx for
NetApp ONTAP file system. For more information, see [Accessing FSx for ONTAP file systems](../../../datasync/latest/userguide/create-ontap-location.md#create-ontap-location-access).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MountOptions" : NfsMountOptions
}

```

### YAML

```yaml

  MountOptions:
    NfsMountOptions

```

## Properties

`MountOptions`

Specifies how DataSync can access a location using the NFS protocol.

_Required_: Yes

_Type_: [NfsMountOptions](aws-properties-datasync-locationfsxontap-nfsmountoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ManagedSecretConfig

NfsMountOptions

All content copied from https://docs.aws.amazon.com/.
