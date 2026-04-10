This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Workflow

The `AWS::Glue::Workflow` is an AWS Glue resource type that manages AWS Glue workflows. A workflow is a container for a set of related jobs, crawlers, and triggers in AWS Glue. Using a workflow, you can design a complex multi-job extract, transform, and load (ETL) activity that AWS Glue can execute and track as single entity.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Glue::Workflow",
  "Properties" : {
      "DefaultRunProperties" : Json,
      "Description" : String,
      "MaxConcurrentRuns" : Integer,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Glue::Workflow
Properties:
  DefaultRunProperties: Json
  Description: String
  MaxConcurrentRuns: Integer
  Name: String
  Tags:
    - Tag

```

## Properties

`DefaultRunProperties`

A collection of properties to be used as part of each execution of the workflow

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the workflow

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxConcurrentRuns`

You can use this parameter to prevent unwanted multiple updates to data, to control costs, or in some cases, to prevent exceeding the maximum number of concurrent runs of any of the component jobs. If you leave this parameter blank, there is no limit to the number of concurrent workflow runs.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the workflow representing the flow

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags to use with this workflow.

_Required_: No

_Type_: Array of [`Tag`](aws-properties-resource-tags.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the workflow name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Next

All content copied from https://docs.aws.amazon.com/.
