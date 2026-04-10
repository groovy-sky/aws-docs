This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ResourceExplorer2::Index

Turns on Resource Explorer in the AWS Region in which you called this
operation by creating an index. Resource Explorer begins discovering the resources in
this Region and stores the details about the resources in the index so that they can be
queried by using the [Search](../../../../reference/resource-explorer/latest/apireference/api-search.md)
operation.

You can create either a local index that returns search results from only the AWS Region in which the index exists, or you can create an aggregator index
that returns search results from all AWS Regions in the AWS account.

For more details about what happens when you turn on Resource Explorer in an AWS Region, see [Turning on\
Resource Explorer to index your resources in an AWS Region](../../../resource-explorer/latest/userguide/manage-service-activate.md)
in the _AWS Resource Explorer User Guide._

If this is the first AWS Region in which you've created an index for
Resource Explorer, this operation also creates a service-linked role in your AWS account that allows Resource Explorer to search for your resources and
populate the index.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ResourceExplorer2::Index",
  "Properties" : {
      "Tags" : {Key: Value, ...},
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::ResourceExplorer2::Index
Properties:
  Tags:
    Key: Value
  Type: String

```

## Properties

`Tags`

The specified tags are attached to only the index created in this AWS Region. The tags don't attach to any of the resources listed in the
index.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Specifies the type of the index in this Region. For information about the aggregator
index and how it differs from a local index, see [Turning on\
cross-Region search by creating an aggregator index](../../../resource-explorer/latest/userguide/manage-aggregator-region.md) in the _AWS Resource Explorer User Guide._.

_Required_: Yes

_Type_: String

_Allowed values_: `LOCAL | AGGREGATOR`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the new index. For example:

`arn:aws:resource-explorer-2:us-east-1:123456789012:index/EXAMPLE8-90ab-cdef-fedc-EXAMPLE22222`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the new index for the AWS Region. For example:

`arn:aws:resource-explorer-2:us-east-1:123456789012:index/EXAMPLE8-90ab-cdef-fedc-EXAMPLE22222`

`IndexState`

Indicates the current state of the index. For example:

`CREATING`

## Examples

### Turning on Resource Explorer in a Region by creating an index

#### JSON

```json

{
    "Description": "Sample stack template that creates a Resource Explorer aggregator index",
    "Resources": {
        "SampleIndex": {
            "Type": "AWS::ResourceExplorer2::Index",
            "Properties": {
                "Type": "AGGREGATOR",
                "Tags": {
                    "Purpose": "ResourceExplorer Sample Stack"
                }
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: "2010-09-09"
Description: A sample template that creates a Resource Explorer aggregator index
Resources:
  SampleIndex:
    Type: 'AWS::ResourceExplorer2::Index'
    Properties:
      Type: AGGREGATOR
      Tags:
        Purpose: ResourceExplorer Sample Stack
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ResourceExplorer2::DefaultViewAssociation

AWS::ResourceExplorer2::View

All content copied from https://docs.aws.amazon.com/.
