This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket ReplicationRuleFilter

A filter that identifies the subset of objects to which the replication rule applies. A
`Filter` must specify exactly one `Prefix`, `TagFilter`, or
an `And` child element.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "And" : ReplicationRuleAndOperator,
  "Prefix" : String,
  "TagFilter" : TagFilter
}

```

### YAML

```yaml

  And:
    ReplicationRuleAndOperator
  Prefix: String
  TagFilter:
    TagFilter

```

## Properties

`And`

A container for specifying rule filters. The filters determine the subset of objects to
which the rule applies. This element is required only if you specify more than one filter. For
example:

- If you specify both a `Prefix` and a `TagFilter`, wrap these
filters in an `And` tag.

- If you specify a filter based on multiple tags, wrap the `TagFilter`
elements in an `And` tag.

_Required_: No

_Type_: [ReplicationRuleAndOperator](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-replicationruleandoperator.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Prefix`

An object key name prefix that identifies the subset of objects to which the rule applies.

###### Important

Replacement must be made for object keys containing special characters (such as carriage returns) when using
XML requests. For more information, see [XML related object key constraints](../../../s3/latest/userguide/object-keys.md#object-key-xml-related-constraints).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TagFilter`

A container for specifying a tag key and value.

The rule applies only to objects that have the tag in their tag set.

_Required_: No

_Type_: [TagFilter](aws-properties-s3-bucket-tagfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ReplicationRuleAndOperator

ReplicationTime
