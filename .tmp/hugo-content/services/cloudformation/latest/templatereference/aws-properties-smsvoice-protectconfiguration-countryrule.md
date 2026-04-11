This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SMSVOICE::ProtectConfiguration CountryRule

Specifies the type of protection to use for a country.

For example, to set Canada as allowed, the `CountryRule` would be formatted as follows:

```

{
    "CountryCode": "CA",
    "ProtectStatus": "ALLOW"
}
```

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CountryCode" : String,
  "ProtectStatus" : String
}

```

### YAML

```yaml

  CountryCode: String
  ProtectStatus: String

```

## Properties

`CountryCode`

The two-character code, in ISO 3166-1 alpha-2 format, for the country or region.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Z]{2}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProtectStatus`

The types of protection that can be used.

_Required_: Yes

_Type_: String

_Allowed values_: `ALLOW | BLOCK | MONITOR | FILTER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SMSVOICE::ProtectConfiguration

CountryRuleSet

All content copied from https://docs.aws.amazon.com/.
