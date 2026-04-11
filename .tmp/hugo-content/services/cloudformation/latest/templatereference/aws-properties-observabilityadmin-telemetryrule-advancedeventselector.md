This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ObservabilityAdmin::TelemetryRule AdvancedEventSelector

Advanced event selectors let you create fine-grained selectors for management, data, and
network activity events.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FieldSelectors" : [ AdvancedFieldSelector, ... ],
  "Name" : String
}

```

### YAML

```yaml

  FieldSelectors:
    - AdvancedFieldSelector
  Name: String

```

## Properties

`FieldSelectors`

Contains all selector statements in an advanced event selector.

_Required_: Yes

_Type_: Array of [AdvancedFieldSelector](aws-properties-observabilityadmin-telemetryrule-advancedfieldselector.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

An optional, descriptive name for an advanced event selector, such as "Log data events for
only two S3 buckets".

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ActionCondition

AdvancedFieldSelector

All content copied from https://docs.aws.amazon.com/.
