This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::Integration S3SourceProperties

The properties that are applied when Amazon S3 is being used as the flow
source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketName" : String,
  "BucketPrefix" : String
}

```

### YAML

```yaml

  BucketName: String
  BucketPrefix: String

```

## Properties

`BucketName`

The Amazon S3 bucket name where the source files are stored.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BucketPrefix`

The object key for the Amazon S3 bucket in which the source files are stored.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ObjectTypeMapping

SalesforceSourceProperties

All content copied from https://docs.aws.amazon.com/.
