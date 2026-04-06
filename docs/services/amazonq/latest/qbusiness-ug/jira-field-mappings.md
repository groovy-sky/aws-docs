# Jira data source connector field mappings

To improve retrieved results and customize the end user chat experience, Amazon Q Business enables you to map document attributes from your data sources to fields
in your Amazon Q index.

Amazon Q offers two kinds of attributes to map to index fields:

- **Reserved or default** – Reserved attributes are
based on document attributes that commonly occur in most data. You can use reserved
attributes to map commonly occurring document attributes in your data source to
Amazon Q index fields.

- **Custom** – You can create custom attributes to map
document attributes that are unique to your data to Amazon Q index
fields.

When you connect Amazon Q to a data source, Amazon Q automatically
maps specific data source document attributes to fields within an Amazon Q index.
If a document attribute in your data source doesn't have a attribute mapping already
available, or if you want to map additional document attributes to index fields, use the
custom field mappings to specify how a data source attribute maps to an Amazon Q
index field. You create field mappings by editing your data source after your application
and retriever are created.

To learn more about document attributes and how they work in Amazon Q, see
[Document attributes and types in Amazon Q](doc-attributes.md).

###### Important

Filtering using document attributes in chat is only supported through the API.

The Amazon Q
Jira connector supports the following entities and the associated reserved
and custom attributes.

###### Supported entities and field mappings

- [Projects](#jira-field-mappings-projects)

- [Issues](#jira-field-mappings-ticket-issues)

- [Comments](#jira-field-mappings-comments)

- [Attachments](#jira-field-mappings-attachments)

- [Worklogs](#jira-field-mappings-worklogs)

## Projects

Jira field nameIndex field nameDescriptionData typetitlej\_titleCustomStringproject\_keyj\_project\_keyCustomStringleadj\_leadCustomString listurl\_source\_uriDefaultString

## Issues

Jira field nameIndex field nameDescriptionData typetitlej\_titleCustomStringissue\_keyj\_issue\_keyCustomStringstatusj\_statusCustomStringproject\_namej\_project\_nameCustomStringprojectKeyj\_project\_keyCustomStringauthors\_authorsDefaultString listassigneej\_assigneeCustomStringcreated\_at\_created\_atDefaultDateupdated\_at\_last\_updated\_atDefaultDateurl\_source\_uriDefaultStringissue\_typej\_issue\_typeCustomStringpriorityj\_priorityCustomStringresolutionj\_resolutionCustomStringaffects\_versionj\_affects\_versionCustomStringfix\_versionj\_fix\_versionCustomStringlabelsj\_labelsCustomStringenvironmentj\_environmentCustomStringreporterj\_reporterCustomStringvotesj\_votesCustomStringwatchersj\_watchersCustomStringduej\_dueCustomStringresolvedj\_resolvedCustomString

## Comments

Jira field nameIndex field nameDescriptionData typeauthors\_authorsDefaultString listtitlej\_titleCustomStringcreatedAt\_created\_atDefaultDateupdatedAt\_last\_updated\_atDefaultDateproject\_namej\_project\_nameCustomStringproject\_keyj\_project\_keyCustomStringissue\_keyj\_issue\_keyCustomStringurl\_source\_uriDefaultString

## Attachments

Jira field nameIndex field nameDescriptionData typetitlej\_titleCustomStringauthors\_authorsDefaultString listsizej\_sizeCustomStringcreatedAt\_created\_atDefaultDateurl\_source\_uriDefaultStringproject\_namej\_project\_nameCustomStringproject\_keyj\_project\_keyCustomStringissue\_keyj\_issue\_keyCustomString

## Worklogs

Jira field nameIndex field nameDescriptionData typetitlej\_titleCustomStringauthors\_authorsDefaultString listcreatedAt\_created\_atDefaultDateupdatedAt\_last\_updated\_atDefaultDateurl\_source\_uriDefaultStringproject\_namej\_project\_nameCustomStringproject\_keyj\_project\_keyCustomStringissue\_keyj\_issue\_keyCustomString

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ACL crawling

IAM role
