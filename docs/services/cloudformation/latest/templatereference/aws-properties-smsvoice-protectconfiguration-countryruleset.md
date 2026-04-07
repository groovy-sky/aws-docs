This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SMSVOICE::ProtectConfiguration CountryRuleSet

The set of `CountryRules` you specify to control which countries AWS End User Messaging SMS can send your messages to.

###### Note

If you don't specify all available ISO country codes in the `CountryRuleSet` for each number capability, the CloudFormation drift detection feature
will detect drift. This is because AWS End User Messaging SMS always returns all country codes.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MMS" : [ CountryRule, ... ],
  "SMS" : [ CountryRule, ... ],
  "VOICE" : [ CountryRule, ... ]
}

```

### YAML

```yaml

  MMS:
    - CountryRule
  SMS:
    - CountryRule
  VOICE:
    - CountryRule

```

## Properties

`MMS`

The set of `CountryRule` s to control which destination countries AWS End User Messaging SMS can send your MMS messages to.

_Required_: No

_Type_: Array of [CountryRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-smsvoice-protectconfiguration-countryrule.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SMS`

The set of `CountryRule` s to control which destination countries AWS End User Messaging SMS can send your SMS messages to.

_Required_: No

_Type_: Array of [CountryRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-smsvoice-protectconfiguration-countryrule.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VOICE`

The set of `CountryRule` s to control which destination countries AWS End User Messaging SMS can send your VOICE messages to.

_Required_: No

_Type_: Array of [CountryRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-smsvoice-protectconfiguration-countryrule.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CountryRule

Tag
