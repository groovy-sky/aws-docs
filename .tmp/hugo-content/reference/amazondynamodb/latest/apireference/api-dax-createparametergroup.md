# CreateParameterGroup

Creates a new parameter group. A parameter group is a collection of parameters that
you apply to all of the nodes in a DAX cluster.

## Request Syntax

```nohighlight

{
   "Description": "string",
   "ParameterGroupName": "string"
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[ParameterGroupName](#API_dax_CreateParameterGroup_RequestSyntax)**

The name of the parameter group to apply to all of the clusters in this replication
group.

Type: String

Required: Yes

**[Description](#API_dax_CreateParameterGroup_RequestSyntax)**

A description of the parameter group.

Type: String

Required: No

## Response Syntax

```nohighlight

{
   "ParameterGroup": {
      "Description": "string",
      "ParameterGroupName": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[ParameterGroup](#API_dax_CreateParameterGroup_ResponseSyntax)**

Represents the output of a _CreateParameterGroup_
action.

Type: [ParameterGroup](api-dax-parametergroup.md) object

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

**ParameterGroupAlreadyExistsFault**

The specified parameter group already exists.

HTTP Status Code: 400

**ParameterGroupQuotaExceededFault**

You have attempted to exceed the maximum number of parameter groups.

HTTP Status Code: 400

**ServiceLinkedRoleNotFoundFault**

The specified service linked role (SLR) was not found.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dax-2017-04-19/createparametergroup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dax-2017-04-19/createparametergroup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dax-2017-04-19/createparametergroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dax-2017-04-19/createparametergroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dax-2017-04-19/createparametergroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dax-2017-04-19/createparametergroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dax-2017-04-19/createparametergroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dax-2017-04-19/createparametergroup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dax-2017-04-19/createparametergroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dax-2017-04-19/createparametergroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateCluster

CreateSubnetGroup

All content copied from https://docs.aws.amazon.com/.
