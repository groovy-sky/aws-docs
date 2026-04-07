This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ResourceExplorer2::DefaultViewAssociation

Sets the specified view as the default for the AWS Region in which
you call this operation. If a user makes a search query that doesn't explicitly specify
the view to use, Resource Explorer chooses this default view automatically for searches
performed in this AWS Region.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ResourceExplorer2::DefaultViewAssociation",
  "Properties" : {
      "ViewArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::ResourceExplorer2::DefaultViewAssociation
Properties:
  ViewArn: String

```

## Properties

`ViewArn`

The ARN of the view to set as the default for the AWS Region and
AWS account in which you call this operation. The specified view
must already exist in the specified Region.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1011`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the identity of the principal that the view is now
associated with in the specified AWS Region. For example:

`123456789012`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AssociatedAwsPrincipal`

The unique identifier of the principal for which the specified view was made the
default for the AWS Region that contains the view. For example:

`123456789012`

## Examples

### Designating the default view for the Region

#### JSON

```json

{
    "Description": "Sample stack template that designates SampleView as the default for its Region",
    "Resources": {
        "SampleDefaultViewAssociation": {
            "Type": "AWS::ResourceExplorer2::DefaultViewAssociation",
            "Properties": {
                "ViewArn": {
                    "Ref": "SampleView"
                }
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: "2010-09-09"
Description: Sample stack template that creates a resource-based policy for SampleView
  SampleDefaultViewAssociation:
    Type: 'AWS::ResourceExplorer2::DefaultViewAssociation'
    Properties:
      ViewArn: !Ref SampleView
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS Resource Explorer

AWS::ResourceExplorer2::Index
