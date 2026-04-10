This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DLM::LifecyclePolicy EventSource

**\[Event-based policies only\]** Specifies an event that activates an event-based policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Parameters" : EventParameters,
  "Type" : String
}

```

### YAML

```yaml

  Parameters:
    EventParameters
  Type: String

```

## Properties

`Parameters`

Information about the event.

_Required_: No

_Type_: [EventParameters](aws-properties-dlm-lifecyclepolicy-eventparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The source of the event. Currently only managed Amazon EventBridge (formerly known as Amazon CloudWatch) events are supported.

_Required_: Yes

_Type_: String

_Allowed values_: `MANAGED_CWE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EventParameters

Exclusions

All content copied from https://docs.aws.amazon.com/.
