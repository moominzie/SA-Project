// Code generated by entc, DO NOT EDIT.

package employee

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/moominzie/user-record/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// EmployeeEmail applies equality check predicate on the "employeeEmail" field. It's identical to EmployeeEmailEQ.
func EmployeeEmail(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmployeeEmail), v))
	})
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassword), v))
	})
}

// EmployeeEmailEQ applies the EQ predicate on the "employeeEmail" field.
func EmployeeEmailEQ(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmployeeEmail), v))
	})
}

// EmployeeEmailNEQ applies the NEQ predicate on the "employeeEmail" field.
func EmployeeEmailNEQ(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEmployeeEmail), v))
	})
}

// EmployeeEmailIn applies the In predicate on the "employeeEmail" field.
func EmployeeEmailIn(vs ...string) predicate.Employee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Employee(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEmployeeEmail), v...))
	})
}

// EmployeeEmailNotIn applies the NotIn predicate on the "employeeEmail" field.
func EmployeeEmailNotIn(vs ...string) predicate.Employee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Employee(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEmployeeEmail), v...))
	})
}

// EmployeeEmailGT applies the GT predicate on the "employeeEmail" field.
func EmployeeEmailGT(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEmployeeEmail), v))
	})
}

// EmployeeEmailGTE applies the GTE predicate on the "employeeEmail" field.
func EmployeeEmailGTE(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEmployeeEmail), v))
	})
}

// EmployeeEmailLT applies the LT predicate on the "employeeEmail" field.
func EmployeeEmailLT(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEmployeeEmail), v))
	})
}

// EmployeeEmailLTE applies the LTE predicate on the "employeeEmail" field.
func EmployeeEmailLTE(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEmployeeEmail), v))
	})
}

// EmployeeEmailContains applies the Contains predicate on the "employeeEmail" field.
func EmployeeEmailContains(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEmployeeEmail), v))
	})
}

// EmployeeEmailHasPrefix applies the HasPrefix predicate on the "employeeEmail" field.
func EmployeeEmailHasPrefix(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEmployeeEmail), v))
	})
}

// EmployeeEmailHasSuffix applies the HasSuffix predicate on the "employeeEmail" field.
func EmployeeEmailHasSuffix(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEmployeeEmail), v))
	})
}

// EmployeeEmailEqualFold applies the EqualFold predicate on the "employeeEmail" field.
func EmployeeEmailEqualFold(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEmployeeEmail), v))
	})
}

// EmployeeEmailContainsFold applies the ContainsFold predicate on the "employeeEmail" field.
func EmployeeEmailContainsFold(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEmployeeEmail), v))
	})
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassword), v))
	})
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPassword), v))
	})
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.Employee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Employee(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPassword), v...))
	})
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.Employee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Employee(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPassword), v...))
	})
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPassword), v))
	})
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPassword), v))
	})
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPassword), v))
	})
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPassword), v))
	})
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPassword), v))
	})
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPassword), v))
	})
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPassword), v))
	})
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPassword), v))
	})
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPassword), v))
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Employee) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Employee) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
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
func Not(p predicate.Employee) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		p(s.Not())
	})
}
