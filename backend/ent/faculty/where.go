// Code generated by entc, DO NOT EDIT.

package faculty

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/moominzie/user-record/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Fname applies equality check predicate on the "fname" field. It's identical to FnameEQ.
func Fname(v string) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFname), v))
	})
}

// FnameEQ applies the EQ predicate on the "fname" field.
func FnameEQ(v string) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFname), v))
	})
}

// FnameNEQ applies the NEQ predicate on the "fname" field.
func FnameNEQ(v string) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFname), v))
	})
}

// FnameIn applies the In predicate on the "fname" field.
func FnameIn(vs ...string) predicate.Faculty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Faculty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFname), v...))
	})
}

// FnameNotIn applies the NotIn predicate on the "fname" field.
func FnameNotIn(vs ...string) predicate.Faculty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Faculty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFname), v...))
	})
}

// FnameGT applies the GT predicate on the "fname" field.
func FnameGT(v string) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFname), v))
	})
}

// FnameGTE applies the GTE predicate on the "fname" field.
func FnameGTE(v string) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFname), v))
	})
}

// FnameLT applies the LT predicate on the "fname" field.
func FnameLT(v string) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFname), v))
	})
}

// FnameLTE applies the LTE predicate on the "fname" field.
func FnameLTE(v string) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFname), v))
	})
}

// FnameContains applies the Contains predicate on the "fname" field.
func FnameContains(v string) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFname), v))
	})
}

// FnameHasPrefix applies the HasPrefix predicate on the "fname" field.
func FnameHasPrefix(v string) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFname), v))
	})
}

// FnameHasSuffix applies the HasSuffix predicate on the "fname" field.
func FnameHasSuffix(v string) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFname), v))
	})
}

// FnameEqualFold applies the EqualFold predicate on the "fname" field.
func FnameEqualFold(v string) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFname), v))
	})
}

// FnameContainsFold applies the ContainsFold predicate on the "fname" field.
func FnameContainsFold(v string) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFname), v))
	})
}

// HasUserInformations applies the HasEdge predicate on the "user_informations" edge.
func HasUserInformations() predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInformationsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserInformationsTable, UserInformationsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserInformationsWith applies the HasEdge predicate on the "user_informations" edge with a given conditions (other predicates).
func HasUserInformationsWith(preds ...predicate.User) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInformationsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserInformationsTable, UserInformationsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Faculty) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Faculty) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Faculty) predicate.Faculty {
	return predicate.Faculty(func(s *sql.Selector) {
		p(s.Not())
	})
}
