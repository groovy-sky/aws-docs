This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudTrail::EventDataStore

###### Important

CloudTrail Lake will no longer be open to new customers starting May 31, 2026. If you would like to use CloudTrail Lake, sign up prior to that date. Existing customers can continue to use the service as normal. For more information, see [CloudTrail Lake availability change](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-lake-service-availability-change.html).

Creates a new event data store.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudTrail::EventDataStore",
  "Properties" : {
      "AdvancedEventSelectors" : [ AdvancedEventSelector, ... ],
      "BillingMode" : String,
      "ContextKeySelectors" : [ ContextKeySelector, ... ],
      "FederationEnabled" : Boolean,
      "FederationRoleArn" : String,
      "IngestionEnabled" : Boolean,
      "InsightsDestination" : String,
      "InsightSelectors" : [ InsightSelector, ... ],
      "KmsKeyId" : String,
      "MaxEventSize" : String,
      "MultiRegionEnabled" : Boolean,
      "Name" : String,
      "OrganizationEnabled" : Boolean,
      "RetentionPeriod" : Integer,
      "Tags" : [ Tag, ... ],
      "TerminationProtectionEnabled" : Boolean
    }
}

```

### YAML

```yaml

Type: AWS::CloudTrail::EventDataStore
Properties:
  AdvancedEventSelectors:
    - AdvancedEventSelector
  BillingMode: String
  ContextKeySelectors:
    - ContextKeySelector
  FederationEnabled: Boolean
  FederationRoleArn: String
  IngestionEnabled: Boolean
  InsightsDestination: String
  InsightSelectors:
    - InsightSelector
  KmsKeyId: String
  MaxEventSize: String
  MultiRegionEnabled: Boolean
  Name: String
  OrganizationEnabled: Boolean
  RetentionPeriod: Integer
  Tags:
    - Tag
  TerminationProtectionEnabled: Boolean

```

## Properties

`AdvancedEventSelectors`

The advanced event selectors to use to select the events for the data store. You can
configure up to five advanced event selectors for each event data store.

For more information about how to use advanced event selectors to log CloudTrail
events, see [Log events by using advanced event selectors](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html#creating-data-event-selectors-advanced) in the CloudTrail User Guide.

For more information about how to use advanced event selectors to include AWS Config configuration items in your event data store, see [Create an event data store for AWS Config configuration\
items](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/lake-eds-cli.html#lake-cli-create-eds-config) in the CloudTrail User Guide.

For more information about how to use advanced event selectors to include events outside of AWS events in your event data store, see [Create an integration to log events from outside AWS](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/lake-integrations-cli.html#lake-cli-create-integration) in the CloudTrail User Guide.

_Required_: No

_Type_: Array of [AdvancedEventSelector](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudtrail-eventdatastore-advancedeventselector.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BillingMode`

The billing mode for the event data store determines the cost for ingesting events and the default and maximum retention period for the event data store.

The following are the possible values:

- `EXTENDABLE_RETENTION_PRICING` \- This billing mode is generally recommended if you want a flexible retention period of up to 3653 days (about 10 years).
The default retention period for this billing mode is 366 days.

- `FIXED_RETENTION_PRICING` \- This billing mode is recommended if you expect to ingest more than 25 TB of event data per month and need a retention period of up to 2557 days (about 7 years).
The default retention period for this billing mode is 2557 days.

The default value is `EXTENDABLE_RETENTION_PRICING`.

For more information about CloudTrail pricing,
see [AWS CloudTrail Pricing](https://aws.amazon.com/cloudtrail/pricing) and
[Managing CloudTrail Lake costs](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-lake-manage-costs.html).

_Required_: No

_Type_: String

_Allowed values_: `EXTENDABLE_RETENTION_PRICING | FIXED_RETENTION_PRICING`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContextKeySelectors`

The list of context key selectors that are configured for the event data store.

_Required_: No

_Type_: Array of [ContextKeySelector](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudtrail-eventdatastore-contextkeyselector.html)

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FederationEnabled`

Indicates if [Lake query federation](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/query-federation.html) is enabled. By default, Lake query federation is disabled.
You cannot delete an event data store if Lake query federation is enabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FederationRoleArn`

If Lake query federation is enabled, provides the ARN of the federation role used to access the resources for the federated event data store.

The federation role must exist in your account and provide the [required minimum permissions](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/query-federation.html#query-federation-permissions-role).

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9._/\-:@=\+,\.]+$`

_Minimum_: `3`

_Maximum_: `125`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IngestionEnabled`

Specifies whether the event data store should start ingesting live events. The default is true.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InsightsDestination`

The ARN (or ID suffix of the ARN) of the destination event data store that logs Insights events. For more information, see
[Create an event data store for CloudTrail Insights events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/query-event-data-store-insights.html).

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9._/\-:]+$`

_Minimum_: `3`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InsightSelectors`

A JSON string that contains the Insights types you want to log on an event data store.
`ApiCallRateInsight` and `ApiErrorRateInsight` are valid Insight
types.

The `ApiCallRateInsight` Insights type analyzes write-only
management API calls that are aggregated per minute against a baseline API call volume.

The `ApiErrorRateInsight` Insights type analyzes management
API calls that result in error codes. The error is shown if the API call is
unsuccessful.

_Required_: No

_Type_: Array of [InsightSelector](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudtrail-eventdatastore-insightselector.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyId`

Specifies the AWS KMS key ID to use to encrypt the events delivered by
CloudTrail. The value can be an alias name prefixed by `alias/`, a
fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique
identifier.

###### Important

Disabling or deleting the KMS key, or removing CloudTrail
permissions on the key, prevents CloudTrail from logging events to the event data
store, and prevents users from querying the data in the event data store that was
encrypted with the key. After you associate an event data store with a KMS key, the KMS key cannot be removed or changed. Before you
disable or delete a KMS key that you are using with an event data store,
delete or back up your event data store.

CloudTrail also supports AWS KMS multi-Region keys. For more
information about multi-Region keys, see [Using multi-Region\
keys](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html) in the _AWS Key Management Service Developer Guide_.

Examples:

- `alias/MyAliasName`

- `arn:aws:kms:us-east-2:123456789012:alias/MyAliasName`

- `arn:aws:kms:us-east-2:123456789012:key/12345678-1234-1234-1234-123456789012`

- `12345678-1234-1234-1234-123456789012`

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9._/\-:]+$`

_Minimum_: `1`

_Maximum_: `350`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxEventSize`

The maximum allowed size for events to be stored in the specified event data store. If you are using context key selectors, MaxEventSize must be set to Large.

_Required_: No

_Type_: String

_Allowed values_: `Standard | Large`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MultiRegionEnabled`

Specifies whether the event data store includes events from all Regions, or only from
the Region in which the event data store is created.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the event data store.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9._\-]+$`

_Minimum_: `3`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OrganizationEnabled`

Specifies whether an event data store collects events logged for an organization in
AWS Organizations.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetentionPeriod`

The retention period of the event data store, in days. If `BillingMode` is set to `EXTENDABLE_RETENTION_PRICING`, you can set a retention period of
up to 3653 days, the equivalent of 10 years. If `BillingMode` is set to `FIXED_RETENTION_PRICING`, you can set a retention period of
up to 2557 days, the equivalent of seven years.

CloudTrail Lake determines whether to retain an event by checking if the `eventTime`
of the event is within the specified retention period. For example, if you set a retention period of 90 days, CloudTrail will remove events
when the `eventTime` is older than 90 days.

###### Note

If you plan to copy trail events to this event data store, we recommend
that you consider both the age of the events that you
want to copy as well as how long you want to keep the copied events
in your event data store. For example, if you copy trail events that are 5 years old
and specify a retention period of 7 years, the event data store
will retain those events for two years.

_Required_: No

_Type_: Integer

_Minimum_: `7`

_Maximum_: `3653`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of tags.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudtrail-eventdatastore-tag.html)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TerminationProtectionEnabled`

Specifies whether termination protection is enabled for the event data store. If
termination protection is enabled, you cannot delete the event data store until termination
protection is disabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When the logical ID of this resource is provided to the Ref intrinsic function, `Ref` returns the resource name.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedTimestamp`

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the time stamp of the creation of the event data store, such as `1248496624`.

`EventDataStoreArn`

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the CloudTrail event data store, such as
`arn:aws:cloudtrail:us-east-1:12345678910:eventdatastore/EXAMPLE-f852-4e8f-8bd1-bcf6cEXAMPLE`.

`Status`

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the status of the event data store, such as `ENABLED`.

`UpdatedTimestamp`

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the time stamp that updates were made to an event data store, such as `1598296624`.

## Examples

- [Example: Create an event data store for management events](#aws-resource-cloudtrail-eventdatastore--examples--Example:_Create_an_event_data_store_for_management_events)

- [Example: Create a destination event data store to collect Insights events](#aws-resource-cloudtrail-eventdatastore--examples--Example:_Create_a_destination_event_data_store_to_collect_Insights_events)

- [Example: Create a source event data store that enables Insights events](#aws-resource-cloudtrail-eventdatastore--examples--Example:_Create_a_source_event_data_store_that_enables_Insights_events)

### Example: Create an event data store for management events

The following example creates an event data store that logs management events in all Regions.
For information about creating an event data store for management events, see
[Create an event data store for CloudTrail events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/query-event-data-store-cloudtrail.html)
in the _AWS CloudTrail User Guide_.

In this example, `FederationEnabled` is set to `true` to enable [Lake query federation](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/query-federation.html) and `FederationRoleArn` is set to the
ARN of the federation role used to access the federated resources. The federation role must exist in your account and provide the [required minimum permissions](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/query-federation.html#query-federation-permissions-role).

#### JSON

```json

{
    "Parameters": {
        "Name": {
            "Type": "String"
        }
    },
    "Conditions": {
        "IsOrganizationsSupported": {
            "Fn::Not": [
                {
                    "Fn::Equals": [
                        {
                            "Ref": "AWS::Partition"
                        },
                        "aws-cn"
                    ]
                }
            ]
        }
    },
    "Resources": {
        "myEventDataStore": {
            "Type": "AWS::CloudTrail::EventDataStore",
            "Properties": {
                "Name": {
                    "Ref": "Name"
                },
                "MultiRegionEnabled": true,
                "IngestionEnabled": true,
                "FederationEnabled": true,
                "FederationRoleArn": "arn:aws:iam::account-id:role/federation-role-name",
                "BillingMode": "EXTENDABLE_RETENTION_PRICING",
                "RetentionPeriod": 366,
                "OrganizationEnabled": {
                    "Fn::If": [
                        "IsOrganizationsSupported",
                        true,
                        {
                            "Ref": "AWS::NoValue"
                        }
                    ]
                },
                "TerminationProtectionEnabled": false,
                "KmsKeyId": {
                    "Fn::ImportValue": "EventDataStoreKeyTest"
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
                "AdvancedEventSelectors": [
                    {
                        "Name": "AdvancedEventSelectors1",
                        "FieldSelectors": [
                            {
                                "Field": "eventCategory",
                                "Equals": [
                                    "Management"
                                ]
                            }
                        ]
                    }
                ]
            }
        }
    },
    "Outputs": {
        "myEventDataStoreArn": {
            "Description": "The event data store ARN",
            "Value": {
                "Fn::GetAtt": [
                    "myEventDataStore",
                    "EventDataStoreArn"
                ]
            }
        }
    }
}
```

#### YAML

```yaml

Parameters:
  Name:
    Type: String
Conditions:
  IsOrganizationsSupported:
    !Not [!Equals [Ref: "AWS::Partition", "aws-cn"]]
Resources:
  myEventDataStore:
    Type: AWS::CloudTrail::EventDataStore
    Properties:
      Name: !Ref Name
      MultiRegionEnabled: true
      IngestionEnabled: true
      FederationEnabled: true
      FederationRoleArn: "arn:aws:iam::account-id:role/federation-role-name"
      BillingMode: "EXTENDABLE_RETENTION_PRICING"
      RetentionPeriod: 366
      OrganizationEnabled:
        Fn::If:
          - IsOrganizationsSupported
          - true
          - !Ref "AWS::NoValue"
      TerminationProtectionEnabled: false
      KmsKeyId:
        Fn::ImportValue: EventDataStoreKeyTest
      Tags:
        - Key: "TagKeyIntTest"
          Value: "TagValueIntTest"
        - Key: "TagKeyIntTest2"
          Value: "TagValueIntTest2"
      AdvancedEventSelectors:
        - Name: "AdvancedEventSelectors1"
          FieldSelectors:
          - Field: "eventCategory"
            Equals: [ "Management" ]
Outputs:
  myEventDataStoreArn:
    Description: The eventdatastore ARN
    Value:
      'Fn::GetAtt':
        - myEventDataStore
        - EventDataStoreArn

```

### Example: Create a destination event data store to collect Insights events

The following example creates a destination event data store that collects Insights on unusual management event activity in the source event data store.
For information about creating event data stores for CloudTrail Insights, see
[Create an event data store for CloudTrail Insights events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/query-event-data-store-insights.html)
in the _AWS CloudTrail User Guide_.

#### JSON

```json

{
    "Parameters": {
        "Name": {
            "Type": "String"
        }
    },
    "Conditions": {
        "IsOrganizationsSupported": {
            "Fn::Not": [
                {
                    "Fn::Equals": [
                        {
                            "Ref": "AWS::Partition"
                        },
                        "aws-cn"
                    ]
                }
            ]
        }
    },
    "Resources": {
        "myEventDataStore": {
            "Type": "AWS::CloudTrail::EventDataStore",
            "Properties": {
                "Name": {
                    "Ref": "Name"
                },
                "MultiRegionEnabled": false,
                "IngestionEnabled": true,
                "BillingMode": "EXTENDABLE_RETENTION_PRICING",
                "RetentionPeriod": 90,
                "OrganizationEnabled": {
                    "Fn::If": [
                        "IsOrganizationsSupported",
                        true,
                        {
                            "Ref": "AWS::NoValue"
                        }
                    ]
                },
                "TerminationProtectionEnabled": false,
                "KmsKeyId": {
                    "Fn::ImportValue": "EventDataStoreKeyTest"
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
                "AdvancedEventSelectors": [
                    {
                        "Name": "AdvancedEventSelectors1",
                        "FieldSelectors": [
                            {
                                "Field": "eventCategory",
                                "Equals": [
                                    "Insight"
                                ]
                            }
                        ]
                    }
                ]
            }
        }
    },
    "Outputs": {
        "myEventDataStoreArn": {
            "Description": "The event data store ARN",
            "Value": {
                "Fn::GetAtt": [
                    "myEventDataStore",
                    "EventDataStoreArn"
                ]
            }
        }
    }
}
```

#### YAML

```yaml

Parameters:
  Name:
    Type: String
Conditions:
  IsOrganizationsSupported:
    !Not [!Equals [Ref: "AWS::Partition", "aws-cn"]]
Resources:
  myEventDataStore:
    Type: AWS::CloudTrail::EventDataStore
    Properties:
      Name: !Ref Name
      MultiRegionEnabled: false
      IngestionEnabled: true
      BillingMode: "EXTENDABLE_RETENTION_PRICING"
      RetentionPeriod: 90
      OrganizationEnabled:
        Fn::If:
          - IsOrganizationsSupported
          - true
          - !Ref "AWS::NoValue"
      TerminationProtectionEnabled: false
      KmsKeyId:
        Fn::ImportValue: EventDataStoreKeyTest
      Tags:
        - Key: "TagKeyIntTest"
          Value: "TagValueIntTest"
        - Key: "TagKeyIntTest2"
          Value: "TagValueIntTest2"
      AdvancedEventSelectors:
        - Name: "AdvancedEventSelectors1"
          FieldSelectors:
          - Field: "eventCategory"
            Equals: [ "Insight" ]
Outputs:
  myEventDataStoreArn:
    Description: The eventdatastore ARN
    Value:
      'Fn::GetAtt':
        - myEventDataStore
        - EventDataStoreArn

```

### Example: Create a source event data store that enables Insights events

The following example creates a source event data store that logs management events and enables Insights. Specify the ARN (or ID suffix of the ARN)
of the destination event data store created in the preceding example as the value of the `InsighsDestination` parameter.
For information about creating event data stores for CloudTrail Insights, see
[Create an event data store for CloudTrail Insights events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/query-event-data-store-insights.html)
in the _AWS CloudTrail User Guide_.

#### JSON

```json

{
    "Parameters": {
        "Name": {
            "Type": "String"
        }
    },
    "Conditions": {
        "IsOrganizationsSupported": {
            "Fn::Not": [
                {
                    "Fn::Equals": [
                        {
                            "Ref": "AWS::Partition"
                        },
                        "aws-cn"
                    ]
                }
            ]
        }
    },
    "Resources": {
        "myEventDataStore": {
            "Type": "AWS::CloudTrail::EventDataStore",
            "Properties": {
                "Name": {
                    "Ref": "Name"
                },
                "MultiRegionEnabled": true,
                "IngestionEnabled": true,
                "BillingMode": "EXTENDABLE_RETENTION_PRICING",
                "RetentionPeriod": 90,
                "OrganizationEnabled": {
                    "Fn::If": [
                        "IsOrganizationsSupported",
                        true,
                        {
                            "Ref": "AWS::NoValue"
                        }
                    ]
                },
                "TerminationProtectionEnabled": false,
                "KmsKeyId": {
                    "Fn::ImportValue": "EventDataStoreKeyTest"
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
                "InsightsDestination": "DestinationEventDataStoreArn",
                "InsightSelectors": [
                  {
                        "InsightType": "ApiCallRateInsight"
                  },
                  {
                        "InsightType": "ApiErrorRateInsight"
                  }
               ],
                "AdvancedEventSelectors": [
                    {
                        "Name": "AdvancedEventSelectors1",
                        "FieldSelectors": [
                            {
                                "Field": "eventCategory",
                                "Equals": [
                                    "Management"
                                ]
                            }
                        ]
                    }
                ]
            }
        }
    },
    "Outputs": {
        "myEventDataStoreArn": {
            "Description": "The event data store ARN",
            "Value": {
                "Fn::GetAtt": [
                    "myEventDataStore",
                    "EventDataStoreArn"
                ]
            }
        }
    }
}
```

#### YAML

```yaml

Parameters:
  Name:
    Type: String
Conditions:
  IsOrganizationsSupported:
    !Not [!Equals [Ref: "AWS::Partition", "aws-cn"]]
Resources:
  myEventDataStore:
    Type: AWS::CloudTrail::EventDataStore
    Properties:
      Name: !Ref Name
      MultiRegionEnabled: true
      IngestionEnabled: true
      BillingMode: "EXTENDABLE_RETENTION_PRICING"
      RetentionPeriod: 90
      OrganizationEnabled:
        Fn::If:
          - IsOrganizationsSupported
          - true
          - !Ref "AWS::NoValue"
      TerminationProtectionEnabled: false
      KmsKeyId:
        Fn::ImportValue: EventDataStoreKeyTest
      Tags:
        - Key: "TagKeyIntTest"
          Value: "TagValueIntTest"
        - Key: "TagKeyIntTest2"
          Value: "TagValueIntTest2"
      InsightsDestination: DestinationEventDataStoreArn
      InsightSelectors:
         - InsightType: "ApiCallRateInsight"
         - InsightType: "ApiErrorRateInsight"
      AdvancedEventSelectors:
        - Name: "AdvancedEventSelectors1"
          FieldSelectors:
          - Field: "eventCategory"
            Equals: [ "Management" ]
Outputs:
  myEventDataStoreArn:
    Description: The eventdatastore ARN
    Value:
      'Fn::GetAtt':
        - myEventDataStore
        - EventDataStoreArn

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Widget

AdvancedEventSelector
