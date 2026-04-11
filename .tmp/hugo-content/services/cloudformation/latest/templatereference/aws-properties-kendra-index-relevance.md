This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::Index Relevance

Provides information for tuning the relevance of a field in a search. When a query
includes terms that match the field, the results are given a boost in the response based
on these tuning parameters.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Duration" : String,
  "Freshness" : Boolean,
  "Importance" : Integer,
  "RankOrder" : String,
  "ValueImportanceItems" : [ ValueImportanceItem, ... ]
}

```

### YAML

```yaml

  Duration: String
  Freshness: Boolean
  Importance: Integer
  RankOrder: String
  ValueImportanceItems:
    - ValueImportanceItem

```

## Properties

`Duration`

Specifies the time period that the boost applies to. For example, to make the boost
apply to documents with the field value within the last month, you would use "2628000s".
Once the field value is beyond the specified range, the effect of the boost drops off.
The higher the importance, the faster the effect drops off. If you don't specify a
value, the default is 3 months. The value of the field is a numeric string followed by
the character "s", for example "86400s" for one day, or "604800s" for one week.

Only applies to `DATE` fields.

_Required_: No

_Type_: String

_Pattern_: `[0-9]+[s]`

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Freshness`

Indicates that this field determines how "fresh" a document is. For example, if
document 1 was created on November 5, and document 2 was created on October 31, document
1 is "fresher" than document 2. Only applies to `DATE` fields.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Importance`

The relative importance of the field in the search. Larger numbers provide more of a
boost than smaller numbers.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RankOrder`

Determines how values should be interpreted.

When the `RankOrder` field is `ASCENDING`, higher numbers are
better. For example, a document with a rating score of 10 is higher ranking than a
document with a rating score of 1.

When the `RankOrder` field is `DESCENDING`, lower numbers are
better. For example, in a task tracking application, a priority 1 task is more important
than a priority 5 task.

Only applies to `LONG` fields.

_Required_: No

_Type_: String

_Allowed values_: `ASCENDING | DESCENDING`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValueImportanceItems`

An array of key-value pairs for different boosts when they appear in the search result
list. For example, if you want to boost query terms that match the "department" field in
the result, query terms that match this field are boosted in the result. You can add
entries from the department field to boost documents with those values higher.

For example, you can add entries to the map with names of departments. If you add
"HR", 5 and "Legal",3 those departments are given special attention when they appear in
the metadata of a document.

_Required_: No

_Type_: Array of [ValueImportanceItem](aws-properties-kendra-index-valueimportanceitem.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

JwtTokenTypeConfiguration

Search

All content copied from https://docs.aws.amazon.com/.
