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
instance, use [DeleteTrafficPolicyInstance](https://docs.aws.amazon.com/Route53/latest/APIReference/API_DeleteTrafficPolicyInstance.html). Amazon Route 53 will delete the
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

Type: [ResourceRecordSet](https://docs.aws.amazon.com/Route53/latest/APIReference/API_ResourceRecordSet.html) object

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/route53-2013-04-01/Change)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/route53-2013-04-01/Change)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/route53-2013-04-01/Change)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AliasTarget

ChangeBatch
