This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RTBFabric::Link ResponderErrorMaskingForHttpCode

Describes the masking for HTTP error codes.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : String,
  "HttpCode" : String,
  "LoggingTypes" : [ String, ... ],
  "ResponseLoggingPercentage" : Number
}

```

### YAML

```yaml

  Action: String
  HttpCode: String
  LoggingTypes:
    - String
  ResponseLoggingPercentage: Number

```

## Properties

`Action`

The action for the error..

_Required_: Yes

_Type_: [String](aws-properties-rtbfabric-link-action.md)

_Allowed values_: `NO_BID | PASSTHROUGH`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`HttpCode`

The HTTP error code.

_Required_: Yes

_Type_: String

_Pattern_: `^DEFAULT|4XX|5XX|\d{3}$`

_Minimum_: `3`

_Maximum_: `7`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`LoggingTypes`

The error log type.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `2`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`ResponseLoggingPercentage`

The percentage of response logging.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OpenRtbAttributeModuleParameters

Tag

All content copied from https://docs.aws.amazon.com/.
