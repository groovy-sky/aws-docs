This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::FunctionDefinitionVersion

The `AWS::Greengrass::FunctionDefinitionVersion` resource represents a
function definition version for AWS IoT Greengrass. A function definition version contains
contain a list of functions.

###### Note

To create a function definition version, you must specify the ID of the function
definition that you want to associate with the version. For information about creating a
function definition, see [`AWS::Greengrass::FunctionDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-functiondefinition.html).

After you create a function definition version that contains the functions you want
to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Greengrass::FunctionDefinitionVersion",
  "Properties" : {
      "DefaultConfig" : DefaultConfig,
      "FunctionDefinitionId" : String,
      "Functions" : [ Function, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Greengrass::FunctionDefinitionVersion
Properties:
  DefaultConfig:
    DefaultConfig
  FunctionDefinitionId: String
  Functions:
    - Function

```

## Properties

`DefaultConfig`

The default configuration that applies to all Lambda functions in the
group. Individual Lambda functions can override these settings.

_Required_: No

_Type_: [DefaultConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-greengrass-functiondefinitionversion-defaultconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FunctionDefinitionId`

The ID of the function definition associated with this version. This value is a
GUID.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Functions`

The functions in this version.

_Required_: Yes

_Type_: Array of [Function](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-greengrass-functiondefinitionversion-function.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the function definition
version, such as `arn:aws:greengrass:us-east-1:123456789012:/greengrass/definition/functions/1234a5b6-78cd-901e-2fgh-3i45j6k178l9/versions/9876ac30-4bdb-4f9d-95af-b5fdb66be1a2`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## Examples

### Function Definition Version Snippet

The following snippet defines function definition and function definition version
resources. The function definition version references the function definition and
contains a function. In this example, the Lambda function is created in
another stack and is referenced using the `ImportValue` function.

For an example of a complete template, see the [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) resource.

#### JSON

```json

"TestFunctionDefinition": {
    "Type": "AWS::Greengrass::FunctionDefinition",
    "Properties": {
        "Name": "DemoTestFunctionDefinition"
    }
},
"TestFunctionDefinitionVersion": {
    "Type": "AWS::Greengrass::FunctionDefinitionVersion",
    "Properties": {
        "FunctionDefinitionId": {
            "Fn::GetAtt": [
                "TestFunctionDefinition",
                "Id"
            ]
        },
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
                    "Pinned": "true",
                    "Executable": "run.exe",
                    "ExecArgs": "argument1",
                    "MemorySize": "512",
                    "Timeout": "2000",
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
                        "AccessSysfs": "false",
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
```

#### YAML

```yaml

TestFunctionDefinition:
  Type: 'AWS::Greengrass::FunctionDefinition'
  Properties:
    Name: DemoTestFunctionDefinition
TestFunctionDefinitionVersion:
  Type: 'AWS::Greengrass::FunctionDefinitionVersion'
  Properties:
    FunctionDefinitionId: !GetAtt
      - TestFunctionDefinition
      - Id
    DefaultConfig:
      Execution:
        IsolationMode: GreengrassContainer
    Functions:
      - Id: TestLambda1
        FunctionArn: !ImportValue TestCanaryLambdaVersionArn
        FunctionConfiguration:
          Pinned: 'true'
          Executable: run.exe
          ExecArgs: argument1
          MemorySize: '512'
          Timeout: '2000'
          EncodingType: binary
          Environment:
            Variables:
              variable1: value1
            ResourceAccessPolicies:
              - ResourceId: ResourceId1
                Permission: ro
              - ResourceId: ResourceId2
                Permission: rw
            AccessSysfs: 'false'
            Execution:
              RunAs:
                Uid: '1'
                Gid: '10'
```

## See also

- [CreateFunctionDefinitionVersion](https://docs.aws.amazon.com/greengrass/v1/apireference/createfunctiondefinitionversion-post.html) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](https://docs.aws.amazon.com/greengrass/v1/developerguide)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RunAs

DefaultConfig
