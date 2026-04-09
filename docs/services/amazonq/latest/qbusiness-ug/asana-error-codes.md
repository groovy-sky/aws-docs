# Understand error codes in the Amazon Q Business Asana connector (Preview)

The following table provides information about error codes you may see for the
Asana connector and suggested resolutions.

Error codeError messageSuggested resolutionASN-5101The token in your data source configuration is missing.Enter valid token (specific to PAT or ServiceAccount) and try
again.ASN-5102The auth type in your data source configuration is invalid. Enter valid auth type (PAT or ServiceAccount) and try again.ASN-5103The auth type in your data source configuration is missing.Enter valid auth type (PAT or ServiceAccount) and try again.ASN-5104The crawl type in your data source configuration is missing.Enter valid crawlType (ex. FULL\_CRAWL) and try again.ASN-5105Only String, String List, Date and Long formats are supported for the
indexFieldType in all the field mappings.Please provide the supported format only for the indexFieldType in
all the fieldMappings.ASN-5106There was an error parsing the field value. The size has exceeded the
maximum allowable limit.The maximum size permitted is 1000 for the field {field name}. Ex –
field name can be token, workspaceId, projectIdASN-5107The connection configuration in your data source configuration is
missing.Enter valid connection configuration details and try again.ASN-5108The repository credentials in your data source configuration is
missing.Enter valid repository credentials details and try again.ASN-5109The repository endpoint metadata in your data source configuration is
missing.Enter valid repository endpoint metadata details and try
again.ASN-5110The additional properties in your data source configuration is
missing.Error messageEnter valid additional properties and try again.ASN-5111Invalid Workspace Id.Provide numeric value for workspace Id.ASN-5112Invalid Project Id.Provide numeric value for Project Id.ASN-5113The workspaceIds field has exceeded the limit. Allowed entries are
limited to 1.Correct the number of entries and try again.ASN-5114The projectIds field has exceeded the limit. Maximum allowed entries
are 20.Correct the number of entries and try again.ASN-5115The regex pattern provided in {pattern field} Patterns is
invalid.Provide valid regex pattern.ASN-5116The workspace Id in your data source configuration is
missing.Enter a single workspace Id and try again.ASN-5200IO Exception occurred while reading contents from Asana.Rerun the connector once again, if the issue persists contact
support.ASN-5201Unknown exception occurred.Rerun the connector once again, if the issue persists contact
support.ASN-5202The user details are not found.Verify the user data using the user API curl command or in postman
for the given authentication type.ASN-5204Could not find Project(s) using the filter provided in the Asana
configuration. Either Workspace(s) or Project(s) access is
forbidden.Check the Workspace and Project(s) access and provide valid inclusion
and exclusion project name filter inputs.ASN-5500Asana connection successful.This is to convey the connection to Asana data source is
successful.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IAM role

Box

All content copied from https://docs.aws.amazon.com/.
