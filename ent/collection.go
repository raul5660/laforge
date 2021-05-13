// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (as *AgentStatusQuery) CollectFields(ctx context.Context, satisfies ...string) *AgentStatusQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		as = as.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return as
}

func (as *AgentStatusQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *AgentStatusQuery {
	return as
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (au *AuthUserQuery) CollectFields(ctx context.Context, satisfies ...string) *AuthUserQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		au = au.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return au
}

func (au *AuthUserQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *AuthUserQuery {
	return au
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (b *BuildQuery) CollectFields(ctx context.Context, satisfies ...string) *BuildQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		b = b.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return b
}

func (b *BuildQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *BuildQuery {
	return b
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (c *CommandQuery) CollectFields(ctx context.Context, satisfies ...string) *CommandQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		c = c.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return c
}

func (c *CommandQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *CommandQuery {
	return c
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (c *CompetitionQuery) CollectFields(ctx context.Context, satisfies ...string) *CompetitionQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		c = c.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return c
}

func (c *CompetitionQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *CompetitionQuery {
	return c
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (d *DNSQuery) CollectFields(ctx context.Context, satisfies ...string) *DNSQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		d = d.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return d
}

func (d *DNSQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *DNSQuery {
	return d
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (dr *DNSRecordQuery) CollectFields(ctx context.Context, satisfies ...string) *DNSRecordQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		dr = dr.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return dr
}

func (dr *DNSRecordQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *DNSRecordQuery {
	return dr
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (d *DiskQuery) CollectFields(ctx context.Context, satisfies ...string) *DiskQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		d = d.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return d
}

func (d *DiskQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *DiskQuery {
	return d
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (e *EnvironmentQuery) CollectFields(ctx context.Context, satisfies ...string) *EnvironmentQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		e = e.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return e
}

func (e *EnvironmentQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *EnvironmentQuery {
	return e
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (fd *FileDeleteQuery) CollectFields(ctx context.Context, satisfies ...string) *FileDeleteQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		fd = fd.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return fd
}

func (fd *FileDeleteQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *FileDeleteQuery {
	return fd
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (fd *FileDownloadQuery) CollectFields(ctx context.Context, satisfies ...string) *FileDownloadQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		fd = fd.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return fd
}

func (fd *FileDownloadQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *FileDownloadQuery {
	return fd
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (fe *FileExtractQuery) CollectFields(ctx context.Context, satisfies ...string) *FileExtractQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		fe = fe.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return fe
}

func (fe *FileExtractQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *FileExtractQuery {
	return fe
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (f *FindingQuery) CollectFields(ctx context.Context, satisfies ...string) *FindingQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		f = f.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return f
}

func (f *FindingQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *FindingQuery {
	return f
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (gfm *GinFileMiddlewareQuery) CollectFields(ctx context.Context, satisfies ...string) *GinFileMiddlewareQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		gfm = gfm.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return gfm
}

func (gfm *GinFileMiddlewareQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *GinFileMiddlewareQuery {
	return gfm
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (h *HostQuery) CollectFields(ctx context.Context, satisfies ...string) *HostQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		h = h.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return h
}

func (h *HostQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *HostQuery {
	return h
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (hd *HostDependencyQuery) CollectFields(ctx context.Context, satisfies ...string) *HostDependencyQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		hd = hd.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return hd
}

func (hd *HostDependencyQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *HostDependencyQuery {
	return hd
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (i *IdentityQuery) CollectFields(ctx context.Context, satisfies ...string) *IdentityQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		i = i.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return i
}

func (i *IdentityQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *IdentityQuery {
	return i
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (in *IncludedNetworkQuery) CollectFields(ctx context.Context, satisfies ...string) *IncludedNetworkQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		in = in.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return in
}

func (in *IncludedNetworkQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *IncludedNetworkQuery {
	return in
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (n *NetworkQuery) CollectFields(ctx context.Context, satisfies ...string) *NetworkQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		n = n.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return n
}

func (n *NetworkQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *NetworkQuery {
	return n
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (pl *PlanQuery) CollectFields(ctx context.Context, satisfies ...string) *PlanQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		pl = pl.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return pl
}

func (pl *PlanQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *PlanQuery {
	return pl
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (ph *ProvisionedHostQuery) CollectFields(ctx context.Context, satisfies ...string) *ProvisionedHostQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		ph = ph.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return ph
}

func (ph *ProvisionedHostQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *ProvisionedHostQuery {
	return ph
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (pn *ProvisionedNetworkQuery) CollectFields(ctx context.Context, satisfies ...string) *ProvisionedNetworkQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		pn = pn.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return pn
}

func (pn *ProvisionedNetworkQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *ProvisionedNetworkQuery {
	return pn
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (ps *ProvisioningStepQuery) CollectFields(ctx context.Context, satisfies ...string) *ProvisioningStepQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		ps = ps.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return ps
}

func (ps *ProvisioningStepQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *ProvisioningStepQuery {
	return ps
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (s *ScriptQuery) CollectFields(ctx context.Context, satisfies ...string) *ScriptQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		s = s.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return s
}

func (s *ScriptQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *ScriptQuery {
	return s
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (s *StatusQuery) CollectFields(ctx context.Context, satisfies ...string) *StatusQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		s = s.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return s
}

func (s *StatusQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *StatusQuery {
	return s
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (t *TagQuery) CollectFields(ctx context.Context, satisfies ...string) *TagQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		t = t.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return t
}

func (t *TagQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *TagQuery {
	return t
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (t *TeamQuery) CollectFields(ctx context.Context, satisfies ...string) *TeamQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		t = t.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return t
}

func (t *TeamQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *TeamQuery {
	return t
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (t *TokenQuery) CollectFields(ctx context.Context, satisfies ...string) *TokenQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		t = t.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return t
}

func (t *TokenQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *TokenQuery {
	return t
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) *UserQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		u = u.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return u
}

func (u *UserQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *UserQuery {
	return u
}
