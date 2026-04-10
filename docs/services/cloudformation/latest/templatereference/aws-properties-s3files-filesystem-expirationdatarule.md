This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Files::FileSystem ExpirationDataRule

Specifies a rule that controls when cached data expires from the file system based on
last access time.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DaysAfterLastAccess" : Integer
}

```

### YAML

```yaml

  DaysAfterLastAccess: Integer

```

## Properties

`DaysAfterLastAccess`

The number of days after last access before cached data expires from the file
system.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `365`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::S3Files::FileSystem

ImportDataRule

All content copied from https://docs.aws.amazon.com/.
