This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::FunctionDefinition

The `AWS::Greengrass::FunctionDefinition` resource represents a function
definition for AWS IoT Greengrass. Function definitions are used to organize your function
definition versions.

Function definitions can reference multiple function definition versions. All function
definition versions must be associated with a function definition. Each function definition
version can contain one or more functions.

###### Note

When you create a function definition, you can optionally include an initial function
definition version. To associate a function definition version later, create an [`AWS::Greengrass::FunctionDefinitionVersion`](../userguide/aws-resource-greengrass-functiondefinitionversion.md) resource and
specify the ID of this function definition.

After you create the function definition version that contains the functions you want
to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](../userguide/aws-resource-greengrass-group.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Greengrass::FunctionDefinition",
  "Properties" : {
      "InitialVersion" : FunctionDefinitionVersion,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Greengrass::FunctionDefinition
Properties:
  InitialVersion:
    FunctionDefinitionVersion
  Name: String
  Tags:
    - Tag

```

## Properties

`InitialVersion`

The function definition version to include when the function definition is created. A
function definition version contains a list of [`function`](../userguide/aws-properties-greengrass-functiondefinition-function.md) property types.

###### Note

To associate a function definition version after the function definition is created,
create an [`AWS::Greengrass::FunctionDefinitionVersion`](../userguide/aws-resource-greengrass-functiondefinitionversion.md) resource and
specify the ID of this function definition.

_Required_: No

_Type_: [FunctionDefinitionVersion](aws-properties-greengrass-functiondefinition-functiondefinitionversion.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the function definition.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Application-specific metadata to attach to the function definition. You can use tags in
IAM policies to control access to AWS IoT Greengrass resources. You
can also use tags to categorize your resources. For more information, see [Tagging Your\
AWS IoT Greengrass Resources](../../../greengrass/v1/developerguide/tagging.md) in the _AWS IoT Greengrass Version 1 Developer Guide_.

This `Json` property type is processed as a map of key-value pairs. It uses
the following format, which is different from most `Tags` implementations in
CloudFormation templates.

```json

"Tags": {
    "KeyName0": "value",
    "KeyName1": "value",
    "KeyName2": "value"
}
```

_Required_: No

_Type_: Array of [`Tag`](aws-properties-resource-tags.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the function definition, such as
`1234a5b6-78cd-901e-2fgh-3i45j6k178l9`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the `FunctionDefinition`, such as
`arn:aws:greengrass:us-east-1:123456789012:/greengrass/definition/functions/1234a5b6-78cd-901e-2fgh-3i45j6k178l9`.

`Id`

The ID of the `FunctionDefinition`, such as
`1234a5b6-78cd-901e-2fgh-3i45j6k178l9`.

`LatestVersionArn`

The ARN of the last `FunctionDefinitionVersion` that was added to the
`FunctionDefinition`, such as `arn:aws:greengrass:us-east-1:123456789012:/greengrass/definition/functions/1234a5b6-78cd-901e-2fgh-3i45j6k178l9/versions/9876ac30-4bdb-4f9d-95af-b5fdb66be1a2`.

`Name`

The name of the `FunctionDefinition`, such as
`MyFunctionDefinition`.

## Examples

### Function Definition Snippet

The following snippet defines a function definition resource with an initial
version that contains a function. In this example, the Lambda function
is created in another stack and is referenced using the `ImportValue`
function.

For an example of a complete template, see the [`AWS::Greengrass::Group`](../userguide/aws-resource-greengrass-group.md) resource.

#### JSON

```json

"TestFunctionDefinition": {
    "Type": "AWS::Greengrass::FunctionDefinition",
    "Properties": {
        "Name": "DemoTestFunctionDefinition",
        "InitialVersion": {
            "DefaultConfig": {
                "Execution": {
                    "IsolationMode": "GreengrassContainer"
                }
            },
            "Functions": [
                {
                    "Id": "TestLambda1",
                    "FunctionArn": {
                        "Fn::ImportValue": "TestCanaryLambdaVersionArn"
                    },
                    "FunctionConfiguration": {
                        "Pinned": "false",
                        "Executable": "run.exe",
                        "ExecArgs": "argument1",
                        "MemorySize": "256",
                        "Timeout": "3000",
                        "EncodingType": "binary",
                        "Environment": {
                            "Variables": {
                                "variable1": "value1"
                            },
                            "ResourceAccessPolicies": [
                                {
                                    "ResourceId": "ResourceId1",
                                    "Permission": "ro"
                                },
                                {
                                    "ResourceId": "ResourceId2",
                                    "Permission": "rw"
                                }
                            ],
                            "AccessSysfs": "true",
                            "Execution": {
                                "RunAs": {
                                    "Uid": "1",
                                    "Gid": "10"
                                }
                            }
                        }
                    }
                }
            ]
        }
    }
}
```

#### YAML

```yaml

TestFunctionDefinition:
  Type: 'AWS::Greengrass::FunctionDefinition'
  Properties:
    Name: DemoTestFunctionDefinition
    InitialVersion:
      DefaultConfig:
        Execution:
          IsolationMode: GreengrassContainer
      Functions:
        - Id: TestLambda1
          FunctionArn: !ImportValue TestCanaryLambdaVersionArn
          FunctionConfiguration:
            Pinned: 'false'
            Executable: run.exe
            ExecArgs: argument1
            MemorySize: '256'
            Timeout: '3000'
            EncodingType: binary
            Environment:
              Variables:
                variable1: value1
              ResourceAccessPolicies:
                - ResourceId: ResourceId1
                  Permission: ro
                - ResourceId: ResourceId2
                  Permission: rw
              AccessSysfs: 'true'
              Execution:
                RunAs:
                  Uid: '1'
                  Gid: '10'
```

## See also

- [CreateFunctionDefinition](../../../../reference/greengrass/v1/apireference/createfunctiondefinition-post.md) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](../../../greengrass/v1/developerguide.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Device

DefaultConfig

All content copied from https://docs.aws.amazon.com/.
