This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSMContacts::Plan ContactTargetInfo

The contact that Incident Manager is engaging during an incident.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContactId" : String,
  "IsEssential" : Boolean
}

```

### YAML

```yaml

  ContactId: String
  IsEssential: Boolean

```

## Properties

`ContactId`

The Amazon Resource Name (ARN) of the contact.

_Required_: Yes

_Type_: String

_Pattern_: `arn:(aws|aws-cn|aws-us-gov):ssm-contacts:[-\w+=\/,.@]*:[0-9]+:([\w+=\/,.@:-])*`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsEssential`

A Boolean value determining if the contact's acknowledgement stops the progress of
stages in the plan.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ChannelTargetInfo

Stage

All content copied from https://docs.aws.amazon.com/.
