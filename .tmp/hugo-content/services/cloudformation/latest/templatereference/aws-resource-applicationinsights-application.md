This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationInsights::Application

The `AWS::ApplicationInsights::Application` resource adds an application that is created from a resource group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApplicationInsights::Application",
  "Properties" : {
      "AttachMissingPermission" : Boolean,
      "AutoConfigurationEnabled" : Boolean,
      "ComponentMonitoringSettings" : [ ComponentMonitoringSetting, ... ],
      "CustomComponents" : [ CustomComponent, ... ],
      "CWEMonitorEnabled" : Boolean,
      "GroupingType" : String,
      "LogPatternSets" : [ LogPatternSet, ... ],
      "OpsCenterEnabled" : Boolean,
      "OpsItemSNSTopicArn" : String,
      "ResourceGroupName" : String,
      "SNSNotificationArn" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ApplicationInsights::Application
Properties:
  AttachMissingPermission: Boolean
  AutoConfigurationEnabled: Boolean
  ComponentMonitoringSettings:
    - ComponentMonitoringSetting
  CustomComponents:
    - CustomComponent
  CWEMonitorEnabled: Boolean
  GroupingType: String
  LogPatternSets:
    - LogPatternSet
  OpsCenterEnabled: Boolean
  OpsItemSNSTopicArn: String
  ResourceGroupName: String
  SNSNotificationArn: String
  Tags:
    - Tag

```

## Properties

`AttachMissingPermission`

If set to true, the managed policies for SSM and CW will be attached to the instance roles if they are missing.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutoConfigurationEnabled`

If set to `true`, the application components will be configured with the monitoring configuration recommended by Application Insights.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComponentMonitoringSettings`

The monitoring settings of the components. Not required to set up default monitoring for all components. To set up default monitoring for all components, set `AutoConfigurationEnabled` to `true`.

_Required_: No

_Type_: Array of [ComponentMonitoringSetting](aws-properties-applicationinsights-application-componentmonitoringsetting.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomComponents`

Describes a custom component by grouping similar standalone instances to monitor.

_Required_: No

_Type_: Array of [CustomComponent](aws-properties-applicationinsights-application-customcomponent.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CWEMonitorEnabled`

Indicates whether Application Insights can listen to CloudWatch events for the application resources, such as `instance terminated`, `failed deployment`, and others.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GroupingType`

Application Insights can create applications based on a resource group or on an account.
To create an account-based application using all of the resources in the account, set this
parameter to `ACCOUNT_BASED`.

_Required_: No

_Type_: String

_Allowed values_: `ACCOUNT_BASED`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LogPatternSets`

The log pattern sets.

_Required_: No

_Type_: Array of [LogPatternSet](aws-properties-applicationinsights-application-logpatternset.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OpsCenterEnabled`

Indicates whether Application Insights will create OpsItems for any problem that is
detected by Application Insights for an application.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OpsItemSNSTopicArn`

The SNS topic provided to Application Insights that is associated with the created
OpsItems to receive SNS notifications for opsItem updates.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws(-[\w]+)*:[\w\d-]+:([\w\d-]*)?:[\w\d_-]*([:/].+)*$`

_Minimum_: `20`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceGroupName`

The name of the resource group used for the application.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9.-_]*`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SNSNotificationArn`

The SNS topic ARN that is associated with SNS notifications for updates or issues.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws(-[\w]+)*:[\w\d-]+:([\w\d-]*)?:[\w\d_-]*([:/].+)*$`

_Minimum_: `20`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of `Tags`.

_Required_: No

_Type_: Array of [Tag](aws-properties-applicationinsights-application-tag.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the application, such as
`arn:aws:applicationinsights:us-east-1:123456789012:application/resource-group/my_resource_group`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ApplicationARN`

Returns the Amazon Resource Name (ARN) of the application, such as
` arn:aws:applicationinsights:us-east-1:123456789012:application/resource-group/my_resource_group`.

## Examples

- [The following example template creates an Application Insights application with all components configured with recommended monitoring settings](#aws-resource-applicationinsights-application--examples--The_following_example_template_creates_an_Application_Insights_application_with_all_components_configured_with_recommended_monitoring_settings)

- [The following example template creates an Application Insights application that includes all of the components in the account](#aws-resource-applicationinsights-application--examples--The_following_example_template_creates_an_Application_Insights_application_that_includes_all_of_the_components_in_the_account)

- [The following example template creates an Application Insights application with detailed settings](#aws-resource-applicationinsights-application--examples--The_following_example_template_creates_an_Application_Insights_application_with_detailed_settings)

- [The following example template creates an Application Insights application with CUSTOM mode component configuration](#aws-resource-applicationinsights-application--examples--The_following_example_template_creates_an_Application_Insights_application_with_CUSTOM_mode_component_configuration)

- [The following example template creates an Application Insights application with DEFAULT mode component configuration](#aws-resource-applicationinsights-application--examples--The_following_example_template_creates_an_Application_Insights_application_with_DEFAULT_mode_component_configuration)

- [The following example template creates an Application Insights application with DEFAULT\_WITH\_OVERWRITE mode component configuration](#aws-resource-applicationinsights-application--examples--The_following_example_template_creates_an_Application_Insights_application_with_DEFAULT_WITH_OVERWRITE_mode_component_configuration)

### The following example template creates an Application Insights application with all components configured with recommended monitoring settings

The following example template performs the following actions:

- Creates an Application Insights application. For more information, see
[CreateApplication](../../../../reference/cloudwatch/latest/apireference/api-createapplication.md) in the _Amazon CloudWatch_
_Application Insights API Reference_.

- Sets `AutoConfigurationEnabled` to `true`, which
configures all components of the application with the recommended
monitoring settings for the `DEFAULT` tier. For more
information, see [DescribeComponentConfigurationRecommendation](../../../../reference/cloudwatch/latest/apireference/api-describecomponentconfigurationrecommendation.md) in the
_Amazon CloudWatch Application Insights API_
_Reference_.

#### YAML

```yaml

---
Type: AWS::ApplicationInsights::Application
Properties:
  ResourceGroupName: my_resource_group
  AutoConfigurationEnabled: true
```

#### JSON

```json

{
    "Type": "AWS::ApplicationInsights::Application",
    "Properties": {
        "ResourceGroupName": "my_resource_group",
        "AutoConfigurationEnabled": true
    }
}
```

### The following example template creates an Application Insights application that includes all of the components in the account

The following example template performs the following actions:

- Creates an Application Insights application. For more information, see
[CreateApplication](../../../../reference/cloudwatch/latest/apireference/api-createapplication.md) in the _Amazon CloudWatch_
_Application Insights API Reference_.

- Sets `GroupingType` to `ACCOUNT_BASED`, which
creates an account level-based application that includes all of the supported resources in the account.

- Sets `AutoConfigurationEnabled` to `true`, which
configures all components of the application with the recommended
monitoring settings for the `DEFAULT` tier. For more
information, see [DescribeComponentConfigurationRecommendation](../../../../reference/cloudwatch/latest/apireference/api-describecomponentconfigurationrecommendation.md) in the
_Amazon CloudWatch Application Insights API_
_Reference_.

- You can specify any name for the `ResourceGroupName`.

#### YAML

```yaml

---
Type: AWS::ApplicationInsights::Application
Properties:
  AutoConfigurationEnabled: true
  GroupingType: ACCOUNT_BASED
  ResourceGroupName: my_resource_group
```

#### JSON

```json

{
    "Type": "AWS::ApplicationInsights::Application",
    "Properties": {
        "AutoConfigurationEnabled": true,
        "GroupingType": ACCOUNT_BASED,
        "ResourceGroupName": "my_resource_group"
    }
}
```

### The following example template creates an Application Insights application with detailed settings

The following example template performs the following actions:

- Creates an Application Insights application with CloudWatch Events notification and OpsCenter
enabled. For more information, see [CreateApplication](../../../../reference/cloudwatch/latest/apireference/api-createapplication.md) in the _Amazon CloudWatch Application_
_Insights API Reference_.

- Tags the application with two tags, one of which has no tag values. For more information, see
[TagResource](../../../../reference/cloudwatch/latest/apireference/api-tagresource.md) in the _Amazon CloudWatch Application Insights_
_API Reference_.

- Creates two custom instance group components. For more information, see [CreateComponent](../../../../reference/cloudwatch/latest/apireference/api-createcomponent.md) in the _Amazon CloudWatch Application_
_Insights API Reference_.

- Creates two log pattern sets. For more information, see [CreateLogPattern](../../../../reference/cloudwatch/latest/apireference/api-createlogpattern.md) in the _Amazon CloudWatch Application_
_Insights API Reference_.

- Sets `AutoConfigurationEnabled` to `true`, which configures all components of the application with the recommended monitoring settings for the `DEFAULT` tier. For more information, see [DescribeComponentConfigurationRecommendation](../../../../reference/cloudwatch/latest/apireference/api-describecomponentconfigurationrecommendation.md) in the _Amazon CloudWatch Application_
_Insights API Reference_.

#### YAML

```yaml

---
Type: AWS::ApplicationInsights::Application
Properties:
  ResourceGroupName: my_resource_group
  CWEMonitorEnabled: true
  OpsCenterEnabled: true
  OpsItemSNSTopicArn: arn:aws:sns:us-east-1:123456789012:my_topic
  AutoConfigurationEnabled: true
  Tags:
  - Key: key1
    Value: value1
  - Key: key2
    Value: ''
  CustomComponents:
  - ComponentName: test_component_1
    ResourceList:
    - arn:aws:ec2:us-east-1:123456789012:instance/i-XXXXX
  - ComponentName: test_component_2
    ResourceList:
    - arn:aws:ec2:us-east-1:123456789012:instance/i-YYYYY
    - arn:aws:ec2:us-east-1:123456789012:instance/i-ZZZZZ
  LogPatternSets:
  - PatternSetName: pattern_set_1
    LogPatterns:
    - PatternName: deadlock_pattern
      Pattern: ".*\\sDeadlocked\\sSchedulers(([^\\w].*)|($))"
      Rank: 1
  - PatternSetName: pattern_set_2
    LogPatterns:
    - PatternName: error_pattern
      Pattern: ".*[\\s\\[]ERROR[\\s\\]].*"
      Rank: 1
    - PatternName: warning_pattern
      Pattern: ".*[\\s\\[]WARN(ING)?[\\s\\]].*"
      Rank: 10
```

#### JSON

```json

{
    "Type": "AWS::ApplicationInsights::Application",
    "Properties": {
        "ResourceGroupName": "my_resource_group",
        "CWEMonitorEnabled": true,
        "OpsCenterEnabled": true,
        "OpsItemSNSTopicArn": "arn:aws:sns:us-east-1:123456789012:my_topic",
        "AutoConfigurationEnabled": true,
        "Tags": [
            {
                "Key": "key1",
                "Value": "value1"
            },
            {
                "Key": "key2",
                "Value": ""
            }
        ],
        "CustomComponents": [
            {
                "ComponentName": "test_component_1",
                "ResourceList": [
                    "arn:aws:ec2:us-east-1:123456789012:instance/i-XXXXX"
                ]
            },
            {
                "ComponentName": "test_component_2",
                "ResourceList": [
                    "arn:aws:ec2:us-east-1:123456789012:instance/i-YYYYY",
                    "arn:aws:ec2:us-east-1:123456789012:instance/i-ZZZZZ"
                ]
            }
        ],
        "LogPatternSets": [
            {
                "PatternSetName": "pattern_set_1",
                "LogPatterns": [
                    {
                        "PatternName": "deadlock_pattern",
                        "Pattern": ".*\\sDeadlocked\\sSchedulers(([^\\w].*)|($))",
                        "Rank": 1
                    }
                ]
            },
            {
                "PatternSetName": "pattern_set_2",
                "LogPatterns": [
                    {
                        "PatternName": "error_pattern",
                        "Pattern": ".*[\\s\\[]ERROR[\\s\\]].*",
                        "Rank": 1
                    },
                    {
                        "PatternName": "warning_pattern",
                        "Pattern": ".*[\\s\\[]WARN(ING)?[\\s\\]].*",
                        "Rank": 10
                    }
                ]
            }
        ]
    }
}
```

### The following example template creates an Application Insights application with CUSTOM mode component configuration

The following example template performs the following actions:

- Creates an Application Insights application. For more information, see [CreateApplication](../../../../reference/cloudwatch/latest/apireference/api-createapplication.md) in the _Amazon CloudWatch Application_
_Insights API Reference_.

- Component `my_component` sets `ComponentConfigurationMode` to
`CUSTOM`, which causes this component to be configured as
specified in `CustomComponentConfiguration`. For more
information, see [UpdateComponentConfiguration](../../../../reference/cloudwatch/latest/apireference/api-updatecomponentconfiguration.md) in the _Amazon_
_CloudWatch Application Insights API Reference_.

#### YAML

```yaml

---
Type: AWS::ApplicationInsights::Application
Properties:
  ResourceGroupName: my_resource_group
  ComponentMonitoringSettings:
  - ComponentARN: my_component
    Tier: SQL_SERVER
    ComponentConfigurationMode: CUSTOM
    CustomComponentConfiguration:
      ConfigurationDetails:
        AlarmMetrics:
        - AlarmMetricName: StatusCheckFailed
        ...
        Logs:
        - LogGroupName: my_log_group_1
          LogPath: C:\LogFolder_1\*
          LogType: DOT_NET_CORE
          Encoding: utf-8
          PatternSet: my_pattern_set_1
        ...
        WindowsEvents:
        - LogGroupName: my_windows_event_log_group_1
          EventName: Application
          EventLevels:
          - ERROR
          - WARNING
          ...
          Encoding: utf-8
          PatternSet: my_pattern_set_2
        ...
        Alarms:
        - AlarmName: my_alarm_name
          Severity: HIGH
        ...
      SubComponentTypeConfigurations:
      - SubComponentType: EC2_INSTANCE
        SubComponentConfigurationDetails:
          AlarmMetrics:
          - AlarmMetricName: DiskReadOps
          ...
          Logs:
          - LogGroupName: my_log_group_2
            LogPath: C:\LogFolder_2\*
            LogType: IIS
            Encoding: utf-8
            PatternSet: my_pattern_set_3
          ...
          WindowsEvents:
          - LogGroupName: my_windows_event_log_group_2
            EventName: Application
            EventLevels:
            - ERROR
            - WARNING
            ...
            Encoding: utf-8
            PatternSet: my_pattern_set_4
          ...
```

#### JSON

```json

{
    "Type": "AWS::ApplicationInsights::Application",
    "Properties": {
        "ResourceGroupName": "my_resource_group",
        "ComponentMonitoringSettings": [
            {
                "ComponentARN": "my_component",
                "Tier": "SQL_SERVER",
                "ComponentConfigurationMode": "CUSTOM",
                "CustomComponentConfiguration": {
                    "ConfigurationDetails": {
                        "AlarmMetrics": [
                            {
                                "AlarmMetricName": "StatusCheckFailed"
                            },
                            ...
                        ],
                        "Logs": [
                            {
                                "LogGroupName": "my_log_group_1",
                                "LogPath": "C:\\LogFolder_1\\*",
                                "LogType": "DOT_NET_CORE",
                                "Encoding": "utf-8",
                                "PatternSet": "my_pattern_set_1"
                            },
                            ...
                        ],
                        "WindowsEvents": [
                            {
                                "LogGroupName": "my_windows_event_log_group_1",
                                "EventName": "Application",
                                "EventLevels": [
                                    "ERROR",
                                    "WARNING",
                                    ...
                                ],
                                "Encoding": "utf-8",
                                "PatternSet": "my_pattern_set_2"
                            },
                            ...
                        ],
                        "Alarms": [
                            {
                                "AlarmName": "my_alarm_name",
                                "Severity": "HIGH"
                            },
                            ...
                        ]
                    },
                    "SubComponentTypeConfigurations": [
                        {
                            "SubComponentType": "EC2_INSTANCE",
                            "SubComponentConfigurationDetails": {
                                "AlarmMetrics": [
                                    {
                                        "AlarmMetricName": "DiskReadOps"
                                    },
                                    ...
                                ],
                                "Logs": [
                                    {
                                        "LogGroupName": "my_log_group_2",
                                        "LogPath": "C:\\LogFolder_2\\*",
                                        "LogType": "IIS",
                                        "Encoding": "utf-8",
                                        "PatternSet": "my_pattern_set_3"
                                    },
                                    ...
                                ],
                                "WindowsEvents": [
                                    {
                                        "LogGroupName": "my_windows_event_log_group_2",
                                        "EventName": "Application",
                                        "EventLevels": [
                                            "ERROR",
                                            "WARNING",
                                            ...
                                        ],
                                        "Encoding": "utf-8",
                                        "PatternSet": "my_pattern_set_4"
                                    },
                                    ...
                                ]
                            }
                        }
                    ]
                }
            }
        ]
    }
}
```

### The following example template creates an Application Insights application with DEFAULT mode component configuration

The following example template performs the following actions:

- Creates an Application Insights application. For more information, see [CreateApplication](../../../../reference/cloudwatch/latest/apireference/api-createapplication.md) in the _Amazon CloudWatch Application_
_Insights API Reference_.

- Component `my_component` sets `ComponentConfigurationMode` to
`DEFAULT` and `Tier` to
`SQL_SERVER`, which causes this component to be
configured with the configuration settings that Application Insights
recommends for the `SQL_Server` tier. For more information,
see [DescribeComponentConfiguration](../../../../reference/cloudwatch/latest/apireference/api-describecomponentconfiguration.md) and [UpdateComponentConfiguration](../../../../reference/cloudwatch/latest/apireference/api-updatecomponentconfiguration.md) in the _Amazon_
_CloudWatch Application Insights API Reference_.

#### YAML

```yaml

---
Type: AWS::ApplicationInsights::Application
Properties:
  ResourceGroupName: my_resource_group
  ComponentMonitoringSettings:
  - ComponentARN: my_component
    Tier: SQL_SERVER
    ComponentConfigurationMode: DEFAULT
```

#### JSON

```json

{
    "Type": "AWS::ApplicationInsights::Application",
    "Properties": {
        "ResourceGroupName": "my_resource_group",
        "ComponentMonitoringSettings": [
            {
                "ComponentARN": "my_component",
                "Tier": "SQL_SERVER",
                "ComponentConfigurationMode": "DEFAULT"
            }
        ]
    }
}
```

### The following example template creates an Application Insights application with DEFAULT\_WITH\_OVERWRITE mode component configuration

The following example template performs the following actions:

- Creates an Application Insights application. For more information, see [CreateApplication](../../../../reference/cloudwatch/latest/apireference/api-createapplication.md) in the _Amazon CloudWatch Application_
_Insights API Reference_.

- Component `my_component` sets `ComponentConfigurationMode` to
`DEFAULT_WITH_OVERWRITE` and `tier` to
`DOT_NET_CORE`, which causes this component to be
configured with the configuration settings that Application Insights
recommends for the `DOT_NET_CORE` tier. Overwritten
configuration settings are specified in the
`DefaultOverwriteComponentConfiguration`:

- At the component level, `AlarmMetrics` settings are overwritten.

- At the sub-component level, for the `EC2_Instance` type sub-components, `Logs` settings are overwritten.

For more
information, see [UpdateComponentConfiguration](../../../../reference/cloudwatch/latest/apireference/api-updatecomponentconfiguration.md) in the _Amazon CloudWatch_
_Application Insights API Reference_.

#### YAML

```yaml

---
Type: AWS::ApplicationInsights::Application
Properties:
  ResourceGroupName: my_resource_group
  ComponentMonitoringSettings:
  - ComponentName: my_component
    Tier: DOT_NET_CORE
    ComponentConfigurationMode: DEFAULT_WITH_OVERWRITE
    DefaultOverwriteComponentConfiguration:
      ConfigurationDetails:
        AlarmMetrics:
        - AlarmMetricName: StatusCheckFailed
      SubComponentTypeConfigurations:
      - SubComponentType: EC2_INSTANCE
        SubComponentConfigurationDetails:
          Logs:
          - LogGroupName: my_log_group
            LogPath: C:\LogFolder\*
            LogType: IIS
            Encoding: utf-8
            PatternSet: my_pattern_set
```

#### JSON

```json

{
    "Type": "AWS::ApplicationInsights::Application",
    "Properties": {
        "ResourceGroupName": "my_resource_group",
        "ComponentMonitoringSettings": [
            {
                "ComponentName": "my_component",
                "Tier": "DOT_NET_CORE",
                "ComponentConfigurationMode": "DEFAULT_WITH_OVERWRITE",
                "DefaultOverwriteComponentConfiguration": {
                    "ConfigurationDetails": {
                        "AlarmMetrics": [
                            {
                                "AlarmMetricName": "StatusCheckFailed"
                            }
                        ]
                    },
                    "SubComponentTypeConfigurations": [
                        {
                            "SubComponentType": "EC2_INSTANCE",
                            "SubComponentConfigurationDetails": {
                                "Logs": [
                                    {
                                        "LogGroupName": "my_log_group",
                                        "LogPath": "C:\\LogFolder\\*",
                                        "LogType": "IIS",
                                        "Encoding": "utf-8",
                                        "PatternSet": "my_pattern_set"
                                    }
                                ]
                            }
                        }
                    ]
                }
            }
        ]
    }
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon CloudWatch Application Insights

Alarm

All content copied from https://docs.aws.amazon.com/.
