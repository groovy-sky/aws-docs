This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::KnowledgeBase SupplementalDataStorageConfiguration

Specifies configurations for the storage location of multimedia content (images, audio, and video) extracted from multimodal documents in your data source. This content can be retrieved and returned to the end user with timestamp references for audio and video segments.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SupplementalDataStorageLocations" : [ SupplementalDataStorageLocation, ... ]
}

```

### YAML

```yaml

  SupplementalDataStorageLocations:
    - SupplementalDataStorageLocation

```

## Properties

`SupplementalDataStorageLocations`

Property description not available.

_Required_: Yes

_Type_: Array of [SupplementalDataStorageLocation](aws-properties-bedrock-knowledgebase-supplementaldatastoragelocation.md)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StorageConfiguration

SupplementalDataStorageLocation

All content copied from https://docs.aws.amazon.com/.
