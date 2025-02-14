// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/moominzie/user-record/ent/faculty"
	"github.com/moominzie/user-record/ent/user"
)

// FacultyCreate is the builder for creating a Faculty entity.
type FacultyCreate struct {
	config
	mutation *FacultyMutation
	hooks    []Hook
}

// SetFname sets the fname field.
func (fc *FacultyCreate) SetFname(s string) *FacultyCreate {
	fc.mutation.SetFname(s)
	return fc
}

// AddUserInformationIDs adds the user_informations edge to User by ids.
func (fc *FacultyCreate) AddUserInformationIDs(ids ...int) *FacultyCreate {
	fc.mutation.AddUserInformationIDs(ids...)
	return fc
}

// AddUserInformations adds the user_informations edges to User.
func (fc *FacultyCreate) AddUserInformations(u ...*User) *FacultyCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return fc.AddUserInformationIDs(ids...)
}

// Mutation returns the FacultyMutation object of the builder.
func (fc *FacultyCreate) Mutation() *FacultyMutation {
	return fc.mutation
}

// Save creates the Faculty in the database.
func (fc *FacultyCreate) Save(ctx context.Context) (*Faculty, error) {
	if _, ok := fc.mutation.Fname(); !ok {
		return nil, &ValidationError{Name: "fname", err: errors.New("ent: missing required field \"fname\"")}
	}
	var (
		err  error
		node *Faculty
	)
	if len(fc.hooks) == 0 {
		node, err = fc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FacultyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fc.mutation = mutation
			node, err = fc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fc.hooks) - 1; i >= 0; i-- {
			mut = fc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FacultyCreate) SaveX(ctx context.Context) *Faculty {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fc *FacultyCreate) sqlSave(ctx context.Context) (*Faculty, error) {
	f, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	f.ID = int(id)
	return f, nil
}

func (fc *FacultyCreate) createSpec() (*Faculty, *sqlgraph.CreateSpec) {
	var (
		f     = &Faculty{config: fc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: faculty.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: faculty.FieldID,
			},
		}
	)
	if value, ok := fc.mutation.Fname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: faculty.FieldFname,
		})
		f.Fname = value
	}
	if nodes := fc.mutation.UserInformationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   faculty.UserInformationsTable,
			Columns: []string{faculty.UserInformationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return f, _spec
}
