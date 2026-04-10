This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudTrail::EventDataStore ContextKeySelector

An object that contains information types to be included in CloudTrail enriched events.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Equals" : [ String, ... ],
  "Type" : String
}

```

### YAML

```yaml

  Equals:
    - String
  Type: String

```

## Properties

`Equals`

A list of keys defined by Type to be included in CloudTrail enriched events.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `128 | 50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Specifies the type of the event record field in ContextKeySelector. Valid values include RequestContext, TagContext.

_Required_: Yes

_Type_: String

_Allowed values_: `RequestContext | TagContext`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AdvancedFieldSelector

InsightSelector

All content copied from https://docs.aws.amazon.com/.
