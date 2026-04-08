# CreateBudget

Creates a budget and, if included, notifications and subscribers.

###### Important

Only one of `BudgetLimit` or `PlannedBudgetLimits` can be present in
the syntax at one time. Use the syntax that matches your use case. The Request Syntax
section shows the `BudgetLimit` syntax. For `PlannedBudgetLimits`,
see the [Examples](api-budgets-createbudget-api-createbudget-examples.md) section.

Similarly, only one set of filter and metric selections can be present in the syntax
at one time. Either `FilterExpression` and `Metrics` or
`CostFilters` and `CostTypes`, not both or a different
combination. We recommend using `FilterExpression` and `Metrics`
as they provide more flexible and powerful filtering capabilities. The Request Syntax
section shows the `FilterExpression`/ `Metrics` syntax.

## Request Syntax

```nohighlight

{
   "AccountId": "string",
   "Budget": {
      "AutoAdjustData": {
         "AutoAdjustType": "string",
         "HistoricalOptions": {
            "BudgetAdjustmentPeriod": number,
            "LookBackAvailablePeriods": number
         },
         "LastAutoAdjustTime": number
      },
      "BillingViewArn": "string",
      "BudgetLimit": {
         "Amount": "string",
         "Unit": "string"
      },
      "BudgetName": "string",
      "BudgetType": "string",
      "CalculatedSpend": {
         "ActualSpend": {
            "Amount": "string",
            "Unit": "string"
         },
         "ForecastedSpend": {
            "Amount": "string",
            "Unit": "string"
         }
      },
      "CostFilters": {
         "string" : [ "string" ]
      },
      "CostTypes": {
         "IncludeCredit": boolean,
         "IncludeDiscount": boolean,
         "IncludeOtherSubscription": boolean,
         "IncludeRecurring": boolean,
         "IncludeRefund": boolean,
         "IncludeSubscription": boolean,
         "IncludeSupport": boolean,
         "IncludeTax": boolean,
         "IncludeUpfront": boolean,
         "UseAmortized": boolean,
         "UseBlended": boolean
      },
      "FilterExpression": {
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
      "HealthStatus": {
         "LastUpdatedTime": number,
         "Status": "string",
         "StatusReason": "string"
      },
      "LastUpdatedTime": number,
      "Metrics": [ "string" ],
      "PlannedBudgetLimits": {
         "string" : {
            "Amount": "string",
            "Unit": "string"
         }
      },
      "TimePeriod": {
         "End": number,
         "Start": number
      },
      "TimeUnit": "string"
   },
   "NotificationsWithSubscribers": [
      {
         "Notification": {
            "ComparisonOperator": "string",
            "NotificationState": "string",
            "NotificationType": "string",
            "Threshold": number,
            "ThresholdType": "string"
         },
         "Subscribers": [
            {
               "Address": "string",
               "SubscriptionType": "string"
            }
         ]
      }
   ],
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

**[AccountId](#API_budgets_CreateBudget_RequestSyntax)**

The `accountId` that is associated with the budget.

Type: String

Length Constraints: Fixed length of 12.

Pattern: `\d{12}`

Required: Yes

**[Budget](#API_budgets_CreateBudget_RequestSyntax)**

The budget object that you want to create.

Type: [Budget](api-budgets-budget.md) object

Required: Yes

**[NotificationsWithSubscribers](#API_budgets_CreateBudget_RequestSyntax)**

A notification that you want to associate with a budget. A budget can have up to five notifications, and each notification can have one SNS subscriber and up to 10 email subscribers. If you include notifications and subscribers in your `CreateBudget` call, AWS creates the notifications and subscribers for you.

Type: Array of [NotificationWithSubscribers](api-budgets-notificationwithsubscribers.md) objects

Array Members: Maximum number of 10 items.

Required: No

**[ResourceTags](#API_budgets_CreateBudget_RequestSyntax)**

An optional list of tags to associate with the specified budget. Each tag consists of a
key and a value, and each key must be unique for the resource.

Type: Array of [ResourceTag](api-budgets-resourcetag.md) objects

Array Members: Minimum number of 0 items. Maximum number of 200 items.

Required: No

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You are not authorized to use this operation with the given parameters.

**Message**

The error message the exception carries.

HTTP Status Code: 400

**BillingViewHealthStatusException**

The billing view status must be HEALTHY to perform this action. Try again when the status is HEALTHY.

**Message**

The error message the exception carries.

HTTP Status Code: 400

**CreationLimitExceededException**

You've exceeded the notification or subscriber limit.

**Message**

The error message the exception carries.

HTTP Status Code: 400

**DuplicateRecordException**

The budget name already exists. Budget names must be unique within an account.

**Message**

The error message the exception carries.

HTTP Status Code: 400

**InternalErrorException**

An error on the server occurred during the processing of your request. Try again later.

**Message**

The error message the exception carries.

HTTP Status Code: 400

**InvalidParameterException**

An error on the client occurred. Typically, the cause is an invalid input value.

**Message**

The error message the exception carries.

HTTP Status Code: 400

**NotFoundException**

We can’t locate the resource that you specified.

**Message**

The error message the exception carries.

HTTP Status Code: 400

**ServiceQuotaExceededException**

You've reached a Service Quota limit on this resource.

**Message**

The error message the exception carries.

HTTP Status Code: 400

**ThrottlingException**

The number of API requests has exceeded the maximum allowed API request throttling limit
for the account.

**Message**

The error message the exception carries.

HTTP Status Code: 400

## Examples

### Example

The following is the `PlannedBudgetLimits` syntax

```

{
   "AccountId": "string",
   "Budget": {
      "PlannedBudgetLimits": {
         "string": {
            "Amount": "string",
            "Unit": "string"
         }
      },
      "BudgetName": "string",
      "BudgetType": "string",
      "CalculatedSpend": {
         "ActualSpend": {
            "Amount": "string",
            "Unit": "string"
         },
         "ForecastedSpend": {
            "Amount": "string",
            "Unit": "string"
         }
      },
      "FilterExpression": {
         "And": [
            "Expression"
         ],
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
      "Metrics": [ "string" ],
      "LastUpdatedTime": number,
      "TimePeriod": {
         "End": number,
         "Start": number
      },
      "TimeUnit": "string"
   },
   "NotificationsWithSubscribers": [
      {
         "Notification": {
            "ComparisonOperator": "string",
            "NotificationState": "string",
            "NotificationType": "string",
            "Threshold": number,
            "ThresholdType": "string"
         },
         "Subscribers": [
            {
               "Address": "string",
               "SubscriptionType": "string"
            }
         ]
      }
   ]
}
```

### Example

The following is a sample request of the `CreateBudget` operation using `BudgetLimit`

#### Sample Request

```

POST / HTTP/1.1
Host: awsbudgets.<region>.<domain>
x-amz-Date: <Date>
Authorization: AWS4-HMAC-SHA256 Credential=<Credential>, SignedHeaders=contenttype;date;host;user-agent;x-amz-date;x-amz-target;x-amzn-requestid,Signature=<Signature>
User-Agent: <UserAgentString>
Content-Type: application/x-amz-json-1.1
Content-Length: <PayloadSizeBytes>
Connection: Keep-Alive
X-Amz-Target: AWSBudgetServiceGateway.CreateBudget

{
   "AccountId": "111122223333",
   "Budget": {
      "BudgetLimit": {
         "Amount": "100",
         "Unit": "USD"
      },
      "BudgetName": "Example Budget",
      "BudgetType": "COST",
      "FilterExpression": {
         "Dimensions": {
            "Key": "AZ",
            "Values": ["us-east-1"]
         }
      },
      "Metrics": ["UNBLENDED_COST"],
      "TimePeriod": {
         "Start": 1712019600,
         "End": 1714611600
      },
      "TimeUnit": "MONTHLY"
   },
   "NotificationsWithSubscribers": [
      {
         "Notification": {
            "ComparisonOperator": "GREATER_THAN",
            "NotificationType": "ACTUAL",
            "Threshold": 80,
            "ThresholdType": "PERCENTAGE"
         },
         "Subscribers": [
            {
               "Address": "budget-alerts@example.com",
               "SubscriptionType": "EMAIL"
            }
         ]
      },
      {
         "Notification": {
            "ComparisonOperator": "GREATER_THAN",
            "NotificationType": "FORECASTED",
            "Threshold": 90,
            "ThresholdType": "PERCENTAGE"
         },
         "Subscribers": [
            {
               "Address": "budget-forecasts@example.com",
               "SubscriptionType": "EMAIL"
            }
         ]
      }
   ]
}
```

### Example

The following is a sample request of the `CreateBudget` operation using
`PlannedBudgetLimits` and using FilterExpression to include multiple
dimensions (Service and Region) on net unblended costs

#### Sample Request

```

POST / HTTP/1.1
Host: awsbudgets.<region>.<domain>
x-amz-Date: <Date>
Authorization: AWS4-HMAC-SHA256 Credential=<Credential>, SignedHeaders=contenttype;date;host;user-agent;x-amz-date;x-amz-target;x-amzn-requestid,Signature=<Signature>
User-Agent: <UserAgentString>
Content-Type: application/x-amz-json-1.1
Content-Length: <PayloadSizeBytes>
Connection: Keep-Alive
X-Amz-Target: AWSBudgetServiceGateway.CreateBudget

{
   "AccountId": "111122223333",
   "Budget": {
      "PlannedBudgetLimits": {
         "1712019600": {
            "Amount": "100",
            "Unit": "USD"
         },
         "1714611600": {
            "Amount": "200",
            "Unit": "USD"
         },
         "1717203600": {
            "Amount": "300",
            "Unit": "USD"
         }
      },
      "BudgetName": "Example Planned Budget",
      "BudgetType": "COST",
      "FilterExpression": {
         "Dimensions": {
            "Key": "SERVICE",
            "Values": ["Amazon Simple Storage Service", "Amazon FSx"]
         },
         "And": [{
            "Dimensions": {
               "Key": "REGION",
               "Values": ["us-east-1", "us-west-2"]
            }
         }]
      },
      "Metrics": ["NET_UNBLENDED_COST"],
      "TimePeriod": {
         "Start": 1712019600,
         "End": 1717203600
      },
      "TimeUnit": "MONTHLY"
   },
   "NotificationsWithSubscribers": [
      {
         "Notification": {
            "ComparisonOperator": "GREATER_THAN",
            "NotificationType": "ACTUAL",
            "Threshold": 50,
            "ThresholdType": "PERCENTAGE"
         },
         "Subscribers": [
            {
               "Address": "planned-budget-alerts-1@example.com",
               "SubscriptionType": "EMAIL"
            }
         ]
      }
   ]
}
```

### Example

The following is a sample request of the `CreateBudget` operation using `BudgetLimit` and filtering for a specific tag.

#### Sample Request

```

POST / HTTP/1.1
Host: awsbudgets.<region>.<domain>
x-amz-Date: <Date>
Authorization: AWS4-HMAC-SHA256 Credential=<Credential>, SignedHeaders=contenttype;date;host;user-agent;x-amz-date;x-amz-target;x-amzn-requestid,Signature=<Signature>
User-Agent: <UserAgentString>
Content-Type: application/x-amz-json-1.1
Content-Length: <PayloadSizeBytes>
Connection: Keep-Alive
X-Amz-Target: AWSBudgetServiceGateway.CreateBudget

{
   "AccountId": "111122223333",
   "Budget": {
      "BudgetLimit": {
         "Amount": "100",
         "Unit": "USD"
      },
      "BudgetName": "Example Tag Budget",
      "BudgetType": "COST",
      "FilterExpression": {
         "Tags": {
            "Key": "user:Key",
            "Values": ["value1", "value2"]
         }
      },
      "Metrics": ["NET_UNBLENDED_COST"],
      "TimePeriod": {
         "Start": 1712019600,
         "End": 1717203600
      },
      "TimeUnit": "MONTHLY"
   },
   "NotificationsWithSubscribers": [
      {
         "Notification": {
            "ComparisonOperator": "GREATER_THAN",
            "NotificationType": "ACTUAL",
            "Threshold": 80,
            "ThresholdType": "PERCENTAGE"
         },
         "Subscribers": [
            {
               "Address": "example@example.com",
               "SubscriptionType": "EMAIL"
            }
         ]
      }
   ]
}
```

### Example

The following is a sample request of the `CreateBudget` operation using
`ResourceTags` and filtering for a specific tag.

#### Sample Request

```

POST / HTTP/1.1
Host: awsbudgets.<region>.<domain>
x-amz-Date: <Date>
Authorization: AWS4-HMAC-SHA256 Credential=<Credential>, SignedHeaders=contenttype;date;host;user-agent;x-amz-date;x-amz-target;x-amzn-requestid,Signature=<Signature>
User-Agent: <UserAgentString>
Content-Type: application/x-amz-json-1.1
Content-Length: <PayloadSizeBytes>
Connection: Keep-Alive
X-Amz-Target: AWSBudgetServiceGateway.CreateBudget

{
  "AccountId": "111122223333",
  "Budget": {
    "BudgetLimit": {
      "Amount": "1000",
      "Unit": "USD"
    },
    "BudgetName": "Tag-Based Budget Example",
    "BudgetType": "COST",
    "FilterExpression": {
      "Tags": {
        "Key": "CostCenter",
        "Values": ["IT-1234", "Marketing-5678"]
      }
    },
    "Metrics": ["NET_UNBLENDED_COST"],
    "TimePeriod": {
      "Start": 1712019600,
      "End": 1714611600
    },
    "TimeUnit": "MONTHLY"
  },
  "NotificationsWithSubscribers": [
    {
      "Notification": {
        "ComparisonOperator": "GREATER_THAN",
        "NotificationType": "ACTUAL",
        "Threshold": 80,
        "ThresholdType": "PERCENTAGE"
      },
      "Subscribers": [
        {
          "Address": "tag-budget-alerts@example.com",
          "SubscriptionType": "EMAIL"
        }
      ]
    }
  ],
  "ResourceTags": [
    {
      "Key": "Project",
      "Value": "CloudMigration"
    },
    {
      "Key": "Owner",
      "Value": "FinanceTeam"
    }
  ]
}
```

### Example

The following is a sample request of the `CreateBudget` operation using
`FilterExpression` to exclude specific services.

#### Sample Request

```

POST / HTTP/1.1
Host: awsbudgets.<region>.<domain>
x-amz-Date: <Date>
Authorization: AWS4-HMAC-SHA256 Credential=<Credential>, SignedHeaders=contenttype;date;host;user-agent;x-amz-date;x-amz-target;x-amzn-requestid,Signature=<Signature>
User-Agent: <UserAgentString>
Content-Type: application/x-amz-json-1.1
Content-Length: <PayloadSizeBytes>
Connection: Keep-Alive
X-Amz-Target: AWSBudgetServiceGateway.CreateBudget

{
   "AccountId": "111122223333",
   "Budget": {
      "BudgetLimit": {
         "Amount": "1000",
         "Unit": "USD"
      },
      "BudgetName": "Example Exclusion Budget",
      "BudgetType": "COST",
      "FilterExpression": {
         "Not": {
            "Dimensions": {
               "Key": "SERVICE",
               "Values": ["AWS Shield", "AWS Support (Enterprise)"]
            }
         }
      },
      "Metrics": ["NET_UNBLENDED_COST"],
      "TimePeriod": {
         "Start": 1712019600,
         "End": 1714611600
      },
      "TimeUnit": "MONTHLY"
   },
   "NotificationsWithSubscribers": [
      {
         "Notification": {
            "ComparisonOperator": "GREATER_THAN",
            "NotificationType": "ACTUAL",
            "Threshold": 70,
            "ThresholdType": "PERCENTAGE"
         },
         "Subscribers": [
            {
               "Address": "exclusion-alerts@example.com",
               "SubscriptionType": "EMAIL"
            }
         ]
      },
      {
         "Notification": {
            "ComparisonOperator": "GREATER_THAN",
            "NotificationType": "FORECASTED",
            "Threshold": 100,
            "ThresholdType": "PERCENTAGE"
         },
         "Subscribers": [
            {
               "Address": "exclusion-forecasts@example.com",
               "SubscriptionType": "EMAIL"
            }
         ]
      }
   ]
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/budgets-2016-10-20/createbudget.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/budgets-2016-10-20/createbudget.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/budgets-2016-10-20/createbudget.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/budgets-2016-10-20/createbudget.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/budgets-2016-10-20/createbudget.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/budgets-2016-10-20/createbudget.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/budgets-2016-10-20/createbudget.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/budgets-2016-10-20/createbudget.md)

- [AWS SDK for Python](../../../../services/goto/boto3/budgets-2016-10-20/createbudget.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/budgets-2016-10-20/createbudget.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AWS Budgets

CreateBudgetAction
