This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Job

The `AWS::Glue::Job` resource specifies an AWS Glue job in the data
catalog. For more information, see [Adding Jobs in AWS Glue](../../../glue/latest/dg/add-job.md) and [Job\
Structure](../../../glue/latest/dg/aws-glue-api-jobs-job.md#aws-glue-api-jobs-job-Job) in the _AWS Glue Developer Guide._

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Glue::Job",
  "Properties" : {
      "AllocatedCapacity" : Number,
      "Command" : JobCommand,
      "Connections" : ConnectionsList,
      "DefaultArguments" : Json,
      "Description" : String,
      "ExecutionClass" : String,
      "ExecutionProperty" : ExecutionProperty,
      "GlueVersion" : String,
      "JobMode" : String,
      "JobRunQueuingEnabled" : Boolean,
      "LogUri" : String,
      "MaintenanceWindow" : String,
      "MaxCapacity" : Number,
      "MaxRetries" : Number,
      "Name" : String,
      "NonOverridableArguments" : Json,
      "NotificationProperty" : NotificationProperty,
      "NumberOfWorkers" : Integer,
      "Role" : String,
      "SecurityConfiguration" : String,
      "Tags" : [ Tag, ... ],
      "Timeout" : Integer,
      "WorkerType" : String
    }
}

```

### YAML

```yaml

Type: AWS::Glue::Job
Properties:
  AllocatedCapacity: Number
  Command:
    JobCommand
  Connections:
    ConnectionsList
  DefaultArguments: Json
  Description: String
  ExecutionClass: String
  ExecutionProperty:
    ExecutionProperty
  GlueVersion: String
  JobMode: String
  JobRunQueuingEnabled: Boolean
  LogUri: String
  MaintenanceWindow: String
  MaxCapacity: Number
  MaxRetries: Number
  Name: String
  NonOverridableArguments: Json
  NotificationProperty:
    NotificationProperty
  NumberOfWorkers: Integer
  Role: String
  SecurityConfiguration: String
  Tags:
    - Tag
  Timeout: Integer
  WorkerType: String

```

## Properties

`AllocatedCapacity`

This parameter is no longer supported. Use `MaxCapacity` instead.

The number of capacity units that are allocated to this job.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Command`

The code that executes a job.

_Required_: Yes

_Type_: [JobCommand](aws-properties-glue-job-jobcommand.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Connections`

The connections used for this job.

_Required_: No

_Type_: [ConnectionsList](aws-properties-glue-job-connectionslist.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultArguments`

The default arguments for this job, specified as name-value pairs.

You can specify arguments here that your own job-execution script consumes, in
addition to arguments that AWS Glue itself consumes.

For information about how to specify and consume your own job arguments, see [Calling AWS Glue APIs in Python](../../../glue/latest/dg/aws-glue-programming-python-calling.md) in the _AWS Glue Developer_
_Guide_.

For information about the key-value pairs that AWS Glue consumes to set up your job,
see [Special Parameters\
Used by AWS Glue](../../../glue/latest/dg/aws-glue-programming-etl-glue-arguments.md) in the _AWS Glue Developer_
_Guide_.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the job.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExecutionClass`

Indicates whether the job is run with a standard or flexible execution class. The standard execution class is ideal for time-sensitive workloads that require fast job startup and dedicated resources.

The flexible execution class is appropriate for time-insensitive jobs whose start and completion times may vary.

Only jobs with AWS Glue version 3.0 and above and command type `glueetl` will be allowed to set `ExecutionClass` to `FLEX`. The flexible execution class is available for Spark jobs.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExecutionProperty`

The maximum number of concurrent runs that are allowed for this job.

_Required_: No

_Type_: [ExecutionProperty](aws-properties-glue-job-executionproperty.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GlueVersion`

Glue version determines the versions of Apache Spark and Python that AWS Glue supports. The Python version indicates the version supported for jobs of type Spark.

For more information about the available AWS Glue versions and corresponding Spark and Python versions, see [Glue version](../../../glue/latest/dg/add-job.md) in the developer guide.

Jobs that are created without specifying a Glue version default to the latest Glue version available.

_Required_: No

_Type_: String

_Pattern_: `^(\w+\.)+\w+$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JobMode`

A mode that describes how a job was created. Valid values are:

- `SCRIPT` \- The job was created using the AWS Glue Studio script editor.

- `VISUAL` \- The job was created using the AWS Glue Studio visual editor.

- `NOTEBOOK` \- The job was created using an interactive sessions notebook.

When the `JobMode` field is missing or null, `SCRIPT` is assigned as the default value.

_Required_: No

_Type_: String

_Allowed values_: `SCRIPT | VISUAL | NOTEBOOK`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JobRunQueuingEnabled`

Specifies whether job run queuing is enabled for the job runs for this job.

A value of true means job run queuing is enabled for the job runs. If false or not populated, the job runs will not be considered for queueing.

If this field does not match the value set in the job run, then the value from the job run field will be used.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogUri`

This field is reserved for future use.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaintenanceWindow`

This field specifies a day of the week and hour for a maintenance window for streaming jobs. AWS Glue periodically performs maintenance activities. During these maintenance windows, AWS Glue will need to restart your streaming jobs.

AWS Glue will restart the job within 3 hours of the specified maintenance window. For instance, if you set up the maintenance window for Monday at 10:00AM GMT, your jobs will be restarted between 10:00AM GMT to 1:00PM GMT.

_Required_: No

_Type_: String

_Pattern_: `^(Sun|Mon|Tue|Wed|Thu|Fri|Sat):([01]?[0-9]|2[0-3])$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxCapacity`

The number of AWS Glue data processing units (DPUs) that can be allocated when this job runs. A DPU is a relative measure
of processing power that consists of 4 vCPUs of compute capacity and 16 GB of memory.

Do not set `Max Capacity` if using `WorkerType` and `NumberOfWorkers`.

The value that can be allocated for `MaxCapacity` depends on whether you are
running a Python shell job or an Apache Spark ETL job:

- When you specify a Python shell job ( `JobCommand.Name` ="pythonshell"), you can
allocate either 0.0625 or 1 DPU. The default is 0.0625 DPU.

- When you specify an Apache Spark ETL job ( `JobCommand.Name` ="glueetl"), you can allocate from 2 to 100 DPUs. The default is 10 DPUs. This job type cannot have a fractional DPU allocation.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxRetries`

The maximum number of times to retry this job after a JobRun fails.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name you assign to this job definition.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NonOverridableArguments`

Non-overridable arguments for this job, specified as name-value pairs.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotificationProperty`

Specifies configuration properties of a notification.

_Required_: No

_Type_: [NotificationProperty](aws-properties-glue-job-notificationproperty.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumberOfWorkers`

The number of workers of a defined `workerType` that are allocated when a job runs.

The maximum number of workers you can define are 299 for `G.1X`, and 149 for `G.2X`.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Role`

The name or Amazon Resource Name (ARN) of the IAM role associated with this
job.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityConfiguration`

The name of the `SecurityConfiguration` structure to be used with this
job.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to use with this job.

_Required_: No

_Type_: Array of [`Tag`](aws-properties-resource-tags.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Timeout`

The job timeout in minutes. This is the maximum time that a job run can consume resources before it is terminated and enters TIMEOUT status. The default is 2,880 minutes (48 hours).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkerType`

The type of predefined worker that is allocated when a job runs.

AWS Glue provides multiple worker types to accommodate different workload requirements:

G Worker Types (General-purpose compute workers):

- G.1X: 1 DPU (4 vCPUs, 16 GB memory, 94GB disk)

- G.2X: 2 DPU (8 vCPUs, 32 GB memory, 138GB disk)

- G.4X: 4 DPU (16 vCPUs, 64 GB memory, 256GB disk)

- G.8X: 8 DPU (32 vCPUs, 128 GB memory, 512GB disk)

- G.12X: 12 DPU (48 vCPUs, 192 GB memory, 768GB disk)

- G.16X: 16 DPU (64 vCPUs, 256 GB memory, 1024GB disk)

R Worker Types (Memory-optimized workers):

- R.1X: 1 M-DPU (4 vCPUs, 32 GB memory)

- R.2X: 2 M-DPU (8 vCPUs, 64 GB memory)

- R.4X: 4 M-DPU (16 vCPUs, 128 GB memory)

- R.8X: 8 M-DPU (32 vCPUs, 256 GB memory)

_Required_: No

_Type_: String

_Allowed values_: `Standard | G.1X | G.2X | G.025X | G.4X | G.8X | Z.2X | G.12X | G.16X | R.1X | R.2X | R.4X | R.8X`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the job name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

The following example creates a job with an associated role. The ScriptLocation is an Amazon S3 location. The example
provided below is a placeholder for your Amazon S3 location.

#### JSON

```json

{
  "Description": "AWS Glue Job Test",
  "Resources": {
    "MyJobRole": {
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
        "Command": {
          "Name": "glueetl",
          "ScriptLocation": "s3://<your-S3-script-uri>"
        },
        "DefaultArguments": {
          "--job-bookmark-option": "job-bookmark-enable"
        },
        "ExecutionProperty": {
          "MaxConcurrentRuns": 2
        },
        "MaxRetries": 0,
        "Name": "cf-job1",
        "Role": {
          "Ref": "MyJobRole"
        }
      }
    }
  }
}
```

#### YAML

```yaml

---
Description: "AWS Glue Job Test"
Resources:
  MyJobRole:
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
      Command:
        Name: glueetl
        ScriptLocation: "s3://<your-S3-script-uri>"
      DefaultArguments:
        "--job-bookmark-option": "job-bookmark-enable"
      ExecutionProperty:
        MaxConcurrentRuns: 2
      MaxRetries: 0
      Name: cf-job1
      Role: !Ref MyJobRole
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TargetProcessingProperties

ConnectionsList

All content copied from https://docs.aws.amazon.com/.
