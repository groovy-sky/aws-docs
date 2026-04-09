# ListTags

List all of the tags for a DAX cluster. You can call
`ListTags` up to 10 times per second, per account.

## Request Syntax

```nohighlight

{
   "NextToken": "string",
   "ResourceName": "string"
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[ResourceName](#API_dax_ListTags_RequestSyntax)**

The name of the DAX resource to which the tags belong.

Type: String

Required: Yes

**[NextToken](#API_dax_ListTags_RequestSyntax)**

An optional token returned from a prior request. Use this token for pagination of
results from this action. If this parameter is specified, the response includes only
results beyond the token.

Type: String

Required: No

## Response Syntax

```nohighlight

{
   "NextToken": "string",
   "Tags": [
      {
         "Key": "string",
         "Value": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[NextToken](#API_dax_ListTags_ResponseSyntax)**

If this value is present, there are additional results to be displayed. To retrieve
them, call `ListTags` again, with `NextToken` set to this
value.

Type: String

**[Tags](#API_dax_ListTags_ResponseSyntax)**

A list of tags currently associated with the DAX cluster.

Type: Array of [Tag](api-dax-tag.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ClusterNotFoundFault**

The requested cluster ID does not refer to an existing DAX
cluster.

HTTP Status Code: 400

**InvalidARNFault**

The Amazon Resource Name (ARN) supplied in the request is not valid.

HTTP Status Code: 400

**InvalidClusterStateFault**

The requested DAX cluster is not in the
_available_ state.

HTTP Status Code: 400

**InvalidParameterCombinationException**

Two or more incompatible parameters were specified.

HTTP Status Code: 400

**InvalidParameterValueException**

The value for a parameter is invalid.

HTTP Status Code: 400

**ServiceLinkedRoleNotFoundFault**

The specified service linked role (SLR) was not found.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dax-2017-04-19/listtags.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dax-2017-04-19/listtags.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dax-2017-04-19/listtags.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dax-2017-04-19/listtags.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dax-2017-04-19/listtags.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dax-2017-04-19/listtags.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dax-2017-04-19/listtags.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dax-2017-04-19/listtags.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dax-2017-04-19/listtags.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dax-2017-04-19/listtags.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IncreaseReplicationFactor

RebootNode

All content copied from https://docs.aws.amazon.com/.
