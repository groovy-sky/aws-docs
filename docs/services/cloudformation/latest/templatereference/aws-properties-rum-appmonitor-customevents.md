This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RUM::AppMonitor CustomEvents

This structure specifies whether this app monitor allows the web client to define and send custom events.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Status" : String
}

```

### YAML

```yaml

  Status: String

```

## Properties

`Status`

Set this to `ENABLED` to allow the web client to send custom events for this app monitor.

Valid values are `ENABLED` and `DISABLED`.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AppMonitorConfiguration

DeobfuscationConfiguration

All content copied from https://docs.aws.amazon.com/.
