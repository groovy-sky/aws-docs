This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KendraRanking::ExecutionPlan

Creates a rescore execution plan. A rescore execution
plan is an Amazon Kendra Intelligent Ranking resource
used for provisioning the `Rescore` API. You set
the number of capacity units that you require for
Amazon Kendra Intelligent Ranking to rescore or re-rank
a search service's results.

For an example of using the
`CreateRescoreExecutionPlan` API, including using
the Python and Java SDKs, see [Semantically \
ranking a search service's results](../../../kendra/latest/dg/search-service-rerank.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::KendraRanking::ExecutionPlan",
  "Properties" : {
      "CapacityUnits" : CapacityUnitsConfiguration,
      "Description" : String,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::KendraRanking::ExecutionPlan
Properties:
  CapacityUnits:
    CapacityUnitsConfiguration
  Description: String
  Name: String
  Tags:
    - Tag

```

## Properties

`CapacityUnits`

You can set additional capacity units to meet the
needs of your rescore execution plan. You are given a single
capacity unit by default. If you want to use the default
capacity, you don't set additional capacity units. For more
information on the default capacity and additional capacity
units, see [Adjusting \
capacity](../../../kendra/latest/dg/adjusting-capacity.md).

_Required_: No

_Type_: [CapacityUnitsConfiguration](aws-properties-kendraranking-executionplan-capacityunitsconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description for the rescore execution plan.

_Required_: No

_Type_: String

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A name for the rescore execution plan.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of key-value pairs that identify or categorize your
rescore execution plan. You can also use tags to help control
access to the rescore execution plan. Tag keys and values can
consist of Unicode letters, digits, white space. They can also
consist of underscore, period, colon, equal, plus, and asperand.

_Required_: No

_Type_: Array of [Tag](aws-properties-kendraranking-executionplan-tag.md)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the rescore execution plan ID. For example:

`{"Ref": "rescore-execution-plan-id"}`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the rescore
execution plan.

`Id`

The identifier of the rescore execution plan.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Kendra Intelligent Ranking

CapacityUnitsConfiguration

All content copied from https://docs.aws.amazon.com/.
