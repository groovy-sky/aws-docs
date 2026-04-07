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

Type: [EnaSrdUdpSpecificationRequest](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_EnaSrdUdpSpecificationRequest.html) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/EnaSrdSpecificationRequest)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/EnaSrdSpecificationRequest)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/EnaSrdSpecificationRequest)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EnaSrdSpecification

EnaSrdUdpSpecification
