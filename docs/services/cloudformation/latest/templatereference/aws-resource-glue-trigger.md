This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Trigger

The `AWS::Glue::Trigger` resource specifies triggers that run AWS Glue
jobs. For more information, see [Triggering Jobs in AWS Glue](https://docs.aws.amazon.com/glue/latest/dg/trigger-job.html) and [Trigger Structure](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-jobs-trigger.html#aws-glue-api-jobs-trigger-Trigger) in the _AWS Glue Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Glue::Trigger",
  "Properties" : {
      "Actions" : [ Action, ... ],
      "Description" : String,
      "EventBatchingCondition" : EventBatchingCondition,
      "Name" : String,
      "Predicate" : Predicate,
      "Schedule" : String,
      "StartOnCreation" : Boolean,
      "Tags" : [ Tag, ... ],
      "Type" : String,
      "WorkflowName" : String
    }
}

```

### YAML

```yaml

Type: AWS::Glue::Trigger
Properties:
  Actions:
    - Action
  Description: String
  EventBatchingCondition:
    EventBatchingCondition
  Name: String
  Predicate:
    Predicate
  Schedule: String
  StartOnCreation: Boolean
  Tags:
    - Tag
  Type: String
  WorkflowName: String

```

## Properties

`Actions`

The actions initiated by this trigger.

_Required_: Yes

_Type_: Array of [Action](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-glue-trigger-action.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of this trigger.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventBatchingCondition`

Batch condition that must be met (specified number of events received or batch time window expired) before EventBridge event trigger fires.

_Required_: No

_Type_: [EventBatchingCondition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-glue-trigger-eventbatchingcondition.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the trigger.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Predicate`

The predicate of this trigger, which defines when it will fire.

_Required_: No

_Type_: [Predicate](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-glue-trigger-predicate.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Schedule`

A `cron` expression used to specify the schedule. For more information, see
[Time-Based Schedules for\
Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html) in the _AWS Glue Developer Guide_. For
example, to run something every day at 12:15 UTC, specify `cron(15 12 * * ?
                *)`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartOnCreation`

Set to true to start `SCHEDULED` and `CONDITIONAL` triggers when created. True is not supported for `ON_DEMAND` triggers.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to use with this trigger.

_Required_: No

_Type_: Array of [`Tag`](aws-properties-resource-tags.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of trigger that this is.

_Required_: Yes

_Type_: String

_Allowed values_: `SCHEDULED | CONDITIONAL | ON_DEMAND | EVENT`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`WorkflowName`

The name of the workflow associated with the trigger.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the trigger name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

- [On-Demand Trigger](#aws-resource-glue-trigger--examples--On-Demand_Trigger)

- [Scheduled Trigger](#aws-resource-glue-trigger--examples--Scheduled_Trigger)

- [Conditional Trigger](#aws-resource-glue-trigger--examples--Conditional_Trigger)

### On-Demand Trigger

The following example creates an on-demand trigger that triggers one
job.

#### JSON

```json

{
  "Resources": {
    "OnDemandJobTrigger": {
      "Type": "AWS::Glue::Trigger",
      "Properties": {
        "Type": "ON_DEMAND",
        "Description": "DESCRIPTION_ON_DEMAND",
        "Actions": [
          {
            "JobName": "prod-job2"
          }
        ],
        "Name": "prod-trigger1-ondemand"
      }
    }
  }
}
```

#### YAML

```yaml

Resources:
  OnDemandJobTrigger:
    Type: AWS::Glue::Trigger
    Properties:
      Type: ON_DEMAND
      Description: DESCRIPTION_ON_DEMAND
      Actions:
        - JobName: prod-job2
      Name: prod-trigger1-ondemand
```

### Scheduled Trigger

The following example creates a scheduled trigger that runs every two hours
and triggers two jobs. Note that it declares an argument for
`prod-job3`.

#### JSON

```json

{
  "Resources": {
    "ScheduledJobTrigger": {
      "Type": "AWS::Glue::Trigger",
      "Properties": {
        "Type": "SCHEDULED",
        "Description": "DESCRIPTION_SCHEDULED",
        "Schedule": "cron(0 */2 * * ? *)",
        "Actions": [
          {
            "JobName": "prod-job2"
          },
          {
            "JobName": "prod-job3",
            "Arguments": {
              "--job-bookmark-option": "job-bookmark-enable"
            }
          }
        ],
        "Name": "prod-trigger1-scheduled"
      }
    }
  }
}
```

#### YAML

```yaml

Resources:
  ScheduledJobTrigger:
    Type: AWS::Glue::Trigger
    Properties:
      Type: SCHEDULED
      Description: DESCRIPTION_SCHEDULED
      Schedule: cron(0 */2 * * ? *)
      Actions:
        - JobName: prod-job2
        - JobName: prod-job3
          Arguments:
            '--job-bookmark-option': job-bookmark-enable
      Name: prod-trigger1-scheduled
```

### Conditional Trigger

The following example creates a conditional trigger that starts a job based on
the successful completion of the job run.

#### JSON

```json

{
  "Description": "AWS Glue trigger test",
  "Resources": {
    "MyJobTriggerRole": {
      "Type": "AWS::IAM::Role",
      "Properties": {
        "AssumeRolePolicyDocument": {
          "Version": "2012-10-17",
          "Statement": [
            {
              "Effect": "Allow",
              "Principal": {
                "Service": [
                  "glue.amazonaws.com"
                ]
              },
              "Action": [
                "sts:AssumeRole"
              ]
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
                  "Action": "*",
                  "Resource": "*"
                }
              ]
            }
          }
        ]
      }
    },
    "MyJob": {
      "Type": "AWS::Glue::Job",
      "Properties": {
        "Name": "MyJobTriggerJob",
        "LogUri": "wikiData",
        "Role": {
          "Ref": "MyJobTriggerRole"
        },
        "Command": {
          "Name": "glueetl",
          "ScriptLocation": "s3://testdata-bucket/s3-target/create-delete-job-xtf-ETL-s3-json-to-csv.py"
        },
        "DefaultArguments": {
          "--job-bookmark-option": "job-bookmark-enable"
        },
        "MaxRetries": 0
      }
    },
    "MyJobTrigger": {
      "Type": "AWS::Glue::Trigger",
      "Properties": {
        "Name": "MyJobTrigger",
        "Type": "CONDITIONAL",
        "Description": "Description for a conditional job trigger",
        "Actions": [
          {
            "JobName": {
              "Ref": "MyJob"
            },
            "Arguments": {
              "--job-bookmark-option": "job-bookmark-enable"
            }
          }
        ],
        "Predicate": {
          "Conditions": [
            {
              "LogicalOperator": "EQUALS",
              "JobName": {
                "Ref": "MyJob"
              },
              "State": "SUCCEEDED"
            }
          ]
        }
      }
    }
  }
}
```

#### YAML

```yaml

---
Description: "AWS Glue trigger test"
Resources:
  MyJobTriggerRole:
    Type: AWS::IAM::Role
    Properties:
      AssumeRolePolicyDocument:
        Version: "2012-10-17"
        Statement:
          -
            Effect: "Allow"
            Principal:
              Service:
                - "glue.amazonaws.com"
            Action:
              - "sts:AssumeRole"
      Path: "/"
      Policies:
        -
          PolicyName: "root"
          PolicyDocument:
            Version: "2012-10-17"
            Statement:
              -
                Effect: "Allow"
                Action: "*"
                Resource: "*"
  MyJob:
    Type: AWS::Glue::Job
    Properties:
      Name: "MyJobTriggerJob"
      LogUri: "wikiData"
      Role: !Ref MyJobTriggerRole
      Command:
        Name: "glueetl"
        ScriptLocation: "s3://testdata-bucket/s3-target/create-delete-job-xtf-ETL-s3-json-to-csv.py"
      DefaultArguments:
        "--job-bookmark-option": "job-bookmark-enable"
      MaxRetries: 0
  MyJobTrigger:
    Type: AWS::Glue::Trigger
    Properties:
      Name: "MyJobTrigger"
      Type: "CONDITIONAL"
      Description: "Description for a conditional job trigger"
      Actions:
        - JobName: !Ref MyJob
          Arguments:
            "--job-bookmark-option": "job-bookmark-enable"
      Predicate:
        Conditions:
          - LogicalOperator: EQUALS
            JobName: !Ref MyJob
            State: SUCCEEDED
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VpcConfiguration

Action
