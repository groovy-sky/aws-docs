This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::LogGroup

The `AWS::Logs::LogGroup` resource specifies a log group. A log group
defines common properties for log streams, such as their retention and access control
rules. Each log stream must belong to one log group.

You can create up to 1,000,000 log groups per Region per account. You must use the
following guidelines when naming a log group:

- Log group names must be unique within a Region for an AWS
account.

- Log group names can be between 1 and 512 characters long.

- Log group names consist of the following characters: a-z, A-Z, 0-9, '\_'
(underscore), '-' (hyphen), '/' (forward slash), and '.' (period).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Logs::LogGroup",
  "Properties" : {
      "BearerTokenAuthenticationEnabled" : Boolean,
      "DataProtectionPolicy" : Json,
      "DeletionProtectionEnabled" : Boolean,
      "FieldIndexPolicies" : [ {Key: Value, ...}, ... ],
      "KmsKeyId" : String,
      "LogGroupClass" : String,
      "LogGroupName" : String,
      "ResourcePolicyDocument" : Json,
      "RetentionInDays" : Integer,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Logs::LogGroup
Properties:
  BearerTokenAuthenticationEnabled: Boolean
  DataProtectionPolicy: Json
  DeletionProtectionEnabled: Boolean
  FieldIndexPolicies:
    -
    Key: Value
  KmsKeyId: String
  LogGroupClass: String
  LogGroupName: String
  ResourcePolicyDocument: Json
  RetentionInDays: Integer
  Tags:
    - Tag

```

## Properties

`BearerTokenAuthenticationEnabled`

Indicates whether bearer token authentication is enabled for this log group. When enabled,
bearer token authentication is allowed on operations until it is explicitly disabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataProtectionPolicy`

Creates a data protection policy and assigns it to the log group. A data protection
policy can help safeguard sensitive data that's ingested by the log group by auditing
and masking the sensitive log data. When a user who does not have permission to view
masked data views a log event that includes masked data, the sensitive data is replaced
by asterisks.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeletionProtectionEnabled`

Indicates whether deletion protection is enabled for this log group. When enabled,
deletion protection blocks all deletion operations until it is explicitly disabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldIndexPolicies`

Creates or updates a _field index policy_ for the specified log
group. Only log groups in the Standard log class support field index policies. For more
information about log classes, see [Log\
classes](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-log-classes.md).

You can use field index policies to create _field indexes_ on
fields found in log events in the log group. Creating field indexes lowers the costs for
CloudWatch Logs Insights queries that reference those field indexes, because
these queries attempt to skip the processing of log events that are known to not match
the indexed field. Good fields to index are fields that you often need to query for and
fields that have high cardinality of values Common examples of indexes include request
ID, session ID, userID, and instance IDs. For more information, see [Create field\
indexes to improve query performance and reduce costs](../../../amazoncloudwatch/latest/logs/cloudwatchlogs-field-indexing.md).

Currently, this array supports only one field index policy object.

_Required_: No

_Type_: Array of Object

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyId`

The Amazon Resource Name (ARN) of the AWS KMS key to use when encrypting
log data.

To associate an AWS KMS key with the log group, specify the ARN of that KMS
key here. If you do so, ingested data is encrypted using this key. This association is
stored as long as the data encrypted with the KMS key is still within CloudWatch Logs. This enables CloudWatch Logs to decrypt this data whenever it is
requested.

If you attempt to associate a KMS key with the log group but the KMS key doesn't
exist or is deactivated, you will receive an `InvalidParameterException`
error.

Log group data is always encrypted in CloudWatch Logs. If you omit this key, the
encryption does not use AWS KMS. For more information, see [Encrypt log data in CloudWatch Logs using AWS Key Management Service](../../../amazoncloudwatch/latest/logs/encrypt-log-data-kms.md)

_Required_: No

_Type_: String

_Pattern_: `^arn:[a-z0-9-]+:kms:[a-z0-9-]+:\d{12}:(key|alias)/.+\Z`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogGroupClass`

Specifies the log group class for this log group. There are two classes:

- The `Standard` log class supports all CloudWatch Logs
features.

- The `Infrequent Access` log class supports a subset of CloudWatch Logs features and incurs lower costs.

For details about the features supported by each class, see [Log\
classes](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-log-classes.md)

_Required_: No

_Type_: String

_Allowed values_: `STANDARD | INFREQUENT_ACCESS | DELIVERY`

_Update requires_: Updates are not supported.

`LogGroupName`

The name of the log group. If you don't specify a name, AWS CloudFormation
generates a unique ID for the log group.

_Required_: No

_Type_: String

_Pattern_: `^[.\-_/#A-Za-z0-9]{1,512}\Z`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourcePolicyDocument`

Creates or updates a resource policy for the specified log group that allows other
services to put log events to this account. A LogGroup can have 1 resource
policy.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetentionInDays`

The number of days to retain the log events in the specified log group. Possible
values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1096, 1827,
2192, 2557, 2922, 3288, and 3653.

To set a log group so that its log events do not expire, do not specify this
property.

_Required_: No

_Type_: Integer

_Allowed values_: `1 | 3 | 5 | 7 | 14 | 30 | 60 | 90 | 120 | 150 | 180 | 365 | 400 | 545 | 731 | 1096 | 1827 | 2192 | 2557 | 2922 | 3288 | 3653`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to the log group.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-logs-loggroup-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the log group, such as
`arn:aws:logs:us-west-1:123456789012:log-group:/mystack-testgroup-12ABC1AB12A1:*`

## Examples

- [Create a log group](#aws-resource-logs-loggroup--examples--Create_a_log_group)

- [Create a log group with a data protection policy](#aws-resource-logs-loggroup--examples--Create_a_log_group_with_a_data_protection_policy)

- [Create a log group with a field index policy](#aws-resource-logs-loggroup--examples--Create_a_log_group_with_a_field_index_policy)

### Create a log group

The following example creates a log group that retains events for seven
days.

#### JSON

```json

"myLogGroup": {
    "Type": "AWS::Logs::LogGroup",
    "Properties": {
        "RetentionInDays": 7
    }
}
```

#### YAML

```yaml

myLogGroup:
  Type: AWS::Logs::LogGroup
  Properties:
    RetentionInDays: 7
```

### Create a log group with a data protection policy

The following example creates a log group that uses a data protection policy
to mask email addresses, and send audit findings to CloudWatch Logs, Firehose, and Amazon S3.

#### JSON

```json

"TestLogGroupDescription": {
  "Type": "AWS::Logs::LogGroup",
  "Properties": {
      "LogGroupName": "my-log-group",
      "DataProtectionPolicy": {
          "Name": "data-protection-policy",
          "Description": "test description",
          "Version": "2021-06-01",
          "Statement": [{
                  "Sid": "audit-policy test",
                  "DataIdentifier": [
                      "arn:aws:dataprotection::aws:data-identifier/EmailAddress",
                      "arn:aws:dataprotection::aws:data-identifier/DriversLicense-US"
                  ],
                  "Operation": {
                      "Audit": {
                          "FindingsDestination": {
                              "CloudWatchLogs": {
                                  "LogGroup": "EXISTING_LOG_GROUP_IN_YOUR_ACCOUNT"
                              },
                              "Firehose": {
                                  "DeliveryStream": "EXISTING_STREAM_IN_YOUR_ACCOUNT"
                              },
                              "S3": {
                                  "Bucket": "EXISTING_BUCKET"
                              }
                          }
                      }
                  }
              },
              {
                  "Sid": "redact-policy",
                  "DataIdentifier": [
                      "arn:aws:dataprotection::aws:data-identifier/EmailAddress",
                      "arn:aws:dataprotection::aws:data-identifier/DriversLicense-US"
                  ],
                  "Operation": {
                      "Deidentify": {
                          "MaskConfig": {}
                      }
                  }
              }
          ]
      }
  }
}
```

#### YAML

```yaml

TestLogGroupDescription:
  Type: AWS::Logs::LogGroup
  Properties:
    LogGroupName: my-log-group
    DataProtectionPolicy:
      Name: data-protection-policy
      Description: test description
      Version: '2021-06-01'
      Statement:
      - Sid: audit-policy test
        DataIdentifier:
        - arn:aws:dataprotection::aws:data-identifier/EmailAddress
        - arn:aws:dataprotection::aws:data-identifier/DriversLicense-US
        Operation:
          Audit:
            FindingsDestination:
              CloudWatchLogs:
                LogGroup: EXISTING_LOG_GROUP_IN_YOUR_ACCOUNT
              Firehose:
                DeliveryStream: EXISTING_STREAM_IN_YOUR_ACCOUNT
              S3:
                Bucket: EXISTING_BUCKET
      - Sid: redact-policy
        DataIdentifier:
        - arn:aws:dataprotection::aws:data-identifier/EmailAddress
        - arn:aws:dataprotection::aws:data-identifier/DriversLicense-US
        Operation:
          Deidentify:
            MaskConfig: {}
```

### Create a log group with a field index policy

The following example creates a log group that uses a field index policy to
index the `RequestId` and `TransactionId` fields in log
events.

#### JSON

```json

"TestFieldIndexingLogGroup": {
    "Type": "AWS::Logs::LogGroup",
    "Properties": {
        "LogGroupName": "my-log-group",
        "FieldIndexPolicies": [{
            "Fields": [
                "RequestId",
                "TransactionId"
            ]
        }]
    }
}
```

#### YAML

```yaml

TestFieldIndexingLogGroup:
  Type: AWS::Logs::LogGroup
  Properties:
    LogGroupName: my-log-group
    FieldIndexPolicies:
        - Fields:
            - RequestId
            - TransactionId
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Logs::LogAnomalyDetector

Tag

All content copied from https://docs.aws.amazon.com/.
