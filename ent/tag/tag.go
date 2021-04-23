// Code generated by entc, DO NOT EDIT.

package tag

const (
	// Label holds the string label denoting the tag type in the database.
	Label = "tag"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"

	// Table holds the table name of the tag in the database.
	Table = "tags"
)

// Columns holds all SQL columns for tag fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldName,
	FieldDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Tag type.
var ForeignKeys = []string{
	"agent_status_agent_status_to_tag",
	"command_command_to_tag",
	"competition_competition_to_tag",
	"dns_dns_to_tag",
	"dns_record_dns_record_to_tag",
	"disk_disk_to_tag",
	"environment_environment_to_tag",
	"file_delete_file_delete_to_tag",
	"file_download_file_download_to_tag",
	"file_extract_file_extract_to_tag",
	"finding_finding_to_tag",
	"host_host_to_tag",
	"included_network_included_network_to_tag",
	"network_network_to_tag",
	"provisioning_step_provisioning_step_to_tag",
	"script_script_to_tag",
	"user_user_to_tag",
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
