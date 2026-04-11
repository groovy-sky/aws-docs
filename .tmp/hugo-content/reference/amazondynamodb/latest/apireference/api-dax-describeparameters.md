# DescribeParameters

Returns the detailed parameter list for a particular parameter group.

## Request Syntax

```nohighlight

{
   "MaxResults": number,
   "NextToken": "string",
   "ParameterGroupName": "string",
   "Source": "string"
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[ParameterGroupName](#API_dax_DescribeParameters_RequestSyntax)**

The name of the parameter group.

Type: String

Required: Yes

**[MaxResults](#API_dax_DescribeParameters_RequestSyntax)**

The maximum number of results to include in the response. If more results exist
than the specified `MaxResults` value, a token is included in the response so
that the remaining results can be retrieved.

The value for `MaxResults` must be between 20 and 100.

Type: Integer

Required: No

**[NextToken](#API_dax_DescribeParameters_RequestSyntax)**

An optional token returned from a prior request. Use this token for pagination of
results from this action. If this parameter is specified, the response includes only
results beyond the token, up to the value specified by
`MaxResults`.

Type: String

Required: No

**[Source](#API_dax_DescribeParameters_RequestSyntax)**

How the parameter is defined. For example, `system` denotes a
system-defined parameter.

Type: String

Required: No

## Response Syntax

```nohighlight

{
   "NextToken": "string",
   "Parameters": [
      {
         "AllowedValues": "string",
         "ChangeType": "string",
         "DataType": "string",
         "Description": "string",
         "IsModifiable": "string",
         "NodeTypeSpecificValues": [
            {
               "NodeType": "string",
               "Value": "string"
            }
         ],
         "ParameterName": "string",
         "ParameterType": "string",
         "ParameterValue": "string",
         "Source": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[NextToken](#API_dax_DescribeParameters_ResponseSyntax)**

Provides an identifier to allow retrieval of paginated results.

Type: String

**[Parameters](#API_dax_DescribeParameters_ResponseSyntax)**

A list of parameters within a parameter group. Each element in the list represents
one parameter.

Type: Array of [Parameter](api-dax-parameter.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterCombinationException**

Two or more incompatible parameters were specified.

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dax-2017-04-19/describeparameters.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dax-2017-04-19/describeparameters.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dax-2017-04-19/describeparameters.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dax-2017-04-19/describeparameters.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dax-2017-04-19/describeparameters.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dax-2017-04-19/describeparameters.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dax-2017-04-19/describeparameters.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dax-2017-04-19/describeparameters.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dax-2017-04-19/describeparameters.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dax-2017-04-19/describeparameters.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeParameterGroups

DescribeSubnetGroups

All content copied from https://docs.aws.amazon.com/.
