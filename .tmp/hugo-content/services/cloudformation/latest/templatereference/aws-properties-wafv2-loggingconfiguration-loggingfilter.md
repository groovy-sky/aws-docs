This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::LoggingConfiguration LoggingFilter

Filtering that specifies which web requests are kept in the logs and which are dropped,
defined for a web ACL's `LoggingConfiguration`.

You can filter on the rule action and on the web request labels that were applied by
matching rules during web ACL evaluation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultBehavior" : String,
  "Filters" : [ Filter, ... ]
}

```

### YAML

```yaml

  DefaultBehavior: String
  Filters:
    - Filter

```

## Properties

`DefaultBehavior`

Default handling for logs that don't match any of the specified filtering conditions.

_Required_: Yes

_Type_: String

_Allowed values_: `KEEP | DROP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Filters`

The filters that you want to apply to the logs.

_Required_: Yes

_Type_: Array of [Filter](aws-properties-wafv2-loggingconfiguration-filter.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LabelNameCondition

SingleHeader

All content copied from https://docs.aws.amazon.com/.
