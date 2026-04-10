This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::EventSourceMapping Filter

A structure within a `FilterCriteria` object that defines an event filtering pattern.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Pattern" : String
}

```

### YAML

```yaml

  Pattern: String

```

## Properties

`Pattern`

A filter pattern. For more information on the syntax of a filter pattern, see
[Filter rule syntax](../../../lambda/latest/dg/invocation-eventfiltering.md#filtering-syntax).

_Required_: No

_Type_: String

_Pattern_: `.*`

_Minimum_: `0`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Endpoints

FilterCriteria

All content copied from https://docs.aws.amazon.com/.
