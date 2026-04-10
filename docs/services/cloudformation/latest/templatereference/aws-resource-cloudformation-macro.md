This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFormation::Macro

The `AWS::CloudFormation::Macro` resource is a CloudFormation
resource type that creates a CloudFormation macro to perform custom processing
on CloudFormation templates.

For more information, see [Perform custom\
processing on CloudFormation templates with template macros](../userguide/template-macros.md) in the
_CloudFormation User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudFormation::Macro",
  "Properties" : {
      "Description" : String,
      "FunctionName" : String,
      "LogGroupName" : String,
      "LogRoleARN" : String,
      "Name" : String
    }
}

```

### YAML

```yaml

Type: AWS::CloudFormation::Macro
Properties:
  Description: String
  FunctionName: String
  LogGroupName: String
  LogRoleARN: String
  Name: String

```

## Properties

`Description`

A description of the macro.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FunctionName`

The Amazon Resource Name (ARN) of the underlying Lambda function that you
want CloudFormation to invoke when the macro is run.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogGroupName`

The CloudWatch Logs group to which CloudFormation sends error logging
information when invoking the macro's underlying Lambda function.

This will be an existing CloudWatch Logs LogGroup. Neither CloudFormation
or Lambda will create the group.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogRoleARN`

The ARN of the role CloudFormation should assume when sending log entries to
CloudWatch Logs.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the macro. The name of the macro must be unique across all macros in the
account.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name. For example:

`{ "Ref": "myMacro" }`

For the macro `myMacro`, `Ref` returns the name of the
macro.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

Returns a unique identifier for the resource.

## See also

- [Perform\
custom processing on CloudFormation templates with template\
macros](../userguide/template-macros.md) in the _CloudFormation User_
_Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TargetFiltersItems

AWS::CloudFormation::ModuleDefaultVersion

All content copied from https://docs.aws.amazon.com/.
