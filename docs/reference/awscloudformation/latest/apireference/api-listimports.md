# ListImports

Lists all stacks that are importing an exported output value. To modify or remove an
exported output value, first use this action to see which stacks are using it. To see the
exported output values in your account, see [ListExports](api-listexports.md).

For more information about importing an exported output value, see the [Fn::ImportValue](../../../../services/cloudformation/latest/templatereference/intrinsic-function-reference-importvalue.md) function.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**ExportName**

The name of the exported output value. CloudFormation returns the stack names that are
importing this value.

Type: String

Required: Yes

**NextToken**

The token for the next set of items to return. (You received this token from a previous
call.)

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

## Response Elements

The following elements are returned by the service.

**Imports.member.N**

A list of stack names that are importing the specified exported output value.

Type: Array of strings

**NextToken**

A string that identifies the next page of exports. If there is no additional page, this
value is null.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

## Examples

### ListExports

This example illustrates one usage of ListImports.

#### Sample Request

```

https://cloudformation.us-east-1.amazonaws.com/
 ?Action=ListImports
 &ExportName=SampleStack-MyExportedValue
 &Version=2010-05-15
 &X-Amz-Algorithm=AWS4-HMAC-SHA256
 &X-Amz-Credential=[Access key ID and scope]
 &X-Amz-Date=20160316T233349Z
 &X-Amz-SignedHeaders=content-type;host
 &X-Amz-Signature=[Signature]
```

#### Sample Response

```

<ListImportsResponse xmlns="http://cloudformation.amazonaws.com/doc/2010-05-15/">
  <ListImportsResult>
    <Imports>
      <member>Import-SampleStack</member>
    </Imports>
  </ListImportsResult>
  <ResponseMetadata>
    <RequestId>a13656a8-a7b9-11e6-964c-41b56747ddb0</RequestId>
  </ResponseMetadata>
</ListImportsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudformation-2010-05-15/listimports.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudformation-2010-05-15/listimports.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/listimports.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudformation-2010-05-15/listimports.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/listimports.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudformation-2010-05-15/listimports.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudformation-2010-05-15/listimports.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudformation-2010-05-15/listimports.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudformation-2010-05-15/listimports.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/listimports.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ListHookResults

ListResourceScanRelatedResources
