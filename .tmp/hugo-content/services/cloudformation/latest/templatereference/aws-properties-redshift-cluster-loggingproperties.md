This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Redshift::Cluster LoggingProperties

Specifies logging information, such as queries and connection attempts, for the
specified Amazon Redshift cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketName" : String,
  "LogDestinationType" : String,
  "LogExports" : [ String, ... ],
  "S3KeyPrefix" : String
}

```

### YAML

```yaml

  BucketName: String
  LogDestinationType: String
  LogExports:
    - String
  S3KeyPrefix: String

```

## Properties

`BucketName`

The name of an existing S3 bucket where the log files are to be stored.

Constraints:

- Must be in the same region as the cluster

- The cluster must have read bucket and put object permissions

_Required_: No

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogDestinationType`

The log destination type. An enum with possible values of `s3` and `cloudwatch`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogExports`

The collection of exported log types. Possible values are `connectionlog`, `useractivitylog`, and `userlog`.

_Required_: No

_Type_: Array of String

_Maximum_: `3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3KeyPrefix`

The prefix applied to the log file names.

Valid characters are any letter from any language, any whitespace character, any numeric character, and the following characters:
underscore ( `_`), period ( `.`), colon ( `:`), slash ( `/`), equal ( `=`), plus ( `+`), backslash ( `\`),
hyphen ( `-`), at symbol ( `@`).

_Required_: No

_Type_: String

_Pattern_: `[\p{L}\p{Z}\p{N}_.:/=+\-@]*`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Endpoint

Tag

All content copied from https://docs.aws.amazon.com/.
