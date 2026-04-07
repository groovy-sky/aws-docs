# ConnectionTrackingSpecificationResponse

A security group connection tracking specification response that enables you to set
the idle timeout for connection tracking on an Elastic network interface. For more
information, see [Connection tracking timeouts](../../../../services/ec2/latest/userguide/security-group-connection-tracking.md#connection-tracking-timeouts) in the
_Amazon EC2 User Guide_.

## Contents

**tcpEstablishedTimeout**

Timeout (in seconds) for idle TCP
connections in an established state. Min: 60 seconds. Max: 432000 seconds (5
days). Default: 432000 seconds. Recommended: Less than 432000 seconds.

Type: Integer

Required: No

**udpStreamTimeout**

Timeout (in seconds) for idle UDP
flows classified as streams which have seen more than one request-response
transaction. Min: 60 seconds. Max: 180 seconds (3 minutes). Default: 180
seconds.

Type: Integer

Required: No

**udpTimeout**

Timeout (in seconds) for idle UDP flows that
have seen traffic only in a single direction or a single request-response
transaction. Min: 30 seconds. Max: 60 seconds. Default: 30 seconds.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/connectiontrackingspecificationresponse.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/connectiontrackingspecificationresponse.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/connectiontrackingspecificationresponse.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ConnectionTrackingSpecificationRequest

ConversionTask
