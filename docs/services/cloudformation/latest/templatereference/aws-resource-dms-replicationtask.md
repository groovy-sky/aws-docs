This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DMS::ReplicationTask

The `AWS::DMS::ReplicationTask` resource creates an AWS DMS replication task.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DMS::ReplicationTask",
  "Properties" : {
      "CdcStartPosition" : String,
      "CdcStartTime" : Number,
      "CdcStopPosition" : String,
      "MigrationType" : String,
      "ReplicationInstanceArn" : String,
      "ReplicationTaskIdentifier" : String,
      "ReplicationTaskSettings" : String,
      "ResourceIdentifier" : String,
      "SourceEndpointArn" : String,
      "TableMappings" : String,
      "Tags" : [ Tag, ... ],
      "TargetEndpointArn" : String,
      "TaskData" : String
    }
}

```

### YAML

```yaml

Type: AWS::DMS::ReplicationTask
Properties:
  CdcStartPosition: String
  CdcStartTime: Number
  CdcStopPosition: String
  MigrationType: String
  ReplicationInstanceArn: String
  ReplicationTaskIdentifier: String
  ReplicationTaskSettings: String
  ResourceIdentifier: String
  SourceEndpointArn: String
  TableMappings: String
  Tags:
    - Tag
  TargetEndpointArn: String
  TaskData: String

```

## Properties

`CdcStartPosition`

Indicates when you want a change data capture (CDC) operation to start. Use either
`CdcStartPosition` or `CdcStartTime` to specify when you want a CDC operation to start.
Specifying both values results in an error.

The value can be in date, checkpoint, log sequence number (LSN), or system change
number (SCN) format.

Here is a date example: `--cdc-start-position "2018-03-08T12:12:12"`

Here is a checkpoint example: `--cdc-start-position "checkpoint:V1#27#mysql-bin-changelog.157832:1975:-1:2002:677883278264080:mysql-bin-changelog.157832:1876#0#0#*#0#93"`

Here is an LSN example: `--cdc-start-position “mysql-bin-changelog.000024:373”`

###### Note

When you use this task setting with a source PostgreSQL database, a logical
replication slot should already be created and associated with the source endpoint.
You can verify this by setting the `slotName` extra connection attribute to the
name of this logical replication slot. For more information, see
[Extra Connection Attributes When Using PostgreSQL as a Source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.PostgreSQL.html#CHAP_Source.PostgreSQL.ConnectionAttrib)
in the _AWS Database Migration Service User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CdcStartTime`

Indicates the start time for a change data capture (CDC) operation.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CdcStopPosition`

Indicates when you want a change data capture (CDC) operation to stop. The value can be
either server time or commit time.

Here is a server time example: `--cdc-stop-position
            "server_time:2018-02-09T12:12:12"`

Here is a commit time example: `--cdc-stop-position "commit_time:
            2018-02-09T12:12:12"`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MigrationType`

The migration type. Valid values: `full-load` \| `cdc` \| `full-load-and-cdc`

_Required_: Yes

_Type_: String

_Allowed values_: `full-load | cdc | full-load-and-cdc`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplicationInstanceArn`

The Amazon Resource Name (ARN) of a replication instance.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ReplicationTaskIdentifier`

An identifier for the replication task.

Constraints:

- Must contain 1-255 alphanumeric characters or hyphens.

- First character must be a letter.

- Cannot end with a hyphen or contain two consecutive hyphens.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplicationTaskSettings`

Overall settings for the task, in JSON format. For more information, see
[Specifying Task Settings for AWS Database Migration Service Tasks](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html)
in the _AWS Database Migration Service User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceIdentifier`

A display name for the resource identifier at the end of the `EndpointArn`
response parameter that is returned in the created `Endpoint` object. The
value for this parameter can have up to 31 characters. It can contain only ASCII
letters, digits, and hyphen ('-'). Also, it can't end with a hyphen or contain two
consecutive hyphens, and can only begin with a letter, such as
`Example-App-ARN1`.

For example, this value might result in the `EndpointArn` value
`arn:aws:dms:eu-west-1:012345678901:rep:Example-App-ARN1`. If you don't
specify a `ResourceIdentifier` value, AWS DMS generates a default
identifier value for the end of `EndpointArn`.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceEndpointArn`

An Amazon Resource Name (ARN) that uniquely identifies the source endpoint.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TableMappings`

The table mappings for the task, in JSON format. For more information, see
[Using Table Mapping to Specify Task Settings](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
in the _AWS Database Migration Service User Guide_.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

One or more tags to be assigned to the replication task.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dms-replicationtask-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetEndpointArn`

An Amazon Resource Name (ARN) that uniquely identifies the target endpoint.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TaskData`

Supplemental information that the task requires to migrate the data for certain source and target endpoints.
For more information, see [Specifying Supplemental Data for Task Settings](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.TaskData.html) in the
_AWS Database Migration Service User Guide._

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the replication task ARN.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

## Examples

- [Create a replication task](#aws-resource-dms-replicationtask--examples--Create_a_replication_task)

- [Create a replication task with full load task settings](#aws-resource-dms-replicationtask--examples--Create_a_replication_task_with_full_load_task_settings)

### Create a replication task

The following example creates a replication task with a selection rule and action.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "myReplicationTask": {
            "Type": "AWS::DMS::ReplicationTask",
            "Properties": {
                "SourceEndpointArn": 11,
                "TargetEndpointArn": "12ff",
                "ReplicationInstanceArn": "ert1",
                "MigrationType": "full-load",
                "TableMappings": "{ \"rules\": [ { \"rule-type\": \"selection\", \"rule-id\": \"1\", \"rule-name\": \"1\", \"object-locator\": { \"schema-name\": \"%\", \"table-name\": \"%\" }, \"rule-action\": \"include\" } ] }"
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  myReplicationTask:
    Type: AWS::DMS::ReplicationTask
    Properties:
      MigrationType: full-load
      ReplicationInstanceArn: ReplicationInstance
      SourceEndpointArn: SourceEndpoint
      TableMappings: "{ \"rules\": [ { \"rule-type\": \"selection\", \"rule-id\": \"1\", \"rule-name\": \"1\", \"object-locator\": { \"schema-name\": \"%\", \"table-name\": \"%\" }, \"rule-action\": \"include\" } ] }"
      TargetEndpointArn: TargetEndpoint

```

### Create a replication task with full load task settings

The following example creates a replication task with a selection rule and full-load task settings. For more information about
full-load task settings, see [Full-load task settings](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.TimeTravel.html).

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "myReplicationTask": {
            "Type": "AWS::DMS::ReplicationTask",
            "Properties": {
                "SourceEndpointArn": "arn:aws:dms:eu-west-1:123456789012:endpoint:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
                "TargetEndpointArn": "arn:aws:dms:eu-west-1:123456789012:endpoint:BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB",
                "ReplicationInstanceArn": "arn:aws:dms:eu-west-1:123456789012:rep:CCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC",
                "MigrationType": "full-load",
                "TableMappings": "{ \"rules\": [ { \"rule-type\": \"selection\", \"rule-id\": \"1\", \"rule-name\": \"1\", \"object-locator\": { \"schema-name\": \"%\", \"table-name\": \"%\" }, \"rule-action\": \"include\" } ] }",
                "ReplicationTaskSettings": "{ \"FullLoadSettings\": { \"CommitRate\": 9000, \"TargetTablePrepMode\": \"DROP_AND_CREATE\", \"CreatePkAfterFullLoad\": true }, \"BeforeImageSettings\": null, \"ControlTablesSettings\": { \"historyTimeslotInMinutes\": 5, \"HistoryTimeslotInMinutes\": 5}}"
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  myReplicationTask:
    Type: AWS::DMS::ReplicationTask
    Properties:
      MigrationType: full-load
      ReplicationInstanceArn: arn:aws:dms:eu-west-1:123456789012:rep:CCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC
      SourceEndpointArn: arn:aws:dms:eu-west-1:123456789012:endpoint:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
      TableMappings: "{ \"rules\": [ { \"rule-type\": \"selection\", \"rule-id\": \"1\", \"rule-name\": \"1\", \"object-locator\": { \"schema-name\": \"%\", \"table-name\": \"%\" }, \"rule-action\": \"include\" } ] }"
      ReplicationTaskSettings: "{ \"FullLoadSettings\": { \"CommitRate\": 9000, \"TargetTablePrepMode\": \"DROP_AND_CREATE\", \"CreatePkAfterFullLoad\": true }, \"BeforeImageSettings\": null, \"ControlTablesSettings\": { \"historyTimeslotInMinutes\": 5, \"HistoryTimeslotInMinutes\": 5}}"
      TargetEndpointArn: arn:aws:dms:eu-west-1:123456789012:endpoint:BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
```

## See also

- [CreateReplicationTask](https://docs.aws.amazon.com/dms/latest/APIReference/API_CreateReplicationTask.html)
in the _AWS Database Migration Service API Reference_

- [Managing AWS resources as a single unit with AWS CloudFormation stacks](../userguide/stacks.md)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
