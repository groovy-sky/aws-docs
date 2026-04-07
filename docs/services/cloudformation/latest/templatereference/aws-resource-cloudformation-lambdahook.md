This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFormation::LambdaHook

The `AWS::CloudFormation::LambdaHook` resource creates and activates a
Lambda Hook. You can use a Lambda Hook to evaluate your
resources before allowing stack operations. This resource forwards requests for resource
evaluation to a Lambda function.

For more information, see [Lambda\
Hooks](https://docs.aws.amazon.com/cloudformation-cli/latest/hooks-userguide/lambda-hooks.html) in the _CloudFormation Hooks User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudFormation::LambdaHook",
  "Properties" : {
      "Alias" : String,
      "ExecutionRole" : String,
      "FailureMode" : String,
      "HookStatus" : String,
      "LambdaFunction" : String,
      "StackFilters" : StackFilters,
      "TargetFilters" : TargetFilters,
      "TargetOperations" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CloudFormation::LambdaHook
Properties:
  Alias: String
  ExecutionRole: String
  FailureMode: String
  HookStatus: String
  LambdaFunction: String
  StackFilters:
    StackFilters
  TargetFilters:
    TargetFilters
  TargetOperations:
    - String

```

## Properties

`Alias`

The type name alias for the Hook. This alias must be unique per account and Region.

The alias must be in the form `Name1::Name2::Name3` and must not begin with
`AWS`. For example, `Private::Lambda::MyTestHook`.

_Required_: Yes

_Type_: String

_Pattern_: `^(?!(?i)aws)[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ExecutionRole`

The IAM role that the Hook assumes to invoke your Lambda function.

_Required_: Yes

_Type_: String

_Pattern_: `arn:.+:iam::[0-9]{12}:role/.+`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FailureMode`

Specifies how the Hook responds when the Lambda function invoked by the
Hook returns a `FAILED` response.

- `FAIL`: Prevents the action from proceeding. This is helpful for
enforcing strict compliance or security policies.

- `WARN`: Issues warnings to users but allows actions to continue.
This is useful for non-critical validations or informational checks.

_Required_: Yes

_Type_: String

_Allowed values_: `FAIL | WARN`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HookStatus`

Specifies if the Hook is `ENABLED` or `DISABLED`.

_Required_: Yes

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LambdaFunction`

Specifies the Lambda function for the Hook. You can use:

- The full Amazon Resource Name (ARN) without a suffix.

- A qualified ARN with a version or alias suffix.

_Required_: Yes

_Type_: String

_Pattern_: `(arn:(aws[a-zA-Z-]*)?:lambda:)?([a-z]{2}(-gov)?(-iso([a-z])?)?-[a-z]+-\d{1}:)?(\d{12}:)?(function:)?([a-zA-Z0-9-_]+)(:(\$LATEST|[a-zA-Z0-9-_]+))?`

_Minimum_: `1`

_Maximum_: `170`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StackFilters`

Specifies the stack level filters for the Hook.

Example stack level filter in JSON:

```json

"StackFilters": {"FilteringCriteria": "ALL", "StackNames": {"Exclude": [ "stack-1", "stack-2"]}}
```

Example stack level filter in YAML:

```yaml

StackFilters:
  FilteringCriteria: ALL
  StackNames:
    Exclude:
      - stack-1
      - stack-2
```

_Required_: No

_Type_: [StackFilters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudformation-lambdahook-stackfilters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetFilters`

Specifies the target filters for the Hook.

Example target filter in JSON:

```json

"TargetFilters": {"Actions": [ "CREATE", "UPDATE", "DELETE" ]}
```

Example target filter in YAML:

```yaml

TargetFilters:
  Actions:
    - CREATE
    - UPDATE
    - DELETE
```

_Required_: No

_Type_: [TargetFilters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudformation-lambdahook-targetfilters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetOperations`

Specifies the list of operations the Hook is run against. For more information, see
[Hook targets](https://docs.aws.amazon.com/cloudformation-cli/latest/hooks-userguide/hooks-concepts.html#hook-terms-hook-target) in the _CloudFormation Hooks User_
_Guide_.

Valid values: `STACK` \| `RESOURCE` \| `CHANGE_SET` \|
`CLOUD_CONTROL`

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Hook Amazon Resource Name (ARN). For example:
`arn:aws:cloudformation:us-west-2:123456789012:type/hook/MyLambdaHook`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`HookArn`

Returns the ARN of a Lambda Hook.

## Examples

### Creating a Lambda Hook in a template

The following example demonstrates how to create a Lambda Hook in
a template.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "Create a Lambda Hook",
    "Parameters": {
        "HookFunctionArn": {
            "Description": "Hook Lambda Function ARN",
            "Type": "String"
        },
        "HookName": {
            "Description": "The name of your Hook",
            "Type": "String",
            "Default": "Test::Lambda::Hook",
            "AllowedPattern": "^(?!(?i)aws)[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}$"
        }
    },
    "Resources": {
        "LambdaInvokerHookRole": {
            "Type": "AWS::IAM::Role",
            "Properties": {
                "AssumeRolePolicyDocument": {
                    "Version": "2012-10-17",
                    "Statement": [
                        {
                            "Effect": "Allow",
                            "Principal": {
                                "Service": ["hooks.cloudformation.amazonaws.com"]
                            },
                            "Action": "sts:AssumeRole"
                        }
                    ]
                },
                "Path": "/",
                "Policies": [
                    {
                        "PolicyName": "LambdaInvokerHookPolicy",
                        "PolicyDocument": {
                            "Version": "2012-10-17",
                            "Statement": [
                                {
                                    "Effect": "Allow",
                                    "Action": ["lambda:InvokeFunction"],
                                    "Resource": {"Ref" : "HookFunctionArn"}
                                }
                            ]
                        }
                    }
                ]
            }
        },
        "MyLambdaHook": {
            "Type": "AWS::CloudFormation::LambdaHook",
            "Properties": {
                "LambdaFunction": {"Ref" : "HookFunctionArn"},
                "HookStatus": "ENABLED",
                "TargetOperations": [
                    "RESOURCE",
                    "STACK"
                ],
                "FailureMode": "WARN",
                "Alias": {"Ref" : "HookName"},
                "ExecutionRole": {
                    "Fn::GetAtt": [
                        "LambdaInvokerHookRole",
                        "Arn"
                    ]
                },
                "TargetFilters": {
                    "Actions": [
                        "CREATE",
                        "UPDATE",
                        "DELETE"
                    ]
                },
                "StackFilters": {
                    "FilteringCriteria": "ALL",
                    "StackNames": {
                        "Exclude": [{"Ref" : "AWS::StackName"}]
                    }
                }
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: Create a Lambda Hook
Parameters:
  HookFunctionArn:
    Description: Hook Lambda Function ARN
    Type: String
  HookName:
    Description: The name of your Hook
    Type: String
    Default: 'Test::Lambda::Hook'
    AllowedPattern: '^(?!(?i)aws)[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}$'
Resources:
  LambdaInvokerHookRole:
    Type: AWS::IAM::Role
    Properties:
      AssumeRolePolicyDocument:
        Version: 2012-10-17
        Statement:
          - Effect: Allow
            Principal:
              Service:
                - hooks.cloudformation.amazonaws.com
            Action: 'sts:AssumeRole'
      Path: /
      Policies:
        - PolicyName: LambdaInvokerHookPolicy
          PolicyDocument:
            Version: 2012-10-17
            Statement:
              - Effect: Allow
                Action:
                  - 'lambda:InvokeFunction'
                Resource: !Ref HookFunctionArn
  MyLambdaHook:
    Type: AWS::CloudFormation::LambdaHook
    Properties:
      LambdaFunction: !Ref HookFunctionArn
      HookStatus: ENABLED
      TargetOperations:
        - RESOURCE
        - STACK
      FailureMode: WARN
      Alias: !Ref HookName
      ExecutionRole: !GetAtt LambdaInvokerHookRole.Arn
      TargetFilters:
        Actions:
          - CREATE
          - UPDATE
          - DELETE
      StackFilters:
        FilteringCriteria: ALL
        StackNames:
          Exclude:
            - !Ref AWS::StackName
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

LoggingConfig

StackFilters
