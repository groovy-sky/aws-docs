# DescribeBatchDeleteConfigurationTask

###### Important

AWS Application Discovery Service is no longer open to new customers. Existing customers can continue to use the service as normal. For more information, see
[AWS Application Discovery Service availability change](../../../../services/application-discovery/latest/userguide/application-discovery-service-availability-change.md).

Takes a unique deletion task identifier as input and returns metadata about a
configuration deletion task.

## Request Syntax

```nohighlight

{
   "taskId": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](https://docs.aws.amazon.com/application-discovery/latest/APIReference/CommonParameters.html).

The request accepts the following data in JSON format.

**[taskId](#API_DescribeBatchDeleteConfigurationTask_RequestSyntax)**

The ID of the task to delete.

Type: String

Pattern: `[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`

Required: Yes

## Response Syntax

```nohighlight

{
   "task": {
      "configurationType": "string",
      "deletedConfigurations": [ "string" ],
      "deletionWarnings": [
         {
            "configurationId": "string",
            "warningCode": number,
            "warningText": "string"
         }
      ],
      "endTime": number,
      "failedConfigurations": [
         {
            "configurationId": "string",
            "errorMessage": "string",
            "errorStatusCode": number
         }
      ],
      "requestedConfigurations": [ "string" ],
      "startTime": number,
      "status": "string",
      "taskId": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[task](#API_DescribeBatchDeleteConfigurationTask_ResponseSyntax)**

The `BatchDeleteConfigurationTask` that represents the deletion task being
executed.

Type: [BatchDeleteConfigurationTask](https://docs.aws.amazon.com/application-discovery/latest/APIReference/API_BatchDeleteConfigurationTask.html) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](https://docs.aws.amazon.com/application-discovery/latest/APIReference/CommonErrors.html).

**AuthorizationErrorException**

The user does not have permission to perform the action. Check the IAM policy
associated with this user.

HTTP Status Code: 400

**HomeRegionNotSetException**

###### Important

AWS Application Discovery Service is no longer open to new customers. Existing customers can continue to use the service as normal. For more information, see
[AWS Application Discovery Service availability change](../../../../services/application-discovery/latest/userguide/application-discovery-service-availability-change.md).

The home Region is not set. Set the home Region to continue.

HTTP Status Code: 400

**InvalidParameterValueException**

The value of one or more parameters are either invalid or out of range. Verify the
parameter values and try again.

HTTP Status Code: 400

**ServerInternalErrorException**

The server experienced an internal error. Try again.

HTTP Status Code: 500

## Examples

The following example shows the request syntax for a deletion task that was
previously started using [StartBatchDeleteConfigurationTask](https://docs.aws.amazon.com/application-discovery/latest/APIReference/API_StartBatchDeleteConfigurationTask.html) specified by
the value passed to the required `taskId` parameter.

#### Sample Request

```

{
    "taskId": "b941cc54-b0df-4cdd-90fc-70ef4293dfce"
}
```

The following example shows the response for a successful
`DescribeBatchDeleteConfigurationTask` API call.

#### Sample Response

```

{
      "task": {
        "configurationType": "SERVER",
        "deletedConfigurations": ["d-server-029yqlktuw2udm", "d-server-03alnm4z74f77f"],
        "deletionWarnings": [],
        "endTime": 1695765672.384,
        "failedConfigurations": [],
        "requestedConfigurations": ["d-server-029yqlktuw2udm", "d-server-03alnm4z74f77f"],
        "startTime": 1695755672.129,
        "status": "COMPLETED",
        "taskId": "b941cc54-b0df-4cdd-90fc-70ef4293dfce"
    }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/discovery-2015-11-01/DescribeBatchDeleteConfigurationTask)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/discovery-2015-11-01/DescribeBatchDeleteConfigurationTask)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/discovery-2015-11-01/DescribeBatchDeleteConfigurationTask)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/discovery-2015-11-01/DescribeBatchDeleteConfigurationTask)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/discovery-2015-11-01/DescribeBatchDeleteConfigurationTask)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/discovery-2015-11-01/DescribeBatchDeleteConfigurationTask)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/discovery-2015-11-01/DescribeBatchDeleteConfigurationTask)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/discovery-2015-11-01/DescribeBatchDeleteConfigurationTask)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/discovery-2015-11-01/DescribeBatchDeleteConfigurationTask)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/discovery-2015-11-01/DescribeBatchDeleteConfigurationTask)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DescribeAgents

DescribeConfigurations
