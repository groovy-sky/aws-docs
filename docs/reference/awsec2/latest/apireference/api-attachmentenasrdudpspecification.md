# AttachmentEnaSrdUdpSpecification

ENA Express is compatible with both TCP and UDP transport protocols. When it's enabled, TCP traffic
automatically uses it. However, some UDP-based applications are designed to handle network packets that are
out of order, without a need for retransmission, such as live video broadcasting or other near-real-time
applications. For UDP traffic, you can specify whether to use ENA Express, based on your application
environment needs.

## Contents

**enaSrdUdpEnabled**

Indicates whether UDP traffic to and from the instance uses ENA Express. To specify this setting,
you must first enable ENA Express.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/AttachmentEnaSrdUdpSpecification)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/AttachmentEnaSrdUdpSpecification)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/AttachmentEnaSrdUdpSpecification)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AttachmentEnaSrdSpecification

AttributeBooleanValue
