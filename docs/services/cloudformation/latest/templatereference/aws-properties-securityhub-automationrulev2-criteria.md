This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::AutomationRuleV2 Criteria

The filtering type and configuration of the automation rule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OcsfFindingCriteria" : OcsfFindingFilters
}

```

### YAML

```yaml

  OcsfFindingCriteria:
    OcsfFindingFilters

```

## Properties

`OcsfFindingCriteria`

The filtering conditions that align with OCSF standards.

_Required_: No

_Type_: [OcsfFindingFilters](aws-properties-securityhub-automationrulev2-ocsffindingfilters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CompositeFilter

DateFilter

All content copied from https://docs.aws.amazon.com/.
