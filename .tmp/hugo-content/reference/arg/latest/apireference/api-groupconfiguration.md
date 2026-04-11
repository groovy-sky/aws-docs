# GroupConfiguration

A service configuration associated with a resource group. The configuration options
are determined by the AWS service that defines the `Type`, and specifies
which resources can be included in the group. You can add a service configuration when
you create the group by using [CreateGroup](api-creategroup.md), or later by using the [PutGroupConfiguration](api-putgroupconfiguration.md) operation. For details about group service
configuration syntax, see [Service configurations for resource\
groups](about-slg.md).

## Contents

###### Note

In the following list, the required parameters are described first.

**Configuration**

The configuration currently associated with the group and in effect.

Type: Array of [GroupConfigurationItem](api-groupconfigurationitem.md) objects

Array Members: Maximum number of 2 items.

Required: No

**FailureReason**

If present, the reason why a request to update the group configuration failed.

Type: String

Required: No

**ProposedConfiguration**

If present, the new configuration that is in the process of being applied to the
group.

Type: Array of [GroupConfigurationItem](api-groupconfigurationitem.md) objects

Array Members: Maximum number of 2 items.

Required: No

**Status**

The current status of an attempt to update the group configuration.

Type: String

Valid Values: `UPDATING | UPDATE_COMPLETE | UPDATE_FAILED`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/resource-groups-2017-11-27/groupconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/resource-groups-2017-11-27/groupconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/resource-groups-2017-11-27/groupconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Group

GroupConfigurationItem

All content copied from https://docs.aws.amazon.com/.
