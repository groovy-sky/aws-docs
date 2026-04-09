# ConnectorEntity

The high-level entity that can be queried in Amazon AppFlow. For example, a
Salesforce entity might be an _Account_ or
_Opportunity_, whereas a ServiceNow entity might be an
_Incident_.

## Contents

**name**

The name of the connector entity.

Type: String

Length Constraints: Maximum length of 128.

Pattern: `\S+`

Required: Yes

**hasNestedEntities**

Specifies whether the connector entity is a parent or a category and has more entities
nested underneath it. If another call is made with `entitiesPath =
        "the_current_entity_name_with_hasNestedEntities_true"`, then it returns the nested
entities underneath it. This provides a way to retrieve all supported entities in a recursive
fashion.

Type: Boolean

Required: No

**label**

The label applied to the connector entity.

Type: String

Length Constraints: Maximum length of 128.

Pattern: `.*`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/connectorentity.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/connectorentity.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/connectorentity.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConnectorDetail

ConnectorEntityField

All content copied from https://docs.aws.amazon.com/.
