This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Macie::CustomDataIdentifier

The `AWS::Macie::CustomDataIdentifier` resource specifies a custom data
identifier. A _custom data identifier_ is a set of custom criteria
for Amazon Macie to use when it inspects data sources for sensitive data.
The criteria consist of a regular expression ( _regex_) that defines a
text pattern to match and, optionally, character sequences and a proximity rule that
refine the results. The character sequences can be:

- _Keywords_, which are words or phrases that must be in
proximity of text that matches the regex, or

- _Ignore words_, which are words or phrases to exclude from
the results.

By using custom data identifiers, you can supplement the managed data identifiers that
Macie provides and detect sensitive data that reflects your
particular scenarios, intellectual property, or proprietary data. For more information,
see [Building custom data identifiers](../../../macie/latest/user/custom-data-identifiers.md) in the _Amazon Macie_
_User Guide_.

An `AWS::Macie::Session` resource must exist for an AWS account before you can create an
`AWS::Macie::CustomDataIdentifier` resource for the account. Use a [DependsOn\
attribute](../userguide/aws-attribute-dependson.md) to ensure that an `AWS::Macie::Session` resource is
created before other Macie resources are created for an account. For
example, `"DependsOn": "Session"`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Macie::CustomDataIdentifier",
  "Properties" : {
      "Description" : String,
      "IgnoreWords" : [ String, ... ],
      "Keywords" : [ String, ... ],
      "MaximumMatchDistance" : Integer,
      "Name" : String,
      "Regex" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Macie::CustomDataIdentifier
Properties:
  Description: String
  IgnoreWords:
    - String
  Keywords:
    - String
  MaximumMatchDistance: Integer
  Name: String
  Regex: String
  Tags:
    - Tag

```

## Properties

`Description`

A custom description of the custom data identifier. The description can contain 1-512
characters.

Avoid including sensitive data in the description. Users of the account might be able
to see the description, depending on the actions that they're allowed to perform in
Amazon Macie.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IgnoreWords`

An array of character sequences ( _ignore words_) to exclude from
the results. If text matches the regular expression ( `Regex`) but it contains
a string in this array, Amazon Macie ignores the text and doesn't include it
in the results.

The array can contain 1-10 ignore words. Each ignore word can contain 4-90 UTF-8
characters. Ignore words are case sensitive.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Keywords`

An array of character sequences ( _keywords_), one of which must
precede and be in proximity ( `MaximumMatchDistance`) of the regular
expression ( `Regex`) to match.

The array can contain 1-50 keywords. Each keyword can contain 3-90 UTF-8 characters.
Keywords aren't case sensitive.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaximumMatchDistance`

The maximum number of characters that can exist between the end of at least one
complete character sequence specified by the `Keywords` array and the end of
text that matches the regular expression ( `Regex`). If a complete keyword
precedes all the text that matches the regular expression and the keyword is within the
specified distance, Amazon Macie includes the result.

The distance can be 1-300 characters. The default value is 50.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

A custom name for the custom data identifier. The name can contain 1-128
characters.

Avoid including sensitive data in the name of a custom data identifier. Users of the
account might be able to see the name, depending on the actions that they're allowed to
perform in Amazon Macie.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Regex`

The regular expression ( _regex_) that defines the text pattern to
match. The expression can contain 1-512 characters.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to the custom data identifier.

For more information, see [Resource tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-macie-customdataidentifier-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the `CustomDataIdentifier`. For
example, `{ "Ref": "CustomDataIdentifier" }`

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the custom data identifier.

`Id`

The unique identifier for the custom data identifier.

## Examples

The following example demonstrates how to declare an
`AWS::Macie::CustomDataIdentifier` resource.

### Creating a custom data identifier

This example creates a custom data identifier that detects six-digit character
sequences that are in proximity of certain keywords, as specified by the
`Keywords` array. If a match is a sample value, as specified by
the `IgnoreWords` array, Amazon Macie excludes that match
from the results.

#### JSON

```json

{
    "Type": "AWS::Macie::CustomDataIdentifier",
    "DependsOn": "Session",
    "Properties": {
        "Description": "My custom data identifier",
        "IgnoreWords": [
            "000000",
            "123456"
        ],
        "Keywords": [
            "employeeID",
            "employee ID"
        ],
        "MaximumMatchDistance": 20,
        "Name": "EmployeeIDCustomDataIdentifier",
        "Regex": "\\d{6}",
        "Tags": [
            {
                "Key": "Stack",
                "Value": "Production"
            }
        ]
    }
}
```

#### YAML

```yaml

Type: 'AWS::Macie::CustomDataIdentifier'
DependsOn: Session
Properties:
  Description: My custom data identifier
  IgnoreWords:
    - '000000'
    - '123456'
  Keywords:
    - 'employeeID'
    - 'employee ID'
  MaximumMatchDistance: 20
  Name: EmployeeIDCustomDataIdentifier
  Regex: '\\d{6}'
  Tags:
    - Key: Stack
      Value: Production
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
