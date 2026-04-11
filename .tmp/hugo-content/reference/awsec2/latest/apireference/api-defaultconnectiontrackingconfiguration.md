# DefaultConnectionTrackingConfiguration

Indicates default conntrack information for the instance type. For more
information, see
[Connection tracking timeouts](../../../../services/ec2/latest/userguide/security-group-connection-tracking.md#connection-tracking-timeouts) in the Amazon EC2 User Guide.

## Contents

**defaultTcpEstablishedTimeout**

Default timeout (in seconds) for idle TCP connections in an established state.

Type: Integer

Required: No

**defaultUdpStreamTimeout**

Default timeout (in seconds) for idle UDP flows classified as streams which have seen more than one request-response transaction.

Type: Integer

Required: No

**defaultUdpTimeout**

Default timeout (in seconds) for idle UDP flows that have seen traffic only in a single direction or a single request-response transaction.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/defaultconnectiontrackingconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/defaultconnectiontrackingconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/defaultconnectiontrackingconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeclarativePoliciesReport

DeleteFleetError
