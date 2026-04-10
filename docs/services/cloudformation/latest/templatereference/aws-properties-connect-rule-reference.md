This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::Rule Reference

Information about the reference when the `referenceType` is
`URL`. Otherwise, null. (Supports variable injection in the
`Value` field.)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : String,
  "Value" : String
}

```

### YAML

```yaml

  Type: String
  Value: String

```

## Properties

`Type`

The type of the reference. `DATE` must be of type Epoch timestamp.

_Allowed values_: `URL` \| `ATTACHMENT` \|
`NUMBER` \| `STRING` \| `DATE` \|
`EMAIL`

_Required_: Yes

_Type_: String

_Allowed values_: `URL | ATTACHMENT | NUMBER | STRING | DATE | EMAIL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

A valid value for the reference. For example, for a URL reference, a formatted URL
that is displayed to an agent in the Contact Control Panel (CCP).

_Required_: Yes

_Type_: String

_Pattern_: `^(/|https:)`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NotificationRecipientType

RuleTriggerEventSource

All content copied from https://docs.aws.amazon.com/.
