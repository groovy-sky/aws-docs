This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RUM::AppMonitor AppMonitorConfiguration

This structure contains much of the configuration data for the app monitor.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowCookies" : Boolean,
  "EnableXRay" : Boolean,
  "ExcludedPages" : [ String, ... ],
  "FavoritePages" : [ String, ... ],
  "GuestRoleArn" : String,
  "IdentityPoolId" : String,
  "IncludedPages" : [ String, ... ],
  "MetricDestinations" : [ MetricDestination, ... ],
  "SessionSampleRate" : Number,
  "Telemetries" : [ String, ... ]
}

```

### YAML

```yaml

  AllowCookies: Boolean
  EnableXRay: Boolean
  ExcludedPages:
    - String
  FavoritePages:
    - String
  GuestRoleArn: String
  IdentityPoolId: String
  IncludedPages:
    - String
  MetricDestinations:
    - MetricDestination
  SessionSampleRate: Number
  Telemetries:
    - String

```

## Properties

`AllowCookies`

If you set this to `true`, the CloudWatch RUM web client sets two cookies, a session
cookie and a user cookie. The cookies allow the CloudWatch RUM web client to collect data relating to
the number of users an application has and the behavior of the application across a
sequence of events. Cookies are stored in the top-level domain of the current page.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableXRay`

If you set this to `true`, CloudWatch RUM sends client-side traces to
X-Ray for each sampled session. You can then see traces and
segments from these user sessions
in the RUM dashboard and the CloudWatch ServiceLens console. For more information,
see [What is AWS X-Ray?](../../../xray/latest/devguide/aws-xray.md)

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExcludedPages`

A list of URLs in your website or application to exclude from RUM data collection.

You can't include both `ExcludedPages` and `IncludedPages` in the same app monitor.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FavoritePages`

A list of pages in your application that are to be displayed with a "favorite" icon
in the CloudWatch RUM console.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GuestRoleArn`

The ARN of the guest IAM role that is attached to the Amazon Cognito identity pool
that is used to authorize the sending of data to CloudWatch RUM.

_Required_: No

_Type_: String

_Pattern_: `arn:[^:]*:[^:]*:[^:]*:[^:]*:.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdentityPoolId`

The ID of the Amazon Cognito identity pool
that is used to authorize the sending of data to CloudWatch RUM.

_Required_: No

_Type_: String

_Pattern_: `[\w-]+:[0-9a-f-]+`

_Minimum_: `1`

_Maximum_: `55`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludedPages`

If this app monitor is to collect data from only certain pages in your application, this structure lists those pages.

You can't include both `ExcludedPages` and `IncludedPages` in the same app monitor.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricDestinations`

An array of structures that each define a destination that this app monitor will send extended metrics to.

_Required_: No

_Type_: Array of [MetricDestination](aws-properties-rum-appmonitor-metricdestination.md)

_Minimum_: `0`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SessionSampleRate`

Specifies the portion of user sessions to use for CloudWatch RUM data collection. Choosing a higher portion gives you
more data but also incurs more costs.

The range for this value is 0 to 1 inclusive. Setting this to 1 means that 100% of user sessions are sampled, and setting
it to 0.1 means that 10% of user sessions are sampled.

If you omit this parameter, the default of 0.1 is used, and 10% of sessions will be sampled.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Telemetries`

An array that lists the types of telemetry data that this app monitor is to collect.

- `errors` indicates that RUM collects data about unhandled JavaScript errors raised
by your application.

- `performance` indicates that RUM collects performance data about how your application
and its resources are loaded and rendered. This includes Core Web Vitals.

- `http` indicates that RUM collects data about HTTP errors thrown by your application.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::RUM::AppMonitor

CustomEvents

All content copied from https://docs.aws.amazon.com/.
