# PatchOperation

For more information about supported patch operations, see [Patch Operations](patch-operations.md).

## Contents

**from**

The copy update operation's source as identified by a JSON-Pointer value referencing
the location within the targeted resource to copy the value from. For example, to
promote a canary deployment, you copy the canary deployment ID to the affiliated
deployment ID by calling a PATCH request on a Stage resource with "op":"copy",
"from":"/canarySettings/deploymentId" and "path":"/deploymentId".

Type: String

Required: No

**op**

An update operation to be performed with this PATCH request. The valid value can be
add, remove, replace or copy. Not all valid operations are supported for a given
resource. Support of the operations depends on specific operational contexts. Attempts
to apply an unsupported operation on a resource will return an error message..

Type: String

Valid Values: `add | remove | replace | move | copy | test`

Required: No

**path**

The op operation's target, as identified by a JSON Pointer value that references a
location within the targeted resource. For example, if the target resource has an
updateable property of {"name":"value"}, the path for this property is /name. If the
name property value is a JSON object (e.g., {"name": {"child/name": "child-value"}}),
the path for the child/name property will be /name/child~1name. Any slash ("/")
character appearing in path names must be escaped with "~1", as shown in the example
above. Each op operation can have only one path associated with it.

Type: String

Required: No

**value**

The new target value of the update operation. It is applicable for the add or replace
operation. When using AWS CLI to update a property of a JSON value, enclose the JSON
object with a pair of single quotes in a Linux shell, e.g., '{"a": ...}'.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/patchoperation.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/patchoperation.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/patchoperation.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MutualTlsAuthenticationInput

QuotaSettings

All content copied from https://docs.aws.amazon.com/.
