This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodePipeline::Pipeline ArtifactStore

The S3 bucket where artifacts for the pipeline are stored.

###### Note

You must include either `artifactStore` or
`artifactStores` in your pipeline, but you cannot use both. If you
create a cross-region action in your pipeline, you must use
`artifactStores`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EncryptionKey" : EncryptionKey,
  "Location" : String,
  "Type" : String
}

```

### YAML

```yaml

  EncryptionKey:
    EncryptionKey
  Location: String
  Type: String

```

## Properties

`EncryptionKey`

The encryption key used to encrypt the data in the artifact store, such as an AWS Key Management Service (AWS KMS) key. If this is undefined,
the default key for Amazon S3 is used. To see an example artifact store encryption key field,
see the example structure here: [AWS::CodePipeline::Pipeline](../userguide/aws-resource-codepipeline-pipeline.md).

_Required_: No

_Type_: [EncryptionKey](aws-properties-codepipeline-pipeline-encryptionkey.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Location`

The S3 bucket used for storing the artifacts for a pipeline. You can specify the
name of an S3 bucket but not a folder in the bucket. A folder to contain the pipeline
artifacts is created for you based on the name of the pipeline. You can use any S3
bucket in the same AWS Region as the pipeline to store your pipeline
artifacts.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9\-\.]+`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the artifact store, such as S3.

_Required_: Yes

_Type_: String

_Allowed values_: `S3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ActionTypeId

ArtifactStoreMap

All content copied from https://docs.aws.amazon.com/.
