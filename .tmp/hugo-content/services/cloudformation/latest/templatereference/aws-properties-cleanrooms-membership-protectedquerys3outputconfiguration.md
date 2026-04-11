This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::Membership ProtectedQueryS3OutputConfiguration

Contains the configuration to write the query results to S3.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Bucket" : String,
  "KeyPrefix" : String,
  "ResultFormat" : String,
  "SingleFileOutput" : Boolean
}

```

### YAML

```yaml

  Bucket: String
  KeyPrefix: String
  ResultFormat: String
  SingleFileOutput: Boolean

```

## Properties

`Bucket`

The S3 bucket to unload the protected query results.

_Required_: Yes

_Type_: String

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KeyPrefix`

The S3 prefix to unload the protected query results.

_Required_: No

_Type_: String

_Pattern_: `[\w!.=*/-]*`

_Minimum_: `0`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResultFormat`

Intended file format of the result.

_Required_: Yes

_Type_: String

_Allowed values_: `CSV | PARQUET`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SingleFileOutput`

Indicates whether files should be output as a single file ( `TRUE`) or output
as multiple files ( `FALSE`). This parameter is only supported for analyses with
the Spark analytics engine.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ProtectedJobS3OutputConfigurationInput

Tag

All content copied from https://docs.aws.amazon.com/.
