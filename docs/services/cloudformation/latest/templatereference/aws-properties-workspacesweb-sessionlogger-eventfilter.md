This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WorkSpacesWeb::SessionLogger EventFilter

The filter that specifies the events to monitor.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "All" : Json,
  "Include" : [ String, ... ]
}

```

### YAML

```yaml

  All: Json
  Include:
    - String

```

## Properties

`All`

The filter that monitors all of the available events, including any new events emitted
in the future. The `All` and `Include` properties are not
required, but one of them should be present. `{}` is a valid input.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Include`

The filter that monitors only the listed set of events. New events are not
auto-monitored. The `All` and `Include` properties are not
required, but one of them should be present.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::WorkSpacesWeb::SessionLogger

LogConfiguration

All content copied from https://docs.aws.amazon.com/.
