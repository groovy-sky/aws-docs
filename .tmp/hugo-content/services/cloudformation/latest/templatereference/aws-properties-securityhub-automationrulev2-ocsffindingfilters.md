This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::AutomationRuleV2 OcsfFindingFilters

Specifies the filtering criteria for security findings using OCSF.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CompositeFilters" : [ CompositeFilter, ... ],
  "CompositeOperator" : String
}

```

### YAML

```yaml

  CompositeFilters:
    - CompositeFilter
  CompositeOperator: String

```

## Properties

`CompositeFilters`

Enables the creation of complex filtering conditions by combining filter criteria.

_Required_: No

_Type_: Array of [CompositeFilter](aws-properties-securityhub-automationrulev2-compositefilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CompositeOperator`

The logical operators used to combine the filtering on multiple `CompositeFilters`.

_Required_: No

_Type_: String

_Allowed values_: `AND | OR`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OcsfDateFilter

OcsfMapFilter

All content copied from https://docs.aws.amazon.com/.
