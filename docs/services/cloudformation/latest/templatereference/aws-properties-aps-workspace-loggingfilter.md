This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::APS::Workspace LoggingFilter

Filtering criteria that determine which queries are logged.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "QspThreshold" : Integer
}

```

### YAML

```yaml

  QspThreshold: Integer

```

## Properties

`QspThreshold`

The Query Samples Processed (QSP) threshold above which queries will be logged.
Queries processing more samples than this threshold will be captured in logs.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LoggingDestination

QueryLoggingConfiguration

All content copied from https://docs.aws.amazon.com/.
