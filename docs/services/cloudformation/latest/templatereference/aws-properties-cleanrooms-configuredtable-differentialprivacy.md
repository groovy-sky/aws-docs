This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::ConfiguredTable DifferentialPrivacy

The analysis method allowed for the configured tables.

`DIRECT_QUERY` allows SQL queries to be run directly on this table.

`DIRECT_JOB` allows PySpark jobs to be run directly on this table.

`MULTIPLE` allows both SQL queries and PySpark jobs to be run directly on this table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Columns" : [ DifferentialPrivacyColumn, ... ]
}

```

### YAML

```yaml

  Columns:
    - DifferentialPrivacyColumn

```

## Properties

`Columns`

The name of the column, such as user\_id, that contains the unique identifier of your users, whose privacy you want to protect. If you want to turn on differential privacy for two or more tables in a collaboration, you must configure the same column as the user identifier column in both analysis rules.

_Required_: Yes

_Type_: Array of [DifferentialPrivacyColumn](aws-properties-cleanrooms-configuredtable-differentialprivacycolumn.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConfiguredTableAnalysisRulePolicyV1

DifferentialPrivacyColumn

All content copied from https://docs.aws.amazon.com/.
