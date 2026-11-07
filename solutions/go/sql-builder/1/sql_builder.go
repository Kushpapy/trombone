package sqlbuilder
import "strings"


// ColumnValue represents a single database column and its corresponding value
// to be used in SQL INSERT statements.
type ColumnValue struct {
	Column string
	Value  string
}

// BuildSelectQuery generates a valid SQL SELECT query for the specified table and columns.
// If the columns slice is empty or nil, it defaults to the wildcard "*".
func BuildSelectQuery(table string, columns []string) string {
	cols := "*"
    if len(columns) > 0 {
        cols = strings.Join(columns, ", ")
    }

    var sb strings.Builder 
    sb.WriteString("SELECT ")
    sb.WriteString(cols)
    sb.WriteString(" FROM ")
    sb.WriteString(table)
    sb.WriteString(";")

    return sb.String()
}

// BuildBatchInserts generates a slice of individual SQL INSERT statements for the provided rows.
// Empty rows are skipped, and values must be enclosed in single quotes.
func BuildBatchInserts(table string, rows [][]ColumnValue) []string {
	var result [] string
    for _, row := range rows {
        if len(row) == 0{
            continue
        }
        
    var cols, vals [] string  
    for _, cv := range row {
        cols = append(cols, cv.Column)
        vals = append(vals, "'" + cv.Value + "'")
    }

    var sb strings.Builder 
    sb.WriteString("INSERT INTO ")
    sb.WriteString(table)
    sb.WriteString(" (")
    sb.WriteString(strings.Join(cols, ", "))
    sb.WriteString(") VALUES (")
    sb.WriteString(strings.Join(vals, ", "))
    sb.WriteString(")")
    sb.WriteString(";")

    result = append(result, sb.String())
    }
    

    
    return result
}
