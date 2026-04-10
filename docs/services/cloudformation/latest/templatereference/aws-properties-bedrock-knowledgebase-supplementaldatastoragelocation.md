This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::KnowledgeBase SupplementalDataStorageLocation

Contains information about a storage location for multimedia content (images, audio, and video) extracted from multimodal documents in your data source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3Location" : S3Location,
  "SupplementalDataStorageLocationType" : String
}

```

### YAML

```yaml

  S3Location:
    S3Location
  SupplementalDataStorageLocationType: String

```

## Properties

`S3Location`

Contains information about the Amazon S3 location for the extracted multimedia content.

_Required_: No

_Type_: [S3Location](aws-properties-bedrock-knowledgebase-s3location.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SupplementalDataStorageLocationType`

Property description not available.

_Required_: Yes

_Type_: String

_Allowed values_: `S3`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SupplementalDataStorageConfiguration

VectorKnowledgeBaseConfiguration

All content copied from https://docs.aws.amazon.com/.
