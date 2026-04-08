# StartBatchDeleteConfigurationTask

###### Important

AWS Application Discovery Service is no longer open to new customers. Existing customers can continue to use the service as normal. For more information, see
[AWS Application Discovery Service availability change](../../../../services/application-discovery/latest/userguide/application-discovery-service-availability-change.md).

Takes a list of configurationId as input and starts an asynchronous deletion task to
remove the configurationItems. Returns a unique deletion task identifier.

## Request Syntax

```nohighlight

{
   "configurationIds": [ "string" ],
   "configurationType": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[configurationIds](#API_StartBatchDeleteConfigurationTask_RequestSyntax)**

The list of configuration IDs that will be deleted by the task.

Type: Array of strings

Length Constraints: Maximum length of 200.

Pattern: `\S*`

Required: Yes

**[configurationType](#API_StartBatchDeleteConfigurationTask_RequestSyntax)**

The type of configuration item to delete. Supported types are: SERVER.

Type: String

Valid Values: `SERVER`

Required: Yes

## Response Syntax

```nohighlight

{
   "taskId": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[taskId](#API_StartBatchDeleteConfigurationTask_ResponseSyntax)**

The unique identifier associated with the newly started deletion task.

Type: String

Pattern: `[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

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

**InvalidParameterException**

One or more parameters are not valid. Verify the parameters and try again.

HTTP Status Code: 400

**InvalidParameterValueException**

The value of one or more parameters are either invalid or out of range. Verify the
parameter values and try again.

HTTP Status Code: 400

**LimitExceededException**

###### Important

AWS Application Discovery Service is no longer open to new customers. Existing customers can continue to use the service as normal. For more information, see
[AWS Application Discovery Service availability change](../../../../services/application-discovery/latest/userguide/application-discovery-service-availability-change.md).

The limit of 200 configuration IDs per request has been exceeded.

HTTP Status Code: 400

**OperationNotPermittedException**

This operation is not permitted.

HTTP Status Code: 400

**ServerInternalErrorException**

The server experienced an internal error. Try again.

HTTP Status Code: 500

## Examples

The following example shows the request syntax to start a deletion task for a batch
of SERVER configurations, specified by the value passed to the required parameters of
`configurationIds` and `configurationType`.

#### Sample Request

```

{
    "configurationType": "SERVER",
    "configurationIds": ["d-server-029yqlktuw2udm", "d-server-03alnm4z74f77f"]
}
```

The following example shows the response for a successful
`StartBatchDeleteConfigurationTask` API call.

#### Sample Response

```

{
    "taskId": "b941cc54-b0df-4cdd-90fc-70ef4293dfce"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/discovery-2015-11-01/startbatchdeleteconfigurationtask.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/discovery-2015-11-01/startbatchdeleteconfigurationtask.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/discovery-2015-11-01/startbatchdeleteconfigurationtask.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/discovery-2015-11-01/startbatchdeleteconfigurationtask.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/discovery-2015-11-01/startbatchdeleteconfigurationtask.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/discovery-2015-11-01/startbatchdeleteconfigurationtask.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/discovery-2015-11-01/startbatchdeleteconfigurationtask.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/discovery-2015-11-01/startbatchdeleteconfigurationtask.md)

- [AWS SDK for Python](../../../../services/goto/boto3/discovery-2015-11-01/startbatchdeleteconfigurationtask.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/discovery-2015-11-01/startbatchdeleteconfigurationtask.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ListServerNeighbors

StartContinuousExport
