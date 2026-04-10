This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::Campaign CampaignCustomMessage

Specifies the contents of a message that's sent through a custom channel to
recipients of a campaign.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Data" : String
}

```

### YAML

```yaml

  Data: String

```

## Properties

`Data`

The raw, JSON-formatted string to use as the payload for the message. The
maximum size is 5 KB.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Pinpoint::Campaign

CampaignEmailMessage

All content copied from https://docs.aws.amazon.com/.
