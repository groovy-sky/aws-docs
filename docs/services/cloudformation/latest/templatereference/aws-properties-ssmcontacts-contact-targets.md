This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSMContacts::Contact Targets

The contact or contact channel that's being engaged.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ChannelTargetInfo" : ChannelTargetInfo,
  "ContactTargetInfo" : ContactTargetInfo
}

```

### YAML

```yaml

  ChannelTargetInfo:
    ChannelTargetInfo
  ContactTargetInfo:
    ContactTargetInfo

```

## Properties

`ChannelTargetInfo`

Information about the contact channel that Incident Manager engages.

_Required_: No

_Type_: [ChannelTargetInfo](aws-properties-ssmcontacts-contact-channeltargetinfo.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContactTargetInfo`

The contact that Incident Manager is engaging during an incident.

_Required_: No

_Type_: [ContactTargetInfo](aws-properties-ssmcontacts-contact-contacttargetinfo.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::SSMContacts::ContactChannel

All content copied from https://docs.aws.amazon.com/.
