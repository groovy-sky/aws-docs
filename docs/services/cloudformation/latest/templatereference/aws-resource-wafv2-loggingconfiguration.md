This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::LoggingConfiguration

Defines an association between logging destinations and a web ACL resource, for logging
from AWS WAF. As part of the association, you can specify parts of the standard logging
fields to keep out of the logs and you can specify filters so that you log only a subset of
the logging records.

If you configure data protection for the web ACL, the protection applies to the data that AWS WAF sends to the logs.

###### Note

You can define one logging destination per web ACL.

You can access information about the traffic that AWS WAF inspects using the following
steps:

1. Create your logging destination. You can use an Amazon CloudWatch Logs log group, an Amazon Simple Storage Service (Amazon S3) bucket, or an Amazon Kinesis Data Firehose.

The name that you give the destination must start with `aws-waf-logs-`. Depending on the type of destination, you might need to configure additional settings or permissions.

For configuration requirements and pricing information for each destination type, see
    [Logging web ACL traffic](https://docs.aws.amazon.com/waf/latest/developerguide/logging.html)
    in the _AWS WAF Developer Guide_.

2. Associate your logging destination to your web ACL using a
    `PutLoggingConfiguration` request.

When you successfully enable logging using a `PutLoggingConfiguration`
request, AWS WAF creates an additional role or policy that is required to write
logs to the logging destination. For an Amazon CloudWatch Logs log group, AWS WAF creates a resource policy on the log group.
For an Amazon S3 bucket, AWS WAF creates a bucket policy. For an Amazon Kinesis Data Firehose, AWS WAF creates a service-linked role.

For additional information about web ACL logging, see
[Logging web ACL traffic information](https://docs.aws.amazon.com/waf/latest/developerguide/logging.html)
in the _AWS WAF Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::WAFv2::LoggingConfiguration",
  "Properties" : {
      "LogDestinationConfigs" : [ String, ... ],
      "LoggingFilter" : LoggingFilter,
      "RedactedFields" : [ FieldToMatch, ... ],
      "ResourceArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::WAFv2::LoggingConfiguration
Properties:
  LogDestinationConfigs:
    - String
  LoggingFilter:
    LoggingFilter
  RedactedFields:
    - FieldToMatch
  ResourceArn: String

```

## Properties

`LogDestinationConfigs`

The logging destination configuration that you want to associate with the web
ACL.

###### Note

You can associate one logging destination to a web ACL.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoggingFilter`

Filtering that specifies which web requests are kept in the logs and which are dropped.
You can filter on the rule action and on the web request labels that were applied by
matching rules during web ACL evaluation.

_Required_: No

_Type_: [LoggingFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wafv2-loggingconfiguration-loggingfilter.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RedactedFields`

The parts of the request that you want to keep out of the logs.

For example, if you
redact the `SingleHeader` field, the `HEADER` field in the logs will
be `REDACTED` for all rules that use the `SingleHeader` `FieldToMatch` setting.

If you configure data protection for the web ACL, the protection applies to the data that AWS WAF sends to the logs.

Redaction applies only to the component that's specified in the rule's `FieldToMatch` setting, so the `SingleHeader` redaction
doesn't apply to rules that use the `Headers` `FieldToMatch`.

###### Note

You can specify only the following fields for redaction: `UriPath`,
`QueryString`, `SingleHeader`, and `Method`.

###### Note

This setting has no impact on request sampling. You can only exclude fields from request sampling by disabling sampling in the web ACL visibility configuration
or by configuring data protection for the web ACL.

_Required_: No

_Type_: Array of [FieldToMatch](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wafv2-loggingconfiguration-fieldtomatch.html)

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceArn`

The Amazon Resource Name (ARN) of the web ACL that you want to associate with
`LogDestinationConfigs`.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the web ACL.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ManagedByFirewallManager`

Indicates whether the logging configuration was created by AWS Firewall Manager, as
part of an AWS WAF policy configuration. If true, only Firewall Manager can modify or delete the configuration.

## Examples

- [Create a logging configuration with redacted fields](#aws-resource-wafv2-loggingconfiguration--examples--Create_a_logging_configuration_with_redacted_fields)

- [Create a logging configuration with a filter](#aws-resource-wafv2-loggingconfiguration--examples--Create_a_logging_configuration_with_a_filter)

### Create a logging configuration with redacted fields

The following shows an example logging configuration with redacted fields.

#### YAML

```yaml

  LoggingConfiguration:
    Type: AWS::WAFv2::LoggingConfiguration
    Properties:
      ResourceArn: arn:aws:wafv2:us-east-1:123456789000:regional/webacl/test-webaclv2/abcd1234-123b-1234-1234-123456789abc
      LogDestinationConfigs:
      - arn:aws:firehose:us-east-1:123456789000:deliverystream/aws-waf-logs-firehose
      RedactedFields:
      - Method: {}
      - QueryString: {}
      - SingleHeader:
          Name: password
      - UriPath: {}
```

#### JSON

```json

  "LoggingConfiguration": {
    "Type": "AWS::WAFv2::LoggingConfiguration",
    "Properties": {
      "ResourceArn": "arn:aws:wafv2:us-east-1:123456789000:regional/webacl/test-webaclv2/abcd1234-123b-1234-1234-123456789abc",
      "LogDestinationConfigs": [
        "arn:aws:firehose:us-east-1:123456789000:deliverystream/aws-waf-logs-firehose"
      ],
      "RedactedFields": [
      {
        "Method": {}
      },
      {
        "QueryString": {}
      },
      {
        "SingleHeader": {
          "Name": "password"
        }
      },
      {
        "UriPath": {}
      } ]
    }
  }
```

### Create a logging configuration with a filter

The following shows an example logging configuration with a logging filter.

#### YAML

```yaml

  LoggingConfiguration:
    Type: AWS::WAFv2::LoggingConfiguration
    Properties:
      ResourceArn: arn:aws:wafv2:us-east-1:123456789000:regional/webacl/test-webaclv2/abcd1234-123b-1234-1234-123456789abc
      LogDestinationConfigs:
      - arn:aws:firehose:us-east-1:123456789000:deliverystream/aws-waf-logs-firehose
      RedactedFields:
      - Method: {}
      - QueryString: {}
      - SingleHeader:
          Name: password
      - UriPath: {}
      LoggingFilter:
        DefaultBehavior: KEEP
        Filters:
        - Behavior: KEEP
          Conditions:
          - ActionCondition:
              Action: BLOCK
          - LabelNameCondition:
              LabelName: fooLabel
          Requirement: MEETS_ANY
```

#### JSON

```json

  "LoggingConfiguration": {
      "Type": "AWS::WAFv2::LoggingConfiguration",
      "Properties": {
        "ResourceArn": "arn:aws:wafv2:us-east-1:123456789000:regional/webacl/test-webaclv2/abcd1234-123b-1234-1234-123456789abc",
        "LogDestinationConfigs": [
          "arn:aws:firehose:us-east-1:123456789000:deliverystream/aws-waf-logs-firehose"
        ],
        "RedactedFields": [
          {
            "Method": {}
          },
          {
            "QueryString": {}
          },
          {
            "SingleHeader": {
              "Name": "password"
            }
          },
          {
            "UriPath": {}
          }
        ],
        "LoggingFilter": {
          "DefaultBehavior": "KEEP",
          "Filters": [
            {
              "Behavior": "KEEP",
              "Conditions": [
                {
                  "ActionCondition": {
                    "Action": "BLOCK"
                  }
                },
                {
                  "LabelNameCondition": {
                    "LabelName": "fooLabel"
                  }
                }
              ],
              "Requirement": "MEETS_ANY"
            }
          ]
        }
      }
    }
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

ActionCondition
