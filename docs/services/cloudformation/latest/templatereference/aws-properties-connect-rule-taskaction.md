This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::Rule TaskAction

Information about the task action. This field is required if
`TriggerEventSource` is one of the following values:
`OnZendeskTicketCreate` \| `OnZendeskTicketStatusUpdate` \|
`OnSalesforceCaseCreate`

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContactFlowArn" : String,
  "Description" : String,
  "Name" : String,
  "References" : Reference
}

```

### YAML

```yaml

  ContactFlowArn: String
  Description: String
  Name: String
  References:
    Reference

```

## Properties

`ContactFlowArn`

The Amazon Resource Name (ARN) of the flow.

_Required_: Yes

_Type_: String

_Pattern_: `^$|arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/contact-flow/[-a-zA-Z0-9]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description. Supports variable injection. For more information, see [JSONPath\
reference](../../../connect/latest/adminguide/contact-lens-variable-injection.md) in the _Amazon Connect Administrators_
_Guide_.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name. Supports variable injection. For more information, see [JSONPath\
reference](../../../connect/latest/adminguide/contact-lens-variable-injection.md) in the _Amazon Connect Administrators_
_Guide_.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`References`

Information about the reference when the `referenceType` is
`URL`. Otherwise, null. `URL` is the only accepted type.
(Supports variable injection in the `Value` field.)

_Required_: No

_Type_: [Reference](aws-properties-connect-rule-reference.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

UpdateCaseAction

All content copied from https://docs.aws.amazon.com/.
