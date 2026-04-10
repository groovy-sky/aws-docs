This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ResourceExplorer2::View

Creates a view that users can query by using the [Search](../../../../reference/resource-explorer/latest/apireference/api-search.md)
operation. Results from queries that you make using this view include only resources
that match the view's `Filters`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ResourceExplorer2::View",
  "Properties" : {
      "Filters" : SearchFilter,
      "IncludedProperties" : [ IncludedProperty, ... ],
      "Scope" : String,
      "Tags" : {Key: Value, ...},
      "ViewName" : String
    }
}

```

### YAML

```yaml

Type: AWS::ResourceExplorer2::View
Properties:
  Filters:
    SearchFilter
  IncludedProperties:
    - IncludedProperty
  Scope: String
  Tags:
    Key: Value
  ViewName: String

```

## Properties

`Filters`

An array of strings that include search keywords, prefixes, and operators that filter
the results that are returned for queries made using this view. When you use this view
in a [Search](../../../../reference/resource-explorer/latest/apireference/api-search.md)
operation, the filter string is combined with the search's `QueryString`
parameter using a logical `AND` operator.

For information about the supported syntax, see [Search query\
reference for Resource Explorer](../../../resource-explorer/latest/userguide/using-search-query-syntax.md) in the _AWS_
_Resource Explorer User Guide_.

###### Important

This query string in the context of this operation supports only [filter prefixes](../../../resource-explorer/latest/userguide/using-search-query-syntax.md#query-syntax-filters) with optional [operators](../../../resource-explorer/latest/userguide/using-search-query-syntax.md#query-syntax-operators). It doesn't support free-form text. For example, the string
`region:us* service:ec2 -tag:stage=prod` includes all Amazon EC2 resources in any AWS Region that begin with the
letters `us` and are _not_ tagged with a key
`Stage` that has the value `prod`.

_Required_: No

_Type_: [SearchFilter](aws-properties-resourceexplorer2-view-searchfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludedProperties`

A list of fields that provide additional information about the view.

_Required_: No

_Type_: Array of [IncludedProperty](aws-properties-resourceexplorer2-view-includedproperty.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scope`

The root ARN of the account, an organizational unit (OU), or an organization ARN. If left empty, the default is account.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Tag key and value pairs that are attached to the view.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ViewName`

The name of the new view.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9\-]{1,64}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the new view. For example:

`arn:aws:resource-explorer-2:us-east-1:123456789012:view/CFNStackView2/EXAMPLE8-90ab-cdef-fedc-EXAMPLE22222`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ViewArn`

The ARN of the new view. For example:

`arn:aws:resource-explorer-2:us-east-1:123456789012:view/MyView/EXAMPLE8-90ab-cdef-fedc-EXAMPLE22222`

## Examples

### Creating a view for users to search an index

#### JSON

```json

{
    "Description": "Sample stack template that creates a Resource Explorer view for the SampleIndex",
    "Resources": {
        "SampleView": {
            "Type": "AWS::ResourceExplorer2::View",
            "Properties": {
                "ViewName": "mySampleView",
                "IncludedProperties": [
                    {
                        "Name": "tags"
                    }
                ],
                "Tags": {
                    "Purpose": "ResourceExplorer Sample Stack"
                }
            },
            "DependsOn": "SampleIndex"
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: "2010-09-09"
Description: A sample template that creates a Resource Explorer view for the SampleIndex
  SampleView:
    Type: 'AWS::ResourceExplorer2::View'
    Properties:
      ViewName: mySampleView
      IncludedProperties:
        - Name: tags
      Tags:
        Purpose: ResourceExplorer Sample Stack
    DependsOn: SampleIndex
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ResourceExplorer2::Index

IncludedProperty

All content copied from https://docs.aws.amazon.com/.
