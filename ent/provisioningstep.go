// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/gen0cide/laforge/ent/provisioningstep"
)

// ProvisioningStep is the model entity for the ProvisioningStep schema.
type ProvisioningStep struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ProvisionerType holds the value of the "provisioner_type" field.
	ProvisionerType string `json:"provisioner_type,omitempty"`
	// StepNumber holds the value of the "step_number" field.
	StepNumber int `json:"step_number,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProvisioningStepQuery when eager-loading is set.
	Edges ProvisioningStepEdges `json:"edges"`
}

// ProvisioningStepEdges holds the relations/edges for other nodes in the graph.
type ProvisioningStepEdges struct {
	// ProvisioningStepToTag holds the value of the ProvisioningStepToTag edge.
	ProvisioningStepToTag []*Tag `json:"ProvisioningStepToTag,omitempty"`
	// ProvisioningStepToStatus holds the value of the ProvisioningStepToStatus edge.
	ProvisioningStepToStatus []*Status `json:"ProvisioningStepToStatus,omitempty"`
	// ProvisioningStepToProvisionedHost holds the value of the ProvisioningStepToProvisionedHost edge.
	ProvisioningStepToProvisionedHost []*ProvisionedHost `json:"ProvisioningStepToProvisionedHost,omitempty"`
	// ProvisioningStepToScript holds the value of the ProvisioningStepToScript edge.
	ProvisioningStepToScript []*Script `json:"ProvisioningStepToScript,omitempty"`
	// ProvisioningStepToCommand holds the value of the ProvisioningStepToCommand edge.
	ProvisioningStepToCommand []*Command `json:"ProvisioningStepToCommand,omitempty"`
	// ProvisioningStepToDNSRecord holds the value of the ProvisioningStepToDNSRecord edge.
	ProvisioningStepToDNSRecord []*DNSRecord `json:"ProvisioningStepToDNSRecord,omitempty"`
	// ProvisioningStepToFileDelete holds the value of the ProvisioningStepToFileDelete edge.
	ProvisioningStepToFileDelete []*FileDelete `json:"ProvisioningStepToFileDelete,omitempty"`
	// ProvisioningStepToFileDownload holds the value of the ProvisioningStepToFileDownload edge.
	ProvisioningStepToFileDownload []*FileDownload `json:"ProvisioningStepToFileDownload,omitempty"`
	// ProvisioningStepToFileExtract holds the value of the ProvisioningStepToFileExtract edge.
	ProvisioningStepToFileExtract []*FileExtract `json:"ProvisioningStepToFileExtract,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [9]bool
}

// ProvisioningStepToTagOrErr returns the ProvisioningStepToTag value or an error if the edge
// was not loaded in eager-loading.
func (e ProvisioningStepEdges) ProvisioningStepToTagOrErr() ([]*Tag, error) {
	if e.loadedTypes[0] {
		return e.ProvisioningStepToTag, nil
	}
	return nil, &NotLoadedError{edge: "ProvisioningStepToTag"}
}

// ProvisioningStepToStatusOrErr returns the ProvisioningStepToStatus value or an error if the edge
// was not loaded in eager-loading.
func (e ProvisioningStepEdges) ProvisioningStepToStatusOrErr() ([]*Status, error) {
	if e.loadedTypes[1] {
		return e.ProvisioningStepToStatus, nil
	}
	return nil, &NotLoadedError{edge: "ProvisioningStepToStatus"}
}

// ProvisioningStepToProvisionedHostOrErr returns the ProvisioningStepToProvisionedHost value or an error if the edge
// was not loaded in eager-loading.
func (e ProvisioningStepEdges) ProvisioningStepToProvisionedHostOrErr() ([]*ProvisionedHost, error) {
	if e.loadedTypes[2] {
		return e.ProvisioningStepToProvisionedHost, nil
	}
	return nil, &NotLoadedError{edge: "ProvisioningStepToProvisionedHost"}
}

// ProvisioningStepToScriptOrErr returns the ProvisioningStepToScript value or an error if the edge
// was not loaded in eager-loading.
func (e ProvisioningStepEdges) ProvisioningStepToScriptOrErr() ([]*Script, error) {
	if e.loadedTypes[3] {
		return e.ProvisioningStepToScript, nil
	}
	return nil, &NotLoadedError{edge: "ProvisioningStepToScript"}
}

// ProvisioningStepToCommandOrErr returns the ProvisioningStepToCommand value or an error if the edge
// was not loaded in eager-loading.
func (e ProvisioningStepEdges) ProvisioningStepToCommandOrErr() ([]*Command, error) {
	if e.loadedTypes[4] {
		return e.ProvisioningStepToCommand, nil
	}
	return nil, &NotLoadedError{edge: "ProvisioningStepToCommand"}
}

// ProvisioningStepToDNSRecordOrErr returns the ProvisioningStepToDNSRecord value or an error if the edge
// was not loaded in eager-loading.
func (e ProvisioningStepEdges) ProvisioningStepToDNSRecordOrErr() ([]*DNSRecord, error) {
	if e.loadedTypes[5] {
		return e.ProvisioningStepToDNSRecord, nil
	}
	return nil, &NotLoadedError{edge: "ProvisioningStepToDNSRecord"}
}

// ProvisioningStepToFileDeleteOrErr returns the ProvisioningStepToFileDelete value or an error if the edge
// was not loaded in eager-loading.
func (e ProvisioningStepEdges) ProvisioningStepToFileDeleteOrErr() ([]*FileDelete, error) {
	if e.loadedTypes[6] {
		return e.ProvisioningStepToFileDelete, nil
	}
	return nil, &NotLoadedError{edge: "ProvisioningStepToFileDelete"}
}

// ProvisioningStepToFileDownloadOrErr returns the ProvisioningStepToFileDownload value or an error if the edge
// was not loaded in eager-loading.
func (e ProvisioningStepEdges) ProvisioningStepToFileDownloadOrErr() ([]*FileDownload, error) {
	if e.loadedTypes[7] {
		return e.ProvisioningStepToFileDownload, nil
	}
	return nil, &NotLoadedError{edge: "ProvisioningStepToFileDownload"}
}

// ProvisioningStepToFileExtractOrErr returns the ProvisioningStepToFileExtract value or an error if the edge
// was not loaded in eager-loading.
func (e ProvisioningStepEdges) ProvisioningStepToFileExtractOrErr() ([]*FileExtract, error) {
	if e.loadedTypes[8] {
		return e.ProvisioningStepToFileExtract, nil
	}
	return nil, &NotLoadedError{edge: "ProvisioningStepToFileExtract"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProvisioningStep) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case provisioningstep.FieldID, provisioningstep.FieldStepNumber:
			values[i] = &sql.NullInt64{}
		case provisioningstep.FieldProvisionerType:
			values[i] = &sql.NullString{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type ProvisioningStep", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProvisioningStep fields.
func (ps *ProvisioningStep) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case provisioningstep.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ps.ID = int(value.Int64)
		case provisioningstep.FieldProvisionerType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field provisioner_type", values[i])
			} else if value.Valid {
				ps.ProvisionerType = value.String
			}
		case provisioningstep.FieldStepNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field step_number", values[i])
			} else if value.Valid {
				ps.StepNumber = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryProvisioningStepToTag queries the "ProvisioningStepToTag" edge of the ProvisioningStep entity.
func (ps *ProvisioningStep) QueryProvisioningStepToTag() *TagQuery {
	return (&ProvisioningStepClient{config: ps.config}).QueryProvisioningStepToTag(ps)
}

// QueryProvisioningStepToStatus queries the "ProvisioningStepToStatus" edge of the ProvisioningStep entity.
func (ps *ProvisioningStep) QueryProvisioningStepToStatus() *StatusQuery {
	return (&ProvisioningStepClient{config: ps.config}).QueryProvisioningStepToStatus(ps)
}

// QueryProvisioningStepToProvisionedHost queries the "ProvisioningStepToProvisionedHost" edge of the ProvisioningStep entity.
func (ps *ProvisioningStep) QueryProvisioningStepToProvisionedHost() *ProvisionedHostQuery {
	return (&ProvisioningStepClient{config: ps.config}).QueryProvisioningStepToProvisionedHost(ps)
}

// QueryProvisioningStepToScript queries the "ProvisioningStepToScript" edge of the ProvisioningStep entity.
func (ps *ProvisioningStep) QueryProvisioningStepToScript() *ScriptQuery {
	return (&ProvisioningStepClient{config: ps.config}).QueryProvisioningStepToScript(ps)
}

// QueryProvisioningStepToCommand queries the "ProvisioningStepToCommand" edge of the ProvisioningStep entity.
func (ps *ProvisioningStep) QueryProvisioningStepToCommand() *CommandQuery {
	return (&ProvisioningStepClient{config: ps.config}).QueryProvisioningStepToCommand(ps)
}

// QueryProvisioningStepToDNSRecord queries the "ProvisioningStepToDNSRecord" edge of the ProvisioningStep entity.
func (ps *ProvisioningStep) QueryProvisioningStepToDNSRecord() *DNSRecordQuery {
	return (&ProvisioningStepClient{config: ps.config}).QueryProvisioningStepToDNSRecord(ps)
}

// QueryProvisioningStepToFileDelete queries the "ProvisioningStepToFileDelete" edge of the ProvisioningStep entity.
func (ps *ProvisioningStep) QueryProvisioningStepToFileDelete() *FileDeleteQuery {
	return (&ProvisioningStepClient{config: ps.config}).QueryProvisioningStepToFileDelete(ps)
}

// QueryProvisioningStepToFileDownload queries the "ProvisioningStepToFileDownload" edge of the ProvisioningStep entity.
func (ps *ProvisioningStep) QueryProvisioningStepToFileDownload() *FileDownloadQuery {
	return (&ProvisioningStepClient{config: ps.config}).QueryProvisioningStepToFileDownload(ps)
}

// QueryProvisioningStepToFileExtract queries the "ProvisioningStepToFileExtract" edge of the ProvisioningStep entity.
func (ps *ProvisioningStep) QueryProvisioningStepToFileExtract() *FileExtractQuery {
	return (&ProvisioningStepClient{config: ps.config}).QueryProvisioningStepToFileExtract(ps)
}

// Update returns a builder for updating this ProvisioningStep.
// Note that you need to call ProvisioningStep.Unwrap() before calling this method if this ProvisioningStep
// was returned from a transaction, and the transaction was committed or rolled back.
func (ps *ProvisioningStep) Update() *ProvisioningStepUpdateOne {
	return (&ProvisioningStepClient{config: ps.config}).UpdateOne(ps)
}

// Unwrap unwraps the ProvisioningStep entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ps *ProvisioningStep) Unwrap() *ProvisioningStep {
	tx, ok := ps.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProvisioningStep is not a transactional entity")
	}
	ps.config.driver = tx.drv
	return ps
}

// String implements the fmt.Stringer.
func (ps *ProvisioningStep) String() string {
	var builder strings.Builder
	builder.WriteString("ProvisioningStep(")
	builder.WriteString(fmt.Sprintf("id=%v", ps.ID))
	builder.WriteString(", provisioner_type=")
	builder.WriteString(ps.ProvisionerType)
	builder.WriteString(", step_number=")
	builder.WriteString(fmt.Sprintf("%v", ps.StepNumber))
	builder.WriteByte(')')
	return builder.String()
}

// ProvisioningSteps is a parsable slice of ProvisioningStep.
type ProvisioningSteps []*ProvisioningStep

func (ps ProvisioningSteps) config(cfg config) {
	for _i := range ps {
		ps[_i].config = cfg
	}
}
