This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ResourceGroups::Group TagFilter

Specifies a single tag key and optional values that you can use to specify membership
in a tag-based group. An AWS resource that doesn't have a matching tag
key and value is rejected as a member of the group.

A `TagFilter` object includes two properties: `Key` (a string)
and `Values` (a list of strings). Only resources in the account that are
tagged with a matching key-value pair are members of the group. The `Values`
property of `TagFilter` is optional, but specifying it narrows the query
results.

As an example, suppose the `TagFilters` string is `[{"Key": "Stage",
                "Values": ["Test", "Beta"]}, {"Key": "Storage"}]`. In this case, only
resources with all of the following tags are members of the group:

- `Stage` tag key with a value of either `Test` or
`Beta`

- `Storage` tag key with any value

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Values" : [ String, ... ]
}

```

### YAML

```yaml

  Key: String
  Values:
    - String

```

## Properties

`Key`

A string that defines a tag key. Only resources in the account that are tagged with a
specified tag key are members of the tag-based resource group.

This field is required when the `ResourceQuery` structure's
`Type` property is `TAG_FILTERS_1_0`. You must specify at
least one tag key.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

A list of tag values that can be included in the tag-based resource group. This is
optional. If you don't specify a value or values for a key, then an AWS
resource with any value for that key is a member.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::ResourceGroups::TagSyncTask

All content copied from https://docs.aws.amazon.com/.
