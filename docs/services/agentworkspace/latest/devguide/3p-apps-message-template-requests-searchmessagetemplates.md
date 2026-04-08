# Retrieve message templates that match a search query in Amazon Connect Agent Workspace

Returns the SearchMessageTemplatesResponse object, which contains the matching
message templates and a token to retrieve the next page of results, if available.
The SearchMessageTemplatesParams object is used to configure the search, allowing
you to filter by various criteria, as well as specify the channels and pagination
options. If no filter text is provided, all active message templates for the agent's
routing profile and the channel(s) specified are returned.

**Signature**

```

searchMessageTemplates(request: SearchMessageTemplatesParams): Promise<SearchMessageTemplatesResponse>
```

**SearchMessageTemplatesResponse Properties**

**Parameter****Type****Description**messageTemplateMessageTemplate\[\]List of message templates matching the search criteria specified
in the requestnextTokenstringThe token for the next set of results. Use the value returned in
the previous response in the next request to retrieve the next set
of results.

**MessageTemplate Properties**

**Parameter****Type****Description**messageTemplateArnstringThe ARN of the message template. This contains the active version
qualifier at the end of the ARNmessageTemplateIdstringThe ID of the message template. This does NOT contain a qualifier
with the version of the message template.namestringName of the message templatedescriptionstringDescription of the message template

**SearchMessageTemplatesParams Properties**

**Parameter****Type****Description**channelsMessageTemplateChannel\[\]The channel(s) to return message templates for. If the list is
empty, no message templates will be returned. Supported values:
"EMAIL"queriesMessageTemplateQueryField\[\]Queries are used to filter the returned message templates by name
or description. Leaving the queries empty will return all message
templates associated with the agent's routing profilemaxResultsnumberMaximum number of message templates to returnnextTokenstringThe token for the next set of results. Use the value returned in
the previous response in the next request to retrieve the next set
of results.

**MessageTemplateQueryField Properties**

**Parameter****Type****Description**name"name" \| "description"The message templates will be filtered by the values matching the
text in the name field providedvaluesstring\[\]The values of the attribute to query the message templates
bypriority"HIGH" \| "MEDIUM" \| "LOW"The importance of the attribute field when calculating query
result relevancy scores. The value set for this parameter affects
the ordering of search results.allowFuzzinessbooleanWhether the query expects only exact matches on the attribute
field values. The results of the query will only include exact
matches if this parameter is set to false.operator"CONTAINS" \| "CONTAINS\_AND\_PREFIX"Include all templates that contain the values or only templates
that contain the values as the prefix.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

isEnabled()

QuickResponses

All content copied from https://docs.aws.amazon.com/.
