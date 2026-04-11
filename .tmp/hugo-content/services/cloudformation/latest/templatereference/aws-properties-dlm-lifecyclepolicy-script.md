This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DLM::LifecyclePolicy Script

**\[Custom snapshot policies that target instances only\]** Information about pre and/or post scripts for a
snapshot lifecycle policy that targets instances. For more information, see
[Automating application-consistent snapshots with pre and post scripts](../../../ec2/latest/userguide/automate-app-consistent-backups.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExecuteOperationOnScriptFailure" : Boolean,
  "ExecutionHandler" : String,
  "ExecutionHandlerService" : String,
  "ExecutionTimeout" : Integer,
  "MaximumRetryCount" : Integer,
  "Stages" : [ String, ... ]
}

```

### YAML

```yaml

  ExecuteOperationOnScriptFailure: Boolean
  ExecutionHandler: String
  ExecutionHandlerService: String
  ExecutionTimeout: Integer
  MaximumRetryCount: Integer
  Stages:
    - String

```

## Properties

`ExecuteOperationOnScriptFailure`

Indicates whether Amazon Data Lifecycle Manager should default to crash-consistent snapshots if the
pre script fails.

- To default to crash consistent snapshot if the pre script fails,
specify `true`.

- To skip the instance for snapshot creation if the pre script fails,
specify `false`.

This parameter is supported only if you run a pre script. If you run a post
script only, omit this parameter.

Default: true

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExecutionHandler`

The SSM document that includes the pre and/or post scripts to run.

- If you are automating VSS backups, specify `AWS_VSS_BACKUP`.
In this case, Amazon Data Lifecycle Manager automatically uses the `AWSEC2-CreateVssSnapshot`
SSM document.

- If you are automating application-consistent snapshots for SAP HANA
workloads, specify `AWSSystemsManagerSAP-CreateDLMSnapshotForSAPHANA`.

- If you are using a custom SSM document that you own, specify either
the name or ARN of the SSM document. If you are using a custom SSM
document that is shared with you, specify the ARN of the SSM document.

_Required_: No

_Type_: String

_Pattern_: `^([a-zA-Z0-9_\-.]{3,128}|[a-zA-Z0-9_\-.:/]{3,200}|[A-Z0-9_]+)$`

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExecutionHandlerService`

Indicates the service used to execute the pre and/or post scripts.

- If you are using custom SSM documents or automating
application-consistent snapshots of SAP HANA workloads, specify
`AWS_SYSTEMS_MANAGER`.

- If you are automating VSS Backups, omit this parameter.

Default: AWS\_SYSTEMS\_MANAGER

_Required_: No

_Type_: String

_Allowed values_: `AWS_SYSTEMS_MANAGER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExecutionTimeout`

Specifies a timeout period, in seconds, after which Amazon Data Lifecycle Manager fails the script
run attempt if it has not completed. If a script does not complete within its
timeout period, Amazon Data Lifecycle Manager fails the attempt. The timeout period applies to the pre
and post scripts individually.

If you are automating VSS Backups, omit this parameter.

Default: 10

_Required_: No

_Type_: Integer

_Minimum_: `10`

_Maximum_: `120`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumRetryCount`

Specifies the number of times Amazon Data Lifecycle Manager should retry scripts that fail.

- If the pre script fails, Amazon Data Lifecycle Manager retries the entire snapshot creation
process, including running the pre and post scripts.

- If the post script fails, Amazon Data Lifecycle Manager retries the post script only; in this
case, the pre script will have completed and the snapshot might have
been created.

If you do not want Amazon Data Lifecycle Manager to retry failed scripts, specify `0`.

Default: 0

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Stages`

Indicate which scripts Amazon Data Lifecycle Manager should run on target instances. Pre scripts
run before Amazon Data Lifecycle Manager initiates snapshot creation. Post scripts run after Amazon Data Lifecycle Manager
initiates snapshot creation.

- To run a pre script only, specify `PRE`. In this case,
Amazon Data Lifecycle Manager calls the SSM document with the `pre-script` parameter
before initiating snapshot creation.

- To run a post script only, specify `POST`. In this case,
Amazon Data Lifecycle Manager calls the SSM document with the `post-script` parameter
after initiating snapshot creation.

- To run both pre and post scripts, specify both `PRE` and `POST`. In
this case, Amazon Data Lifecycle Manager calls the SSM document with the `pre-script`
parameter before initiating snapshot creation, and then it calls the SSM
document again with the `post-script` parameter after initiating
snapshot creation.

If you are automating VSS Backups, omit this parameter.

Default: PRE and POST

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Schedule

ShareRule

All content copied from https://docs.aws.amazon.com/.
