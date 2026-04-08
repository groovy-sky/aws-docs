# AccountLimit

A complex type that contains the type of limit that you specified in the request and
the current value for that limit.

## Contents

**Type**

The limit that you requested. Valid values include the following:

- **MAX\_HEALTH\_CHECKS\_BY\_OWNER**: The maximum
number of health checks that you can create using the current account.

- **MAX\_HOSTED\_ZONES\_BY\_OWNER**: The maximum number
of hosted zones that you can create using the current account.

- **MAX\_REUSABLE\_DELEGATION\_SETS\_BY\_OWNER**: The
maximum number of reusable delegation sets that you can create using the current
account.

- **MAX\_TRAFFIC\_POLICIES\_BY\_OWNER**: The maximum
number of traffic policies that you can create using the current account.

- **MAX\_TRAFFIC\_POLICY\_INSTANCES\_BY\_OWNER**: The
maximum number of traffic policy instances that you can create using the current
account. (Traffic policy instances are referred to as traffic flow policy
records in the Amazon Route 53 console.)

Type: String

Valid Values: `MAX_HEALTH_CHECKS_BY_OWNER | MAX_HOSTED_ZONES_BY_OWNER | MAX_TRAFFIC_POLICY_INSTANCES_BY_OWNER | MAX_REUSABLE_DELEGATION_SETS_BY_OWNER | MAX_TRAFFIC_POLICIES_BY_OWNER`

Required: Yes

**Value**

The current value for the limit that is specified by [Type](api-accountlimit-route53-type-accountlimit-type.md).

Type: Long

Valid Range: Minimum value of 1.

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/accountlimit.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/accountlimit.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/accountlimit.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Amazon Route 53

AlarmIdentifier
