This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::APS::Workspace

An Amazon Managed Service for Prometheus workspace is a logical and isolated
Prometheus server dedicated to ingesting, storing, and querying your
Prometheus-compatible metrics.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::APS::Workspace",
  "Properties" : {
      "AlertManagerDefinition" : String,
      "Alias" : String,
      "KmsKeyArn" : String,
      "LoggingConfiguration" : LoggingConfiguration,
      "QueryLoggingConfiguration" : QueryLoggingConfiguration,
      "Tags" : [ Tag, ... ],
      "WorkspaceConfiguration" : WorkspaceConfiguration
    }
}

```

### YAML

```yaml

Type: AWS::APS::Workspace
Properties:
  AlertManagerDefinition: String
  Alias: String
  KmsKeyArn: String
  LoggingConfiguration:
    LoggingConfiguration
  QueryLoggingConfiguration:
    QueryLoggingConfiguration
  Tags:
    - Tag
  WorkspaceConfiguration:
    WorkspaceConfiguration

```

## Properties

`AlertManagerDefinition`

The alert manager definition, a YAML configuration for the alert manager in your
Amazon Managed Service for Prometheus workspace.

For details about the alert manager definition, see [Creating an alert\
manager configuration files](../../../prometheus/latest/userguide/amp-alertmanager-config.md) in the _Amazon Managed Service for Prometheus User_
_Guide_.

The following example shows part of a CloudFormation YAML file with an embedded alert
manager definition (following the `- |-`).

```

  Workspace:
    Type: AWS::APS::Workspace
    ....
    Properties:
      ....
      AlertManagerDefinition:
        Fn::Sub:
          - |-
            alertmanager_config: |
              templates:
                - 'default_template'
              route:
                receiver: example-sns
              receivers:
                - name: example-sns
                  sns_configs:
                    - topic_arn: 'arn:aws:sns:${AWS::Region}:${AWS::AccountId}:${TopicName}'
          -
```

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Alias`

The alias that is assigned to this workspace to help identify it. It does not need to
be unique.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyArn`

(optional) The ARN for a customer managed AWS KMS key to use for encrypting
data within your workspace. For more information about using your own key in your
workspace, see [Encryption at rest](../../../prometheus/latest/userguide/encryption-at-rest-amazon-service-prometheus.md) in the _Amazon Managed Service for Prometheus_
_User Guide_.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[-a-z]*:kms:[-a-z0-9]+:[0-9]{12}:key/.+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LoggingConfiguration`

Contains information about the logging configuration for the workspace.

_Required_: No

_Type_: [LoggingConfiguration](aws-properties-aps-workspace-loggingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryLoggingConfiguration`

The definition of logging configuration in an Amazon Managed Service for Prometheus
workspace.

_Required_: No

_Type_: [QueryLoggingConfiguration](aws-properties-aps-workspace-queryloggingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The list of tag keys and values that are associated with the workspace.

_Required_: No

_Type_: Array of [Tag](aws-properties-aps-workspace-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkspaceConfiguration`

Use this structure to define label sets and the ingestion limits for time series that
match label sets, and to specify the retention period of the workspace.

_Required_: No

_Type_: [WorkspaceConfiguration](aws-properties-aps-workspace-workspaceconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name. For example:

`{ "Ref": "Id" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the workspace. For example,
`arn:aws:aps:<region>:123456789012:workspace/ws-example1-1234-abcd-5678-ef90abcd1234`.

`PrometheusEndpoint`

The Prometheus endpoint available for this workspace. For example,
`https://aps-workspaces.<region>.amazonaws.com/workspaces/ws-example1-1234-abcd-5678-ef90abcd1234/api/v1/`.

`WorkspaceId`

The unique ID for the workspace. For example,
`ws-example1-1234-abcd-5678-ef90abcd1234`.

## Examples

- [Amazon Managed Service for Prometheus workspace example](#aws-resource-aps-workspace--examples--Amazon_Managed_Service_for_Prometheus_workspace_example)

- [Amazon Managed Service for Prometheus logging configuration example](#aws-resource-aps-workspace--examples--Amazon_Managed_Service_for_Prometheus_logging_configuration_example)

### Amazon Managed Service for Prometheus workspace example

The following example creates an Amazon Managed Service for Prometheus workspace with an alias
and one tag.

#### JSON

```json

{ "Resources": { "APSWorkspace": { "Type":
                "AWS::APS::Workspace", "Properties": { "Alias": "TestWorkspace" "Tags": [ { "Key":
                "BusinessPurpose", "Value": "LoadTesting" } ] } } } }
```

#### YAML

```yaml

Resources: APSWorkspace: Type: AWS::APS::Workspace Properties:
                Alias: TestWorkspace Tags: - Key: BusinessPurpose Value: LoadTesting
```

### Amazon Managed Service for Prometheus logging configuration example

The following example creates a new workspace and sets a new logging
configuration. You must replace the `LogGroupArn` with a valid ARN
for your system.

#### JSON

```json

{ "Resources": { "APSWorkspace": { "Type":
                "AWS::APS::Workspace", "Properties": { "Alias": "TestWorkspace",
                "LoggingConfiguration": { "LogGroupArn":
                "arn:aws:logs:{region}:{account}:log-group:test-log-group:*" }, "Tags": [ { "Key":
                "BusinessPurpose", "Value": "LoadTesting" } ] } } } }
```

#### YAML

```yaml

Resources: APSWorkspace: Type: AWS::APS::Workspace Properties:
                Alias: TestWorkspace LoggingConfiguration: LogGroupArn:
                "arn:aws:logs:{region}:{account}:log-group:test-log-group:*" Tags: - Key:
                BusinessPurpose Value: LoadTesting
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VpcConfiguration

CloudWatchLogDestination

All content copied from https://docs.aws.amazon.com/.
