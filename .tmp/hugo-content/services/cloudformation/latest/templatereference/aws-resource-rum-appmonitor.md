This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RUM::AppMonitor

Creates a CloudWatch RUM app monitor, which you can use to collect telemetry data from your application
and send it to CloudWatch RUM. The data includes performance and reliability information such as
page load time, client-side errors,
and user behavior.

After you create an app monitor, sign in to the CloudWatch RUM console to get
the JavaScript code snippet to add to your web application. For more information, see
[How do I find a code snippet\
that I've already generated?](../../../amazoncloudwatch/latest/monitoring/cloudwatch-rum-find-code-snippet.md)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::RUM::AppMonitor",
  "Properties" : {
      "AppMonitorConfiguration" : AppMonitorConfiguration,
      "CustomEvents" : CustomEvents,
      "CwLogEnabled" : Boolean,
      "DeobfuscationConfiguration" : DeobfuscationConfiguration,
      "Domain" : String,
      "DomainList" : [ String, ... ],
      "Name" : String,
      "Platform" : String,
      "ResourcePolicy" : ResourcePolicy,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::RUM::AppMonitor
Properties:
  AppMonitorConfiguration:
    AppMonitorConfiguration
  CustomEvents:
    CustomEvents
  CwLogEnabled: Boolean
  DeobfuscationConfiguration:
    DeobfuscationConfiguration
  Domain: String
  DomainList:
    - String
  Name: String
  Platform: String
  ResourcePolicy:
    ResourcePolicy
  Tags:
    - Tag

```

## Properties

`AppMonitorConfiguration`

A structure that contains much of the configuration data for the app monitor. If you are using
Amazon Cognito for authorization, you must include this structure in your request, and it
must include the ID of the
Amazon Cognito identity pool to use for authorization. If you don't
include `AppMonitorConfiguration`, you must set up your own
authorization method. For more information, see
[Authorize your application\
to send data to AWS](../../../amazoncloudwatch/latest/monitoring/cloudwatch-rum-get-started-authorization.md).

If you omit this argument, the sample rate used for CloudWatch RUM is set to 10% of the user sessions.

_Required_: No

_Type_: [AppMonitorConfiguration](aws-properties-rum-appmonitor-appmonitorconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomEvents`

Specifies whether this app monitor allows the web client to define and send custom events. If you omit
this parameter, custom events are `DISABLED`.

_Required_: No

_Type_: [CustomEvents](aws-properties-rum-appmonitor-customevents.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CwLogEnabled`

Data collected by CloudWatch RUM is kept by RUM for 30 days and then deleted. This parameter specifies whether
CloudWatch RUM
sends a copy of this telemetry data to Amazon CloudWatch Logs
in your account. This enables you to keep the telemetry data for more than 30 days, but it does incur
Amazon CloudWatch Logs charges.

If you omit this parameter, the default is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeobfuscationConfiguration`

A structure that contains the configuration for how an app monitor can deobfuscate stack traces.

_Required_: No

_Type_: [DeobfuscationConfiguration](aws-properties-rum-appmonitor-deobfuscationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Domain`

The top-level internet domain name for which your application has administrative authority. This parameter or the `DomainList` parameter is required.

_Required_: No

_Type_: String

_Pattern_: `^(localhost)$|^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$|(?=^[a-zA-Z0-9\.\*-]{4,253}$)(?!.*\.-)(?!.*-\.)(?!.*\.\.)(?!.*[^\.]{64,})^(\*\.)?(?![-\.\*])[^\*]{1,}\.(\*|(?!.*--)(?=.*[a-zA-Z])[^\*]{1,}[^\*-])$`

_Minimum_: `1`

_Maximum_: `253`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainList`

List the domain names for which your application has administrative authority. This parameter or the `Domain` parameter is required.

You can have a minimum of 1 and a maximum of 5 `Domain` under `DomainList`. Each `Domain` must be a minimum length of 1 and a maximum of 253 characters.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `253 | 5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A name for the app monitor. This parameter is required.

_Required_: Yes

_Type_: String

_Pattern_: `[\.\-_/#A-Za-z0-9]+`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Platform`

Property description not available.

_Required_: No

_Type_: String

_Allowed values_: `Web | Android | iOS`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourcePolicy`

Use this structure to assign a resource-based policy to a CloudWatch RUM app monitor to control access to it. Each app monitor can
have one resource-based policy. The maximum size of the policy is 4 KB. To learn more about using resource policies with RUM, see [Using resource-based policies with CloudWatch RUM](../../../amazoncloudwatch/latest/monitoring/cloudwatch-rum-resource-policies.md).

_Required_: No

_Type_: [ResourcePolicy](aws-properties-rum-appmonitor-resourcepolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Assigns one or more tags (key-value pairs) to the app monitor.

Tags can help you organize and categorize your resources. You can also use them to scope user
permissions by granting a user
permission to access or change only resources with certain tag values.

Tags don't have any semantic meaning to AWS and are interpreted strictly as strings of characters.

You can associate as many as 50 tags with an app monitor.

For more information, see [Tagging AWS resources](../../../../general/latest/gr/aws-tagging.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-rum-appmonitor-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the app monitor.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`Id`

The ID of the app monitor, such as
`123456ab-1234-4ca9-9d2f-a1b2c3456789`.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CloudWatch RUM

AppMonitorConfiguration

All content copied from https://docs.aws.amazon.com/.
