This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Job JobSample

A sample configuration for profile jobs only, which determines the number of rows on which the
profile job is run. If a `JobSample` value isn't provided, the
default is used. The default value is CUSTOM\_ROWS for the mode parameter and
20,000 for the size parameter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Mode" : String,
  "Size" : Integer
}

```

### YAML

```yaml

  Mode: String
  Size: Integer

```

## Properties

`Mode`

A value that determines whether the profile job is run on the entire dataset or a
specified number of rows. This value must be one of the following:

- FULL\_DATASET - The profile job is run on the entire dataset.

- CUSTOM\_ROWS - The profile job is run on the number of rows specified in the
`Size` parameter.

_Required_: No

_Type_: String

_Allowed values_: `FULL_DATASET | CUSTOM_ROWS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Size`

The `Size` parameter is only required when the mode is CUSTOM\_ROWS. The
profile job is run on the specified number of rows. The maximum value for size is
Long.MAX\_VALUE.

Long.MAX\_VALUE = 9223372036854775807

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EntityDetectorConfiguration

Output
