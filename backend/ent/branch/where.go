// Code generated by entc, DO NOT EDIT.

package branch

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/moominzie/user-record/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Brname applies equality check predicate on the "brname" field. It's identical to BrnameEQ.
func Brname(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBrname), v))
	})
}

// BrnameEQ applies the EQ predicate on the "brname" field.
func BrnameEQ(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBrname), v))
	})
}

// BrnameNEQ applies the NEQ predicate on the "brname" field.
func BrnameNEQ(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBrname), v))
	})
}

// BrnameIn applies the In predicate on the "brname" field.
func BrnameIn(vs ...string) predicate.Branch {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Branch(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBrname), v...))
	})
}

// BrnameNotIn applies the NotIn predicate on the "brname" field.
func BrnameNotIn(vs ...string) predicate.Branch {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Branch(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBrname), v...))
	})
}

// BrnameGT applies the GT predicate on the "brname" field.
func BrnameGT(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBrname), v))
	})
}

// BrnameGTE applies the GTE predicate on the "brname" field.
func BrnameGTE(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBrname), v))
	})
}

// BrnameLT applies the LT predicate on the "brname" field.
func BrnameLT(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBrname), v))
	})
}

// BrnameLTE applies the LTE predicate on the "brname" field.
func BrnameLTE(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBrname), v))
	})
}

// BrnameContains applies the Contains predicate on the "brname" field.
func BrnameContains(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBrname), v))
	})
}

// BrnameHasPrefix applies the HasPrefix predicate on the "brname" field.
func BrnameHasPrefix(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBrname), v))
	})
}

// BrnameHasSuffix applies the HasSuffix predicate on the "brname" field.
func BrnameHasSuffix(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBrname), v))
	})
}

// BrnameEqualFold applies the EqualFold predicate on the "brname" field.
func BrnameEqualFold(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBrname), v))
	})
}

// BrnameContainsFold applies the ContainsFold predicate on the "brname" field.
func BrnameContainsFold(v string) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBrname), v))
	})
}

// HasUserInformations applies the HasEdge predicate on the "user_informations" edge.
func HasUserInformations() predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInformationsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserInformationsTable, UserInformationsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserInformationsWith applies the HasEdge predicate on the "user_informations" edge with a given conditions (other predicates).
func HasUserInformationsWith(preds ...predicate.User) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
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
func And(predicates ...predicate.Branch) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Branch) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
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
func Not(p predicate.Branch) predicate.Branch {
	return predicate.Branch(func(s *sql.Selector) {
		p(s.Not())
	})
}
