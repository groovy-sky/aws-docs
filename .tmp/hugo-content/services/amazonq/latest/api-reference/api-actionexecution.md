# ActionExecution

Performs an Amazon Q Business plugin action during a non-streaming chat
conversation.

## Contents

**payload**

A mapping of field names to the field values in input that an end user provides to
Amazon Q Business requests to perform their plugin action.

Type: String to [ActionExecutionPayloadField](api-actionexecutionpayloadfield.md) object map

Key Length Constraints: Minimum length of 1.

Required: Yes

**payloadFieldNameSeparator**

A string used to retain information about the hierarchical contexts within an action
execution event payload.

Type: String

Length Constraints: Fixed length of 1.

Required: Yes

**pluginId**

The identifier of the plugin the action is attached to.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/actionexecution.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/actionexecution.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/actionexecution.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ActionConfiguration

ActionExecutionEvent

All content copied from https://docs.aws.amazon.com/.
