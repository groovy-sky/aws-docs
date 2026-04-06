This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudTrail::Trail

Creates a trail that specifies the settings for delivery of log data to an Amazon S3 bucket.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudTrail::Trail",
  "Properties" : {
      "AdvancedEventSelectors" : [ AdvancedEventSelector, ... ],
      "AggregationConfigurations" : [ AggregationConfiguration, ... ],
      "CloudWatchLogsLogGroupArn" : String,
      "CloudWatchLogsRoleArn" : String,
      "EnableLogFileValidation" : Boolean,
      "EventSelectors" : [ EventSelector, ... ],
      "IncludeGlobalServiceEvents" : Boolean,
      "InsightSelectors" : [ InsightSelector, ... ],
      "IsLogging" : Boolean,
      "IsMultiRegionTrail" : Boolean,
      "IsOrganizationTrail" : Boolean,
      "KMSKeyId" : String,
      "S3BucketName" : String,
      "S3KeyPrefix" : String,
      "SnsTopicName" : String,
      "Tags" : [ Tag, ... ],
      "TrailName" : String
    }
}

```

### YAML

```yaml

Type: AWS::CloudTrail::Trail
Properties:
  AdvancedEventSelectors:
    - AdvancedEventSelector
  AggregationConfigurations:
    - AggregationConfiguration
  CloudWatchLogsLogGroupArn: String
  CloudWatchLogsRoleArn: String
  EnableLogFileValidation: Boolean
  EventSelectors:
    - EventSelector
  IncludeGlobalServiceEvents: Boolean
  InsightSelectors:
    - InsightSelector
  IsLogging: Boolean
  IsMultiRegionTrail: Boolean
  IsOrganizationTrail: Boolean
  KMSKeyId: String
  S3BucketName: String
  S3KeyPrefix: String
  SnsTopicName: String
  Tags:
    - Tag
  TrailName: String

```

## Properties

`AdvancedEventSelectors`

Specifies the settings for advanced event selectors. You can use advanced event selectors to
log management events, data events for all resource types, and network activity events.

You can add advanced event
selectors, and conditions for your advanced event selectors, up to a maximum of 500 values
for all conditions and selectors on a trail. You can use either
`AdvancedEventSelectors` or `EventSelectors`, but not both. If you
apply `AdvancedEventSelectors` to a trail, any existing
`EventSelectors` are overwritten. For more information about advanced event
selectors, see [Logging data events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html) and
[Logging network activity events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-network-events-with-cloudtrail.html)
in the _AWS CloudTrail User Guide_.

_Required_: No

_Type_: Array of [AdvancedEventSelector](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudtrail-trail-advancedeventselector.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AggregationConfigurations`

Property description not available.

_Required_: No

_Type_: Array of [AggregationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudtrail-trail-aggregationconfiguration.html)

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CloudWatchLogsLogGroupArn`

Specifies a log group name using an Amazon Resource Name (ARN), a unique identifier that
represents the log group to which CloudTrail logs are delivered. You must use a log
group that exists in your account.

To enable CloudWatch Logs delivery, you must provide values for `CloudWatchLogsLogGroupArn` and `CloudWatchLogsRoleArn`.

###### Note

If you previously enabled CloudWatch Logs delivery
and want to disable CloudWatch Logs delivery, you must set the values
of the `CloudWatchLogsRoleArn` and `CloudWatchLogsLogGroupArn` fields to `""`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CloudWatchLogsRoleArn`

Specifies the role for the CloudWatch Logs endpoint to assume to write to a user's
log group. You must use a role that exists in your account.

To enable CloudWatch Logs delivery, you must provide values for `CloudWatchLogsLogGroupArn` and `CloudWatchLogsRoleArn`.

###### Note

If you previously enabled CloudWatch Logs delivery
and want to disable CloudWatch Logs delivery, you must set the values
of the `CloudWatchLogsRoleArn` and `CloudWatchLogsLogGroupArn` fields to `""`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableLogFileValidation`

Specifies whether log file validation is enabled. The default is false.

###### Note

When you disable log file integrity validation, the chain of digest files is broken
after one hour. CloudTrail does not create digest files for log files that were
delivered during a period in which log file integrity validation was disabled. For
example, if you enable log file integrity validation at noon on January 1, disable it at
noon on January 2, and re-enable it at noon on January 10, digest files will not be
created for the log files delivered from noon on January 2 to noon on January 10. The
same applies whenever you stop CloudTrail logging or delete a trail.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventSelectors`

Use event selectors to further specify the management and data event settings for your
trail. By default, trails created without specific event selectors will be configured to
log all read and write management events, and no data events. When an event occurs in your
account, CloudTrail evaluates the event selector for all trails. For each trail, if
the event matches any event selector, the trail processes and logs the event. If the event
doesn't match any event selector, the trail doesn't log the event.

You can configure up to five event selectors for a trail.

You cannot apply both event selectors and advanced event selectors to a trail.

_Required_: No

_Type_: Array of [EventSelector](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudtrail-trail-eventselector.html)

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeGlobalServiceEvents`

Specifies whether the trail is publishing events from global services such as IAM to the log files.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InsightSelectors`

A JSON string that contains the Insights types you want to log on a trail.
`ApiCallRateInsight` and `ApiErrorRateInsight` are valid Insight
types.

The `ApiCallRateInsight` Insights type analyzes write-only
management API calls that are aggregated per minute against a baseline API call volume.

The `ApiErrorRateInsight` Insights type analyzes management
API calls that result in error codes. The error is shown if the API call is
unsuccessful.

_Required_: No

_Type_: Array of [InsightSelector](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudtrail-trail-insightselector.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsLogging`

Whether the CloudTrail trail is currently logging AWS API
calls.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsMultiRegionTrail`

Specifies whether the trail applies only to the current Region or to all Regions. The
default is false. If the trail exists only in the current Region and this value is set to
true, shadow trails (replications of the trail) will be created in the other Regions. If
the trail exists in all Regions and this value is set to false, the trail will remain in
the Region where it was created, and its shadow trails in other Regions will be deleted. As
a best practice, consider using trails that log events in all Regions.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsOrganizationTrail`

Specifies whether the trail is applied to all accounts in an organization in AWS Organizations, or only for the current AWS account. The default is false,
and cannot be true unless the call is made on behalf of an AWS account that
is the management account for an organization in AWS Organizations. If the trail is not an organization trail and this is set to
`true`, the trail will be created in all AWS accounts that
belong to the organization. If the trail is an organization trail and this is set to
`false`, the trail will remain in the current AWS account but
be deleted from all member accounts in the organization.

###### Note

Only the management account for the organization can convert an organization trail to a non-organization trail, or convert a non-organization trail to
an organization trail.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KMSKeyId`

Specifies the AWS KMS key ID to use to encrypt the logs and digest files delivered by CloudTrail. The value can be an alias name prefixed by "alias/", a fully specified ARN to
an alias, a fully specified ARN to a key, or a globally unique identifier.

CloudTrail also supports AWS KMS multi-Region keys. For more
information about multi-Region keys, see [Using multi-Region\
keys](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html) in the _AWS Key Management Service Developer Guide_.

Examples:

- alias/MyAliasName

- arn:aws:kms:us-east-2:123456789012:alias/MyAliasName

- arn:aws:kms:us-east-2:123456789012:key/12345678-1234-1234-1234-123456789012

- 12345678-1234-1234-1234-123456789012

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3BucketName`

Specifies the name of the Amazon S3 bucket designated for publishing log files.
See [Amazon S3\
Bucket naming rules](../../../s3/latest/userguide/bucketnamingrules.md).

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3KeyPrefix`

Specifies the Amazon S3 key prefix that comes after the name of the bucket you
have designated for log file delivery. For more information, see [Finding Your CloudTrail Log Files](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/get-and-view-cloudtrail-log-files.html#cloudtrail-find-log-files). The maximum length is 200
characters.

_Required_: No

_Type_: String

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnsTopicName`

Specifies the name or ARN of the Amazon SNS topic defined for notification of log file
delivery. The maximum length is 256 characters.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A custom set of tags (key-value pairs) for this trail.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudtrail-trail-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrailName`

Specifies the name of the trail. The name must meet the following requirements:

- Contain only ASCII letters (a-z, A-Z), numbers (0-9), periods (.), underscores
(\_), or dashes (-)

- Start with a letter or number, and end with a letter or number

- Be between 3 and 128 characters

- Have no adjacent periods, underscores or dashes. Names like
`my-_namespace` and `my--namespace` are not valid.

- Not be in IP address format (for example, 192.168.5.4)

_Required_: No

_Type_: String

_Pattern_: `(^[a-zA-Z0-9]$)|(^[a-zA-Z0-9]([a-zA-Z0-9\._-])*[a-zA-Z0-9]$)`

_Minimum_: `3`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When the logical ID of this resource is provided to the Ref intrinsic function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the CloudTrail trail, such as `arn:aws:cloudtrail:us-east-2:123456789012:trail/myCloudTrail`.

`SnsTopicArn`

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the Amazon SNS topic that's associated with the CloudTrail trail,
such as `arn:aws:sns:us-east-2:123456789012:mySNSTopic`.

## Examples

You can configure a trail using either `AdvancedEventSelectors` or `EventSelectors`. `EventSelectors` are limited to management events and data events for the following
resource types: `AWS::DynamoDB::Table`, `AWS::Lambda::Function`, and `AWS::S3::Object`. In contrast, you can use `AdvancedEventSelectors` to configure data events
for all available resource types as well as management events.

- [Example: Create a trail by using advanced event selectors](#aws-resource-cloudtrail-trail--examples--Example:_Create_a_trail_by_using_advanced_event_selectors)

- [Example: Create a trail by using event selectors](#aws-resource-cloudtrail-trail--examples--Example:_Create_a_trail_by_using_event_selectors)

### Example: Create a trail by using advanced event selectors

The following example creates a trail that logs events in all regions, an Amazon S3 bucket where logs are published, and
an Amazon SNS topic where notifications are sent. This example uses `AdvancedEventSelectors` to log all read-only and write-only management events,
and data events for all Amazon S3 buckets except the S3 bucket for the trail logs. The bucket and topic policies allow CloudTrail (from
the specified regions) to publish logs to the S3 bucket and to send notifications to an
email that you specify. For information about CloudTrail
bucket policies, see
[Amazon S3 Bucket Policy](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/create_trail_bucket_policy.html) in the _AWS CloudTrail User Guide_.

#### JSON

```json

{
    "Parameters": {
        "TrailName": {
            "Type": "String"
        },
        "BucketName": {
            "Type": "String"
        }
    },
    "Conditions": {
        "IsOrganizationsSupported": {
            "Fn::Equals": [
                {
                    "Ref": "AWS::Partition"
                },
                "aws"
            ]
        }
    },
    "Resources": {
        "Trail": {
            "Type": "AWS::CloudTrail::Trail",
            "Properties": {
                "S3BucketName": {
                    "Ref": "BucketName"
                },
                "S3KeyPrefix": "Uluru",
                "IsLogging": true,
                "TrailName": {
                    "Ref": "TrailName"
                },
                "EnableLogFileValidation": true,
                "IncludeGlobalServiceEvents": true,
                "IsMultiRegionTrail": true,
                "CloudWatchLogsLogGroupArn": {
                    "Fn::ImportValue": "TrailLogGroupTestArn"
                },
                "CloudWatchLogsRoleArn": {
                    "Fn::ImportValue": "TrailLogGroupRoleTestArn"
                },
                "KMSKeyId": {
                    "Fn::ImportValue": "TrailKeyTest"
                },
                "Tags": [
                    {
                        "Key": "TagKeyIntTest",
                        "Value": "TagValueIntTest"
                    },
                    {
                        "Key": "TagKeyIntTest2",
                        "Value": "TagValueIntTest2"
                    }
                ],
                "SnsTopicName": {
                    "Fn::ImportValue": "TrailTopicTest"
                },
                "AdvancedEventSelectors": [
                    {
                        {
                           "Name": "Log management events",
                           "FieldSelectors": [
                              {
                                 "Field": "eventCategory",
                                 "Equals": [
                                       "Management"
                                 ]
                              }
                           ]
                        },
                        {
                           "Name": "Exclude S3 data events for the log bucket",
                           "FieldSelectors": [
                              {
                                 "Field": "eventCategory",
                                 "Equals": [
                                       "Data"
                                 ]
                              },
                              {
                                 "Field": "resources.type",
                                 "Equals": [
                                       "AWS::S3::Object"
                                 ]
                              },
                              {
                                 "Field": "resources.ARN",
                                 "NotStartsWith": [
                                       "LogBucketARN"
                                 ]
                              }
                           ]
                        }
                    }
                ]
            }
        }
    },
    "Outputs": {
        "ARN": {
            "Description": "The trail ARN",
            "Value": {
                "Fn::GetAtt": [
                    "Trail",
                    "Arn"
                ]
            }
        },
        "TopicArn": {
            "Description": "The SnS Topic ARN",
            "Value": {
                "Fn::GetAtt": [
                    "Trail",
                    "SnsTopicArn"
                ]
            }
        }
    }
}
```

#### YAML

```yaml

Parameters:
  TrailName:
    Type: String
  BucketName:
    Type: String
Conditions:
  IsOrganizationsSupported:
    Fn::Equals:
      - { Ref: "AWS::Partition" }
      - "aws"
Resources:
  Trail:
    Type: AWS::CloudTrail::Trail
    Properties:
      S3BucketName: !Ref BucketName
      S3KeyPrefix: "Uluru"
      IsLogging: true
      TrailName: !Ref TrailName
      EnableLogFileValidation: true
      IncludeGlobalServiceEvents: true
      IsMultiRegionTrail: true
      CloudWatchLogsLogGroupArn:
        Fn::ImportValue: "TrailLogGroupTestArn"
      CloudWatchLogsRoleArn:
        Fn::ImportValue: "TrailLogGroupRoleTestArn"
      KMSKeyId:
        Fn::ImportValue: TrailKeyTest
      Tags:
        - Key: "TagKeyIntTest"
          Value: "TagValueIntTest"
        - Key: "TagKeyIntTest2"
          Value: "TagValueIntTest2"
      SnsTopicName:
        Fn::ImportValue: TrailTopicTest
      AdvancedEventSelectors:
        - Name: "Log management events"
          FieldSelectors:
          - Field: "eventCategory"
            Equals: [ "Management" ]
        - Name: "Exclude S3 data events for the log bucket"
          FieldSelectors:
          - Field: "eventCategory"
            Equals: [ "Data" ]
          - Field: "resources.type"
            Equals: "AWS::S3::Object"
          - Field: "resources.ARN"
            NotStartsWith: "LogBucketARN"
Outputs:
  ARN:
    Description: The trail ARN
    Value:
      'Fn::GetAtt':
        - Trail
        - Arn
  TopicArn:
    Description: The SnS Topic ARN
    Value:
      'Fn::GetAtt':
        - Trail
        - SnsTopicArn
```

### Example: Create a trail by using event selectors

The following example creates a trail that logs events in all regions, an Amazon S3 bucket where logs are published, and
an Amazon SNS topic where notifications are sent. This example uses `EventSelectors` to log all read-only and write-only management events,
and data events for Amazon S3 buckets. The bucket and topic policies allow CloudTrail (from
the specified regions) to publish logs to the S3 bucket and to send notifications to an
email that you specify. For information about CloudTrail
bucket policies, see
[Amazon S3 Bucket Policy](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/create_trail_bucket_policy.html) in the _AWS CloudTrail User Guide_.

#### JSON

```json

{
    "Parameters": {
        "TrailName": {
            "Type": "String"
        },
        "BucketName": {
            "Type": "String"
        }
    },
    "Conditions": {
        "IsOrganizationsSupported": {
            "Fn::Equals": [
                {
                    "Ref": "AWS::Partition"
                },
                "aws"
            ]
        }
    },
    "Resources": {
        "Trail": {
            "Type": "AWS::CloudTrail::Trail",
            "Properties": {
                "S3BucketName": {
                    "Ref": "BucketName"
                },
                "S3KeyPrefix": "Uluru",
                "IsLogging": true,
                "TrailName": {
                    "Ref": "TrailName"
                },
                "EnableLogFileValidation": true,
                "IncludeGlobalServiceEvents": true,
                "IsMultiRegionTrail": true,
                "CloudWatchLogsLogGroupArn": {
                    "Fn::ImportValue": "TrailLogGroupTestArn"
                },
                "CloudWatchLogsRoleArn": {
                    "Fn::ImportValue": "TrailLogGroupRoleTestArn"
                },
                "KMSKeyId": {
                    "Fn::ImportValue": "TrailKeyTest"
                },
                "Tags": [
                    {
                        "Key": "TagKeyIntTest",
                        "Value": "TagValueIntTest"
                    },
                    {
                        "Key": "TagKeyIntTest2",
                        "Value": "TagValueIntTest2"
                    }
                ],
                "SnsTopicName": {
                    "Fn::ImportValue": "TrailTopicTest"
                },
                "EventSelectors": [
                    {
                        "DataResources": [
                            {
                                "Type": "AWS::S3::Object",
                                "Values": [
                                    {
                                        "Fn::Sub": "arn:${AWS::Partition}:s3"
                                    }
                                ]
                            }
                        ],
                        "IncludeManagementEvents": true,
                        "ReadWriteType": "All"
                    }
                ]
            }
        }
    },
    "Outputs": {
        "ARN": {
            "Description": "The trail ARN",
            "Value": {
                "Fn::GetAtt": [
                    "Trail",
                    "Arn"
                ]
            }
        },
        "TopicArn": {
            "Description": "The SnS Topic ARN",
            "Value": {
                "Fn::GetAtt": [
                    "Trail",
                    "SnsTopicArn"
                ]
            }
        }
    }
}
```

#### YAML

```yaml

Parameters:
  TrailName:
    Type: String
  BucketName:
    Type: String
Conditions:
  IsOrganizationsSupported:
    Fn::Equals:
      - { Ref: "AWS::Partition" }
      - "aws"
Resources:
  Trail:
    Type: AWS::CloudTrail::Trail
    Properties:
      S3BucketName: !Ref BucketName
      S3KeyPrefix: "Uluru"
      IsLogging: true
      TrailName: !Ref TrailName
      EnableLogFileValidation: true
      IncludeGlobalServiceEvents: true
      IsMultiRegionTrail: true
      CloudWatchLogsLogGroupArn:
        Fn::ImportValue: "TrailLogGroupTestArn"
      CloudWatchLogsRoleArn:
        Fn::ImportValue: "TrailLogGroupRoleTestArn"
      KMSKeyId:
        Fn::ImportValue: TrailKeyTest
      Tags:
        - Key: "TagKeyIntTest"
          Value: "TagValueIntTest"
        - Key: "TagKeyIntTest2"
          Value: "TagValueIntTest2"
      SnsTopicName:
        Fn::ImportValue: TrailTopicTest
      EventSelectors:
        - DataResources:
            - Type: AWS::S3::Object
              Values:
                - !Sub "arn:${AWS::Partition}:s3"
          IncludeManagementEvents: true
          ReadWriteType: All
Outputs:
  ARN:
    Description: The trail ARN
    Value:
      'Fn::GetAtt':
        - Trail
        - Arn
  TopicArn:
    Description: The SnS Topic ARN
    Value:
      'Fn::GetAtt':
        - Trail
        - SnsTopicArn
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::CloudTrail::ResourcePolicy

AdvancedEventSelector
