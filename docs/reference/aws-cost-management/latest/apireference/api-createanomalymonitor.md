# CreateAnomalyMonitor

Creates a new cost anomaly detection monitor with the requested type and monitor
specification.

## Request Syntax

```nohighlight

{
   "AnomalyMonitor": {
      "CreationDate": "string",
      "DimensionalValueCount": number,
      "LastEvaluatedDate": "string",
      "LastUpdatedDate": "string",
      "MonitorArn": "string",
      "MonitorDimension": "string",
      "MonitorName": "string",
      "MonitorSpecification": {
         "And": [
            "Expression"
         ],
         "CostCategories": {
            "Key": "string",
            "MatchOptions": [ "string" ],
            "Values": [ "string" ]
         },
         "Dimensions": {
            "Key": "string",
            "MatchOptions": [ "string" ],
            "Values": [ "string" ]
         },
         "Not": "Expression",
         "Or": [
            "Expression"
         ],
         "Tags": {
            "Key": "string",
            "MatchOptions": [ "string" ],
            "Values": [ "string" ]
         }
      },
      "MonitorType": "string"
   },
   "ResourceTags": [
      {
         "Key": "string",
         "Value": "string"
      }
   ]
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[AnomalyMonitor](#API_CreateAnomalyMonitor_RequestSyntax)**

The cost anomaly detection monitor object that you want to create.

Type: [AnomalyMonitor](api-anomalymonitor.md) object

Required: Yes

**[ResourceTags](#API_CreateAnomalyMonitor_RequestSyntax)**

An optional list of tags to associate with the specified [`AnomalyMonitor`](api-anomalymonitor.md). You can use resource tags to control access to your
`monitor` using IAM policies.

Each tag consists of a key and a value, and each key must be unique for the resource. The
following restrictions apply to resource tags:

- Although the maximum number of array members is 200, you can assign a maximum of 50
user-tags to one resource. The remaining are reserved for AWS use

- The maximum length of a key is 128 characters

- The maximum length of a value is 256 characters

- Keys and values can only contain alphanumeric characters, spaces, and any of the
following: `_.:/=+@-`

- Keys and values are case sensitive

- Keys and values are trimmed for any leading or trailing whitespaces

- Don’t use `aws:` as a prefix for your keys. This prefix is reserved for
AWS use

Type: Array of [ResourceTag](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_ResourceTag.html) objects

Array Members: Minimum number of 0 items. Maximum number of 200 items.

Required: No

## Response Syntax

```nohighlight

{
   "MonitorArn": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[MonitorArn](#API_CreateAnomalyMonitor_ResponseSyntax)**

The unique identifier of your newly created cost anomaly detection monitor.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

Pattern: `[\S\s]*`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**LimitExceededException**

You made too many calls in a short period of time. Try again later.

HTTP Status Code: 400

## Examples

The following are sample requests and responses of the
`CreateAnomalyMonitor` operation.

### Example 1

The following example for the `CreateAnomalyMonitor` operation creates
an AWS managed services monitor with resource tags.

#### Sample Request

```

{
    "AnomalyMonitor": {
        "MonitorName": "Monitor name",
        "MonitorType": "DIMENSIONAL",
        "MonitorDimension": "SERVICE"
    },
    "ResourceTags": [
        {
            "Key": "tag key",
            "Value": "tag value"
        }
    ]
}
```

#### Sample Response

```

{
    "MonitorArn": "arn:${partition}:ce::${account-id}:anomalymonitor/${monitor-id}"
}
```

### Example 2

The following example for the `CreateAnomalyMonitor` operation creates
an AWS managed linked account monitor.

#### Sample Request

```

{
    "AnomalyMonitor": {
        "MonitorName": "Monitor name",
        "MonitorType": "DIMENSIONAL",
        "MonitorDimension": "LINKED_ACCOUNT"
    }
}
```

#### Sample Response

```

{
    "MonitorArn": "arn:${partition}:ce::${account-id}:anomalymonitor/${monitor-id}"
}
```

### Example 3

The following example for the `CreateAnomalyMonitor` operation creates
an AWS managed cost allocation tag monitor.

#### Sample Request

```

{
    "AnomalyMonitor": {
        "MonitorName": "Monitor name",
        "MonitorType": "DIMENSIONAL",
        "MonitorDimension": "TAG",
        "MonitorSpecification": {
            "Tags": {
                "Key": "tag key"
            }
        }
    }
}
```

#### Sample Response

```

{
    "MonitorArn": "arn:${partition}:ce::${account-id}:anomalymonitor/${monitor-id}"
}
```

### Example 4

The following example for the `CreateAnomalyMonitor` operation creates
an AWS managed cost category monitor.

#### Sample Request

```

{
    "AnomalyMonitor": {
        "MonitorName": "Monitor name",
        "MonitorType": "DIMENSIONAL",
        "MonitorDimension": "COST_CATEGORY",
        "MonitorSpecification": {
            "CostCategories": {
                "Key": "cost category key"
            }
        }
    }
}
```

#### Sample Response

```

{
    "MonitorArn": "arn:${partition}:ce::${account-id}:anomalymonitor/${monitor-id}"
}
```

### Example 5

The following example for the `CreateAnomalyMonitor` operation creates
a customer managed linked account monitor. You can track up to ten linked accounts with a
single monitor.

#### Sample Request

```

{
    "AnomalyMonitor": {
        "MonitorName": "Monitor name",
        "MonitorType": "CUSTOM",
        "MonitorSpecification": {
            "Dimensions": {
                "Key": "LINKED_ACCOUNT",
                "Values": [
                    "000011112222",
                    "111122223333"
                ]
            }
        }
    }
}
```

#### Sample Response

```

{
    "MonitorArn": "arn:${partition}:ce::${account-id}:anomalymonitor/${monitor-id}"
}
```

### Example 6

The following example for the `CreateAnomalyMonitor` operation creates
a customer managed cost allocation tag monitor. You can track up to ten tag values with
a single monitor.

#### Sample Request

```

{
    "AnomalyMonitor": {
        "MonitorName": "Monitor name",
        "MonitorType": "CUSTOM",
        "MonitorSpecification": {
            "Tags": {
                "Key": "tag key",
                "Values": [
                    "tag value 1",
                    "tag value 2"
                ]
            }
        }
    }
}
```

#### Sample Response

```

{
    "MonitorArn": "arn:${partition}:ce::${account-id}:anomalymonitor/${monitor-id}"
}
```

### Example 7

The following example for the `CreateAnomalyMonitor` operation creates
a customer managed cost category monitor. You can track one value with a single monitor.

#### Sample Request

```

{
    "AnomalyMonitor": {
        "MonitorName": "Monitor name",
        "MonitorType": "CUSTOM",
        "MonitorSpecification": {
            "CostCategories": {
                "Key": "cost category key",
                "Values": [ "cost category value" ]
            }
        }
    }
}
```

#### Sample Response

```

{
    "MonitorArn": "arn:${partition}:ce::${account-id}:anomalymonitor/${monitor-id}"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ce-2017-10-25/CreateAnomalyMonitor)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ce-2017-10-25/CreateAnomalyMonitor)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ce-2017-10-25/CreateAnomalyMonitor)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ce-2017-10-25/CreateAnomalyMonitor)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ce-2017-10-25/CreateAnomalyMonitor)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ce-2017-10-25/CreateAnomalyMonitor)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ce-2017-10-25/CreateAnomalyMonitor)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ce-2017-10-25/CreateAnomalyMonitor)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ce-2017-10-25/CreateAnomalyMonitor)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ce-2017-10-25/CreateAnomalyMonitor)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS Cost Explorer

CreateAnomalySubscription
