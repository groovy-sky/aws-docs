# DescribeOrganizationsAccess

Retrieves information about the account's `OrganizationAccess` status. This API
can be called either by the management account or the delegated administrator by using the
`CallAs` parameter. This API can also be called without the `CallAs`
parameter by the management account.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**CallAs**

\[Service-managed permissions\] Specifies whether you are acting as an account administrator
in the organization's management account or as a delegated administrator in a
member account.

By default, `SELF` is specified.

- If you are signed in to the management account, specify
`SELF`.

- If you are signed in to a delegated administrator account, specify
`DELEGATED_ADMIN`.

Your AWS account must be registered as a delegated administrator in the management account. For more information, see [Register a\
delegated administrator](../../../../services/cloudformation/latest/userguide/stacksets-orgs-delegated-admin.md) in the _AWS CloudFormation User Guide_.

Type: String

Valid Values: `SELF | DELEGATED_ADMIN`

Required: No

## Response Elements

The following element is returned by the service.

**Status**

Presents the status of the `OrganizationAccess`.

Type: String

Valid Values: `ENABLED | DISABLED | DISABLED_PERMANENTLY`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidOperation**

The specified operation isn't valid.

HTTP Status Code: 400

**OperationNotFound**

The specified ID refers to an operation that doesn't exist.

HTTP Status Code: 404

## Examples

### DescribeOrganizationsAccess

This example illustrates one usage of DescribeOrganizationsAccess.

#### Sample Request

```

https://cloudformation.us-east-1.amazonaws.com/
    ?Action=DescribeOrganizationsAccess
```

#### Sample Response

```

<DescribeOrganizationsAccessResponse xmlns="http://internal.amazon.com/coral/com.amazonaws.maestro.service.v20160713/">
  <DescribeOrganizationsAccessResult>
    <Status>ENABLED</Status>
  </DescribeOrganizationsAccessResult>
  <ResponseMetadata>
    <RequestId>240f1cd7-48d6-41a6-b61b-65c7c81893e9</RequestId>
  </ResponseMetadata>
</DescribeOrganizationsAccessResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudformation-2010-05-15/describeorganizationsaccess.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudformation-2010-05-15/describeorganizationsaccess.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/describeorganizationsaccess.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudformation-2010-05-15/describeorganizationsaccess.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/describeorganizationsaccess.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudformation-2010-05-15/describeorganizationsaccess.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudformation-2010-05-15/describeorganizationsaccess.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudformation-2010-05-15/describeorganizationsaccess.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudformation-2010-05-15/describeorganizationsaccess.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/describeorganizationsaccess.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeGeneratedTemplate

DescribePublisher

All content copied from https://docs.aws.amazon.com/.
