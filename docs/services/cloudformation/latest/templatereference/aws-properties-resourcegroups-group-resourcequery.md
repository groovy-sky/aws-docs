This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ResourceGroups::Group ResourceQuery

The query used to dynamically define the members of a group. For more information
about how to construct a query, see [Build queries and groups in\
AWS Resource Groups](../../../arg/latest/userguide/gettingstarted-query.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Query" : Query,
  "Type" : String
}

```

### YAML

```yaml

  Query:
    Query
  Type: String

```

## Properties

`Query`

The query that defines the membership of the group. This is a structure with
properties that depend on the `Type`.

The `Query` structure must be included in the following scenarios:

- When the `Type` is `TAG_FILTERS_1_0`, you must specify a
`Query` structure that contains a `TagFilters` list of
tags. Resources with tags that match those in the `TagFilter` list
become members of the resource group.

- When the `Type` is `CLOUDFORMATION_STACK_1_0` then this
field is required only when you must specify a CloudFormation stack other than
the one you are defining. To do this, the `Query` structure must
contain the `StackIdentifier` property. If you don't specify either a
`Query` structure or a `StackIdentifier` within that
`Query`, then it defaults to the CloudFormation stack that you're
currently constructing.

_Required_: No

_Type_: [Query](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-resourcegroups-group-query.html)

_Pattern_: `[\s\S]*`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Specifies the type of resource query that determines this group's membership. There
are two valid query types:

- `TAG_FILTERS_1_0` indicates that the group is a tag-based group.
To complete the group membership, you must include the `TagFilters`
property to specify the tag filters to use in the query.

- `CLOUDFORMATION_STACK_1_0`, the default, indicates that the
group is a CloudFormation stack-based group. Group membership is based on the
CloudFormation stack. You must specify the `StackIdentifier` property
in the query to define which stack to associate the group with, or leave it
empty to default to the stack where the group is defined.

_Required_: No

_Type_: String

_Allowed values_: `TAG_FILTERS_1_0 | CLOUDFORMATION_STACK_1_0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Query

Tag
