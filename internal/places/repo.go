package places

import (
	"database/sql"
	"fmt"
	"strings"
)

// Repo deals with places repository.
type Repo interface {
	GetByName(name string) ([]*Place, error)
}

type placeStmts struct {
	selectByName *sql.Stmt
}

// PostgresRepo is a repo implementation for working with PostgreSQL.
type PostgresRepo struct {
	db   *sql.DB
	stmt *placeStmts
}

// NewPostgresRepo creates new PostgreSQL repo.
func NewPostgresRepo(db *sql.DB) (*PostgresRepo, error) {
	stmts, err := prepareStmt(db)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statements: %v", err)
	}
	return &PostgresRepo{db, stmts}, nil
}

// GetByName find places by name (or by part of name).
func (r *PostgresRepo) GetByName(name string) ([]*Place, error) {
	name = searchable(name)
	rows, err := r.stmt.selectByName.Query(name)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer rows.Close()

	places := []*Place{}
	for rows.Next() {
		p := &Place{Coordinate: &Point{}}
		err = rows.Scan(&p.ID, &p.OSMID, &p.OSMType, &p.Name, &p.Coordinate.Lat, &p.Coordinate.Lon)
		if err != nil {
			return nil, fmt.Errorf("failed to scan rows: %v", err)
		}
		places = append(places, p)
	}

	return places, nil
}

// prepareStmt prepares all statements.
func prepareStmt(db *sql.DB) (*placeStmts, error) {
	stmts := &placeStmts{}
	var err error

	stmts.selectByName, err = db.Prepare(`
		SELECT id, osm_id, osm_type, name, coordinate[0], coordinate[1]
		FROM place
		WHERE tsv @@ to_tsquery($1)
		LIMIT 10
	`)
	if err != nil {
		return nil, fmt.Errorf("faied to prepare selectByName statement: %v", err)
	}

	return stmts, nil
}

// searchable prepares string for using in Posgres full-text search.
func searchable(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}
	words := strings.Fields(s)
	return strings.Join(words, ":* & ") + ":*"
}
