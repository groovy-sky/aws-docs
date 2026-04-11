This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityLake::Subscriber CustomLogSource

Third-party custom log source that meets the requirements to be added to Amazon Security Lake. For more details, see [Custom log source](../../../security-lake/latest/userguide/custom-sources.md#iam-roles-custom-sources) in the _Amazon Security Lake User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SourceName" : String,
  "SourceVersion" : String
}

```

### YAML

```yaml

  SourceName: String
  SourceVersion: String

```

## Properties

`SourceName`

The name of the custom log source.

_Required_: No

_Type_: String

_Pattern_: `^[\\\w\-_:/.]*$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceVersion`

The source version of the custom log source.

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z0-9\-\.\_]*$`

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AwsLogSource

Source

All content copied from https://docs.aws.amazon.com/.
