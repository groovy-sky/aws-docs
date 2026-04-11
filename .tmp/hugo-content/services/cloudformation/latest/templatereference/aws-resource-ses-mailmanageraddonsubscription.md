This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerAddonSubscription

Creates a subscription for an Add On representing the acceptance of its terms of use
and additional pricing. The subscription can then be used to create an instance for use
in rule sets or traffic policies.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SES::MailManagerAddonSubscription",
  "Properties" : {
      "AddonName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SES::MailManagerAddonSubscription
Properties:
  AddonName: String
  Tags:
    - Tag

```

## Properties

`AddonName`

The name of the Add On to subscribe to. You can only have one subscription for each
Add On name.

Valid Values: `TRENDMICRO_VSAPI | SPAMHAUS_DBL | ABUSIX_MAIL_INTELLIGENCE |
                VADE_ADVANCED_EMAIL_SECURITY`

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags used to organize, track, or control access for the resource. For example, { "tags": {"key1":"value1", "key2":"value2"} }.

_Required_: No

_Type_: Array of [Tag](aws-properties-ses-mailmanageraddonsubscription-tag.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`AddonSubscriptionArn`

The Amazon Resource Name (ARN) of the Add On subscription.

`AddonSubscriptionId`

The unique ID of the Add On subscription.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
