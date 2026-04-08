# ActivateOrganizationsAccess

Activate trusted access with AWS Organizations. With trusted access between StackSets
and Organizations activated, the management account has permissions to create
and manage StackSets for your organization.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidOperation**

The specified operation isn't valid.

HTTP Status Code: 400

**OperationNotFound**

The specified ID refers to an operation that doesn't exist.

HTTP Status Code: 404

## Examples

### ActivateOrganizationsAccess

This example illustrates one usage of ActivateOrganizationsAccess.

#### Sample Request

```

https://cloudformation.us-east-1.amazonaws.com/
    ?Action=ActivateOrganizationsAccess
```

#### Sample Response

```

<ActivateOrganizationsAccessResponse xmlns="http://internal.amazon.com/coral/com.amazonaws.maestro.service.v20160713/">
  <ActivateOrganizationsAccessResult/>
  <ResponseMetadata>
    <RequestId>f2038c6a-2ef6-45d0-a045-f2b4d15647ba</RequestId>
  </ResponseMetadata>
</ActivateOrganizationsAccessResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudformation-2010-05-15/activateorganizationsaccess.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudformation-2010-05-15/activateorganizationsaccess.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/activateorganizationsaccess.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudformation-2010-05-15/activateorganizationsaccess.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/activateorganizationsaccess.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudformation-2010-05-15/activateorganizationsaccess.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudformation-2010-05-15/activateorganizationsaccess.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudformation-2010-05-15/activateorganizationsaccess.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudformation-2010-05-15/activateorganizationsaccess.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/activateorganizationsaccess.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Actions

ActivateType
