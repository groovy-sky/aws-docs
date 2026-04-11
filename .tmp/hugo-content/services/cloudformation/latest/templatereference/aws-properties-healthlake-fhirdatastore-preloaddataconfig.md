This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::HealthLake::FHIRDatastore PreloadDataConfig

An optional parameter to preload (import) open source Synthea FHIR data upon creation of
the data store.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PreloadDataType" : String
}

```

### YAML

```yaml

  PreloadDataType: String

```

## Properties

`PreloadDataType`

The type of preloaded data. Only Synthea preloaded data is supported.

_Required_: Yes

_Type_: String

_Allowed values_: `SYNTHEA`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KmsEncryptionConfig

SseConfiguration

All content copied from https://docs.aws.amazon.com/.
