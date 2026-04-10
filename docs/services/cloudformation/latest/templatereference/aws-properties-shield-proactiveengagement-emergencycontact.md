This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Shield::ProactiveEngagement EmergencyContact

Contact information that the SRT can use to contact you if you have proactive engagement enabled, for escalations to the SRT and to initiate proactive customer support.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContactNotes" : String,
  "EmailAddress" : String,
  "PhoneNumber" : String
}

```

### YAML

```yaml

  ContactNotes: String
  EmailAddress: String
  PhoneNumber: String

```

## Properties

`ContactNotes`

Additional notes regarding the contact.

_Required_: No

_Type_: String

_Pattern_: `^[\w\s\.\-,:/()+@]*$`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmailAddress`

The email address for the contact.

_Required_: Yes

_Type_: String

_Pattern_: `^\S+@\S+\.\S+$`

_Minimum_: `1`

_Maximum_: `150`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PhoneNumber`

The phone number for the contact.

_Required_: No

_Type_: String

_Pattern_: `^\+[1-9]\d{1,14}$`

_Minimum_: `1`

_Maximum_: `16`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Shield::ProactiveEngagement

AWS::Shield::Protection

All content copied from https://docs.aws.amazon.com/.
