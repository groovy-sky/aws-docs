# DeleteParameterGroup

Deletes the specified parameter group. You cannot delete a parameter group if it is
associated with any DAX clusters.

## Request Syntax

```nohighlight

{
   "ParameterGroupName": "string"
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[ParameterGroupName](#API_dax_DeleteParameterGroup_RequestSyntax)**

The name of the parameter group to delete.

Type: String

Required: Yes

## Response Syntax

```nohighlight

{
   "DeletionMessage": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[DeletionMessage](#API_dax_DeleteParameterGroup_ResponseSyntax)**

A user-specified message for this action (i.e., a reason for deleting the parameter
group).

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterCombinationException**

Two or more incompatible parameters were specified.

HTTP Status Code: 400

**InvalidParameterGroupStateFault**

One or more parameters in a parameter group are in an invalid state.

HTTP Status Code: 400

**InvalidParameterValueException**

The value for a parameter is invalid.

HTTP Status Code: 400

**ParameterGroupNotFoundFault**

The specified parameter group does not exist.

HTTP Status Code: 400

**ServiceLinkedRoleNotFoundFault**

The specified service linked role (SLR) was not found.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dax-2017-04-19/deleteparametergroup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dax-2017-04-19/deleteparametergroup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dax-2017-04-19/deleteparametergroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dax-2017-04-19/deleteparametergroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dax-2017-04-19/deleteparametergroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dax-2017-04-19/deleteparametergroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dax-2017-04-19/deleteparametergroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dax-2017-04-19/deleteparametergroup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dax-2017-04-19/deleteparametergroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dax-2017-04-19/deleteparametergroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteCluster

DeleteSubnetGroup

All content copied from https://docs.aws.amazon.com/.
