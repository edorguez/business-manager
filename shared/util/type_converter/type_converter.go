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

func NewSqlNullInt32(v *int32) sql.NullInt32 {
	if v == nil {
		return sql.NullInt32{}
	}

	return sql.NullInt32{
		Int32: *v,
		Valid: true,
	}
}

func NewInt32(v sql.NullInt32) *int32 {
	if !v.Valid {
		return nil
	}

	return &v.Int32
}

func NewSqlNullBool(v *bool) sql.NullBool {
	if v == nil {
		return sql.NullBool{}
	}

	return sql.NullBool{
		Bool:  *v,
		Valid: true,
	}
}

func NewBool(v sql.NullBool) *bool {
	if !v.Valid {
		return nil
	}

	return &v.Bool
}
