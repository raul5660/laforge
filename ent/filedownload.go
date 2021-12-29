// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/gen0cide/laforge/ent/environment"
	"github.com/gen0cide/laforge/ent/filedownload"
	"github.com/google/uuid"
)

// FileDownload is the model entity for the FileDownload schema.
type FileDownload struct {
	config ` json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// HclID holds the value of the "hcl_id" field.
	HclID string `json:"hcl_id,omitempty" hcl:"id,label"`
	// SourceType holds the value of the "source_type" field.
	SourceType string `json:"source_type,omitempty" hcl:"source_type,attr"`
	// Source holds the value of the "source" field.
	Source string `json:"source,omitempty" hcl:"source,attr"`
	// Destination holds the value of the "destination" field.
	Destination string `json:"destination,omitempty" hcl:"destination,attr"`
	// Template holds the value of the "template" field.
	Template bool `json:"template,omitempty" hcl:"template,optional"`
	// Perms holds the value of the "perms" field.
	Perms string `json:"perms,omitempty" hcl:"perms,optional"`
	// Disabled holds the value of the "disabled" field.
	Disabled bool `json:"disabled,omitempty" hcl:"disabled,optional"`
	// Md5 holds the value of the "md5" field.
	Md5 string `json:"md5,omitempty" hcl:"md5,optional"`
	// AbsPath holds the value of the "abs_path" field.
	AbsPath string `json:"abs_path,omitempty" hcl:"abs_path,optional"`
	// Tags holds the value of the "tags" field.
	Tags map[string]string `json:"tags,omitempty" hcl:"tags,optional"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FileDownloadQuery when eager-loading is set.
	Edges FileDownloadEdges `json:"edges"`

	// Edges put into the main struct to be loaded via hcl
	// FileDownloadToEnvironment holds the value of the FileDownloadToEnvironment edge.
	HCLFileDownloadToEnvironment *Environment `json:"FileDownloadToEnvironment,omitempty"`
	//
	environment_environment_to_file_download *uuid.UUID
}

// FileDownloadEdges holds the relations/edges for other nodes in the graph.
type FileDownloadEdges struct {
	// FileDownloadToEnvironment holds the value of the FileDownloadToEnvironment edge.
	FileDownloadToEnvironment *Environment `json:"FileDownloadToEnvironment,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// FileDownloadToEnvironmentOrErr returns the FileDownloadToEnvironment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FileDownloadEdges) FileDownloadToEnvironmentOrErr() (*Environment, error) {
	if e.loadedTypes[0] {
		if e.FileDownloadToEnvironment == nil {
			// The edge FileDownloadToEnvironment was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: environment.Label}
		}
		return e.FileDownloadToEnvironment, nil
	}
	return nil, &NotLoadedError{edge: "FileDownloadToEnvironment"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FileDownload) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case filedownload.FieldTags:
			values[i] = new([]byte)
		case filedownload.FieldTemplate, filedownload.FieldDisabled:
			values[i] = new(sql.NullBool)
		case filedownload.FieldHclID, filedownload.FieldSourceType, filedownload.FieldSource, filedownload.FieldDestination, filedownload.FieldPerms, filedownload.FieldMd5, filedownload.FieldAbsPath:
			values[i] = new(sql.NullString)
		case filedownload.FieldID:
			values[i] = new(uuid.UUID)
		case filedownload.ForeignKeys[0]: // environment_environment_to_file_download
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type FileDownload", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FileDownload fields.
func (fd *FileDownload) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case filedownload.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				fd.ID = *value
			}
		case filedownload.FieldHclID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hcl_id", values[i])
			} else if value.Valid {
				fd.HclID = value.String
			}
		case filedownload.FieldSourceType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source_type", values[i])
			} else if value.Valid {
				fd.SourceType = value.String
			}
		case filedownload.FieldSource:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source", values[i])
			} else if value.Valid {
				fd.Source = value.String
			}
		case filedownload.FieldDestination:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field destination", values[i])
			} else if value.Valid {
				fd.Destination = value.String
			}
		case filedownload.FieldTemplate:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field template", values[i])
			} else if value.Valid {
				fd.Template = value.Bool
			}
		case filedownload.FieldPerms:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field perms", values[i])
			} else if value.Valid {
				fd.Perms = value.String
			}
		case filedownload.FieldDisabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field disabled", values[i])
			} else if value.Valid {
				fd.Disabled = value.Bool
			}
		case filedownload.FieldMd5:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field md5", values[i])
			} else if value.Valid {
				fd.Md5 = value.String
			}
		case filedownload.FieldAbsPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field abs_path", values[i])
			} else if value.Valid {
				fd.AbsPath = value.String
			}
		case filedownload.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &fd.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case filedownload.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field environment_environment_to_file_download", values[i])
			} else if value.Valid {
				fd.environment_environment_to_file_download = new(uuid.UUID)
				*fd.environment_environment_to_file_download = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryFileDownloadToEnvironment queries the "FileDownloadToEnvironment" edge of the FileDownload entity.
func (fd *FileDownload) QueryFileDownloadToEnvironment() *EnvironmentQuery {
	return (&FileDownloadClient{config: fd.config}).QueryFileDownloadToEnvironment(fd)
}

// Update returns a builder for updating this FileDownload.
// Note that you need to call FileDownload.Unwrap() before calling this method if this FileDownload
// was returned from a transaction, and the transaction was committed or rolled back.
func (fd *FileDownload) Update() *FileDownloadUpdateOne {
	return (&FileDownloadClient{config: fd.config}).UpdateOne(fd)
}

// Unwrap unwraps the FileDownload entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fd *FileDownload) Unwrap() *FileDownload {
	tx, ok := fd.config.driver.(*txDriver)
	if !ok {
		panic("ent: FileDownload is not a transactional entity")
	}
	fd.config.driver = tx.drv
	return fd
}

// String implements the fmt.Stringer.
func (fd *FileDownload) String() string {
	var builder strings.Builder
	builder.WriteString("FileDownload(")
	builder.WriteString(fmt.Sprintf("id=%v", fd.ID))
	builder.WriteString(", hcl_id=")
	builder.WriteString(fd.HclID)
	builder.WriteString(", source_type=")
	builder.WriteString(fd.SourceType)
	builder.WriteString(", source=")
	builder.WriteString(fd.Source)
	builder.WriteString(", destination=")
	builder.WriteString(fd.Destination)
	builder.WriteString(", template=")
	builder.WriteString(fmt.Sprintf("%v", fd.Template))
	builder.WriteString(", perms=")
	builder.WriteString(fd.Perms)
	builder.WriteString(", disabled=")
	builder.WriteString(fmt.Sprintf("%v", fd.Disabled))
	builder.WriteString(", md5=")
	builder.WriteString(fd.Md5)
	builder.WriteString(", abs_path=")
	builder.WriteString(fd.AbsPath)
	builder.WriteString(", tags=")
	builder.WriteString(fmt.Sprintf("%v", fd.Tags))
	builder.WriteByte(')')
	return builder.String()
}

// FileDownloads is a parsable slice of FileDownload.
type FileDownloads []*FileDownload

func (fd FileDownloads) config(cfg config) {
	for _i := range fd {
		fd[_i].config = cfg
	}
}