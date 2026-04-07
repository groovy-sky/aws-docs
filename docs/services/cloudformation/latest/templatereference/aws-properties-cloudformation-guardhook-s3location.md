This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFormation::GuardHook S3Location

Specifies the S3 location where your Guard rules or input
parameters are located.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Uri" : String,
  "VersionId" : String
}

```

### YAML

```yaml

  Uri: String
  VersionId: String

```

## Properties

`Uri`

Specifies the S3 path to the file that contains your Guard
rules or input parameters (in the form `s3://<bucket name>/<file
                name>`).

For Guard rules, the object stored in S3 must have one of the
following file extensions: `.guard`, `.zip`, or
`.tar.gz`.

For input parameters, the object stored in S3 must have one of the following file
extensions: `.yaml`, `.json`, `.zip`, or
`.tar.gz`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VersionId`

For S3 buckets with versioning enabled, specifies the unique ID of the S3 object
version to download your Guard rules or input parameters
from.

The Guard Hook downloads files from S3 every time the Hook is
invoked. To prevent accidental changes or deletions, we recommend using a version when
configuring your Guard Hook.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Options

StackFilters
