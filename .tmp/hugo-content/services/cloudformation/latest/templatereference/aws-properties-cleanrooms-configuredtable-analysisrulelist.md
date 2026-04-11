This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::ConfiguredTable AnalysisRuleList

A type of analysis rule that enables row-level analysis.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdditionalAnalyses" : String,
  "AllowedJoinOperators" : [ String, ... ],
  "JoinColumns" : [ String, ... ],
  "ListColumns" : [ String, ... ]
}

```

### YAML

```yaml

  AdditionalAnalyses: String
  AllowedJoinOperators:
    - String
  JoinColumns:
    - String
  ListColumns:
    - String

```

## Properties

`AdditionalAnalyses`

An indicator as to whether additional analyses (such as AWS Clean Rooms ML) can be applied to the output of the direct query.

_Required_: No

_Type_: String

_Allowed values_: `ALLOWED | REQUIRED | NOT_ALLOWED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowedJoinOperators`

The logical operators (if any) that are to be used in an INNER JOIN match condition.
Default is `AND`.

_Required_: No

_Type_: Array of String

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JoinColumns`

Columns that can be used to join a configured table with the table of the member who can
query and other members' configured tables.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ListColumns`

Columns that can be listed in the output.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AnalysisRuleCustom

AthenaTableReference

All content copied from https://docs.aws.amazon.com/.
