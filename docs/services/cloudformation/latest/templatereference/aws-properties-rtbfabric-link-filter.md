This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RTBFabric::Link Filter

Describes the configuration of a filter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Criteria" : [ FilterCriterion, ... ]
}

```

### YAML

```yaml

  Criteria:
    - FilterCriterion

```

## Properties

`Criteria`

Describes the criteria for a filter.

_Required_: Yes

_Type_: Array of [FilterCriterion](aws-properties-rtbfabric-link-filtercriterion.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ApplicationLogs

FilterCriterion

All content copied from https://docs.aws.amazon.com/.
