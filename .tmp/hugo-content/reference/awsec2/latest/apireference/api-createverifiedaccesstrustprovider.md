# CreateVerifiedAccessTrustProvider

A trust provider is a third-party entity that creates, maintains, and manages identity
information for users and devices. When an application request is made, the identity
information sent by the trust provider is evaluated by Verified Access before allowing or
denying the application request.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientToken**

A unique, case-sensitive token that you provide to ensure idempotency of your
modification request. For more information, see [Ensuring idempotency](../../../../services/ec2/latest/devguide/ec2-api-idempotency.md).

Type: String

Required: No

**Description**

A description for the Verified Access trust provider.

Type: String

Required: No

**DeviceOptions**

The options for a device-based trust provider. This parameter is required when the
provider type is `device`.

Type: [CreateVerifiedAccessTrustProviderDeviceOptions](api-createverifiedaccesstrustproviderdeviceoptions.md) object

Required: No

**DeviceTrustProviderType**

The type of device-based trust provider. This parameter is required when the provider
type is `device`.

Type: String

Valid Values: `jamf | crowdstrike | jumpcloud`

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**NativeApplicationOidcOptions**

The OpenID Connect (OIDC) options.

Type: [CreateVerifiedAccessNativeApplicationOidcOptions](api-createverifiedaccessnativeapplicationoidcoptions.md) object

Required: No

**OidcOptions**

The options for a OpenID Connect-compatible user-identity trust provider. This parameter
is required when the provider type is `user`.

Type: [CreateVerifiedAccessTrustProviderOidcOptions](api-createverifiedaccesstrustprovideroidcoptions.md) object

Required: No

**PolicyReferenceName**

The identifier to be used when working with policy rules.

Type: String

Required: Yes

**SseSpecification**

The options for server side encryption.

Type: [VerifiedAccessSseSpecificationRequest](api-verifiedaccessssespecificationrequest.md) object

Required: No

**TagSpecification.N**

The tags to assign to the Verified Access trust provider.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

**TrustProviderType**

The type of trust provider.

Type: String

Valid Values: `user | device`

Required: Yes

**UserTrustProviderType**

The type of user-based trust provider. This parameter is required when the provider type
is `user`.

Type: String

Valid Values: `iam-identity-center | oidc`

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**verifiedAccessTrustProvider**

Details about the Verified Access trust provider.

Type: [VerifiedAccessTrustProvider](api-verifiedaccesstrustprovider.md) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/createverifiedaccesstrustprovider.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/createverifiedaccesstrustprovider.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createverifiedaccesstrustprovider.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createverifiedaccesstrustprovider.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createverifiedaccesstrustprovider.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createverifiedaccesstrustprovider.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createverifiedaccesstrustprovider.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createverifiedaccesstrustprovider.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/createverifiedaccesstrustprovider.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createverifiedaccesstrustprovider.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateVerifiedAccessInstance

CreateVolume
