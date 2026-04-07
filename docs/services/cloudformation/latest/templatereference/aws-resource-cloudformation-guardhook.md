This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFormation::GuardHook

The `AWS::CloudFormation::GuardHook` resource creates and activates a
Guard Hook. Using the Guard domain
specific language (DSL), you can author Guard Hooks to evaluate
your resources before allowing stack operations.

For more information, see [Guard Hooks](https://docs.aws.amazon.com/cloudformation-cli/latest/hooks-userguide/guard-hooks.html) in the _CloudFormation Hooks_
_User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudFormation::GuardHook",
  "Properties" : {
      "Alias" : String,
      "ExecutionRole" : String,
      "FailureMode" : String,
      "HookStatus" : String,
      "LogBucket" : String,
      "Options" : Options,
      "RuleLocation" : S3Location,
      "StackFilters" : StackFilters,
      "TargetFilters" : TargetFilters,
      "TargetOperations" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CloudFormation::GuardHook
Properties:
  Alias: String
  ExecutionRole: String
  FailureMode: String
  HookStatus: String
  LogBucket: String
  Options:
    Options
  RuleLocation:
    S3Location
  StackFilters:
    StackFilters
  TargetFilters:
    TargetFilters
  TargetOperations:
    - String

```

## Properties

`Alias`

The type name alias for the Hook. This alias must be unique per account and
Region.

The alias must be in the form `Name1::Name2::Name3` and must not begin with
`AWS`. For example, `Private::Guard::MyTestHook`.

_Required_: Yes

_Type_: String

_Pattern_: `^(?!(?i)aws)[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ExecutionRole`

The IAM role that the Hook assumes to retrieve your Guard rules
from S3 and optionally write a detailed Guard output report
back.

_Required_: Yes

_Type_: String

_Pattern_: `arn:.+:iam::[0-9]{12}:role/.+`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FailureMode`

Specifies how the Hook responds when rules fail their evaluation.

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

`LogBucket`

Specifies the name of an S3 bucket to store the Guard output
report. This report contains the results of your Guard rule
validations.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Options`

Specifies the S3 location of your input parameters.

_Required_: No

_Type_: [Options](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudformation-guardhook-options.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleLocation`

Specifies the S3 location of your Guard rules.

_Required_: Yes

_Type_: [S3Location](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudformation-guardhook-s3location.html)

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

_Type_: [StackFilters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudformation-guardhook-stackfilters.html)

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

_Type_: [TargetFilters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudformation-guardhook-targetfilters.html)

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
`arn:aws:cloudformation:us-west-2:123456789012:type/hook/MyGuardHook`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`HookArn`

Returns the ARN of a Guard Hook.

## Examples

### Creating a Guard Hook in a template

The following example demonstrates how to create a Guard Hook in a template.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "Create a Guard Hook",
    "Parameters": {
        "GuardRulesS3Bucket": {
            "Description": "S3 bucket where your rules are",
            "Type": "String"
        },
        "GuardRulesS3Path": {
            "Description": "Location within GuardRulesS3Bucket where your Guard rules are",
            "Type": "String"
        },
        "GuardOutputBucket": {
            "Description": "S3 bucket to put Guard output",
            "Type": "String"
        },
        "HookName": {
            "Description": "The name of your Hook",
            "Type": "String",
            "Default": "Test::Guard::Hook",
            "AllowedPattern": "^(?!(?i)aws)[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}$"
        }
    },
    "Resources": {
        "GuardHookRole": {
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
                        "PolicyName": "root",
                        "PolicyDocument": {
                            "Version": "2012-10-17",
                            "Statement": [
                                {
                                    "Effect": "Allow",
                                    "Action": [
                                        "s3:GetObject",
                                        "s3:GetObjectVersion",
                                        "s3:ListBucket"
                                    ],
                                    "Resource": [
                                        {"Fn::Sub": "arn:aws:s3:::${GuardRulesS3Bucket}"},
                                        {"Fn::Sub": "arn:aws:s3:::${GuardRulesS3Bucket}/*"}
                                    ]
                                },
                                {
                                    "Effect": "Allow",
                                    "Action": ["s3:PutObject"],
                                    "Resource": [{"Fn::Sub": "arn:aws:s3:::${GuardOutputBucket}/*"}]
                                }
                            ]
                        }
                    }
                ]
            }
        },
        "GuardHook": {
            "Type": "AWS::CloudFormation::GuardHook",
            "Properties": {
                "TargetOperations": [
                    "RESOURCE",
                    "STACK"
                ],
                "Alias": {"Ref": "HookName"},
                "ExecutionRole": {
                    "Fn::GetAtt": [
                        "GuardHookRole",
                        "Arn"
                    ]
                },
                "FailureMode": "WARN",
                "HookStatus": "ENABLED",
                "LogBucket": {"Ref": "GuardOutputBucket"},
                "RuleLocation": {
                    "Uri": {"Fn::Sub": "s3://${GuardRulesS3Bucket}/${GuardRulesS3Path}"}
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
                        "Exclude": [{"Ref": "AWS::StackName"}]
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
Description: Create a Guard Hook
Parameters:
  GuardRulesS3Bucket:
    Description: S3 bucket where your rules are
    Type: String
  GuardRulesS3Path:
    Description: Location within GuardRulesS3Bucket where your Guard rules are
    Type: String
  GuardOutputBucket:
    Description: S3 bucket to put Guard output
    Type: String
  HookName:
    Description: The name of your Hook
    Type: String
    Default: 'Test::Guard::Hook'
    AllowedPattern: '^(?!(?i)aws)[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}$'
Resources:
  GuardHookRole:
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
        - PolicyName: root
          PolicyDocument:
            Version: "2012-10-17"
            Statement:
              - Effect: Allow
                Action:
                  - s3:GetObject
                  - s3:GetObjectVersion
                  - s3:ListBucket
                Resource:
                  - !Sub arn:aws:s3:::${GuardRulesS3Bucket}
                  - !Sub arn:aws:s3:::${GuardRulesS3Bucket}/*
              - Effect: Allow
                Action:
                  - s3:PutObject
                Resource:
                  - !Sub arn:aws:s3:::${GuardOutputBucket}/*
  GuardHook:
    Type: AWS::CloudFormation::GuardHook
    Properties:
      TargetOperations:
        - RESOURCE
        - STACK
      Alias: !Ref HookName
      ExecutionRole: !GetAtt GuardHookRole.Arn
      FailureMode: WARN
      HookStatus: ENABLED
      LogBucket: !Ref GuardOutputBucket
      RuleLocation:
        Uri: !Sub s3://${GuardRulesS3Bucket}/${GuardRulesS3Path}
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

AWS::CloudFormation::CustomResource

Options
