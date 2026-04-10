This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard StaticFileSource

The source of the static file.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3Options" : StaticFileS3SourceOptions,
  "UrlOptions" : StaticFileUrlSourceOptions
}

```

### YAML

```yaml

  S3Options:
    StaticFileS3SourceOptions
  UrlOptions:
    StaticFileUrlSourceOptions

```

## Properties

`S3Options`

The structure that contains the Amazon S3 location to download the static file from.

_Required_: No

_Type_: [StaticFileS3SourceOptions](aws-properties-quicksight-dashboard-staticfiles3sourceoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UrlOptions`

The structure that contains the URL to download the static file from.

_Required_: No

_Type_: [StaticFileUrlSourceOptions](aws-properties-quicksight-dashboard-staticfileurlsourceoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StaticFileS3SourceOptions

StaticFileUrlSourceOptions

All content copied from https://docs.aws.amazon.com/.
