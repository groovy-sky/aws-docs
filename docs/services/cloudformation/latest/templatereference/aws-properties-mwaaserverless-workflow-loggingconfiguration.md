This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MWAAServerless::Workflow LoggingConfiguration

Configuration for workflow logging that specifies where you should store your workflow execution logs. Amazon Managed Workflows for Apache Airflow Serverless provides comprehensive logging capabilities that capture workflow execution details, task-level information, and system events. Logs are automatically exported to your specified CloudWatch log group using remote logging functionality, providing centralized observability across the distributed, multi-tenant execution environment. This enables effective debugging, monitoring, and compliance auditing of workflow executions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LogGroupName" : String
}

```

### YAML

```yaml

  LogGroupName: String

```

## Properties

`LogGroupName`

The name of the CloudWatch log group where workflow execution logs are stored.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EncryptionConfiguration

NetworkConfiguration

All content copied from https://docs.aws.amazon.com/.
