# Retrieve quick responses that match a search query in Amazon Connect Agent Workspace

Returns the SearchQuickResponsesResult object, which contains the matching quick
response results and a token to retrieve the next page of results, if available. The
SearchQuickResponsesRequest object is used to configure the search, allowing you to
filter by various criteria, as well as specify the channels, user-defined contact
attributes, and pagination options. If no queries are provided, the method will
return all quick responses associated with the agent's routing profile.

**Signature**

```

searchQuickResponses(queryRequest: SearchQuickResponsesRequest): Promise<SearchQuickResponsesResult>
```

**SearchQuickResponsesResult Properties**

**Parameter****Type****Description**resultsQuickResponsesSearchResultData\[\]The results of the quick responses searchnextTokenstringThe token for the next set of results. Use the value returned in
the previous response in the next request to retrieve the next set
of results.

**QuickResponsesSearchResultData Properties**

**Parameter****Type****Description**contentsQuickResponseContentsThe contents of the quick response.knowledgeBaseIdstringThe identifier of the knowledge base.namestringThe name of the quick response.quickResponseArnstringThe Amazon Resource Name (ARN) of the quick response.quickResponseIdstringThe identifier of the quick response.descriptionstringThe description of the quick response.shortcutKeystringThe shortcut key of the quick response. The value should be
unique across the knowledge base.attributesNotInterpolatedstring\[\]The user defined contact attributes that are not resolved when
the search result is returned.

**QuickResponseContents Properties**

**Parameter****Type****Description**markdownstringThe content of the quick response stored in markdownplainTextstringThe content of the quick response stored in plain text

**SearchQuickResponsesRequest Properties**

**Parameter****Type****Description**queriesQuickResponsesQuery\[\]Query used to filter quick responses; if no queries are provided,
the client will return all quick responses associated with the
agent's routing profile.channelsQuickResponseChannel\[\]The channels to filter the request by. Supported values: "Chat"
or "Email"attributesRecord<string, string>The user-defined Amazon Connect contact attributes to be resolved
when search results are returned.debounceMSnumberThe default value is set to 250ms; set it to 0 to disable
debounced input changemaxResultsnumberThe number of results to be returned. Minimum value of 1. Maximum
value of 100.nextTokenstringThe token for the next set of results. Use the value returned in
the previous response in the next request to retrieve the next set
of results.

**QuickResponsesQuery Properties**

**Parameter****Type****Description**nameQuickResponsesQueryFieldNameThe name of the attribute to query the quick responses by.
Supported values: "content", "name", "description", or
"shortcutKey"valuesstring\[\]The values of the attribute to query the quick responses
by.operatorQuickResponsesQueryOperatorThe operator to use for matching attribute field values in the
query. Supported values: "CONTAINS" or "CONTAINS\_AND\_PREFIX"priorityQuickResponsesQueryPriorityThe importance of the attribute field when calculating query
result relevancy scores. The value set for this parameter affects
the ordering of search results. Supported values: "HIGH", "MEDIUM",
or "LOW"allowFuzzinessbooleanWhether the query expects only exact matches on the attribute
field values. The results of the query will only include exact
matches if this parameter is set to false.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

isEnabled()

User

All content copied from https://docs.aws.amazon.com/.
