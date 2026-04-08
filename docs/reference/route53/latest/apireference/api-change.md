# Change

The information for each resource record set that you want to change.

## Contents

**Action**

The action to perform:

- `CREATE`: Creates a resource record set that has the specified
values.

- `DELETE`: Deletes a existing resource record set.

###### Important

To delete the resource record set that is associated with a traffic policy
instance, use [DeleteTrafficPolicyInstance](api-deletetrafficpolicyinstance.md). Amazon Route 53 will delete the
resource record set automatically. If you delete the resource record set by
using `ChangeResourceRecordSets`, Route 53 doesn't automatically
delete the traffic policy instance, and you'll continue to be charged for it
even though it's no longer in use.

- `UPSERT`: If a resource record set doesn't already exist, Route 53
creates it. If a resource record set does exist, Route 53 updates it with the
values in the request.

Type: String

Valid Values: `CREATE | DELETE | UPSERT`

Required: Yes

**ResourceRecordSet**

Information about the resource record set to create, delete, or update.

Type: [ResourceRecordSet](api-resourcerecordset.md) object

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/change.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/change.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/change.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AliasTarget

ChangeBatch
