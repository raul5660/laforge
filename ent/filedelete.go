// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/gen0cide/laforge/ent/filedelete"
)

// FileDelete is the model entity for the FileDelete schema.
type FileDelete struct {
	config `hcl:"-" json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Path holds the value of the "path" field.
	Path string `json:"path,omitempty" hcl:"path,attr"`
	// Tags holds the value of the "tags" field.
	Tags map[string]string `json:"tags,omitempty" hcl:"tags,attr"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FileDeleteQuery when eager-loading is set.
	Edges                                              FileDeleteEdges `json:"edges"`
	provisioning_step_provisioning_step_to_file_delete *int
}

// FileDeleteEdges holds the relations/edges for other nodes in the graph.
type FileDeleteEdges struct {
	// FileDeleteToTag holds the value of the FileDeleteToTag edge.
	FileDeleteToTag []*Tag
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// FileDeleteToTagOrErr returns the FileDeleteToTag value or an error if the edge
// was not loaded in eager-loading.
func (e FileDeleteEdges) FileDeleteToTagOrErr() ([]*Tag, error) {
	if e.loadedTypes[0] {
		return e.FileDeleteToTag, nil
	}
	return nil, &NotLoadedError{edge: "FileDeleteToTag"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FileDelete) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // path
		&[]byte{},         // tags
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*FileDelete) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // provisioning_step_provisioning_step_to_file_delete
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FileDelete fields.
func (fd *FileDelete) assignValues(values ...interface{}) error {
	if m, n := len(values), len(filedelete.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	fd.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field path", values[0])
	} else if value.Valid {
		fd.Path = value.String
	}

	if value, ok := values[1].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field tags", values[1])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &fd.Tags); err != nil {
			return fmt.Errorf("unmarshal field tags: %v", err)
		}
	}
	values = values[2:]
	if len(values) == len(filedelete.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field provisioning_step_provisioning_step_to_file_delete", value)
		} else if value.Valid {
			fd.provisioning_step_provisioning_step_to_file_delete = new(int)
			*fd.provisioning_step_provisioning_step_to_file_delete = int(value.Int64)
		}
	}
	return nil
}

// QueryFileDeleteToTag queries the FileDeleteToTag edge of the FileDelete.
func (fd *FileDelete) QueryFileDeleteToTag() *TagQuery {
	return (&FileDeleteClient{config: fd.config}).QueryFileDeleteToTag(fd)
}

// Update returns a builder for updating this FileDelete.
// Note that, you need to call FileDelete.Unwrap() before calling this method, if this FileDelete
// was returned from a transaction, and the transaction was committed or rolled back.
func (fd *FileDelete) Update() *FileDeleteUpdateOne {
	return (&FileDeleteClient{config: fd.config}).UpdateOne(fd)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (fd *FileDelete) Unwrap() *FileDelete {
	tx, ok := fd.config.driver.(*txDriver)
	if !ok {
		panic("ent: FileDelete is not a transactional entity")
	}
	fd.config.driver = tx.drv
	return fd
}

// String implements the fmt.Stringer.
func (fd *FileDelete) String() string {
	var builder strings.Builder
	builder.WriteString("FileDelete(")
	builder.WriteString(fmt.Sprintf("id=%v", fd.ID))
	builder.WriteString(", path=")
	builder.WriteString(fd.Path)
	builder.WriteString(", tags=")
	builder.WriteString(fmt.Sprintf("%v", fd.Tags))
	builder.WriteByte(')')
	return builder.String()
}

// FileDeletes is a parsable slice of FileDelete.
type FileDeletes []*FileDelete

func (fd FileDeletes) config(cfg config) {
	for _i := range fd {
		fd[_i].config = cfg
	}
}
