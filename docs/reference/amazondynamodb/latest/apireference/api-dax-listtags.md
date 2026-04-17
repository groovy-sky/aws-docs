---
title: "ListTags"
---

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/dax-2017-04-19/ListTags)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/dax-2017-04-19/ListTags)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/dax-2017-04-19/ListTags)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/dax-2017-04-19/ListTags)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/dax-2017-04-19/ListTags)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/dax-2017-04-19/ListTags)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/dax-2017-04-19/ListTags)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/dax-2017-04-19/ListTags)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/dax-2017-04-19/ListTags)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/dax-2017-04-19/ListTags)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IncreaseReplicationFactor

RebootNode

All content copied from https://docs.aws.amazon.com/.
