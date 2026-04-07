# LaunchTemplateInstanceMetadataOptions

The metadata options for the instance. For more information, see [Use\
instance metadata to manage your EC2 instance](../../../../services/ec2/latest/userguide/ec2-instance-metadata.md) in the
_Amazon EC2 User Guide_.

## Contents

**httpEndpoint**

Enables or disables the HTTP metadata endpoint on your instances. If the parameter is
not specified, the default state is `enabled`.

###### Note

If you specify a value of `disabled`, you will not be able to access
your instance metadata.

Type: String

Valid Values: `disabled | enabled`

Required: No

**httpProtocolIpv6**

Enables or disables the IPv6 endpoint for the instance metadata service.

Default: `disabled`

Type: String

Valid Values: `disabled | enabled`

Required: No

**httpPutResponseHopLimit**

The desired HTTP PUT response hop limit for instance metadata requests. The larger the
number, the further instance metadata requests can travel.

Possible values: Integers from 1 to 64

Type: Integer

Required: No

**httpTokens**

Indicates whether IMDSv2 is required.

- `optional` \- IMDSv2 is optional. You can choose whether to send a
session token in your instance metadata retrieval requests. If you retrieve IAM
role credentials without a session token, you receive the IMDSv1 role
credentials. If you retrieve IAM role credentials using a valid session token,
you receive the IMDSv2 role credentials.

- `required` \- IMDSv2 is required. You must send a session token in
your instance metadata retrieval requests. With this option, retrieving the IAM
role credentials always returns IMDSv2 credentials; IMDSv1 credentials are not
available.

Type: String

Valid Values: `optional | required`

Required: No

**instanceMetadataTags**

Set to `enabled` to allow access to instance tags from the instance
metadata. Set to `disabled` to turn off access to instance tags from the
instance metadata. For more information, see [View tags for your EC2\
instances using instance metadata](../../../../services/ec2/latest/userguide/work-with-tags-in-imds.md).

Default: `disabled`

Type: String

Valid Values: `disabled | enabled`

Required: No

**state**

The state of the metadata option changes.

`pending` \- The metadata options are being updated and the instance is not
ready to process metadata traffic with the new selection.

`applied` \- The metadata options have been successfully applied on the
instance.

Type: String

Valid Values: `pending | applied`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/launchtemplateinstancemetadataoptions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/launchtemplateinstancemetadataoptions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/launchtemplateinstancemetadataoptions.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

LaunchTemplateInstanceMarketOptionsRequest

LaunchTemplateInstanceMetadataOptionsRequest
