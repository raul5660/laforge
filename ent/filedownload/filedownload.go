// Code generated by entc, DO NOT EDIT.

package filedownload

const (
	// Label holds the string label denoting the filedownload type in the database.
	Label = "file_download"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSourceType holds the string denoting the source_type field in the database.
	FieldSourceType = "source_type"
	// FieldSource holds the string denoting the source field in the database.
	FieldSource = "source"
	// FieldDestination holds the string denoting the destination field in the database.
	FieldDestination = "destination"
	// FieldTemplate holds the string denoting the template field in the database.
	FieldTemplate = "template"
	// FieldPerms holds the string denoting the perms field in the database.
	FieldPerms = "perms"
	// FieldDisabled holds the string denoting the disabled field in the database.
	FieldDisabled = "disabled"
	// FieldMd5 holds the string denoting the md5 field in the database.
	FieldMd5 = "md5"
	// FieldAbsPath holds the string denoting the abs_path field in the database.
	FieldAbsPath = "abs_path"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"

	// EdgeFileDownloadToTag holds the string denoting the filedownloadtotag edge name in mutations.
	EdgeFileDownloadToTag = "FileDownloadToTag"

	// Table holds the table name of the filedownload in the database.
	Table = "file_downloads"
	// FileDownloadToTagTable is the table the holds the FileDownloadToTag relation/edge.
	FileDownloadToTagTable = "tags"
	// FileDownloadToTagInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	FileDownloadToTagInverseTable = "tags"
	// FileDownloadToTagColumn is the table column denoting the FileDownloadToTag relation/edge.
	FileDownloadToTagColumn = "file_download_file_download_to_tag"
)

// Columns holds all SQL columns for filedownload fields.
var Columns = []string{
	FieldID,
	FieldSourceType,
	FieldSource,
	FieldDestination,
	FieldTemplate,
	FieldPerms,
	FieldDisabled,
	FieldMd5,
	FieldAbsPath,
	FieldTags,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the FileDownload type.
var ForeignKeys = []string{
	"provisioning_step_provisioning_step_to_file_download",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
