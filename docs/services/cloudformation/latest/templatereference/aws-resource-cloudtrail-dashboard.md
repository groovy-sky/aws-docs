This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudTrail::Dashboard

###### Important

CloudTrail Lake will no longer be open to new customers starting May 31, 2026. If you would like to use CloudTrail Lake, sign up prior to that date. Existing customers can continue to use the service as normal. For more information, see [CloudTrail Lake availability change](../../../awscloudtrail/latest/userguide/cloudtrail-lake-service-availability-change.md).

Creates a custom dashboard or the Highlights dashboard.

- **Custom dashboards** \- Custom dashboards allow you to query
events in any event data store type. You can add up to 10 widgets to a custom dashboard. You can manually refresh a custom dashboard, or you can set a refresh schedule.

- **Highlights dashboard** \- You can create
the Highlights dashboard to see a summary of key user activities and API usage across all your event data stores.
CloudTrail Lake manages the Highlights dashboard and refreshes the dashboard every 6 hours. To create the Highlights dashboard, you must set and enable a refresh schedule.

CloudTrail runs queries to populate the dashboard's widgets during a manual or scheduled refresh. CloudTrail must be granted permissions to run the `StartQuery` operation on your behalf. To provide permissions, run the `PutResourcePolicy` operation to attach a resource-based policy to each event data store. For more information,
see [Example: Allow CloudTrail to run queries to populate a dashboard](../../../awscloudtrail/latest/userguide/security-iam-resource-based-policy-examples.md#security_iam_resource-based-policy-examples-eds-dashboard) in the _AWS CloudTrail User Guide_.

To set a refresh schedule, CloudTrail must be granted permissions to run the `StartDashboardRefresh` operation to refresh the dashboard on your behalf. To provide permissions, run the `PutResourcePolicy` operation to attach a resource-based policy to the dashboard. For more information,
see [Resource-based policy example for a dashboard](../../../awscloudtrail/latest/userguide/security-iam-resource-based-policy-examples.md#security_iam_resource-based-policy-examples-dashboards) in the _AWS CloudTrail User Guide_.

For more information about dashboards, see [CloudTrail Lake dashboards](../../../awscloudtrail/latest/userguide/lake-dashboard.md) in the _AWS CloudTrail User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudTrail::Dashboard",
  "Properties" : {
      "Name" : String,
      "RefreshSchedule" : RefreshSchedule,
      "Tags" : [ Tag, ... ],
      "TerminationProtectionEnabled" : Boolean,
      "Widgets" : [ Widget, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CloudTrail::Dashboard
Properties:
  Name: String
  RefreshSchedule:
    RefreshSchedule
  Tags:
    - Tag
  TerminationProtectionEnabled: Boolean
  Widgets:
    - Widget

```

## Properties

`Name`

The name of the dashboard. The name must be unique to your account.

To create the Highlights dashboard, the name must be `AWSCloudTrail-Highlights`.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RefreshSchedule`

The schedule for a dashboard refresh.

_Required_: No

_Type_: [RefreshSchedule](aws-properties-cloudtrail-dashboard-refreshschedule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of tags.

_Required_: No

_Type_: Array of [Tag](aws-properties-cloudtrail-dashboard-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TerminationProtectionEnabled`

Specifies whether termination protection is enabled for the dashboard. If termination protection is enabled, you cannot delete the dashboard until termination protection is disabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Widgets`

An array of widgets for a custom dashboard. A custom dashboard can have a maximum of ten widgets.

You do not need to specify widgets for the Highlights dashboard.

_Required_: No

_Type_: Array of [Widget](aws-properties-cloudtrail-dashboard-widget.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`CreatedTimestamp`

The timestamp that shows when the dashboard was created.

`DashboardArn`

The ARN for the dashboard.

`Status`

The status of the dashboard.

`Type`

The type of dashboard.

`UpdatedTimestamp`

The timestamp that shows when the dashboard was updated.

## Examples

### Example: Create a custom dashboard

The following example creates a custom dashboard named `AccountActivityDashboard` with four widgets.
In this example, a refresh schedule has been set for every 6 hours and termination protection is enabled to prevent the dashboard from being accidentally deleted.

#### JSON

```json

{
    "Resources": {
        "Dashboard": {
            "Type": "AWS::CloudTrail::Dashboard",
            "Properties": {
                "Name": "AccountActivityDashboard",
                "RefreshSchedule": {
                    "Frequency": {
                        "Unit": "HOURS",
                        "Value": 6
                    },
                    "Status": "ENABLED",
                    "TimeOfDay": "00:00"
                },
                "Tags": [{
                        "Key": "k1",
                        "Value": "v1"
                    },
                    {
                        "Key": "k2",
                        "Value": "v2"
                    }
                ],
                "TerminationProtectionEnabled": true,
                "Widgets": [{
                        "ViewProperties": {
                            "Height": "2",
                            "Width": "4",
                            "Title": "TopErrors",
                            "View": "Table"
                        },
                        "QueryStatement": "SELECT errorCode, COUNT(*) AS eventCount FROM eds WHERE eventTime > '?' AND eventTime < '?' AND (errorCode is not null) GROUP BY errorCode ORDER BY eventCount DESC LIMIT 100",
                        "QueryParameters": ["$StartTime$", "$EndTime$"]
                    },
                    {
                        "ViewProperties": {
                            "Height": "2",
                            "Width": "4",
                            "Title": "MostActiveRegions",
                            "View": "PieChart",
                            "LabelColumn": "awsRegion",
                            "ValueColumn": "eventCount",
                            "FilterColumn": "awsRegion"
                        },
                        "QueryStatement": "SELECT awsRegion, COUNT(*) AS eventCount FROM eds where eventTime > '?' and eventTime < '?' GROUP BY awsRegion ORDER BY eventCount LIMIT 100",
                        "QueryParameters": ["$StartTime$", "$EndTime$"]
                    },
                    {
                        "ViewProperties": {
                            "Height": "2",
                            "Width": "4",
                            "Title": "AccountActivity",
                            "View": "LineChart",
                            "YAxisColumn": "eventCount",
                            "XAxisColumn": "eventDate",
                            "FilterColumn": "readOnly"
                        },
                        "QueryStatement": "SELECT DATE_TRUNC('?', eventTime) AS eventDate, IF(readOnly, 'read', 'write') AS readOnly, COUNT(*) as eventCount FROM eds WHERE eventTime > '?' AND eventTime < '?' GROUP BY DATE_TRUNC('?', eventTime), readOnly ORDER BY DATE_TRUNC('?', eventTime), readOnly",
                        "QueryParameters": ["$Period$", "$StartTime$", "$EndTime$", "$Period$", "$Period$"]
                    },
                    {
                        "ViewProperties": {
                            "Height": "2",
                            "Width": "4",
                            "Title": "TopServices",
                            "View": "BarChart",
                            "LabelColumn": "service",
                            "ValueColumn": "eventCount",
                            "FilterColumn": "service",
                            "Orientation": "Horizontal"
                        },
                        "QueryStatement": "SELECT REPLACE(eventSource, '.amazonaws.com') AS service, COUNT(*) AS eventCount FROM eds WHERE eventTime > '?' AND eventTime < '?' GROUP BY eventSource ORDER BY eventCount DESC LIMIT 100",
                        "QueryParameters": ["$StartTime$", "$EndTime$"]
                    }
                ]
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
  Dashboard:
    Type: 'AWS::CloudTrail::Dashboard'
    Properties:
      Name: "AccountActivityDashboard"
      TerminationProtectionEnabled: true
      RefreshSchedule:
        Frequency:
          Unit: "HOURS"
          Value: 6
        Status: "ENABLED"
        TimeOfDay: "00:00"
      Tags:
        - Key: "k1"
          Value: "v1"
        - Key: "k2"
          Value: "v2"
      Widgets:
        - QueryStatement: "SELECT errorCode, COUNT(*) AS eventCount FROM eds WHERE eventTime > '?' AND eventTime < '?' AND (errorCode is not null) GROUP BY errorCode ORDER BY eventCount DESC LIMIT 100"
          QueryParameters:
            - "$StartTime$"
            - "$EndTime$"
          ViewProperties:
            Height: "2"
            Width: "4"
            Title: "TopErrors"
            View: "Table"
        - QueryStatement: "SELECT awsRegion, COUNT(*) AS eventCount FROM eds where eventTime > '?' and eventTime < '?' GROUP BY awsRegion ORDER BY eventCount LIMIT 100"
          QueryParameters:
            - "$StartTime$"
            - "$EndTime$"
          ViewProperties:
            Height: "2"
            Width: "4"
            Title: "MostActiveRegions"
            View: "PieChart"
            LabelColumn: "awsRegion"
            ValueColumn: "eventCount"
            FilterColumn: "awsRegion"
        - QueryStatement: "SELECT DATE_TRUNC('?', eventTime) AS eventDate, IF(readOnly, 'read', 'write') AS readOnly, COUNT(*) as eventCount FROM eds WHERE eventTime > '?' AND eventTime < '?' GROUP BY DATE_TRUNC('?', eventTime), readOnly ORDER BY DATE_TRUNC('?', eventTime), readOnly"
          QueryParameters:
            - "$Period$"
            - "$StartTime$"
            - "$EndTime$"
            - "$Period$"
            - "$Period$"
          ViewProperties:
            Height: "2"
            Width: "4"
            Title: "AccountActivity"
            View: "LineChart"
            YAxisColumn: "eventCount"
            XAxisColumn: "eventDate"
            FilterColumn: "readOnly"
        - QueryStatement: "SELECT REPLACE(eventSource, '.amazonaws.com') AS service, COUNT(*) AS eventCount FROM eds WHERE eventTime > '?' AND eventTime < '?' GROUP BY eventSource ORDER BY eventCount DESC LIMIT 100"
          QueryParameters:
            - "$StartTime$"
            - "$EndTime$"
          ViewProperties:
            Height: "2"
            Width: "4"
            Title: "TopServices"
            View: "BarChart"
            LabelColumn: "service"
            ValueColumn: "eventCount"
            FilterColumn: "service"
            Orientation: "Horizontal"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Frequency

All content copied from https://docs.aws.amazon.com/.
