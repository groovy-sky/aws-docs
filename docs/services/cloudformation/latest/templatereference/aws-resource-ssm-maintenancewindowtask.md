This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSM::MaintenanceWindowTask

The `AWS::SSM::MaintenanceWindowTask` resource defines information about a
task for an AWS Systems Manager maintenance window. For more information, see
[RegisterTaskWithMaintenanceWindow](../../../../reference/systems-manager/latest/apireference/api-registertaskwithmaintenancewindow.md) in the _AWS Systems Manager_
_API Reference_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SSM::MaintenanceWindowTask",
  "Properties" : {
      "CutoffBehavior" : String,
      "Description" : String,
      "LoggingInfo" : LoggingInfo,
      "MaxConcurrency" : String,
      "MaxErrors" : String,
      "Name" : String,
      "Priority" : Integer,
      "ServiceRoleArn" : String,
      "Targets" : [ Target, ... ],
      "TaskArn" : String,
      "TaskInvocationParameters" : TaskInvocationParameters,
      "TaskParameters" : Json,
      "TaskType" : String,
      "WindowId" : String
    }
}

```

### YAML

```yaml

Type: AWS::SSM::MaintenanceWindowTask
Properties:
  CutoffBehavior: String
  Description: String
  LoggingInfo:
    LoggingInfo
  MaxConcurrency: String
  MaxErrors: String
  Name: String
  Priority: Integer
  ServiceRoleArn: String
  Targets:
    - Target
  TaskArn: String
  TaskInvocationParameters:
    TaskInvocationParameters
  TaskParameters: Json
  TaskType: String
  WindowId: String

```

## Properties

`CutoffBehavior`

The specification for whether tasks should continue to run after the cutoff time specified
in the maintenance windows is reached.

_Required_: No

_Type_: String

_Allowed values_: `CONTINUE_TASK | CANCEL_TASK`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the task.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoggingInfo`

###### Note

`LoggingInfo` has been deprecated. To specify an Amazon S3
bucket to contain logs for Run Command tasks, instead use the
`OutputS3BucketName` and `OutputS3KeyPrefix` options in
the `TaskInvocationParameters` structure. For information about how
Systems Manager handles these options for the supported maintenance window
task types, see [AWS::SSM::MaintenanceWindowTask\
MaintenanceWindowRunCommandParameters](../userguide/aws-properties-ssm-maintenancewindowtask-maintenancewindowruncommandparameters.md).

Information about an Amazon S3 bucket to write Run Command task-level logs
to.

_Required_: No

_Type_: [LoggingInfo](aws-properties-ssm-maintenancewindowtask-logginginfo.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxConcurrency`

The maximum number of targets this task can be run for, in parallel.

###### Note

Although this element is listed as "Required: No", a value can be omitted only when you are
registering or updating a [targetless\
task](../../../systems-manager/latest/userguide/maintenance-windows-targetless-tasks.md) You must provide a value in all other cases.

For maintenance window tasks without a target specified, you can't supply a value for this
option. Instead, the system inserts a placeholder value of `1`. This value doesn't
affect the running of your task.

_Required_: No

_Type_: String

_Pattern_: `^([1-9][0-9]*|[1-9][0-9]%|[1-9]%|100%)$`

_Minimum_: `1`

_Maximum_: `7`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxErrors`

The maximum number of errors allowed before this task stops being scheduled.

###### Note

Although this element is listed as "Required: No", a value can be omitted only when you are
registering or updating a [targetless\
task](../../../systems-manager/latest/userguide/maintenance-windows-targetless-tasks.md) You must provide a value in all other cases.

For maintenance window tasks without a target specified, you can't supply a value for this
option. Instead, the system inserts a placeholder value of `1`. This value doesn't
affect the running of your task.

_Required_: No

_Type_: String

_Pattern_: `^([1-9][0-9]*|[0]|[1-9][0-9]%|[0-9]%|100%)$`

_Minimum_: `1`

_Maximum_: `7`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The task name.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-.]{3,128}$`

_Minimum_: `3`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Priority`

The priority of the task in the maintenance window. The lower the number, the higher the
priority. Tasks that have the same priority are scheduled in parallel.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceRoleArn`

The Amazon Resource Name (ARN) of the IAM service role for
AWS Systems Manager to assume when running a maintenance window task. If you do not specify a
service role ARN, Systems Manager uses a service-linked role in your account. If no
appropriate service-linked role for Systems Manager exists in your account, it is created when
you run `RegisterTaskWithMaintenanceWindow`.

However, for an improved security posture, we strongly recommend creating a custom
policy and custom service role for running your maintenance window tasks. The policy
can be crafted to provide only the permissions needed for your particular
maintenance window tasks. For more information, see [Setting up Maintenance Windows](../../../systems-manager/latest/userguide/sysman-maintenance-permissions.md) in the in the
_AWS Systems Manager User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Targets`

The targets, either instances or window target IDs.

- Specify instances using `Key=InstanceIds,Values=instanceid1,instanceid2`.

- Specify window target IDs using `Key=WindowTargetIds,Values=window-target-id-1,window-target-id-2`.

_Required_: No

_Type_: Array of [Target](aws-properties-ssm-maintenancewindowtask-target.md)

_Minimum_: `0`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TaskArn`

The resource that the task uses during execution.

For `RUN_COMMAND` and `AUTOMATION` task types,
`TaskArn` is the SSM document name or Amazon Resource Name (ARN).

For `LAMBDA` tasks, `TaskArn` is the function name or
ARN.

For `STEP_FUNCTIONS` tasks, `TaskArn` is the state machine
ARN.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TaskInvocationParameters`

The parameters to pass to the task when it runs. Populate only the fields that match
the task type. All other fields should be empty.

###### Important

When you update a maintenance window task that has options specified in
`TaskInvocationParameters`, you must provide again all the
`TaskInvocationParameters` values that you want to retain. The values
you do not specify again are removed. For example, suppose that when you registered
a Run Command task, you specified `TaskInvocationParameters` values for
`Comment`, `NotificationConfig`, and
`OutputS3BucketName`. If you update the maintenance window task and
specify only a different `OutputS3BucketName` value, the values for
`Comment` and `NotificationConfig` are removed.

_Required_: No

_Type_: [TaskInvocationParameters](aws-properties-ssm-maintenancewindowtask-taskinvocationparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TaskParameters`

###### Note

`TaskParameters` has been deprecated. To specify parameters to pass to
a task when it runs, instead use the `Parameters` option in the
`TaskInvocationParameters` structure. For information about how
Systems Manager handles these options for the supported maintenance window
task types, see [MaintenanceWindowTaskInvocationParameters](../../../../reference/systems-manager/latest/apireference/api-maintenancewindowtaskinvocationparameters.md).

The parameters to pass to the task when it runs.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TaskType`

The type of task. Valid values: `RUN_COMMAND`, `AUTOMATION`,
`LAMBDA`, `STEP_FUNCTIONS`.

_Required_: Yes

_Type_: String

_Allowed values_: `RUN_COMMAND | AUTOMATION | STEP_FUNCTIONS | LAMBDA`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`WindowId`

The ID of the maintenance window where the task is registered.

_Required_: Yes

_Type_: String

_Pattern_: `^mw-[0-9a-f]{17}$`

_Minimum_: `20`

_Maximum_: `20`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the maintenance window task ID, such as
`12a345b6-bbb7-4bb6-90b0-8c9577a2d2b9`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`WindowTaskId`

The task ID.

## Examples

- [Create a Run Command task that targets instances using a resource group name](#aws-resource-ssm-maintenancewindowtask--examples--Create_a_Run_Command_task_that_targets_instances_using_a_resource_group_name)

- [Create a Run Command task that targets instances using a maintenance window target ID](#aws-resource-ssm-maintenancewindowtask--examples--Create_a_Run_Command_task_that_targets_instances_using_a_maintenance_window_target_ID)

- [Create a Run Command task that runs a PowerShell script](#aws-resource-ssm-maintenancewindowtask--examples--Create_a_Run_Command_task_that_runs_a_PowerShell_script)

- [Create a task that runs an Automation runbook](#aws-resource-ssm-maintenancewindowtask--examples--Create_a_task_that_runs_an_Automation_runbook)

- [Create a Step Functions task that targets a maintenance window target ID](#aws-resource-ssm-maintenancewindowtask--examples--Create_a_Step_Functions_task_that_targets_a_maintenance_window_target_ID)

- [Create a Step Functions task that targets an instance ID](#aws-resource-ssm-maintenancewindowtask--examples--Create_a_Step_Functions_task_that_targets_an_instance_ID)

- [Create a task that runs an AWS Lambda function](#aws-resource-ssm-maintenancewindowtask--examples--Create_a_task_that_runs_an_function)

### Create a Run Command task that targets instances using a resource group name

The following example creates a maintenance window Run Command task that
installs patches on instances using a using a resource group name as the
target.

#### JSON

```json

{
    "Resources": {
        "PatchTask": {
            "Type": "AWS::SSM::MaintenanceWindowTask",
            "Properties": {
                "Description": "Apply OS patches on instances in target",
                "MaxConcurrency": 1,
                "MaxErrors": 1,
                "Priority": 0,
                "TaskType": "RUN_COMMAND",
                "WindowId": {
                    "Ref": "MaintenanceWindow"
                },
                "TaskArn": "AWS-RunPatchBaseline",
                "TaskInvocationParameters": {
                    "MaintenanceWindowRunCommandParameters": {
                        "Parameters": {
                            "Operation": [
                                "Install"
                            ],
                            "RebootOption": [
                                "NoReboot"
                            ]
                        }
                    }
                },
                "Targets": [
                    {
                        "Key": "WindowTargetIds",
                        "Values": [
                            {
                                "Ref": "MaintenanceWindowTarget"
                            }
                        ]
                    }
                ]
            }
        },
        "MaintenanceWindow": {
            "Type": "AWS::SSM::MaintenanceWindow",
            "Properties": {
                "Name": "MaintenanceWindow",
                "AllowUnassociatedTargets": true,
                "Cutoff": 0,
                "Description": "Maintenance window for instances",
                "Duration": 1,
                "Schedule": "cron(20 17 ? * MON-FRI *)"
            }
        },
        "MaintenanceWindowTarget": {
            "Type": "AWS::SSM::MaintenanceWindowTarget",
            "Properties": {
                "ResourceType": "RESOURCE_GROUP",
                "Targets": [
                    {
                        "Key": "resource-groups:Name",
                        "Values": [
                            "TestResourceGroup"
                        ]
                    }
                ],
                "WindowId": {
                    "Ref": "MaintenanceWindow"
                }
            }
        }
    }
}
```

#### YAML

```yaml

---
Resources:
  PatchTask:
    Type: AWS::SSM::MaintenanceWindowTask
    Properties:
      Description: Apply OS patches on instances in target
      MaxConcurrency: 1
      MaxErrors: 1
      Priority: 0
      TaskType: RUN_COMMAND
      WindowId:
        Ref: MaintenanceWindow
      TaskArn: AWS-RunPatchBaseline
      TaskInvocationParameters:
        MaintenanceWindowRunCommandParameters:
          Parameters:
            Operation:
            - Install
            RebootOption:
            - NoReboot
      Targets:
      - Key: WindowTargetIds
        Values:
        - Ref: MaintenanceWindowTarget
  MaintenanceWindow:
    Type: AWS::SSM::MaintenanceWindow
    Properties:
      Name: MaintenanceWindow
      AllowUnassociatedTargets: true
      Cutoff: 0
      Description: Maintenance window for instances
      Duration: 1
      Schedule: cron(20 17 ? * MON-FRI *)
  MaintenanceWindowTarget:
    Type: AWS::SSM::MaintenanceWindowTarget
    Properties:
      ResourceType: RESOURCE_GROUP
      Targets:
      - Key: resource-groups:Name
        Values:
        - TestResourceGroup
      WindowId:
        Ref: MaintenanceWindow

```

### Create a Run Command task that targets instances using a maintenance window target ID

The following example creates a maintenance window Run Command task that
installs patches on instances but does not reboot them. The maintenance window
task targets managed instances using a maintenance window target ID.

#### JSON

```json

{
    "Resources": {
        "MaintenanceWindowRunCommandTask": {
            "Type": "AWS::SSM::MaintenanceWindowTask",
            "Properties": {
                "WindowId": "MaintenanceWindow",
                "Targets": [
                    {
                        "Key": "WindowTargetIds",
                        "Values": [
                            "MaintenanceWindowTarget"
                        ]
                    }
                ],
                "TaskType": "RUN_COMMAND",
                "TaskArn": "AWS-RunPatchBaseline",
                "TaskInvocationParameters": {
                    "MaintenanceWindowRunCommandParameters": {
                        "Parameters": {
                            "Operation": [
                                "Install"
                            ],
                            "RebootOption": [
                                "NoReboot"
                            ]
                        }
                    },
                    "MaxConcurrency": 7,
                    "MaxErrors": 7,
                    "Priority": 5
                },
                "DependsOn": "MaintenanceWindowTarget"
            }
        }
    }
}
```

#### YAML

```yaml

---
Resources:
  MaintenanceWindowRunCommandTask:
    Type: AWS::SSM::MaintenanceWindowTask
    Properties:
      WindowId: MaintenanceWindow
      Targets:
      - Key: WindowTargetIds
        Values:
        - MaintenanceWindowTarget
      TaskType: RUN_COMMAND
      TaskArn: AWS-RunPatchBaseline
      TaskInvocationParameters:
        MaintenanceWindowRunCommandParameters:
          Parameters:
            Operation:
            - Install
            RebootOption:
            - NoReboot
      MaxConcurrency: 7
      MaxErrors: 7
      Priority: 5
      DependsOn: MaintenanceWindowTarget
```

### Create a Run Command task that runs a PowerShell script

The following example demonstrates running a command with
`AWS-RunPowerShellScript`.

#### JSON

```json

{
    "Resources": {
        "MaintenanceWindowRunCommandTask": {
            "Type": "AWS::SSM::MaintenanceWindowTask",
            "Properties": {
                "WindowId": {
                    "Ref": "MaintenanceWindow"
                },
                "Targets": [
                    {
                        "Key": "WindowTargetIds",
                        "Values": [
                            "MaintenanceWindowTarget"
                        ]
                    }
                ],
                "TaskType": "RUN_COMMAND",
                "TaskArn": "AWS-RunPowerShellScript",
                "TaskInvocationParameters": {
                    "MaintenanceWindowRunCommandParameters": {
                        "Comment": "This is a comment",
                        "CloudWatchOutputConfig": {
                            "CloudWatchLogGroupName": "MyLogGroupName",
                            "CloudWatchOutputEnabled": true
                        },
                        "Parameters": {
                            "executionTimeout": [
                                "3600"
                            ],
                            "commands": [
                                "Get-Service myImportantService | Restart-Service\nGet-ExecutionPolicy -List\nSet-ExecutionPolicy -Scope Process AllSigned\n"
                            ]
                        }
                    }
                },
                "MaxConcurrency": 7,
                "MaxErrors": 7,
                "Priority": 5
            },
            "DependsOn": "MaintenanceWindowTarget"
        }
    }
}
```

#### YAML

```yaml

---
Resources:
  MaintenanceWindowRunCommandTask:
    Type: 'AWS::SSM::MaintenanceWindowTask'
    Properties:
      WindowId: !Ref MaintenanceWindow
      Targets:
        - Key: WindowTargetIds
          Values:
            - MaintenanceWindowTarget
      TaskType: RUN_COMMAND
      TaskArn: AWS-RunPowerShellScript
      TaskInvocationParameters:
        MaintenanceWindowRunCommandParameters:
          Comment: This is a comment
          CloudWatchOutputConfig:
            CloudWatchLogGroupName: MyLogGroupName
            CloudWatchOutputEnabled: true
          Parameters:
            executionTimeout:
              - '3600'
            commands:
              - Get-Service myImportantService | Restart-Service
              - Get-ExecutionPolicy -List
              - Set-ExecutionPolicy -Scope Process AllSigned
      MaxConcurrency: 7
      MaxErrors: 7
      Priority: 5
    DependsOn: MaintenanceWindowTarget

```

### Create a task that runs an Automation runbook

The following example creates a Systems Manager maintenance window task
that uses the runbook `AWS-PatchInstanceWithRollback` to patch
instances.

#### JSON

```json

{
    "Resources": {
        "MaintenanceWindowStepFunctionsTask": {
            "Type": "AWS::SSM::MaintenanceWindowTask",
            "Properties": {
                "WindowId": "MaintenanceWindow",
                "Targets": [
                    {
                        "Key": "WindowTargetIds",
                        "Values": [
                            "MaintenanceWindowTarget"
                        ]
                    }
                ],
                "TaskArn": "AWS-PatchInstanceWithRollback",
                "ServiceRoleArn": "arn:aws:iam::111222333444:role/MyMaintenanceWindowServiceRole",
                "TaskType": "AUTOMATION",
                "TaskInvocationParameters": {
                    "MaintenanceWindowAutomationParameters": {
                        "DocumentVersion": "1",
                        "Parameters": {
                            "InstanceId": [
                                "{{RESOURCE_ID}}"
                            ]
                        }
                    }
                },
                "Priority": 1,
                "MaxConcurrency": 5,
                "MaxErrors": 5,
                "Name": "AutomationTask"
            },
            "DependsOn": "MaintenanceWindowTarget"
        }
    }
}
```

#### YAML

```yaml

---
Resources:
  MaintenanceWindowStepFunctionsTask:
    Type: AWS::SSM::MaintenanceWindowTask
    Properties:
      WindowId: MaintenanceWindow
      Targets:
      - Key: WindowTargetIds
        Values:
        - MaintenanceWindowTarget
      TaskArn: AWS-PatchInstanceWithRollback
      ServiceRoleArn: arn:aws:iam::111222333444:role/MyMaintenanceWindowServiceRole
      TaskType: AUTOMATION
      TaskInvocationParameters:
        MaintenanceWindowAutomationParameters:
          DocumentVersion: 1
          Parameters:
            InstanceId:
              - '{{RESOURCE_ID}}'
      Priority: 1
      MaxConcurrency: 5
      MaxErrors: 5
      Name: AutomationTask
    DependsOn: MaintenanceWindowTarget
```

### Create a Step Functions task that targets a maintenance window target ID

The following example creates a Systems Manager maintenance window task
that runs the specified Step Function. The maintenance window task targets
managed instances using a maintenance window target ID.

#### JSON

```json

{
    "Resources": {
        "MaintenanceWindowStepFunctionsTask": {
            "Type": "AWS::SSM::MaintenanceWindowTask",
            "Properties": {
                "WindowId": "MaintenanceWindow",
                "Targets": [
                    {
                        "Key": "WindowTargetIds",
                        "Values": [
                            "MaintenanceWindowTarget"
                        ]
                    }
                ],
                "TaskArn": "SSMStepFunctionDemo",
                "ServiceRoleArn": "StepFunctionRole.Arn",
                "TaskType": "STEP_FUNCTIONS",
                "TaskInvocationParameters": {
                    "MaintenanceWindowStepFunctionsParameters": {
                        "Input": "{\"instanceId\":\"{{TARGET_ID}}\", \"wait_time\": 20}",
                        "Name": "{{INVOCATION_ID}}"
                    }
                },
                "Priority": 1,
                "MaxConcurrency": 5,
                "MaxErrors": 5,
                "Name": "StepFunctionsTask"
            },
            "DependsOn": "MaintenanceWindowTarget"
        }
    }
}
```

#### YAML

```yaml

---
Resources:
  MaintenanceWindowStepFunctionsTask:
    Type: AWS::SSM::MaintenanceWindowTask
    Properties:
      WindowId: MaintenanceWindow
      Targets:
      - Key: WindowTargetIds
        Values:
        - MaintenanceWindowTarget
      TaskArn: SSMStepFunctionDemo
      ServiceRoleArn: StepFunctionRole.Arn
      TaskType: STEP_FUNCTIONS
      TaskInvocationParameters:
        MaintenanceWindowStepFunctionsParameters:
          Input: '{"instanceId":"{{TARGET_ID}}", "wait_time": 20}'
          Name: "{{INVOCATION_ID}}"
      Priority: 1
      MaxConcurrency: 5
      MaxErrors: 5
      Name: StepFunctionsTask
   DependsOn: MaintenanceWindowTarget
```

### Create a Step Functions task that targets an instance ID

The following example creates a Systems Manager maintenance window task
that runs the specified Step Function. The maintenance window task targets the
specified instance IDs.

#### JSON

```json

{
    "Resources": {
        "StepFunctionsTask": {
            "Type": "AWS::SSM::MaintenanceWindowTask",
            "Properties": {
                "WindowId": "MaintenanceWindow",
                "Targets": [
                    {
                        "Key": "InstanceIds",
                        "Values": [
                            "i-012345678912345678"
                        ]
                    }
                ],
                "TaskArn": "SSMStepFunctionDemo",
                "ServiceRoleArn": "StepFunctionRole.Arn",
                "TaskType": "STEP_FUNCTIONS",
                "TaskInvocationParameters": {
                    "MaintenanceWindowStepFunctionsParameters": {
                        "Input": "{\"instanceId\":\"{{TARGET_ID}}\", \"wait_time\": 20}",
                        "Name": "{{INVOCATION_ID}}"
                    }
                },
                "Priority": 1,
                "MaxConcurrency": 5,
                "MaxErrors": 5,
                "Name": "StepFunctionsTask"
            },
            "DependsOn": "MaintenanceWindowTarget"
        }
    }
}
```

#### YAML

```yaml

---
Resources:
  StepFunctionsTask:
    Type: 'AWS::SSM::MaintenanceWindowTask'
    Properties:
      WindowId: MaintenanceWindow
      Targets:
      - Key: InstanceIds
        Values:
        - i-012345678912345678
      TaskArn: SSMStepFunctionDemo
      ServiceRoleArn: StepFunctionRole.Arn
      TaskType: STEP_FUNCTIONS
      TaskInvocationParameters:
        MaintenanceWindowStepFunctionsParameters:
          Input: '{"instanceId":"{{TARGET_ID}}", "wait_time": 20}'
          Name: "{{INVOCATION_ID}}"
      Priority: 1
      MaxConcurrency: 5
      MaxErrors: 5
      Name: StepFunctionsTask
    DependsOn: MaintenanceWindowTarget
```

### Create a task that runs an AWS Lambda function

The following example runs an AWS Lambda function to restart
instances.

###### Note

The value for `Payload` in
`MaintenanceWindowLambdaParameters` must be formatted as a
Base64-encoded binary data object.

#### JSON

```json

{
   "Resources": {
      "LambdaTask": {
         "Type": "AWS::SSM::MaintenanceWindowTask",
         "Properties": {
            "WindowId": "mw-04fd6f19dfEXAMPLE",
            "TaskArn": "arn:aws:lambda:us-east-2:111222333444:function:MyLambdaTaskArn",
            "ServiceRoleArn": "arn:aws:iam::111222333444:role/aws-service-role/ssm.amazonaws.com/AWSServiceRoleForAmazonSSM",
            "TaskType": "LAMBDA",
            "TaskInvocationParameters": {
               "MaintenanceWindowLambdaParameters": {
                  "ClientContext": "eyJ0ZXN0Q29udGV4dCI6Ik5vdGhp==trucated==EXAMPLE",
                  "Qualifier": "$LATEST",
                  "Payload": "eyJJbnN0YW5jZUlkIjoie3tSRVNPVVJDRV9JRH19IiwidGFyZ2V0VHlwZSI6Int7VEFSR0VUX1RZUEV9fSJ9"
               }
            },
            "Priority": 1,
            "Name": "UpdateLambdaTaskEXAMPLE"
         }
      }
   }
}
```

#### YAML

```yaml

---
Resources:
  LambdaTask:
    Type: 'AWS::SSM::MaintenanceWindowTask'
    Properties:
      WindowId: mw-04fd6f19dfEXAMPLE
      TaskArn: >-
        arn:aws:lambda:us-east-2:111222333444:function:MyLambdaTaskArn
      ServiceRoleArn: >-
        arn:aws:iam::111222333444:role/aws-service-role/ssm.amazonaws.com/AWSServiceRoleForAmazonSSM
      TaskType: LAMBDA
      TaskInvocationParameters:
        MaintenanceWindowLambdaParameters:
          ClientContext: eyJ0ZXN0Q29udGV4dCI6Ik5vdGhp==trucated==EXAMPLE
          Qualifier: $LATEST
          Payload: >-
            eyJJbnN0YW5jZUlkIjoie3tSRVNPVVJDRV9JRH19IiwidGFyZ2V0VHlwZSI6Int7VEFSR0VUX1RZUEV9fSJ9
      Priority: 1
      Name: UpdateLambdaTaskEXAMPLE
```

## See also

- [AWS::SSM::MaintenanceWindow](../userguide/aws-resource-ssm-maintenancewindow.md)

- [AWS::SSM::MaintenanceWindowTarget](../userguide/aws-resource-ssm-maintenancewindowtarget.md)

- [RegisterTaskWithMaintenanceWindow](../../../../reference/systems-manager/latest/apireference/api-registertaskwithmaintenancewindow.md) in the _AWS Systems Manager API Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Targets

CloudWatchOutputConfig

All content copied from https://docs.aws.amazon.com/.
