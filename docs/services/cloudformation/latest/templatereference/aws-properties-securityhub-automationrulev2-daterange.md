This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::AutomationRuleV2 DateRange

A date range for the date filter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Unit" : String,
  "Value" : Number
}

```

### YAML

```yaml

  Unit: String
  Value: Number

```

## Properties

`Unit`

A date range unit for the date filter.

_Required_: Yes

_Type_: String

_Allowed values_: `DAYS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

A date range value for the date filter.

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DateFilter

ExternalIntegrationConfiguration

All content copied from https://docs.aws.amazon.com/.
