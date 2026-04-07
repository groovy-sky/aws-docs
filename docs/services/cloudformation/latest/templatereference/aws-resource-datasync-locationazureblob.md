This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::LocationAzureBlob

Creates a transfer _location_ for a Microsoft Azure Blob Storage
container. AWS DataSync can use this location as a transfer source or destination.
You can make transfers with or without a [DataSync agent](https://docs.aws.amazon.com/datasync/latest/userguide/creating-azure-blob-location.html#azure-blob-creating-agent) that connects to your
container.

Before you begin, make sure you know [how DataSync accesses Azure Blob Storage](https://docs.aws.amazon.com/datasync/latest/userguide/creating-azure-blob-location.html#azure-blob-access) and works with [access tiers](https://docs.aws.amazon.com/datasync/latest/userguide/creating-azure-blob-location.html#azure-blob-access-tiers) and [blob types](https://docs.aws.amazon.com/datasync/latest/userguide/creating-azure-blob-location.html#blob-types).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataSync::LocationAzureBlob",
  "Properties" : {
      "AgentArns" : [ String, ... ],
      "AzureAccessTier" : String,
      "AzureBlobAuthenticationType" : String,
      "AzureBlobContainerUrl" : String,
      "AzureBlobSasConfiguration" : AzureBlobSasConfiguration,
      "AzureBlobType" : String,
      "CmkSecretConfig" : CmkSecretConfig,
      "CustomSecretConfig" : CustomSecretConfig,
      "Subdirectory" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DataSync::LocationAzureBlob
Properties:
  AgentArns:
    - String
  AzureAccessTier: String
  AzureBlobAuthenticationType: String
  AzureBlobContainerUrl: String
  AzureBlobSasConfiguration:
    AzureBlobSasConfiguration
  AzureBlobType: String
  CmkSecretConfig:
    CmkSecretConfig
  CustomSecretConfig:
    CustomSecretConfig
  Subdirectory: String
  Tags:
    - Tag

```

## Properties

`AgentArns`

(Optional) Specifies the Amazon Resource Name (ARN) of the DataSync agent that
can connect with your Azure Blob Storage container. If you are setting up an agentless
cross-cloud transfer, you do not need to specify a value for this parameter.

You can specify more than one agent. For more information, see [Using multiple\
agents for your transfer](https://docs.aws.amazon.com/datasync/latest/userguide/multiple-agents.html).

###### Note

Make sure you configure this parameter correctly when you first create your storage
location. You cannot add or remove agents from a storage location after you create
it.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `128 | 4`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AzureAccessTier`

Specifies the access tier that you want your objects or files transferred into. This only
applies when using the location as a transfer destination. For more information, see [Access tiers](https://docs.aws.amazon.com/datasync/latest/userguide/creating-azure-blob-location.html#azure-blob-access-tiers).

_Required_: No

_Type_: String

_Allowed values_: `HOT | COOL | ARCHIVE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AzureBlobAuthenticationType`

Specifies the authentication method DataSync uses to access your Azure Blob
Storage. DataSync can access blob storage using a shared access signature
(SAS).

_Required_: Yes

_Type_: String

_Allowed values_: `SAS | NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AzureBlobContainerUrl`

Specifies the URL of the Azure Blob Storage container involved in your transfer.

_Required_: No

_Type_: String

_Pattern_: `^https://[A-Za-z0-9]((.|-+)?[A-Za-z0-9]){0,252}/[a-z0-9](-?[a-z0-9]){2,62}$`

_Maximum_: `325`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AzureBlobSasConfiguration`

Specifies the SAS configuration that allows DataSync to access your Azure Blob
Storage.

###### Note

If you provide an authentication token using `SasConfiguration`, but do not
provide secret configuration details using `CmkSecretConfig` or
`CustomSecretConfig`, then DataSync stores the token using your
AWS account's secrets manager secret.

_Required_: No

_Type_: [AzureBlobSasConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datasync-locationazureblob-azureblobsasconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AzureBlobType`

Specifies the type of blob that you want your objects or files to be when transferring
them into Azure Blob Storage. Currently, DataSync only supports moving data into
Azure Blob Storage as block blobs. For more information on blob types, see the [Azure Blob Storage documentation](https://learn.microsoft.com/en-us/rest/api/storageservices/understanding-block-blobs--append-blobs--and-page-blobs).

_Required_: No

_Type_: String

_Allowed values_: `BLOCK`

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

_Type_: [CmkSecretConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datasync-locationazureblob-cmksecretconfig.html)

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

_Type_: [CustomSecretConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datasync-locationazureblob-customsecretconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Subdirectory`

Specifies path segments if you want to limit your transfer to a virtual directory in your
container (for example, `/my/images`).

_Required_: No

_Type_: String

_Pattern_: `^[\p{L}\p{M}\p{Z}\p{S}\p{N}\p{P}\p{C}]*$`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Specifies labels that help you categorize, filter, and search for your AWS
resources. We recommend creating at least a name tag for your transfer location.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datasync-locationazureblob-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the location resource ARN. For example:

`arn:aws:datasync:us-east-2:111222333444:location/loc-07db7abfc326c50s3`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CmkSecretConfig.SecretArn`

Property description not available.

`LocationArn`

The ARN of the Azure Blob Storage transfer location that you created.

`LocationUri`

The URI of the Azure Blob Storage transfer location that you created.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AzureBlobSasConfiguration
