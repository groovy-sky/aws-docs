This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::StorageLensGroup Filter

This resource sets the criteria for the Storage Lens group data that is displayed. For
multiple filter conditions, the `AND` or `OR` logical operator is
used.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "And" : And,
  "MatchAnyPrefix" : [ String, ... ],
  "MatchAnySuffix" : [ String, ... ],
  "MatchAnyTag" : [ Tag, ... ],
  "MatchObjectAge" : MatchObjectAge,
  "MatchObjectSize" : MatchObjectSize,
  "Or" : Or
}

```

### YAML

```yaml

  And:
    And
  MatchAnyPrefix:
    - String
  MatchAnySuffix:
    - String
  MatchAnyTag:
    - Tag
  MatchObjectAge:
    MatchObjectAge
  MatchObjectSize:
    MatchObjectSize
  Or:
    Or

```

## Properties

`And`

This property contains the `And` logical operator, which allows multiple filter
conditions to be joined for more complex comparisons of Storage Lens group data. Objects must
match all of the listed filter conditions that are joined by the `And` logical
operator. Only one of each filter condition is allowed.

_Required_: No

_Type_: [And](aws-properties-s3-storagelensgroup-and.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MatchAnyPrefix`

This property contains a list of prefixes. At least one prefix must be specified. Up to 10
prefixes are allowed.

_Required_: No

_Type_: Array of String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MatchAnySuffix`

This property contains a list of suffixes. At least one suffix must be specified. Up to 10
suffixes are allowed.

_Required_: No

_Type_: Array of String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MatchAnyTag`

This property contains the list of S3 object tags. At least one object tag must be
specified. Up to 10 object tags are allowed.

_Required_: No

_Type_: Array of [Tag](aws-properties-s3-storagelensgroup-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MatchObjectAge`

This property contains `DaysGreaterThan` and `DaysLessThan` to
define the object age range (minimum and maximum number of days).

_Required_: No

_Type_: [MatchObjectAge](aws-properties-s3-storagelensgroup-matchobjectage.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MatchObjectSize`

This property contains `BytesGreaterThan` and `BytesLessThan` to
define the object size range (minimum and maximum number of Bytes).

_Required_: No

_Type_: [MatchObjectSize](aws-properties-s3-storagelensgroup-matchobjectsize.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Or`

This property contains the `Or` logical operator, which allows multiple filter
conditions to be joined. Objects can match any of the listed filter conditions, which are
joined by the `Or` logical operator. Only one of each filter condition is
allowed.

_Required_: No

_Type_: [Or](aws-properties-s3-storagelensgroup-or.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

And

MatchObjectAge

All content copied from https://docs.aws.amazon.com/.
