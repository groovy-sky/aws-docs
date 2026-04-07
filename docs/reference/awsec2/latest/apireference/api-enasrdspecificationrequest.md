# EnaSrdSpecificationRequest

Launch instances with ENA Express settings configured from your launch
template.

## Contents

**EnaSrdEnabled** (request), **EnaSrdEnabled** (response)

Specifies whether ENA Express is enabled for the network interface when you launch an
instance.

Type: Boolean

Required: No

**EnaSrdUdpSpecification** (request), **EnaSrdUdpSpecification** (response)

Contains ENA Express settings for UDP network traffic for the network interface
attached to the instance.

Type: [EnaSrdUdpSpecificationRequest](api-enasrdudpspecificationrequest.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/enasrdspecificationrequest.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/enasrdspecificationrequest.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/enasrdspecificationrequest.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

EnaSrdSpecification

EnaSrdUdpSpecification
