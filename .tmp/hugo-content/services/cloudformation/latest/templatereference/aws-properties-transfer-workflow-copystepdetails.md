This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Transfer::Workflow CopyStepDetails

Details for a step that performs a file copy.

Consists of the following values:

- A description

- An Amazon S3 location for the destination of the file copy.

- A flag that indicates whether to overwrite an existing file of the same name.
The default is `FALSE`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DestinationFileLocation" : S3FileLocation,
  "Name" : String,
  "OverwriteExisting" : String,
  "SourceFileLocation" : String
}

```

### YAML

```yaml

  DestinationFileLocation:
    S3FileLocation
  Name: String
  OverwriteExisting: String
  SourceFileLocation: String

```

## Properties

`DestinationFileLocation`

Specifies the location for the file being copied. Use
`${Transfer:UserName}` or `${Transfer:UploadDate}` in this
field to parametrize the destination prefix by username or uploaded date.

- Set the value of `DestinationFileLocation` to
`${Transfer:UserName}` to copy uploaded files to an Amazon S3 bucket that is prefixed with the name of the Transfer Family
user that uploaded the file.

- Set the value of `DestinationFileLocation` to
`${Transfer:UploadDate}` to copy uploaded files to an Amazon S3 bucket that is prefixed with the date of the upload.

###### Note

The system resolves `UploadDate` to a date format of
_YYYY-MM-DD_, based on the date the file is uploaded
in UTC.

_Required_: No

_Type_: [S3FileLocation](aws-properties-transfer-workflow-s3filelocation.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the step, used as an identifier.

_Required_: No

_Type_: String

_Pattern_: `^[\w-]*$`

_Minimum_: `0`

_Maximum_: `30`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OverwriteExisting`

A flag that indicates whether to overwrite an existing file of the same name.
The default is `FALSE`.

If the workflow is processing a file that has the same name as an existing file, the behavior is as follows:

- If `OverwriteExisting` is `TRUE`, the existing file is replaced with the file being processed.

- If `OverwriteExisting` is `FALSE`, nothing happens, and the workflow processing stops.

_Required_: No

_Type_: String

_Allowed values_: `TRUE | FALSE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceFileLocation`

Specifies which file to use as input to the workflow step: either the output from the previous step, or the originally uploaded file
for the workflow.

- To use the previous file as the input, enter `${previous.file}`.
In this case, this workflow step uses the output file from the previous workflow step as input.
This is the default value.

- To use the originally uploaded file location as input for this step, enter `${original.file}`.

_Required_: No

_Type_: String

_Pattern_: `^\$\{(\w+.)+\w+\}$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Transfer::Workflow

CustomStepDetails

All content copied from https://docs.aws.amazon.com/.
