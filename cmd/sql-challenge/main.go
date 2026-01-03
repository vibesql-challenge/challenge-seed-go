// SQL Database Challenge - Go Starter
//
// This is a minimal REPL skeleton that implements the protocol expected by
// the SQLLogicTest runner. Your job is to replace the "not implemented"
// responses with actual SQL execution.
//
// Protocol:
// - Read SQL from stdin (one statement at a time, ending with semicolon)
// - After receiving a blank line, execute the accumulated SQL
// - Output results as tab-separated values, one row per line
// - Output a blank line to signal end of results
// - For errors, output "Error: <message>" then a blank line
//
// See CLAUDE.md for implementation tips and suggested approach.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var sqlBuffer strings.Builder

	for scanner.Scan() {
		line := scanner.Text()

		// Empty line signals end of statement - time to execute
		if line == "" {
			if sqlBuffer.Len() > 0 {
				executeSQL(strings.TrimSpace(sqlBuffer.String()))
				sqlBuffer.Reset()
			}
			continue
		}

		// Accumulate SQL
		if sqlBuffer.Len() > 0 {
			sqlBuffer.WriteString(" ")
		}
		sqlBuffer.WriteString(line)
	}

	// Handle any remaining SQL
	if sqlBuffer.Len() > 0 {
		executeSQL(strings.TrimSpace(sqlBuffer.String()))
	}
}

// executeSQL parses and executes a SQL statement.
//
// This is where you implement your database!
//
// For successful queries, output:
// - Result rows as tab-separated values
// - One row per line
// - Blank line to signal end
//
// For successful statements (CREATE, INSERT, etc.):
// - Just output a blank line
//
// For errors:
// - Output "Error: <message>"
// - Then a blank line
func executeSQL(sql string) {
	// Remove trailing semicolon for parsing
	sql = strings.TrimSpace(sql)
	sql = strings.TrimSuffix(sql, ";")
	sql = strings.TrimSpace(sql)

	if sql == "" {
		// Empty statement - just acknowledge
		fmt.Println()
		return
	}

	// TODO: Implement your SQL database here!
	//
	// Suggested packages to create:
	//
	// pkg/parser/    - Parse SQL into AST
	// pkg/types/     - SQL types: Integer, Real, Text, Null
	// pkg/storage/   - In-memory table storage
	// pkg/executor/  - Execute queries
	//
	// Start with these milestones:
	//
	// 1. SELECT <literal>
	//    - "SELECT 1" -> outputs "1"
	//    - "SELECT 'hello'" -> outputs "hello"
	//
	// 2. CREATE TABLE, INSERT, basic SELECT
	//    - Store tables in map[string]*Table
	//    - Table = []Row, Row = []Value
	//
	// 3. WHERE clauses
	//    - Filter rows based on predicates
	//
	// 4. JOINs
	//    - Start with CROSS JOIN, then INNER JOIN
	//
	// 5. Aggregates
	//    - COUNT, SUM, AVG, MIN, MAX
	//    - GROUP BY, HAVING
	//
	// 6. Subqueries
	//    - Scalar subqueries in SELECT
	//    - IN (subquery)
	//    - EXISTS
	//
	// For now, we just return "not implemented" for everything:

	fmt.Printf("Error: not implemented - %s\n", firstWord(sql))
	fmt.Println()
}

// firstWord returns the first word of a SQL statement (for error messages)
func firstWord(sql string) string {
	sql = strings.TrimSpace(sql)
	if idx := strings.Index(sql, " "); idx != -1 {
		return sql[:idx]
	}
	return sql
}
