This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GuardDuty::Filter Condition

Specifies the condition to apply to a single field when filtering through GuardDuty findings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Eq" : [ String, ... ],
  "Equals" : [ String, ... ],
  "GreaterThan" : Integer,
  "GreaterThanOrEqual" : Integer,
  "Gt" : Integer,
  "Gte" : Integer,
  "LessThan" : Integer,
  "LessThanOrEqual" : Integer,
  "Lt" : Integer,
  "Lte" : Integer,
  "Neq" : [ String, ... ],
  "NotEquals" : [ String, ... ]
}

```

### YAML

```yaml

  Eq:
    - String
  Equals:
    - String
  GreaterThan: Integer
  GreaterThanOrEqual: Integer
  Gt: Integer
  Gte: Integer
  LessThan: Integer
  LessThanOrEqual: Integer
  Lt: Integer
  Lte: Integer
  Neq:
    - String
  NotEquals:
    - String

```

## Properties

`Eq`

Represents the equal condition to apply to a single field when querying for
findings.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Equals`

Represents an _equal_ condition to be applied to
a single field when querying for findings.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GreaterThan`

Represents a _greater than_ condition to be applied to a single field
when querying for findings.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GreaterThanOrEqual`

Represents a _greater than or equal_ condition to be applied to a
single field when querying for findings.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Gt`

Represents a _greater than_ condition to be applied to a single field
when querying for findings.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Gte`

Represents the greater than or equal condition to apply to a single field when querying
for findings.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LessThan`

Represents a _less than_ condition to be applied to a single field when
querying for findings.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LessThanOrEqual`

Represents a _less than or equal_ condition to be applied to a single
field when querying for findings.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Lt`

Represents the less than condition to apply to a single field when querying for
findings.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Lte`

Represents the less than or equal condition to apply to a single field when querying for
findings.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Neq`

Represents the not equal condition to apply to a single field when querying for
findings.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotEquals`

Represents a _not equal_ condition to be applied
to a single field when querying for findings.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::GuardDuty::Filter

FindingCriteria

All content copied from https://docs.aws.amazon.com/.
