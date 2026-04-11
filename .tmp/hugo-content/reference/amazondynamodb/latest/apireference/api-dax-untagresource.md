# UntagResource

Removes the association of tags from a DAX resource. You can call
`UntagResource` up to 5 times per second, per account.

## Request Syntax

```nohighlight

{
   "ResourceName": "string",
   "TagKeys": [ "string" ]
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[ResourceName](#API_dax_UntagResource_RequestSyntax)**

The name of the DAX resource from which the tags should be
removed.

Type: String

Required: Yes

**[TagKeys](#API_dax_UntagResource_RequestSyntax)**

A list of tag keys. If the DAX cluster has any tags with these keys,
then the tags are removed from the cluster.

Type: Array of strings

Required: Yes

## Response Syntax

```nohighlight

{
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

**[Tags](#API_dax_UntagResource_ResponseSyntax)**

The tag keys that have been removed from the cluster.

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

**TagNotFoundFault**

The tag does not exist.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dax-2017-04-19/untagresource.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dax-2017-04-19/untagresource.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dax-2017-04-19/untagresource.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dax-2017-04-19/untagresource.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dax-2017-04-19/untagresource.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dax-2017-04-19/untagresource.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dax-2017-04-19/untagresource.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dax-2017-04-19/untagresource.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dax-2017-04-19/untagresource.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dax-2017-04-19/untagresource.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TagResource

UpdateCluster

All content copied from https://docs.aws.amazon.com/.
