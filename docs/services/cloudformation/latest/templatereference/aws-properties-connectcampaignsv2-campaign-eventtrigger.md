This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ConnectCampaignsV2::Campaign EventTrigger

The event trigger of the campaign.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomerProfilesDomainArn" : String
}

```

### YAML

```yaml

  CustomerProfilesDomainArn: String

```

## Properties

`CustomerProfilesDomainArn`

The Amazon Resource Name (ARN) of the Customer Profiles domain.

_Required_: No

_Type_: String

_Pattern_: `^arn:.*$`

_Minimum_: `20`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EntryLimitsConfig

LocalTimeZoneConfig

All content copied from https://docs.aws.amazon.com/.
