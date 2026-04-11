This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MWAA::Environment

The `AWS::MWAA::Environment` resource creates an Amazon Managed Workflows for Apache Airflow (MWAA) environment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MWAA::Environment",
  "Properties" : {
      "AirflowConfigurationOptions" : Json,
      "AirflowVersion" : String,
      "DagS3Path" : String,
      "EndpointManagement" : String,
      "EnvironmentClass" : String,
      "ExecutionRoleArn" : String,
      "KmsKey" : String,
      "LoggingConfiguration" : LoggingConfiguration,
      "MaxWebservers" : Integer,
      "MaxWorkers" : Integer,
      "MinWebservers" : Integer,
      "MinWorkers" : Integer,
      "Name" : String,
      "NetworkConfiguration" : NetworkConfiguration,
      "PluginsS3ObjectVersion" : String,
      "PluginsS3Path" : String,
      "RequirementsS3ObjectVersion" : String,
      "RequirementsS3Path" : String,
      "Schedulers" : Integer,
      "SourceBucketArn" : String,
      "StartupScriptS3ObjectVersion" : String,
      "StartupScriptS3Path" : String,
      "Tags" : [ Tag, ... ],
      "WebserverAccessMode" : String,
      "WeeklyMaintenanceWindowStart" : String,
      "WorkerReplacementStrategy" : String
    }
}

```

### YAML

```yaml

Type: AWS::MWAA::Environment
Properties:
  AirflowConfigurationOptions: Json
  AirflowVersion: String
  DagS3Path: String
  EndpointManagement: String
  EnvironmentClass: String
  ExecutionRoleArn: String
  KmsKey: String
  LoggingConfiguration:
    LoggingConfiguration
  MaxWebservers: Integer
  MaxWorkers: Integer
  MinWebservers: Integer
  MinWorkers: Integer
  Name: String
  NetworkConfiguration:
    NetworkConfiguration
  PluginsS3ObjectVersion: String
  PluginsS3Path: String
  RequirementsS3ObjectVersion: String
  RequirementsS3Path: String
  Schedulers: Integer
  SourceBucketArn: String
  StartupScriptS3ObjectVersion: String
  StartupScriptS3Path: String
  Tags:
    - Tag
  WebserverAccessMode: String
  WeeklyMaintenanceWindowStart: String
  WorkerReplacementStrategy: String

```

## Properties

`AirflowConfigurationOptions`

A list of key-value pairs containing the Airflow configuration options for your environment. For example, `core.default_timezone: utc`. To learn more, see [Apache Airflow configuration options](../../../mwaa/latest/userguide/configuring-env-variables.md).

_Required_: No

_Type_: Json

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`AirflowVersion`

The version of Apache Airflow to use for the environment. If no value is specified, defaults to the latest version.

If you specify a newer version number for an existing environment, the version update requires some service interruption before taking effect.

_Allowed Values_: `2.7.2` \| `2.8.1` \| `2.9.2` \| `2.10.1` \| `2.10.3` \| `3.0.6` (latest)

_Required_: No

_Type_: String

_Pattern_: `^[0-9a-z.]+$`

_Maximum_: `32`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`DagS3Path`

The relative path to the DAGs folder on your Amazon S3 bucket. For example, `dags`. To learn more, see [Adding or updating DAGs](../../../mwaa/latest/userguide/configuring-dag-folder.md).

_Required_: No

_Type_: String

_Pattern_: `.*`

_Maximum_: `1024`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`EndpointManagement`

Defines whether the VPC endpoints configured for the environment are created, and managed, by the customer or by Amazon MWAA. If set to `SERVICE`, Amazon MWAA will create and manage the required VPC endpoints in
your VPC. If set to `CUSTOMER`, you must create, and manage, the VPC endpoints in your VPC.

_Required_: No

_Type_: String

_Allowed values_: `CUSTOMER | SERVICE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EnvironmentClass`

The environment class type. Valid values: `mw1.micro`, `mw1.small`, `mw1.medium`, `mw1.large`, `mw1.1large`, and `mw1.2large`. To learn more, see [Amazon MWAA environment class](../../../mwaa/latest/userguide/environment-class.md).

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`ExecutionRoleArn`

The Amazon Resource Name (ARN) of the execution role in IAM that allows MWAA to access AWS resources in your environment. For example, `arn:aws:iam::123456789:role/my-execution-role`. To learn more, see [Amazon MWAA Execution role](../../../mwaa/latest/userguide/mwaa-create-role.md).

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b)(-[a-z]+)?:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+$`

_Maximum_: `1224`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`KmsKey`

The AWS Key Management Service (KMS) key to encrypt and decrypt the data in your environment. You can use an AWS KMS key managed by MWAA, or a customer-managed KMS key (advanced).

_Required_: No

_Type_: String

_Pattern_: `^(((arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b)(-[a-z]+)?:kms:[a-z]{2}-[a-z]+-\d:\d+:)?key\/)?[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}|(arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b):kms:[a-z]{2}-[a-z]+-\d:\d+:)?alias/.+)$`

_Maximum_: `1224`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LoggingConfiguration`

The Apache Airflow logs being sent to CloudWatch Logs: `DagProcessingLogs`, `SchedulerLogs`, `TaskLogs`, `WebserverLogs`, `WorkerLogs`.

_Required_: No

_Type_: [LoggingConfiguration](aws-properties-mwaa-environment-loggingconfiguration.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`MaxWebservers`

The maximum number of web servers that you want to run in your environment. Amazon MWAA scales the number of Apache Airflow web servers up to the number you specify for `MaxWebservers` when you interact with your Apache Airflow environment using Apache Airflow REST API, or the Apache Airflow CLI. For example, in scenarios where your workload requires network calls to the Apache Airflow REST API with a high transaction-per-second (TPS) rate, Amazon MWAA will increase the number of web servers up to the number set in `MaxWebservers`. As TPS rates decrease Amazon MWAA disposes of the additional web servers, and scales down to the number set in `MinWebservers`.

Valid values: For environments larger than mw1.micro, accepts values from `2` to `5`. Defaults to `2` for all environment sizes except mw1.micro, which defaults to `1`.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`MaxWorkers`

The maximum number of workers that you want to run in your environment. MWAA scales the number of Apache Airflow workers up to the number you specify in the `MaxWorkers` field. For example, `20`. When there are no more tasks running, and no more in the queue, MWAA disposes of the extra workers leaving the one worker that is included with your environment, or the number you specify in `MinWorkers`.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`MinWebservers`

The minimum number of web servers that you want to run in your environment. Amazon MWAA scales the number of Apache Airflow web servers up to the number you specify for `MaxWebservers` when you interact with your Apache Airflow environment using Apache Airflow REST API, or the Apache Airflow CLI. As the transaction-per-second rate, and the network load, decrease, Amazon MWAA disposes of the additional web servers, and scales down to the number set in `MinWebservers`.

Valid values: For environments larger than mw1.micro, accepts values from `2` to `5`. Defaults to `2` for all environment sizes except mw1.micro, which defaults to `1`.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`MinWorkers`

The minimum number of workers that you want to run in your environment. MWAA scales the number of Apache Airflow workers up to the number you specify in the `MaxWorkers` field. When there are no more tasks running, and no more in the queue, MWAA disposes of the extra workers leaving the worker count you specify in the `MinWorkers` field. For example, `2`.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Name`

The name of your Amazon MWAA environment.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z][0-9a-zA-Z\-_]*$`

_Minimum_: `1`

_Maximum_: `80`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetworkConfiguration`

The VPC networking components used to secure and enable network traffic between the AWS resources for your environment. To learn more, see [About networking on Amazon MWAA](../../../mwaa/latest/userguide/networking-about.md).

_Required_: No

_Type_: [NetworkConfiguration](aws-properties-mwaa-environment-networkconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PluginsS3ObjectVersion`

The version of the plugins.zip file on your Amazon S3 bucket. To learn more, see [Installing custom plugins](../../../mwaa/latest/userguide/configuring-dag-import-plugins.md).

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`PluginsS3Path`

The relative path to the `plugins.zip` file on your Amazon S3 bucket. For example, `plugins.zip`. To learn more, see [Installing custom plugins](../../../mwaa/latest/userguide/configuring-dag-import-plugins.md).

_Required_: No

_Type_: String

_Pattern_: `.*`

_Maximum_: `1024`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`RequirementsS3ObjectVersion`

The version of the requirements.txt file on your Amazon S3 bucket. To learn more, see [Installing Python dependencies](../../../mwaa/latest/userguide/working-dags-dependencies.md).

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`RequirementsS3Path`

The relative path to the `requirements.txt` file on your Amazon S3 bucket. For example, `requirements.txt`. To learn more, see [Installing Python dependencies](../../../mwaa/latest/userguide/working-dags-dependencies.md).

_Required_: No

_Type_: String

_Pattern_: `.*`

_Maximum_: `1024`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Schedulers`

The number of schedulers that you want to run in your environment. Valid values:

- **v2** \- For environments larger than mw1.micro, accepts values from 2 to 5. Defaults to 2 for all environment sizes except mw1.micro, which defaults to 1.

- **v1** \- Accepts 1.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`SourceBucketArn`

The Amazon Resource Name (ARN) of the Amazon S3 bucket where your DAG code and supporting files are stored. For example, `arn:aws:s3:::my-airflow-bucket-unique-name`. To learn more, see [Create an Amazon S3 bucket for Amazon MWAA](../../../mwaa/latest/userguide/mwaa-s3-bucket.md).

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b)(-[a-z]+)?:s3:::[a-z0-9.\-]+$`

_Minimum_: `1`

_Maximum_: `1224`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`StartupScriptS3ObjectVersion`

The version of the startup shell script in your Amazon S3 bucket. You must specify the [version ID](../../../s3/latest/userguide/versioning-workflows.md) that Amazon S3 assigns to the file
every time you update the script.

Version IDs are Unicode, UTF-8 encoded, URL-ready, opaque strings that are no more than 1,024 bytes long. The following is an example:

`3sL4kqtJlcpXroDTDmJ+rmSpXd3dIbrHY+MTRCxf3vjVBH40Nr8X8gdRQBpUMLUo`

For more information, see [Using a startup script](../../../mwaa/latest/userguide/using-startup-script.md).

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`StartupScriptS3Path`

The relative path to the startup shell script in your Amazon S3 bucket. For example, `s3://mwaa-environment/startup.sh`.

Amazon MWAA runs the script as your environment starts, and before running the Apache Airflow process.
You can use this script to install dependencies, modify Apache Airflow configuration options, and set environment variables. For more information, see
[Using a startup script](../../../mwaa/latest/userguide/using-startup-script.md).

_Required_: No

_Type_: String

_Pattern_: `.*`

_Maximum_: `1024`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Tags`

The key-value tag pairs associated to your environment. For example, `"Environment": "Staging"`. To learn more, see [Tagging](../../../../general/latest/gr/aws-tagging.md).

If you specify new tags for an existing environment, the update requires service interruption before taking effect.

_Required_: No

_Type_: Array of [`Tag`](aws-properties-resource-tags.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`WebserverAccessMode`

The Apache Airflow _Web server_ access mode. To learn more, see [Apache Airflow access modes](../../../mwaa/latest/userguide/configuring-networking.md). Valid values: `PRIVATE_ONLY` or `PUBLIC_ONLY`.

_Required_: No

_Type_: String

_Allowed values_: `PRIVATE_ONLY | PUBLIC_ONLY`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`WeeklyMaintenanceWindowStart`

The day and time of the week to start weekly maintenance updates of your environment in the following format: `DAY:HH:MM`. For example: `TUE:03:30`. You can specify a start time in 30 minute increments only. Supported input includes the following:

- MON\|TUE\|WED\|THU\|FRI\|SAT\|SUN:(\[01\]\\\d\|2\[0-3\]):(00\|30)

_Required_: No

_Type_: String

_Pattern_: `(MON|TUE|WED|THU|FRI|SAT|SUN):([01]\d|2[0-3]):(00|30)`

_Maximum_: `9`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`WorkerReplacementStrategy`

The worker replacement strategy for your Amazon MWAA environment.

_Required_: No

_Type_: String

_Allowed values_: `FORCED | GRACEFUL`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the environment details.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN for the Amazon MWAA environment.

`CeleryExecutorQueue`

The queue ARN for the environment's [Celery Executor](https://airflow.apache.org/docs/apache-airflow/stable/core-concepts/executor/celery.html). Amazon MWAA uses a Celery Executor
to distribute tasks across multiple workers. When you create an environment in a shared VPC, you must provide access to the Celery Executor queue from your VPC.

`DatabaseVpcEndpointService`

The VPC endpoint for the environment's Amazon RDS database.

`LoggingConfiguration.DagProcessingLogs.CloudWatchLogGroupArn`

The ARN for the CloudWatch Logs group where the Apache Airflow DAG processing logs are published.

`LoggingConfiguration.SchedulerLogs.CloudWatchLogGroupArn`

The ARN for the CloudWatch Logs group where the Apache Airflow Scheduler logs are published.

`LoggingConfiguration.TaskLogs.CloudWatchLogGroupArn`

The ARN for the CloudWatch Logs group where the Apache Airflow task logs are published.

`LoggingConfiguration.WebserverLogs.CloudWatchLogGroupArn`

The ARN for the CloudWatch Logs group where the Apache Airflow Web server logs are published.

`LoggingConfiguration.WorkerLogs.CloudWatchLogGroupArn`

The ARN for the CloudWatch Logs group where the Apache Airflow Worker logs are published.

`WebserverUrl`

The URL of your Apache Airflow UI.

`WebserverVpcEndpointService`

The VPC endpoint for the environment's web server.

## Examples

- [Create a MWAA environment - JSON](#aws-resource-mwaa-environment--examples--Create_a_MWAA_environment_-_JSON)

- [Create a MWAA environment - YAML](#aws-resource-mwaa-environment--examples--Create_a_MWAA_environment_-_YAML)

### Create a MWAA environment - JSON

The following example shows how to create a MWAA environment:

#### JSON

```json

{
    "Environment": {
        "Type": "AWS::MWAA::Environment",
        "Properties": {
            "Name": "my-airflow-environment",
            "AirflowConfigurationOptions": {
                "logging.logging_level": "INFO",
                "core.default_timezone": "utc"
            },
            "Tags": {
                "Environment": "Staging",
                "Team": "Analytics"
            },
            "NetworkConfiguration": {
                "SubnetIds": [
                    "subnet-123456",
                    "subnet-789011"
                ],
                "SecurityGroupIds": [
                    "sg-0101010"
                ]
            },
            "LoggingConfiguration": {
                "DagProcessingLogs": {
                    "Enabled": true,
                    "LogLevel": "INFO"
                },
                "SchedulerLogs": {
                    "Enabled": false,
                    "LogLevel": "INFO"
                },
                "TaskLogs": {
                    "Enabled": true,
                    "LogLevel": "INFO"
                },
                "WebserverLogs": {
                    "Enabled": false,
                    "LogLevel": "INFO"
                },
                "WorkerLogs": {
                    "Enabled": false,
                    "LogLevel": "INFO"
                }
            },
            "SourceBucketArn": "arn:aws:s3:::my-dags-bucket",
            "ExecutionRoleArn": "arn:aws:iam::012345678900:role/service-role/my-execution-role",
            "MaxWorkers": 1,
            "DagS3Path": "dags",
            "EnvironmentClass": "mw1.small"
        }
    }
}
```

### Create a MWAA environment - YAML

The following example shows how to create a MWAA environment:

#### YAML

```yaml

Environment:
    Properties:
      AirflowConfigurationOptions:
        core.default_timezone: utc
        logging.logging_level: INFO
      DagS3Path: dags
      EnvironmentClass: mw1.small
      ExecutionRoleArn: "arn:aws:iam::012345678900:role/service-role/my-execution-role"
      LoggingConfiguration:
        DagProcessingLogs:
          Enabled: true
          LogLevel: INFO
        SchedulerLogs:
          Enabled: false
          LogLevel: INFO
        TaskLogs:
          Enabled: true
          LogLevel: INFO
        WebserverLogs:
          Enabled: false
          LogLevel: INFO
        WorkerLogs:
          Enabled: false
          LogLevel: INFO
      MaxWorkers: 1
      Name: my-airflow-environment
      NetworkConfiguration:
        SecurityGroupIds:
          - sg-0101010
        SubnetIds:
          - subnet-123456
          - subnet-789011
      SourceBucketArn: "arn:aws:s3:::my-dags-bucket"
      Tags:
        Environment: Staging
        Team: Analytics
    Type: "AWS::MWAA::Environment"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Managed Workflows for Apache Airflow

LoggingConfiguration

All content copied from https://docs.aws.amazon.com/.
