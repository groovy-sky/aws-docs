This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Deadline::Limit

Creates a limit that manages the distribution of shared resources, such as floating
licenses. A limit can throttle work assignments, help manage workloads, and track current
usage. Before you use a limit, you must associate the limit with one or more queues.

You must add the `amountRequirementName` to a step in a job template to
declare the limit requirement.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Deadline::Limit",
  "Properties" : {
      "AmountRequirementName" : String,
      "Description" : String,
      "DisplayName" : String,
      "FarmId" : String,
      "MaxCount" : Integer
    }
}

```

### YAML

```yaml

Type: AWS::Deadline::Limit
Properties:
  AmountRequirementName: String
  Description: String
  DisplayName: String
  FarmId: String
  MaxCount: Integer

```

## Properties

`AmountRequirementName`

The value that you specify as the `name` in the `amounts` field of
the `hostRequirements` in a step of a job template to declare the limit
requirement.

_Required_: Yes

_Type_: String

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

A description of the limit. A clear description helps you identify the purpose of the
limit.

###### Important

This field can store any content. Escape or encode this content before displaying it
on a webpage or any other system that might interpret the content of this field.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayName`

The name of the limit used in lists to identify the limit.

###### Important

This field can store any content. Escape or encode this content before displaying it on a webpage or any other system that might interpret the content of this field.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FarmId`

The unique identifier of the farm that contains the limit.

_Required_: Yes

_Type_: String

_Pattern_: `^farm-[0-9a-f]{32}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaxCount`

The maximum number of resources constrained by this limit. When all of the resources are
in use, steps that require the limit won't be scheduled until the resource is
available.

The `maxValue` must not be 0. If the value is -1, there is no restriction on
the number of resources that can be acquired for this limit.

_Required_: Yes

_Type_: Integer

_Minimum_: `-1`

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the unique identifier of the limit.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CurrentCount`

The number of resources from the limit that are being used by jobs. The result is
delayed and may not be the count at the time that you called the operation.

`LimitId`

The unique identifier of the limit.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::Deadline::MeteredProduct

All content copied from https://docs.aws.amazon.com/.
