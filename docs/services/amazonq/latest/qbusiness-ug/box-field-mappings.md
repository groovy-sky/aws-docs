# Box data source connector field mappings

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
Box connector supports the following entities and the associated reserved and
custom attributes.

###### Supported entities and field mappings

- [Files and folders](#box-field-mappings-files-folders)

- [Comments](#box-field-mappings-comments)

- [Tasks](#box-field-mappings-tasks)

- [Web links](#box-field-mappings-web-links)

## Files and folders

Box field nameIndex field nameDescriptionData typebx\_createdAt\_created\_atDefaultDatebx\_modifiedAt\_last\_updated\_atDefaultDatebx\_authors\_authorsDefaultString listbx\_uri\_source\_uriDefaultStringbx\_sizebx\_file\_sizeCustomStringbx\_category\_categoryDefaultString

## Comments

Box field nameIndex field nameDescriptionData typebx\_createdAt\_created\_atDefaultDatebx\_modifiedAt\_last\_updated\_atDefaultDatebx\_author\_authorsCustomStringbx\_parentFilebx\_comment\_itemCustomStringbx\_category\_categoryDefaultString

## Tasks

Box field nameIndex field nameDescriptionData typebx\_createdAt\_created\_atDefaultDatebx\_actionbx\_task\_actionCustomStringbx\_taskCompletebx\_task\_completedCustomStringbx\_taskItembx\_task\_itemCustomStringbx\_taskAssignedbx\_task\_assigned\_toCustomStringbx\_authorbx\_authorCustomStringbx\_category\_categoryDefaultStringbx\_uri\_source\_uriDefaultString

## Web links

Box field nameIndex field nameDescriptionData typebx\_createdAt\_created\_atDefaultDatebx\_authorbx\_authorCustomStringbx\_category\_categoryDefaultStringbx\_uri\_source\_uriDefaultString

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ACL crawling

IAM role
