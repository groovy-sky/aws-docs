This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Domain SharingSettings

Specifies options when sharing an Amazon SageMaker Studio notebook. These settings are specified as part of
`DefaultUserSettings` when the [CreateDomain](../../../../reference/sagemaker/latest/apireference/api-createdomain.md) API is called, and as part of
`UserSettings` when the [CreateUserProfile](../../../../reference/sagemaker/latest/apireference/api-createuserprofile.md) API is called.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NotebookOutputOption" : String,
  "S3KmsKeyId" : String,
  "S3OutputPath" : String
}

```

### YAML

```yaml

  NotebookOutputOption: String
  S3KmsKeyId: String
  S3OutputPath: String

```

## Properties

`NotebookOutputOption`

Whether to include the notebook cell output when sharing the notebook. The default is
`Disabled`.

_Required_: No

_Type_: String

_Allowed values_: `Allowed | Disabled`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3KmsKeyId`

When `NotebookOutputOption` is `Allowed`, the AWS Key
Management Service (KMS) encryption key ID used to encrypt the notebook cell output in the
Amazon S3 bucket.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3OutputPath`

When `NotebookOutputOption` is `Allowed`, the Amazon S3
bucket used to store the shared notebook snapshots.

_Required_: No

_Type_: String

_Pattern_: `^(https|s3)://([^/]+)/?(.*)$`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3FileSystemConfig

StudioWebPortalSettings

All content copied from https://docs.aws.amazon.com/.
