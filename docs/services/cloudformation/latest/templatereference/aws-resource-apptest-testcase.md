This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppTest::TestCase

Creates a test case for an application.

For more information about test cases, see [Test cases](../../../m2/latest/userguide/testing-test-cases.md) and [Application Testing\
concepts](../../../m2/latest/userguide/concepts-apptest.md) in the _AWS Mainframe Modernization User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppTest::TestCase",
  "Properties" : {
      "Description" : String,
      "Name" : String,
      "Steps" : [ Step, ... ],
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::AppTest::TestCase
Properties:
  Description: String
  Name: String
  Steps:
    - Step
  Tags:
    Key: Value

```

## Properties

`Description`

The description of the test case.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the test case.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z][A-Za-z0-9_\-]{1,59}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Steps`

The steps in the test case.

_Required_: Yes

_Type_: Array of [Step](aws-properties-apptest-testcase-step.md)

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The specified tags of the test case.

_Required_: No

_Type_: Object of String

_Pattern_: `^(?!aws:).+$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the test case Amazon Resource Name (ARN), such as the
following:

`{ "Ref": "SampleTestCase" }`

Returns a value similar to the following:

`arn:aws:apptest:us-east-1:123456789012:testcase/y3ca6bhaife2bcvxar3lpivfou`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreationTime`

The creation time of the test case.

`LastUpdateTime`

The last update time of the test case.

`Status`

The status of the test case.

`TestCaseArn`

The Amazon Resource Name (ARN) of the test case.

`TestCaseId`

The response test case ID of the test case.

`TestCaseVersion`

The version of the test case.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Mainframe Modernization Application Testing

Batch

All content copied from https://docs.aws.amazon.com/.
