# BatchDisassociateUserStack

Disassociates the specified users from the specified stacks.

## Request Syntax

```nohighlight

{
   "UserStackAssociations": [
      {
         "AuthenticationType": "string",
         "SendEmailNotification": boolean,
         "StackName": "string",
         "UserName": "string"
      }
   ]
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[UserStackAssociations](#API_BatchDisassociateUserStack_RequestSyntax)**

The list of UserStackAssociation objects.

Type: Array of [UserStackAssociation](api-userstackassociation.md) objects

Array Members: Minimum number of 1 item. Maximum number of 25 items.

Required: Yes

## Response Syntax

```nohighlight

{
   "errors": [
      {
         "ErrorCode": "string",
         "ErrorMessage": "string",
         "UserStackAssociation": {
            "AuthenticationType": "string",
            "SendEmailNotification": boolean,
            "StackName": "string",
            "UserName": "string"
         }
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[errors](#API_BatchDisassociateUserStack_ResponseSyntax)**

The list of UserStackAssociationError objects.

Type: Array of [UserStackAssociationError](api-userstackassociationerror.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterCombinationException**

Indicates an incorrect combination of parameters, or a missing parameter.

**Message**

The error message in the exception.

HTTP Status Code: 400

**OperationNotPermittedException**

The attempted operation is not permitted.

**Message**

The error message in the exception.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appstream-2016-12-01/batchdisassociateuserstack.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appstream-2016-12-01/batchdisassociateuserstack.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/batchdisassociateuserstack.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appstream-2016-12-01/batchdisassociateuserstack.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/batchdisassociateuserstack.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appstream-2016-12-01/batchdisassociateuserstack.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appstream-2016-12-01/batchdisassociateuserstack.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appstream-2016-12-01/batchdisassociateuserstack.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appstream-2016-12-01/batchdisassociateuserstack.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/batchdisassociateuserstack.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BatchAssociateUserStack

CopyImage

All content copied from https://docs.aws.amazon.com/.
