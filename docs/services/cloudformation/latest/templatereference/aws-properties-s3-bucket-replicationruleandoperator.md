This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket ReplicationRuleAndOperator

A container for specifying rule filters. The filters determine the subset of objects to
which the rule applies. This element is required only if you specify more than one filter.

For example:

- If you specify both a `Prefix` and a `TagFilter`, wrap these
filters in an `And` tag.

- If you specify a filter based on multiple tags, wrap the `TagFilter`
elements in an `And` tag

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Prefix" : String,
  "TagFilters" : [ TagFilter, ... ]
}

```

### YAML

```yaml

  Prefix: String
  TagFilters:
    - TagFilter

```

## Properties

`Prefix`

An object key name prefix that identifies the subset of objects to which the rule applies.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TagFilters`

An array of tags containing key and value pairs.

_Required_: No

_Type_: Array of [TagFilter](aws-properties-s3-bucket-tagfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReplicationRule

ReplicationRuleFilter

All content copied from https://docs.aws.amazon.com/.
