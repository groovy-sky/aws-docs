This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WorkSpacesWeb::UserSettings ImageMetadata

Metadata information about an uploaded image file.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FileExtension" : String,
  "LastUploadTimestamp" : String,
  "MimeType" : String
}

```

### YAML

```yaml

  FileExtension: String
  LastUploadTimestamp: String
  MimeType: String

```

## Properties

`FileExtension`

The file extension of the image.

_Required_: Yes

_Type_: String

_Pattern_: `^[\s\S]*$`

_Minimum_: `0`

_Maximum_: `131072`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LastUploadTimestamp`

The timestamp when the image was last uploaded.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MimeType`

The MIME type of the image.

_Required_: Yes

_Type_: String

_Allowed values_: `image/png | image/jpeg | image/x-icon`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CookieSynchronizationConfiguration

LocalizedBrandingStrings

All content copied from https://docs.aws.amazon.com/.
