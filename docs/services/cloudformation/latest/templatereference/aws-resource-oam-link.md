This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Oam::Link

Creates a link between a source account and a sink that you have created in a monitoring account.

Before you create a link, you must create a sink in the monitoring account. The sink must have a sink policy
that permits the source account to link to it. You can grant permission to source accounts by granting permission
to an entire organization, an organizational unit, or to individual accounts.

For more information, see
[CreateSink](../../../../reference/oam/latest/apireference/api-createsink.md) and
[PutSinkPolicy](../../../../reference/oam/latest/apireference/api-putsinkpolicy.md).

Each monitoring account can be linked to as many as 100,000 source accounts.

Each source account can be linked to as many as five monitoring accounts.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Oam::Link",
  "Properties" : {
      "LabelTemplate" : String,
      "LinkConfiguration" : LinkConfiguration,
      "ResourceTypes" : [ String, ... ],
      "SinkIdentifier" : String,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::Oam::Link
Properties:
  LabelTemplate: String
  LinkConfiguration:
    LinkConfiguration
  ResourceTypes:
    - String
  SinkIdentifier: String
  Tags:
    Key: Value

```

## Properties

`LabelTemplate`

Specify a friendly human-readable name to use to identify this source account when you are viewing data from it in the monitoring
account.

You can include the following variables in your template:

- `$AccountName` is the name of the account

- `$AccountEmail` is a globally-unique email address, which includes the
email domain, such as `mariagarcia@example.com`

- `$AccountEmailNoDomain` is an email address without the domain name,
such as `mariagarcia`

###### Note

In the AWS GovCloud (US-East) and AWS GovCloud (US-West) Regions, the only supported option is to use custom labels, and the `$AccountName`, `$AccountEmail`,
and `$AccountEmailNoDomain` variables
all resolve as _account-id_ instead of the specified variable.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LinkConfiguration`

Use this structure to optionally create filters that specify that only some metric namespaces or log groups are to be shared from
the source account to the monitoring account.

_Required_: No

_Type_: [LinkConfiguration](aws-properties-oam-link-linkconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceTypes`

An array of strings that define which types of data that the source account shares with the monitoring
account. Valid values are `AWS::CloudWatch::Metric | AWS::Logs::LogGroup | AWS::XRay::Trace | AWS::ApplicationInsights::Application | AWS::InternetMonitor::Monitor`.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SinkIdentifier`

The ARN of the sink in the monitoring account that you want to link to.
You can use [ListSinks](../../../../reference/oam/latest/apireference/api-listsinks.md) to
find the ARNs of sinks.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to the link.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Object of String

_Pattern_: `^(?!aws:.*).{1,128}$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the link. For example,
`arn:aws:oam:us-west-1:111111111111:link:abcd1234-a123-456a-a12b-a123b456c789`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the link. For example, `arn:aws:oam:us-west-1:111111111111:link:abcd1234-a123-456a-a12b-a123b456c789`

`Label`

The friendly human-readable name used to identify this source account when it is viewed from the monitoring
account. For example, `my-account1`.

## Examples

- [Sample CloudWatch cross-account observability link](#aws-resource-oam-link--examples--Sample_CloudWatch_cross-account_observability_link)

- [Sample link for CloudWatch Application Insights applications support](#aws-resource-oam-link--examples--Sample_link_for_applications_support)

- [Cross-account observability link with resource filtering.](#aws-resource-oam-link--examples--Cross-account_observability_link_with_resource_filtering.)

### Sample CloudWatch cross-account observability link

This example creates a link from the current source account to the sink created
in the `111111111111` account. Over this link, logs and traces are shared
but metrics are not.

#### JSON

```json

{
    "LabelTemplate": "$AccountEmail",
    "ResourceTypes": [
        "AWS::Logs::LogGroup",
        "AWS::XRay::Trace"
    ],
     "SinkIdentifier": "arn:aws:oam:eu-north-1:111111111111:sink/EXAMPLE-206d-4daf-9b42-1e17d5f145ef"
}
```

#### YAML

```yaml

LabelTemplate: "$AccountEmail"
ResourceTypes:
   - "AWS::Logs::LogGroup"
   - "AWS::XRay::Trace"
 SinkIdentifier: "arn:aws:oam:eu-north-1:111111111111:sink/EXAMPLE-206d-4daf-9b42-1e17d5f145ef"

```

### Sample link for CloudWatch Application Insights applications support

This example creates a link from the current source account to the sink created
in the `111111111111` account.
To properly view Application Insights applications with dashboards,
logs, metrics, traces, and Application Insights applications must be shared.

#### JSON

```json

{
     "LabelTemplate": "$AccountEmail",
     "ResourceTypes": [
         "AWS::Logs::LogGroup",
         "AWS::CloudWatch::Metric",
         "AWS::XRay::Trace",
         "AWS::ApplicationInsights::Application"
     ],
     "SinkIdentifier": "arn:aws:oam:eu-north-1:111111111111:sink/EXAMPLE-206d-4daf-9b42-1e17d5f145ef"
 }
```

#### YAML

```yaml

LabelTemplate: "$AccountEmail"
 ResourceTypes:
    - "AWS::Logs::LogGroup"
    - "AWS::CloudWatch::Metric"
    - "AWS::XRay::Trace"
    - "AWS::ApplicationInsights::Application"
 SinkIdentifier: "arn:aws:oam:eu-north-1:111111111111:sink/EXAMPLE-206d-4daf-9b42-1e17d5f145ef"
SinkIdentifier: "arn:aws:oam:eu-north-1:1111111111111111:sink/EXAMPLE-206d-4daf-9b42-1e17d5f145ef"

```

### Cross-account observability link with resource filtering.

This example creates a link that shares only one metric namespace and one log group
from the source account to the monitoring account.

#### JSON

```json

{
    "TestLink": {

        "Type": "AWS::Oam::Link",
        "Properties": {
            "LabelTemplate": "$AccountEmail",
            "ResourceTypes": [
                "AWS::CloudWatch::Metric",
                "AWS::Logs::LogGroup"
            ],
            "SinkIdentifier": {
                "Fn::ImportValue": "export-canary-sinkarn"
            },
            "LinkConfiguration": {
                "MetricConfiguration": {
                    "Filter": "Namespace = 'TestNamespace'"
                },
                "LogGroupConfiguration": {
                    "Filter": "LogGroupName = 'TestLogGroupName'"
                }
            }
        }
    }
}
```

#### YAML

```yaml

TestLink:
  DependsOn: WaiterCustomResource
  Type: AWS::Oam::Link
  Properties:
    LabelTemplate: "$AccountEmail"
    ResourceTypes:
      - "AWS::CloudWatch::Metric"
      - "AWS::Logs::LogGroup"
    SinkIdentifier: !ImportValue export-canary-sinkarn
    LinkConfiguration:
      MetricConfiguration:
        Filter: "Namespace = 'TestNamespace'"
      LogGroupConfiguration:
        Filter: "LogGroupName = 'TestLogGroupName'"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Observability Access Manager (OAM)

LinkConfiguration

All content copied from https://docs.aws.amazon.com/.
