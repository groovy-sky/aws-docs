This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::LocationAzureBlob AzureBlobSasConfiguration

The shared access signature (SAS) configuration that allows AWS DataSync to
access your Microsoft Azure Blob Storage.

For more information, see [SAS\
tokens](../../../datasync/latest/userguide/creating-azure-blob-location.md#azure-blob-sas-tokens) for accessing your Azure Blob Storage.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AzureBlobSasToken" : String
}

```

### YAML

```yaml

  AzureBlobSasToken: String

```

## Properties

`AzureBlobSasToken`

Specifies a SAS token that provides permissions to access your Azure Blob
Storage.

The token is part of the SAS URI string that comes after the storage resource URI and
a question mark. A token looks something like this:

`sp=r&st=2023-12-20T14:54:52Z&se=2023-12-20T22:54:52Z&spr=https&sv=2021-06-08&sr=c&sig=aBBKDWQvyuVcTPH9EBp%2FXTI9E%2F%2Fmq171%2BZU178wcwqU%3D`

_Required_: Yes

_Type_: String

_Pattern_: `(^.+$)`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::DataSync::LocationAzureBlob

CmkSecretConfig

All content copied from https://docs.aws.amazon.com/.
