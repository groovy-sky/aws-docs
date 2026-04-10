This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::DataSource ImageExtractionConfiguration

The configuration for extracting semantic meaning from images in documents. For more information, see [Extracting semantic meaning from images and visuals](../../../amazonq/latest/qbusiness-ug/extracting-meaning-from-images.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ImageExtractionStatus" : String
}

```

### YAML

```yaml

  ImageExtractionStatus: String

```

## Properties

`ImageExtractionStatus`

Specify whether to extract semantic meaning from images and visuals from documents.

_Required_: Yes

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HookConfiguration

InlineDocumentEnrichmentConfiguration

All content copied from https://docs.aws.amazon.com/.
