This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Macie::FindingsFilter

The `AWS::Macie::FindingsFilter` resource specifies a findings filter. In Amazon Macie, a
_findings filter_, also referred to as a _filter_
_rule_, is a set of custom criteria that specifies which findings to
include or exclude from the results of a query for findings. The criteria can help you
identify and focus on findings that have specific characteristics, such as severity,
type, or the name of an affected AWS resource. You can also configure a
findings filter to suppress (automatically archive) findings that match the filter's
criteria. For more information, see [Filtering Macie findings](../../../macie/latest/user/findings-filter-overview.md) in
the _Amazon Macie User Guide_.

An `AWS::Macie::Session` resource must exist for an AWS account before you can create an
`AWS::Macie::FindingsFilter` resource for the account. Use a [DependsOn\
attribute](../userguide/aws-attribute-dependson.md) to ensure that an `AWS::Macie::Session` resource is
created before other Macie resources are created for an account. For
example, `"DependsOn": "Session"`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Macie::FindingsFilter",
  "Properties" : {
      "Action" : String,
      "Description" : String,
      "FindingCriteria" : FindingCriteria,
      "Name" : String,
      "Position" : Integer,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Macie::FindingsFilter
Properties:
  Action: String
  Description: String
  FindingCriteria:
    FindingCriteria
  Name: String
  Position: Integer
  Tags:
    - Tag

```

## Properties

`Action`

The action to perform on findings that match the filter criteria
( `FindingCriteria`). Valid values are:

- `ARCHIVE` \- Suppress (automatically archive) the findings.

- `NOOP` \- Don't perform any action on the findings.

_Required_: No

_Type_: String

_Allowed values_: `ARCHIVE | NOOP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A custom description of the findings filter. The description can contain 1-512
characters.

Avoid including sensitive data in the description. Users of the account might be able
to see the description, depending on the actions that they're allowed to perform in
Amazon Macie.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FindingCriteria`

The criteria to use to filter findings.

_Required_: Yes

_Type_: [FindingCriteria](aws-properties-macie-findingsfilter-findingcriteria.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A custom name for the findings filter. The name can contain 3-64 characters.

Avoid including sensitive data in the name. Users of the account might be able to see
the name, depending on the actions that they're allowed to perform in Amazon Macie.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Position`

The position of the findings filter in the list of saved filter rules on the Amazon Macie console. This value also determines the order in which the filter
is applied to findings, relative to other filters that are also applied to
findings.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to the findings filter.

For more information, see [Resource tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-macie-findingsfilter-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the `FindingsFilter`. For example,
`{ "Ref": "FindingsFilter" }`.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the findings filter.

`Id`

The unique identifier for the findings filter.

## Examples

The following example demonstrates how to declare an
`AWS::Macie::FindingsFilter` resource.

### Creating a findings filter that filters by account ID

This example creates a findings filter that suppresses (automatically
archives) findings for AWS resources that are owned by a
specific account ( `123456789012`).

#### JSON

```json

{
    "Type": "AWS::Macie::FindingsFilter",
    "DependsOn": "Session",
    "Properties": {
        "Action": "ARCHIVE",
        "Description": "My custom findings filter",
        "FindingCriteria": {
            "Criterion": {
                "accountId": {
                    "eq": [
                        "123456789012"
                    ]
                }
            }
        },
        "Name": "MyFilterName",
        "Position": 1,
        "Tags": [
            {
                "Key": "CostCenter",
                "Value": "CC12345"
            }
        ]

    }
}
```

#### YAML

```yaml

Type: 'AWS::Macie::FindingsFilter'
DependsOn: Session
Properties:
  Action: ARCHIVE
  Description: My custom findings filter
  FindingCriteria:
    Criterion:
      accountId:
        eq:
          - '123456789012'
  Name: MyFilterName
  Position: 1
  Tags:
    - Key: CostCenter
      Value: CC12345
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

CriterionAdditionalProperties

All content copied from https://docs.aws.amazon.com/.
