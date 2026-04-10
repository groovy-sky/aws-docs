This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Backup::RestoreTestingPlan RestoreTestingRecoveryPointSelection

`RecoveryPointSelection` has five parameters (three required and two
optional). The values you specify determine which recovery point is included in the restore
test. You must indicate with `Algorithm` if you want the latest recovery point
within your `SelectionWindowDays` or if you want a random recovery point, and
you must indicate through `IncludeVaults` from which vaults the recovery points
can be chosen.

`Algorithm` ( _required_) Valid values:
" `LATEST_WITHIN_WINDOW`" or " `RANDOM_WITHIN_WINDOW`".

`Recovery point types` ( _required_) Valid values:
" `SNAPSHOT`" and/or " `CONTINUOUS`". Include `SNAPSHOT`
to restore only snapshot recovery points; include `CONTINUOUS` to restore
continuous recovery points (point in time restore / PITR); use both to restore either a
snapshot or a continuous recovery point. The recovery point will be determined by the value
for `Algorithm`.

`IncludeVaults` ( _required_). You must include one or more
backup vaults. Use the wildcard \["\*"\] or specific ARNs.

`SelectionWindowDays` ( _optional_) Value must be an
integer (in days) from 1 to 365. If not included, the value defaults to
`30`.

`ExcludeVaults` ( _optional_). You can choose to input one
or more specific backup vault ARNs to exclude those vaults' contents from restore
eligibility. Or, you can include a list of selectors. If this parameter and its value are
not included, it defaults to empty list.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Algorithm" : String,
  "ExcludeVaults" : [ String, ... ],
  "IncludeVaults" : [ String, ... ],
  "RecoveryPointTypes" : [ String, ... ],
  "SelectionWindowDays" : Integer
}

```

### YAML

```yaml

  Algorithm: String
  ExcludeVaults:
    - String
  IncludeVaults:
    - String
  RecoveryPointTypes:
    - String
  SelectionWindowDays: Integer

```

## Properties

`Algorithm`

Acceptable values include "LATEST\_WITHIN\_WINDOW" or
"RANDOM\_WITHIN\_WINDOW"

_Required_: Yes

_Type_: String

_Allowed values_: `LATEST_WITHIN_WINDOW | RANDOM_WITHIN_WINDOW`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExcludeVaults`

Accepted values include specific ARNs or list of selectors.
Defaults to empty list if not listed.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeVaults`

Accepted values include wildcard \["\*"\] or by specific ARNs or
ARN wilcard replacement
\["arn:aws:backup:us-west-2:123456789012:backup-vault:asdf", ...\]
\["arn:aws:backup:\*:\*:backup-vault:asdf-\*", ...\]

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecoveryPointTypes`

These are the types of recovery points.

Include `SNAPSHOT`
to restore only snapshot recovery points; include `CONTINUOUS` to restore
continuous recovery points (point in time restore / PITR); use both to restore either a
snapshot or a continuous recovery point. The recovery point will be determined by the value
for `Algorithm`.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelectionWindowDays`

Accepted values are integers from 1 to 365.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Backup::RestoreTestingPlan

Tag

All content copied from https://docs.aws.amazon.com/.
