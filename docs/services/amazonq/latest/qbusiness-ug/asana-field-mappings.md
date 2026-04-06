# Amazon Q Business Asana data source connector field mappings (Preview)

To improve retrieved results and customize the end user chat experience, Amazon Q Business enables you to map document attributes from your data sources to
fields in your Amazon Q index.

Amazon Q offers two kinds of attributes to map to index fields:

- **Reserved or default** – Reserved attributes are
based on document attributes that commonly occur in most data. You can use
reserved attributes to map commonly occurring document attributes in your data
source to Amazon Q index fields.

- **Custom** – You can create custom attributes to
map document attributes that are unique to your data to Amazon Q
index fields.

When you connect Amazon Q to a data source, Amazon Q
automatically maps specific data source document attributes to fields within an Amazon Q index. If a document attribute in your data source doesn't have a
attribute mapping already available, or if you want to map additional document
attributes to index fields, use the custom field mappings to specify how a data source
attribute maps to an Amazon Q index field. You create field mappings by
editing your data source after your application environment and retriever are created.

To learn more about document attributes and how they work in Amazon Q, see
[Document attributes and types in Amazon Q](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/doc-attributes-types.html).

###### Important

Filtering using document attributes in chat is only supported through the
API.

The Amazon Q Asana connector supports the following entities
and the associated reserved and custom attributes.

###### Supported entities and field mappings

- [Projects](#Asana-field-mappings-projects)

- [Tasks](#Asana-field-mappings-tasks)

- [Comments](#Asana-field-mappings-comments)

## Projects

Amazon Q supports crawling Asana Projects and offers the
following project field mappings.

Asana field nameIndex field nameDescriptionData typeprojectPermaLink\_source\_uriDefaultStringprojectCreatedDate\_created\_atDefaultDateprojectModifiedDate\_last\_updated\_atDefaultDatecatagory\_catagoryDefaultStringisArchivedasana\_archivedCustomCustomdueOnasana\_dueOnCustomDatestartOnasana\_startOnCustomStringisPublicProjectasana\_isPublicProjectCustomStringownerIdasana\_ownerIdCustomStringownerNameasana\_ownerNameCustomStringteamIdasana\_teamIdCustomString listteamNameasana\_teamNameCustomString listworkspaceIdasana\_workspaceIdCustomStringworkspaceNameasana\_workspaceNameCustomStringisOrganizationasana\_isOrganizationCustomString

## Tasks

Amazon Q supports crawling Asana Tasks and offers the
following project field mappings.

Asana field nameIndex field nameDescriptionData typetaskPermaLink\_source\_uriDefaultStringtaskCreatedDate\_created\_atDefaultDatetaskModifiedDate\_last\_update\_atDefaultDatecategory\_categoryDefaultStringassigneeIdasana\_assigneeIdCustomStringassigneeNameasana\_assigneeNameCustomStringisCompletedasana\_isCompletedCustomStringdueOnasana\_dueOnCustomStringisSubtaskasana\_isSubtaskCustomStringtopLevelTaskIdasana\_topLevelTaskIdCustomStringtopLevelTaskNameasana\_topLevelTaskNameCustomStringsection Idasana\_sectionIdCustomStringsectionNameasana\_sectionNameCustomStringprojectIdasana\_projectIdCustomStringprojectNameasana\_projectNameCustomStringworkspaceIdasana\_workspaceIdCustomStringworkspaceNameasana\_workspaceNameCustomStringisOrganizationasana\_isOrganizationCustomString

## Comments

Amazon Q supports crawling Asana Comments and offers the
following project field mappings.

Asana field nameIndex field nameDescriptionData typecommentPermaLink\_source\_uriDefaultStringcategory\_categoryDefaultStringtaskIdasana\_taskIdCustomStringtaskNameasana\_taskNameCustomStringteamIdasana\_teamIdCustomStringteamNameasana\_teamNameCustomStringsectionIdasana\_sectionIdCustomStringsectionNameasana\_sectionNameCustomStringprojectIdasana\_projectIdCustomStringprojectNameasana\_projectNameCustomStringworkspaceIdasana\_workspaceIdCustomStringworkspaceNameasana\_workspaceNameCustomStringisOrganizationasana\_isOrganizationCustomString

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Using the API

IAM role
