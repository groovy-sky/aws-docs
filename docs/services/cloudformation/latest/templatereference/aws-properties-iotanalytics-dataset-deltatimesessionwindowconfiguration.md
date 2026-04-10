This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Dataset DeltaTimeSessionWindowConfiguration

A structure that contains the configuration information of a delta time session
window.

[`DeltaTime`](../../../../reference/iotanalytics/latest/apireference/api-deltatime.md) specifies a time interval. You can use
`DeltaTime` to create dataset contents with data that has arrived in the data
store since the last execution. For an example of `DeltaTime`, see [Creating\
a SQL dataset with a delta window (CLI)](../../../iotanalytics/latest/userguide/automate-create-dataset.md#automate-example6) in the
_AWS IoT Analytics User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TimeoutInMinutes" : Integer
}

```

### YAML

```yaml

  TimeoutInMinutes: Integer

```

## Properties

`TimeoutInMinutes`

A time interval. You can use `timeoutInMinutes` so that AWS IoT Analytics can batch up late
data notifications that have been generated since the last execution. AWS IoT Analytics sends one batch of
notifications to Amazon CloudWatch Events at one time.

For more information about how to write a timestamp expression, see [Date and Time Functions and\
Operators](https://prestodb.io/docs/current/functions/datetime.html), in the _Presto 0.172 Documentation_.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `60`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeltaTime

Filter

All content copied from https://docs.aws.amazon.com/.
