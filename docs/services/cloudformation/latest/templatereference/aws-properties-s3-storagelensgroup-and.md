This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::StorageLensGroup And

This resource is a logical operator that allows multiple filter conditions to be joined
for more complex comparisons of Storage Lens group data. Objects must match all of the listed
filter conditions that are joined by the `And` logical operator. Only one of each
filter condition is allowed.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MatchAnyPrefix" : [ String, ... ],
  "MatchAnySuffix" : [ String, ... ],
  "MatchAnyTag" : [ Tag, ... ],
  "MatchObjectAge" : MatchObjectAge,
  "MatchObjectSize" : MatchObjectSize
}

```

### YAML

```yaml

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

```

## Properties

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

This property contains the list of object tags. At least one object tag must be specified.
Up to 10 object tags are allowed.

_Required_: No

_Type_: Array of [Tag](aws-properties-s3-storagelensgroup-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MatchObjectAge`

This property contains `DaysGreaterThan` and `DaysLessThan`
properties to define the object age range (minimum and maximum number of days).

_Required_: No

_Type_: [MatchObjectAge](aws-properties-s3-storagelensgroup-matchobjectage.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MatchObjectSize`

This property contains `BytesGreaterThan` and `BytesLessThan` to
define the object size range (minimum and maximum number of Bytes).

_Required_: No

_Type_: [MatchObjectSize](aws-properties-s3-storagelensgroup-matchobjectsize.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::S3::StorageLensGroup

Filter

All content copied from https://docs.aws.amazon.com/.
