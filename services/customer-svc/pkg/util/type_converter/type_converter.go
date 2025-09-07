package type_converter

import "database/sql"

func NewSqlNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{}
	}

	if len(*s) == 0 {
		return sql.NullString{}
	}

	return sql.NullString{
		String: *s,
		Valid:  true,
	}
}

func NewString(v sql.NullString) *string {
	if v.String == "" || !v.Valid {
		return nil
	}

	return &v.String
}
