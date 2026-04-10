This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::EventSourceMapping FilterCriteria

An object that contains the filters for an event source.

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

A list of filters.

_Required_: No

_Type_: Array of [Filter](aws-properties-lambda-eventsourcemapping-filter.md)

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Filter

LoggingConfig

All content copied from https://docs.aws.amazon.com/.
