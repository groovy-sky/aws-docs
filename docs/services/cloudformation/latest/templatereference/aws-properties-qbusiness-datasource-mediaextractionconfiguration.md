This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::DataSource MediaExtractionConfiguration

The configuration for extracting information from media in documents.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AudioExtractionConfiguration" : AudioExtractionConfiguration,
  "ImageExtractionConfiguration" : ImageExtractionConfiguration,
  "VideoExtractionConfiguration" : VideoExtractionConfiguration
}

```

### YAML

```yaml

  AudioExtractionConfiguration:
    AudioExtractionConfiguration
  ImageExtractionConfiguration:
    ImageExtractionConfiguration
  VideoExtractionConfiguration:
    VideoExtractionConfiguration

```

## Properties

`AudioExtractionConfiguration`

Configuration settings for extracting and processing audio content from media files.

_Required_: No

_Type_: [AudioExtractionConfiguration](aws-properties-qbusiness-datasource-audioextractionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageExtractionConfiguration`

The configuration for extracting semantic meaning from images in documents.
For more information, see [Extracting semantic meaning from images and visuals](../../../amazonq/latest/qbusiness-ug/extracting-meaning-from-images.md).

_Required_: No

_Type_: [ImageExtractionConfiguration](aws-properties-qbusiness-datasource-imageextractionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VideoExtractionConfiguration`

Configuration settings for extracting and processing video content from media files.

_Required_: No

_Type_: [VideoExtractionConfiguration](aws-properties-qbusiness-datasource-videoextractionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InlineDocumentEnrichmentConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
