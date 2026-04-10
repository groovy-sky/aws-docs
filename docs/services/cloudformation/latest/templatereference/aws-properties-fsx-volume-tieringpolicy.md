This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::Volume TieringPolicy

Describes the data tiering policy for an ONTAP volume. When enabled, Amazon FSx for ONTAP's intelligent
tiering automatically transitions a volume's data between the file system's primary storage and capacity
pool storage based on your access patterns.

Valid tiering policies are the following:

- `SNAPSHOT_ONLY` \- (Default value) moves cold snapshots to the capacity pool storage tier.

- `AUTO` \- moves cold user data and snapshots to the capacity pool storage tier based on your access patterns.

- `ALL` \- moves all user data blocks in both the active file system and Snapshot copies to the storage pool tier.

- `NONE` \- keeps a volume's data in the primary storage tier, preventing it from being moved to the capacity pool tier.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CoolingPeriod" : Integer,
  "Name" : String
}

```

### YAML

```yaml

  CoolingPeriod: Integer
  Name: String

```

## Properties

`CoolingPeriod`

Specifies the number of days that user data in a volume must remain inactive before it is considered "cold"
and moved to the capacity pool. Used with the `AUTO` and `SNAPSHOT_ONLY` tiering policies.
Enter a whole number between 2 and 183. Default values are 31 days for `AUTO` and 2 days for
`SNAPSHOT_ONLY`.

_Required_: No

_Type_: Integer

_Minimum_: `2`

_Maximum_: `183`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

Specifies the tiering policy used to transition data. Default value is `SNAPSHOT_ONLY`.

- `SNAPSHOT_ONLY` \- moves cold snapshots to the capacity pool storage tier.

- `AUTO` \- moves cold user data and snapshots to the capacity pool storage tier
based on your access patterns.

- `ALL` \- moves all user data blocks in both the active file system and Snapshot copies to the
storage pool tier.

- `NONE` \- keeps a volume's data in the primary storage tier, preventing it from being moved to
the capacity pool tier.

_Required_: No

_Type_: String

_Allowed values_: `SNAPSHOT_ONLY | AUTO | ALL | NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

UserAndGroupQuotas

All content copied from https://docs.aws.amazon.com/.
