# DeleteOptionGroup

Deletes an existing option group.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**OptionGroupName**

The name of the option group to be deleted.

###### Note

You can't delete default option groups.

Type: String

Required: Yes

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidOptionGroupStateFault**

The option group isn't in the _available_ state.

HTTP Status Code: 400

**OptionGroupNotFoundFault**

The specified option group could not be found.

HTTP Status Code: 404

## Examples

### Example

This example illustrates one usage of DeleteOptionGroup.

#### Sample Request

```

https://rds.us-east-1.amazonaws.com/
   ?Action=DeleteOptionGroup
   &OptionGroupName=myawsuser-og00
   &SignatureMethod=HmacSHA256
   &SignatureVersion=4
   &Version=2014-10-31
   &X-Amz-Algorithm=AWS4-HMAC-SHA256
   &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20140425/us-east-1/rds/aws4_request
   &X-Amz-Date=20140425T181205Z
   &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
   &X-Amz-Signature=8a684aebe6d5219bb3572316a341963324d6ef339bd0dcfa5854f1a01d401214

```

#### Sample Response

```

<DeleteOptionGroupResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <ResponseMetadata>
    <RequestId>0ac9cda2-bbf4-11d3-f92b-31fa5e8dbc99</RequestId>
  </ResponseMetadata>
</DeleteOptionGroupResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/deleteoptiongroup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/deleteoptiongroup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/deleteoptiongroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/deleteoptiongroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/deleteoptiongroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/deleteoptiongroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/deleteoptiongroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/deleteoptiongroup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/deleteoptiongroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/deleteoptiongroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteIntegration

DeleteTenantDatabase

All content copied from https://docs.aws.amazon.com/.
