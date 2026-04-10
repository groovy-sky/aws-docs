This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmplifyUIBuilder::Form FileUploaderFieldConfig

Describes the configuration for the file uploader field.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AcceptedFileTypes" : [ String, ... ],
  "AccessLevel" : String,
  "IsResumable" : Boolean,
  "MaxFileCount" : Number,
  "MaxSize" : Number,
  "ShowThumbnails" : Boolean
}

```

### YAML

```yaml

  AcceptedFileTypes:
    - String
  AccessLevel: String
  IsResumable: Boolean
  MaxFileCount: Number
  MaxSize: Number
  ShowThumbnails: Boolean

```

## Properties

`AcceptedFileTypes`

The file types that are allowed to be uploaded by the file uploader. Provide this
information in an array of strings specifying the valid file extensions.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AccessLevel`

The access level to assign to the uploaded files in the Amazon S3 bucket where
they are stored. The valid values for this property are `private`,
`protected`, or `public`. For detailed information about the
permissions associated with each access level, see [File access\
levels](https://docs.amplify.aws/lib/storage/configureaccess/q/platform/js) in the _Amplify documentation_.

_Required_: Yes

_Type_: String

_Allowed values_: `public | protected | private`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsResumable`

Allows the file upload operation to be paused and resumed. The default value is
`false`.

When `isResumable` is set to `true`, the file uploader uses a
multipart upload to break the files into chunks before upload. The progress of the upload
isn't continuous, because the file uploader uploads a chunk at a time.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxFileCount`

Specifies the maximum number of files that can be selected to upload. The default value is
an unlimited number of files.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxSize`

The maximum file size in bytes that the file uploader will accept. The default value is an
unlimited file size.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShowThumbnails`

Specifies whether to display or hide the image preview after selecting a file for upload.
The default value is `true` to display the image preview.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FieldValidationConfiguration

FormButton

All content copied from https://docs.aws.amazon.com/.
