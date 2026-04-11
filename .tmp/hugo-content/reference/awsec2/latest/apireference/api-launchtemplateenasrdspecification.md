# LaunchTemplateEnaSrdSpecification

ENA Express uses AWS Scalable Reliable Datagram (SRD) technology to increase the
maximum bandwidth used per stream and minimize tail latency of network traffic between EC2 instances.
With ENA Express, you can communicate between two EC2 instances in the same subnet within the same
account, or in different accounts. Both sending and receiving instances must have ENA Express enabled.

To improve the reliability of network packet delivery, ENA Express reorders network packets on the
receiving end by default. However, some UDP-based applications are designed to handle network packets
that are out of order to reduce the overhead for packet delivery at the network layer. When ENA Express
is enabled, you can specify whether UDP network traffic uses it.

## Contents

**enaSrdEnabled**

Indicates whether ENA Express is enabled for the network interface.

Type: Boolean

Required: No

**enaSrdUdpSpecification**

Configures ENA Express for UDP network traffic.

Type: [LaunchTemplateEnaSrdUdpSpecification](api-launchtemplateenasrdudpspecification.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/launchtemplateenasrdspecification.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/launchtemplateenasrdspecification.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/launchtemplateenasrdspecification.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

LaunchTemplateElasticInferenceAcceleratorResponse

LaunchTemplateEnaSrdUdpSpecification
