# Data Model

When using Amazon SimpleDB, you organize your structured data in domains within which you can put
data, get data, or run queries.

Domains consist of items which are described by _attribute_ name-value
pairs. Consider the spreadsheet model shown in the following
image.

![Spreadsheet model showing items with attributes in columns and query domains in tabs.](https://docs.aws.amazon.com/images/AmazonSimpleDB/latest/DeveloperGuide/images/spreadsheet3.png)

The components correspond to each part of a spreadsheet:

- Customer Account—Represented by the entire
spreadsheet, it refers to the Amazon Web Services _account_ to
which all domains are assigned.

- Domains—Represented by the domain worksheet
tabs at the bottom of the spreadsheet, domains are similar to tables that contain
similar data.

You can execute queries against a domain, but cannot execute queries across
different domains.

- Items—Represented by rows, items represent
individual objects that contain one or more attribute name-value pairs.

- Attributes—Represented by columns,
attributes represent categories of data that can be assigned to items.

- Values—Represented by cells, values
represent instances of attributes for items. An attribute can have multiple
values.

Unlike a spreadsheet, however, multiple values can be associated with a cell. For example,
an item can have both the color value _red_ and _blue_.
Additionally, Amazon SimpleDB does not require the presence of specific attributes. You can create a
single domain that contains completely different product types. For example, the following
table contains clothing, automotive parts, and motorcycle parts.

IDCategorySubcat.NameColorSizeMakeModelItem\_01ClothesSweaterCathair SweaterSiameseSmall, Medium, LargeItem\_02ClothesPantsDesigner JeansPaisley Acid Wash30x32, 32x32, 32x34Item\_03ClothesPantsSweatpantsBlue, Yellow, PinkLargeItem\_04Car PartsEngineTurbosAudiS4Item\_05Car PartsEmissions02 SensorAudiS4Item\_06Motorcycle PartsBodyworkFender EliminatorBlueYamahaR1Item\_07Motorcycle Parts, ClothingClothingLeather PantsBlackSmall, Medium, Large

Regardless of how you store your data, Amazon SimpleDB automatically indexes your data for quick
and accurate retrieval.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon SimpleDB Concepts

Operations

All content copied from https://docs.aws.amazon.com/.
