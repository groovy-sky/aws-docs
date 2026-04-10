This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe FilterCriteria

The collection of event patterns used to filter events.

To remove a filter, specify a `FilterCriteria` object with an empty array of `Filter` objects.

For more information, see [Events and Event\
Patterns](../../../eventbridge/latest/userguide/eventbridge-and-event-patterns.md) in the _Amazon EventBridge User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Filters" : [ Filter, ... ]
}

```

### YAML

```yaml

  Filters:
    - Filter

```

## Properties

`Filters`

The event patterns.

_Required_: No

_Type_: Array of [Filter](aws-properties-pipes-pipe-filter.md)

_Minimum_: `0`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Filter

FirehoseLogDestination

All content copied from https://docs.aws.amazon.com/.
