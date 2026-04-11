This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AccessAnalyzer::Analyzer Filter

The criteria that defines the archive rule.

To learn about filter keys that you can use to create an archive rule, see [AWS Identity and Access Management Access Analyzer filter keys](../../../iam/latest/userguide/access-analyzer-reference-filter-keys.md) in the
_AWS Identity and Access Management User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Contains" : [ String, ... ],
  "Eq" : [ String, ... ],
  "Exists" : Boolean,
  "Neq" : [ String, ... ],
  "Property" : String
}

```

### YAML

```yaml

  Contains:
    - String
  Eq:
    - String
  Exists: Boolean
  Neq:
    - String
  Property: String

```

## Properties

`Contains`

A "contains" condition to match for the rule.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Eq`

An "equals" condition to match for the rule.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Exists`

An "exists" condition to match for the rule.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Neq`

A "not equal" condition to match for the rule.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Property`

The property used to define the criteria in the filter for the rule.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ArchiveRule

InternalAccessAnalysisRule

All content copied from https://docs.aws.amazon.com/.
