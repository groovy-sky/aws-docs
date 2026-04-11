This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::Task ManifestConfig

Configures a manifest, which is a list of files or objects that you want AWS DataSync to transfer. For more information and configuration examples, see [Specifying what DataSync transfers by using a manifest](../../../datasync/latest/userguide/transferring-with-manifest.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : String,
  "Format" : String,
  "Source" : Source
}

```

### YAML

```yaml

  Action: String
  Format: String
  Source:
    Source

```

## Properties

`Action`

Specifies what DataSync uses the manifest for.

_Required_: No

_Type_: String

_Allowed values_: `TRANSFER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Format`

Specifies the file format of your manifest. For more information, see [Creating a manifest](../../../datasync/latest/userguide/transferring-with-manifest.md#transferring-with-manifest-create).

_Required_: No

_Type_: String

_Allowed values_: `CSV`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

Specifies the manifest that you want DataSync to use and where it's
hosted.

###### Note

You must specify this parameter if you're configuring a new manifest on or after
February 7, 2024.

If you don't, you'll get a 400 status code and `ValidationException` error
stating that you're missing the IAM role for DataSync to access the
S3 bucket where you're hosting your manifest. For more information, see [Providing DataSync access to your manifest](../../../datasync/latest/userguide/transferring-with-manifest.md#transferring-with-manifest-access).

_Required_: Yes

_Type_: [Source](aws-properties-datasync-task-source.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FilterRule

ManifestConfigSourceS3

All content copied from https://docs.aws.amazon.com/.
