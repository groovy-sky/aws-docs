This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ResourceGroups::Group Query

Specifies details within a `ResourceQuery` structure that determines the
membership of the resource group. The contents required in the `Query`
structure are determined by the `Type` property of the containing
`ResourceQuery` structure.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ResourceTypeFilters" : [ String, ... ],
  "StackIdentifier" : String,
  "TagFilters" : [ TagFilter, ... ]
}

```

### YAML

```yaml

  ResourceTypeFilters:
    - String
  StackIdentifier: String
  TagFilters:
    - TagFilter

```

## Properties

`ResourceTypeFilters`

Specifies limits to the types of resources that can be included in the resource group.
For example, if `ResourceTypeFilters` is `["AWS::EC2::Instance",
                "AWS::DynamoDB::Table"]`, only EC2 instances or DynamoDB tables can be members
of this resource group. The default value is `["AWS::AllSupported"]`.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StackIdentifier`

Specifies the ARN of a CloudFormation stack. All supported resources of the
CloudFormation stack are members of the resource group. If you don't specify an ARN,
this parameter defaults to the current stack that you are defining, which means that all
the resources of the current stack are grouped.

You can specify a value for `StackIdentifier` only when the
`ResourceQuery.Type` property is
`CLOUDFORMATION_STACK_1_0.`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TagFilters`

A list of key-value pair objects that limit which resources can be members of the
resource group. This property is required when the `ResourceQuery.Type`
property is `TAG_FILTERS_1_0`.

A resource must have a tag that matches every filter that is provided in the
`TagFilters` list.

_Required_: Conditional

_Type_: Array of [TagFilter](aws-properties-resourcegroups-group-tagfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConfigurationParameter

ResourceQuery

All content copied from https://docs.aws.amazon.com/.
