# ServiceNow Online data source connector field mappings

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
ServiceNow connector supports the following entities and the associated
reserved and custom attributes.

###### Supported entities and field mappings

- [Knowledge articles](#servicenow-field-mappings-ka)

- [Service catalog](#servicenow-field-mappings-sc)

- [Attachments](#servicenow-field-mappings-attachment)

- [Incidents](#servicenow-field-mappings-incidents)

## Knowledge articles

Amazon Q supports crawling [ServiceNow Online Knowledge articles](https://docs.servicenow.com/bundle/xanadu-servicenow-platform/page/product/knowledge-management/task/create-knowledge-article.html) and offers the following knowledge
article field mappings.

ServiceNow field nameIndex field nameDescriptionData typetextsn\_ka\_textCustomStringshort\_descriptionsn\_ka\_short\_descriptionCustomStringsys\_created\_on\_created\_atDefaultDatesys\_updated\_on\_last\_updated\_atDefaultDatekb\_category\_name\_categoryDefaultStringsys\_created\_by\_authorsDefaultStringsys\_updated\_bysn\_updatedByCustomStringsys\_idsn\_sys\_idCustomStringpublishedsn\_ka\_publish\_dateCustomDateworkflow\_statesn\_ka\_workflow\_stateCustomStringkb\_categorysn\_ka\_categoryCustomStringarticle\_typesn\_ka\_article\_typeCustomStringfirst\_namesn\_ka\_first\_nameCustomStringlast\_namesn\_ka\_last\_nameCustomStringuser\_namesn\_ka\_user\_nameCustomStringvalid\_tosn\_ka\_valid\_toCustomDatekb\_knowledge\_basesn\_ka\_knowledge\_baseCustomStringnumbersn\_ka\_numberCustomStringurlsn\_urlCustomStringdiplayUrl\_source\_uriDefaultStringreplacement\_articlesn\_ka\_replacement\_articleCustomStringdescriptionsn\_ka\_descriptionCustomStringwikisn\_ka\_wikiCustomStringratingsn\_ka\_ratingCustomStringratingsn\_ka\_ratingCustomStringview\_as\_allowedsn\_ka\_view\_as\_allowedCustomStringsourcesn\_ka\_sourceCustomStringimagesn\_ka\_imageCustomStringauthorsn\_ka\_authorCustomStringactivesn\_ka\_activeCustomStringhelpful\_countsn\_ka\_helpful\_countCustomStringmeta\_descriptionsn\_ka\_meta\_descriptionCustomStringmetasn\_ka\_metaCustomStringtopicsn\_ka\_topicCustomStringrolessn\_ka\_rolesCustomStringdisable\_suggestingsn\_ka\_disable\_suggestingCustomStringuse\_countsn\_ka\_use\_countCustomStringflaggedsn\_ka\_flaggedCustomStringdisable\_commentingsn\_ka\_disable\_commentingCustomStringretiredsn\_ka\_retiredCustomStringdisplay\_attachmentssn\_ka\_display\_attachmentsCustomStringtaxonomy\_topicsn\_ka\_taxonomy\_topicCustomString

## Service catalog

Amazon Q supports crawling [ServiceNow Online service catalogs](https://docs.servicenow.com/bundle/vancouver-servicenow-platform/page/product/service-catalog-management/concept/service-catalog.html) and offers the following service
catalog field mappings.

ServiceNow field nameIndex field nameDescriptionData typedescriptionsn\_sc\_descriptionCustomStringshort\_descriptionsn\_sc\_short\_descriptionCustomStringsys\_created\_on\_created\_atDefaultDatesys\_updated\_on\_last\_updated\_atDefaultDatecategory\_name\_categoryDefaultStringsys\_created\_by\_authorsDefaultString listsys\_updated\_bysn\_updated\_byCustomStringsys\_idsn\_sys\_idCustomStringsc\_catalogssn\_sc\_catalogsCustomStringsc\_catalogs\_namesn\_sc\_catalogs\_nameCustomStringcategorysn\_sc\_categoryCustomStringcategory\_full\_namesn\_sc\_categoryCustomStringurlsn\_urlCustomStringdisplayUrl\_source\_uriDefaultStringshow\_variable\_help\_on\_loadsn\_sc\_show\_var\_help\_on\_loadCustomStringno\_order\_nowsn\_sc\_no\_order\_nowCustomStringsc\_ic\_versionsn\_sc\_sc\_ic\_versionCustomStringdelivery\_timesn\_sc\_deliver\_timeCustomStringpublished\_refsn\_sc\_published\_refCustomStringpricesn\_sc\_priceCustomStringrecurring\_frequencysn\_sc\_recurring\_frequencyCustomStringsys\_namesn\_sc\_sys\_nameCustomStringmodelsn\_sc\_modelCustomStringstatesn\_sc\_stateCustomStringno\_cartsn\_sc\_no\_cartCustomStringgroupsn\_sc\_groupCustomStringhide\_spsn\_sc\_hide\_spCustomStringordersn\_sc\_orderCustomStringstart\_closedsn\_sc\_start\_closedCustomStringimagesn\_sc\_imageCustomStringno\_quantitysn\_sc\_no\_quantityCustomStringdelivery\_plansn\_sc\_delivery\_planCustomStringactivesn\_sc\_activeCustomStringchecked\_outsn\_sc\_checked\_outCustomStringcustom\_cartsn\_sc\_custom\_cartCustomStringno\_cart\_v2sn\_sc\_no\_cart\_v2CustomStringno\_proceed\_checkoutsn\_sc\_no\_proceed\_checkoutCustomStringignore\_pricesn\_sc\_ignore\_priceCustomStringsys\_update\_namesn\_sc\_sys\_update\_nameCustomStringmetasn\_sc\_metaCustomStringomit\_pricesn\_sc\_omit\_priceCustomStringnamesn\_sc\_nameCustomStringmobile\_hide\_pricesn\_sc\_mobile\_hide\_priceCustomStringno\_wishlist\_v2sn\_sc\_no\_wishlist\_v2CustomStringpreviewsn\_sc\_previewCustomStringtypesn\_sc\_typeCustomStringaccess\_typesn\_sc\_access\_typeCustomStringrolessn\_sc\_rolesCustomStringiconsn\_sc\_iconCustomStringmobile\_picturesn\_sc\_mobile\_pictureCustomStringavailabilitysn\_sc\_availabilityCustomStringmandatory\_attachmentsn\_sc\_mandatory\_attachmentCustomStringrequest\_methodsn\_sc\_request\_methodCustomStringvisible\_guidesn\_sc\_visible\_guideCustomStringvisible\_standalonesn\_sc\_visible\_standaloneCustomStringno\_ordersn\_sc\_no\_orderCustomStringvendorsn\_sc\_vendorCustomStringno\_attachment\_v2sn\_sc\_no\_attachment\_v2CustomStringmobile\_picture\_typesn\_sc\_mobile\_picture\_typeCustomStringvisible\_bundlesn\_sc\_visible\_bundleCustomStringordered\_item\_linksn\_sc\_ordered\_item\_linkCustomStringownersn\_sc\_ownerCustomStringno\_delivery\_time\_v2sn\_sc\_no\_delivery\_time\_v2CustomStringcostsn\_sc\_costCustomStringno\_quantity\_v2sn\_sc\_no\_quantity\_v2CustomStringrecurring\_pricesn\_sc\_recurring\_priceCustomStringlist\_pricesn\_sc\_list\_priceCustomStringsyst\_tagssn\_sc\_sys\_tagsCustomStringbillablesn\_sc\_billableCustomStringpicturesn\_sc\_pictureCustomStringdisplay\_price\_propertysn\_sc\_display\_price\_propertyCustomStringtaxonomy\_topicsn\_sc\_taxonomy\_topicCustomStringdelivery\_plain\_scriptsn\_sc\_delivery\_plain\_scriptCustomStringlocationsn\_sc\_locationCustomString

## Attachments

Amazon Q supports crawling [ServiceNow Online attachments](https://docs.servicenow.com/bundle/tokyo-platform-user-interface/page/use/using-forms/task/t_AddingAnAttachment.html) and offers the following attachment field
mappings.

ServiceNow field nameIndex field nameDescriptionData typesize\_bytessn\_file\_sizeCustomLong (numeric)file\_namesn\_file\_nameCustomStringsys\_mod\_countsn\_sys\_mod\_countCustomStringaverage\_image\_colorsn\_average\_image\_colorCustomStringimage\_widthsn\_image\_widthCustomStringsys\_updated\_on\_last\_updated\_atDefaultDatesys\_tagssn\_sys\_tagsCustomStringtable\_namesn\_table\_nameCustomStringsys\_idsn\_sys\_idCustomStringimage\_heightsn\_image\_heightCustomStringsys\_updated\_bysn\_updated\_byCustomStringcontent\_typesn\_content\_typeCustomStringsys\_created\_on\_created\_atDefaultDatesize\_compressedsn\_size\_compressedCustomStringcompressedsn\_compressedCustomStringstatesn\_stateCustomStringtable\_sys\_idsn\_table\_sys\_idCustomStringchunk\_size\_bytessn\_chunk\_size\_bytesCustomStringhashsn\_hashCustomStringsys\_created\_by\_authorsDefaultString listsys\_updated\_bysn\_updated\_byCustomStringurlsn\_urlCustomStringdisplayUrl\_source\_uriDefaultString

## Incidents

Amazon Q supports crawling [ServiceNow Online incidents](https://docs.servicenow.com/bundle/tokyo-it-service-management/page/product/incident-management/concept/c_IncidentManagement.html) and offers the following incident field
mappings.

ServiceNow field nameIndex field nameDescriptionData typeshort\_descriptionsn\_inc\_short\_descriptionCustomStringdescriptionsn\_inc\_descriptionCustomStringsys\_updated\_on\_last\_updated\_atDefaultDatenumbersn\_inc\_numberCustomStringsys\_updated\_bysn\_updatedByCustomStringdisplayUrl\_source\_uriDefaultStringopened\_bysn\_inc\_opened\_byCustomStringsys\_created\_on\_created\_atDefaultDatestatesn\_inc\_stateCustomStringsys\_created\_by\_authorsDefaultString listbusiness\_impactsn\_inc\_business\_impactDefaultStringimpactsn\_inc\_business\_impactCustomStringprioritysn\_inc\_priorityCustomStringurgencysn\_inc\_urgencyCustomStringopened\_atan\_inc\_opened\_atCustomStringbusiness\_durationsn\_inc\_business\_durationCustomStringcaller\_idsn\_inc\_caller\_idCustomStringresolved\_atsn\_inc\_resolved\_atCustomStringcategorysn\_inc\_categoryCustomStringsubcategorysn\_inc\_subcategoryCustomStringclose\_codesn\_inc\_close\_codeCustomStringassignment\_groupsn\_inc\_assignment\_groupCustomStringclose\_notessn\_inc\_close\_notesCustomStringdisplayUrl\_source\_uriDefaultStringsys\_class\_namesn\_inc\_sys\_class\_nameCustomStringparent\_incidentan\_inc\_parent\_incidentCustomStringincident\_statesn\_incident\_stateCustomStringcompanysn\_inc\_companyCustomStringassigned\_tosn\_inc\_assigned\_toCustomStringhold\_reasonan\_inc\_hold\_reasonCustomStringwork\_notessn\_inc\_work\_notesCustomStringcomments\_and\_work\_notessn\_inc\_comments\_and\_work\_notesCustomStringwork\_notes\_listsn\_work\_notes\_listCustomStringcommentssn\_inc\_commentsCustomStringsys\_idsn\_inc\_sys\_idCustomStringurlsn\_urlCustomStringactivesn\_inc\_activeCustomStringactivity\_duesn\_inc\_activity\_dueCustomStringadditional\_assignee\_listsn\_inc\_additional\_assign\_listCustomStringapprovalsn\_inc\_approvalCustomStringapproval\_historysn\_inc\_approval\_historyCustomStringapproval\_setsn\_inc\_approval\_setCustomDatebusiness\_servicesn\_inc\_business\_serviceCustomStringclosed\_bysn\_inc\_closed\_byCustomStringcmdb\_cisn\_inc\_cmdb\_idCustomStringresolved\_bysn\_inc\_resolved\_byCustomStringsys\_domainsn\_inc\_sys\_domainCustomStringbusiness\_stcsn\_inc\_business\_stcCustomStringcalendar\_durationsn\_inc\_calendar\_durationCustomStringcalendar\_stcsn\_inc\_calendar\_stcCustomStringcausesn\_inc\_causeCustomStringcaused\_bysn\_inc\_caused\_byCustomStringchild\_incidentssn\_inc\_child\_incidentsCustomStringclosed\_atsn\_inc\_closed\_atCustomStringcontact\_typesn\_inc\_contact\_typeCustomStringcontractsn\_inc\_contractCustomStringcorrelation\_displaysn\_inc\_correlation\_displayCustomStringdelivery\_plansn\_inc\_delivery\_planCustomStringdelivery\_tasksn\_inc\_delivery\_taskCustomStringdue\_datesn\_inc\_due\_dateCustomStringescalationsn\_inc\_escalationCustomStringexpected\_startsn\_inc\_expected\_startCustomStringfollow\_upsn\_inc\_follow\_upCustomStringgroup\_listsn\_inc\_group\_listCustomStringknowledgesn\_inc\_knowledgeCustomStringlocationsn\_inc\_locationCustomStringmade\_slasn\_inc\_made\_slaCustomStringnotifysn\_inc\_notifyCustomStringordersn\_inc\_orderCustomStringorigin\_idsn\_inc\_origin\_idCustomStringorigin\_tablesn\_inc\_origin\_tableCustomStringparentsn\_inc\_parentCustomStringproblem\_idsn\_inc\_problem\_idCustomStringreassignment\_countsn\_inc\_reassignment\_countCustomStringrepoen\_countsn\_inc\_reopen\_countCustomStringreopened\_bysn\_inc\_reopened\_byCustomStringreopened\_timesn\_inc\_reopened\_timeCustomStringrfcsn\_inc\_rfcCustomStringroute\_reasonsn\_inc\_route\_reasonCustomStringservice\_offeringsn\_inc\_service\_offeringCustomStringseveritysn\_inc\_severityCustomStringsla\_duesn\_inc\_sla\_dueCustomDatetask\_effective\_numbersn\_inc\_task\_effective\_numberCustomStringtime\_workedsn\_inc\_time\_workedCustomDateuniversal\_requestsn\_inc\_universal\_requestCustomStringupon\_approvalsn\_inc\_upon\_approvalCustomStringupon\_rejectsn\_inc\_upon\_rejectCustomStringuser\_inputsn\_inc\_user\_inputCustomStringwatch\_listsn\_inc\_watch\_listCustomStringwork\_endsn\_inc\_work\_endCustomStringwork\_startsn\_inc\_work\_startCustomString

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ACL crawling

IAM role

All content copied from https://docs.aws.amazon.com/.
