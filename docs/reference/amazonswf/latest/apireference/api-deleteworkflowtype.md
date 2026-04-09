# DeleteWorkflowType

Deletes the specified _workflow type_.

Note: Prior to deletion, workflow types must first be **deprecated**.

After a workflow type has been deleted, you cannot create new executions of that type. Executions that
started before the type was deleted will continue to run.

**Access Control**

You can use IAM policies to control this action's access to Amazon SWF resources as follows:

- Use a `Resource` element with the domain name to limit the action to
only specified domains.

- Use an `Action` element to allow or deny permission to call this
action.

- Constrain the following parameters by using a `Condition` element with
the appropriate keys.

- `workflowType.name`: String constraint. The key is
`swf:workflowType.name`.

- `workflowType.version`: String constraint. The key is
`swf:workflowType.version`.

If the caller doesn't have sufficient permissions to invoke the action, or the
parameter values fall outside the specified constraints, the action fails. The associated
event attribute's `cause` parameter is set to `OPERATION_NOT_PERMITTED`.
For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF\
Workflows](../../../../services/amazonswf/latest/developerguide/swf-dev-iam.md) in the _Amazon SWF Developer Guide_.

## Request Syntax

```nohighlight

{
   "domain": "string",
   "workflowType": {
      "name": "string",
      "version": "string"
   }
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[domain](#API_DeleteWorkflowType_RequestSyntax)**

The name of the domain in which the workflow type is registered.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

**[workflowType](#API_DeleteWorkflowType_RequestSyntax)**

The workflow type to delete.

Type: [WorkflowType](api-workflowtype.md) object

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**OperationNotPermittedFault**

Returned when the caller doesn't have sufficient permissions to invoke the action.

**message**

A description that may help with diagnosing the cause of the fault.

HTTP Status Code: 400

**TypeNotDeprecatedFault**

Returned when the resource type has not been deprecated.

HTTP Status Code: 400

**UnknownResourceFault**

Returned when the named resource cannot be found with in the scope of this operation (region or domain). This could happen if the named resource was never created or is no longer available for this operation.

**message**

A description that may help with diagnosing the cause of the fault.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/swf-2012-01-25/deleteworkflowtype.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/swf-2012-01-25/deleteworkflowtype.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/deleteworkflowtype.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/swf-2012-01-25/deleteworkflowtype.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/deleteworkflowtype.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/swf-2012-01-25/deleteworkflowtype.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/swf-2012-01-25/deleteworkflowtype.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/swf-2012-01-25/deleteworkflowtype.md)

- [AWS SDK for Python](../../../../services/goto/boto3/swf-2012-01-25/deleteworkflowtype.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/deleteworkflowtype.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteActivityType

DeprecateActivityType

All content copied from https://docs.aws.amazon.com/.
