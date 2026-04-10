This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Schedule

Specifies a new schedule for one or more AWS Glue DataBrew jobs. Jobs can be
run at a specific date and time, or at regular intervals.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataBrew::Schedule",
  "Properties" : {
      "CronExpression" : String,
      "JobNames" : [ String, ... ],
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DataBrew::Schedule
Properties:
  CronExpression: String
  JobNames:
    - String
  Name: String
  Tags:
    - Tag

```

## Properties

`CronExpression`

The dates and times when the job is to run. For more information, see [Working with cron\
expressions for recipe jobs](../../../databrew/latest/dg/jobs-recipe.md#jobs.cron) in the _AWS Glue DataBrew Developer_
_Guide_.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JobNames`

A list of jobs to be run, according to the schedule.

_Required_: No

_Type_: Array of String

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the schedule.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Metadata tags that have been applied to the schedule.

_Required_: No

_Type_: Array of [Tag](aws-properties-databrew-schedule-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref`
function, `Ref` returns the resource name. For example:

`{ "Ref": "mySchedule" }`

For an AWS Glue DataBrew schedule named `mySchedule`,
`Ref` returns the name of the schedule.

## Examples

### Creating schedules

The following examples create new DataBrew schedules.

#### YAML

```yaml

Resources:
  TestDataBrewSchedule:
    Type: AWS::DataBrew::Schedule
    Properties:
      JobNames: ["job-name"]
      Name: schedule-name
      CronExpression: "cron(0 0/1 ? * * *)"
      Tags: [{Key: key00AtCreate, Value: value001AtCreate}]

```

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "This CloudFormation template specifies a DataBrew Schedule",
    "Resources": {
        "MyDataBrewSchedule": {
            "Type": "AWS::DataBrew::Schedule",
            "Properties": {
                "JobNames": ["job-test"],
                "Name": "cf-test-schedule1",
                "CronExpression": "cron(0 0/1 ? * * *)"
            },
            "Tags": [
                    {
                        "Key": "key00AtCreate",
                        "Value": "value001AtCreate"
                    }
                ]
        }
    }
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Threshold

Tag

All content copied from https://docs.aws.amazon.com/.
