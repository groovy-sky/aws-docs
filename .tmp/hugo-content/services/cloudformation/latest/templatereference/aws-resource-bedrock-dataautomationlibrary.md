This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataAutomationLibrary

A data automation library.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Bedrock::DataAutomationLibrary",
  "Properties" : {
      "EncryptionConfiguration" : EncryptionConfiguration,
      "LibraryDescription" : String,
      "LibraryName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Bedrock::DataAutomationLibrary
Properties:
  EncryptionConfiguration:
    EncryptionConfiguration
  LibraryDescription: String
  LibraryName: String
  Tags:
    - Tag

```

## Properties

`EncryptionConfiguration`

Encryption settings for an invocation.

_Required_: No

_Type_: [EncryptionConfiguration](aws-properties-bedrock-dataautomationlibrary-encryptionconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LibraryDescription`

The library's description.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LibraryName`

The library's name.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9-_]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](aws-properties-bedrock-dataautomationlibrary-tag.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`CreationTime`

When the library was created.

`EntityTypes`

The entity types supported by the library.

`LibraryArn`

The library's ARN.

`Status`

The library's status.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

EncryptionConfiguration

All content copied from https://docs.aws.amazon.com/.
