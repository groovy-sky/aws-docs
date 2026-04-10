This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::QuotaShare QuotaShareResourceSharingConfiguration

Specifies whether a quota share reserves, lends, or both lends and borrows idle compute capacity.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BorrowLimit" : Integer,
  "Strategy" : String
}

```

### YAML

```yaml

  BorrowLimit: Integer
  Strategy: String

```

## Properties

`BorrowLimit`

The maximum percentage of additional capacity that the quota share can borrow from other shares. `borrowLimit` can only be applied to quota shares with a strategy of `LEND_AND_BORROW`.
This value is expressed as a percentage of the quota share's configured [CapacityLimits](../../../../reference/batch/latest/apireference/api-quotasharecapacitylimit.md).

The `borrowLimit` is applied uniformly across all capacity units.
For example, if the `borrowLimit` is 200, the quota
share can borrow up to 200% of its configured `maxCapacity` for each capacity unit. The default `borrowLimit` is -1, which indicates unlimited borrowing.

_Required_: No

_Type_: Integer

_Minimum_: `-1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Strategy`

The resource sharing strategy for the quota share. The `RESERVE` strategy allows a quota share to reserve idle capacity for itself.
`LEND` configures the share to lend its idle capacity to another share in need of capacity.
The `LEND_AND_BORROW` strategy configures the share to borrow idle capacity from an underutilized share, as well as lend to another share.

_Required_: Yes

_Type_: String

_Allowed values_: `RESERVE | LEND | LEND_AND_BORROW`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

QuotaSharePreemptionConfiguration

AWS::Batch::SchedulingPolicy

All content copied from https://docs.aws.amazon.com/.
