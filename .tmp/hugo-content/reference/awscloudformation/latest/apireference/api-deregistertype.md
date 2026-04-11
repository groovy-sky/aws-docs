# DeregisterType

Marks an extension or extension version as `DEPRECATED` in the CloudFormation
registry, removing it from active use. Deprecated extensions or extension versions cannot be
used in CloudFormation operations.

To deregister an entire extension, you must individually deregister all active versions of
that extension. If an extension has only a single active version, deregistering that version
results in the extension itself being deregistered and marked as deprecated in the
registry.

You can't deregister the default version of an extension if there are other active version
of that extension. If you do deregister the default version of an extension, the extension
type itself is deregistered as well and marked as deprecated.

To view the deprecation status of an extension or extension version, use [DescribeType](api-describetype.md).

For more information, see [Remove\
third-party private extensions from your account](../../../../services/cloudformation/latest/userguide/registry-private-deregister-extension.md) in the
_AWS CloudFormation User Guide_.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**Arn**

The Amazon Resource Name (ARN) of the extension.

Conditional: You must specify either `TypeName` and `Type`, or
`Arn`.

Type: String

Length Constraints: Maximum length of 1024.

Pattern: `arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:[0-9]{12}:type/.+`

Required: No

**Type**

The kind of extension.

Conditional: You must specify either `TypeName` and `Type`, or
`Arn`.

Type: String

Valid Values: `RESOURCE | MODULE | HOOK`

Required: No

**TypeName**

The name of the extension.

Conditional: You must specify either `TypeName` and `Type`, or
`Arn`.

Type: String

Length Constraints: Minimum length of 10. Maximum length of 204.

Pattern: `[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}(::MODULE){0,1}`

Required: No

**VersionId**

The ID of a specific version of the extension. The version ID is the value at the end of
the Amazon Resource Name (ARN) assigned to the extension version when it is registered.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `[A-Za-z0-9-]+`

Required: No

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**CFNRegistry**

An error occurred during a CloudFormation registry operation.

**Message**

A message with details about the error that occurred.

HTTP Status Code: 400

**TypeNotFound**

The specified extension doesn't exist in the CloudFormation registry.

HTTP Status Code: 404

## Examples

### Deregistering an extension version

The following example removes a specific version of the
`My::Resource::Example` resource type from active use in the CloudFormation
registry.

#### Sample Request

```

https://cloudformation.us-east-1.amazonaws.com/
 ?Action=DeregisterType
 &Version=2010-05-15
 &TypeName=My::Resource::Example
 &Type=RESOURCE
 &VersionId=00000002
 &X-Amz-Algorithm=AWS4-HMAC-SHA256
 &X-Amz-Credential=[Access key ID and scope]
 &X-Amz-Date=20191204T181601Z
 &X-Amz-SignedHeaders=content-type;host
 &X-Amz-Signature=[Signature]
```

#### Sample Response

```

<DeregisterTypeResponse xmlns="http://cloudformation.amazonaws.com/doc/2010-05-15/">
  <DeregisterTypeResult/>
  <ResponseMetadata>
    <RequestId>78c291d1-4463-4845-a600-29221example</RequestId>
  </ResponseMetadata>
</DeregisterTypeResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudformation-2010-05-15/deregistertype.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudformation-2010-05-15/deregistertype.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/deregistertype.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudformation-2010-05-15/deregistertype.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/deregistertype.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudformation-2010-05-15/deregistertype.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudformation-2010-05-15/deregistertype.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudformation-2010-05-15/deregistertype.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudformation-2010-05-15/deregistertype.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/deregistertype.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteStackSet

DescribeAccountLimits

All content copied from https://docs.aws.amazon.com/.
