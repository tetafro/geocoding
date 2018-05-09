package places

import (
	"database/sql"
	"fmt"
	"strings"
)

const (
	// Muximum number of places for one request.
	placesLimit = 10
)

// Repo deals with places repository.
type Repo interface {
	GetByFullname(fullname string) ([]*Place, error)
}

type placeStmts struct {
	selectByFullname *sql.Stmt
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

// GetByFullname find places by fullname (or by part of fullname).
func (r *PostgresRepo) GetByFullname(fullname string) ([]*Place, error) {
	fullname = searchable(fullname)
	rows, err := r.stmt.selectByFullname.Query(fullname, placesLimit)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer rows.Close()

	places := []*Place{}
	for rows.Next() {
		p := &Place{Coordinate: &Point{}}
		err = rows.Scan(
			&p.ID, &p.OSMID,
			&p.Country, &p.City, &p.Street, &p.Housenumber, &p.Name,
			&p.Fullname,
			&p.Coordinate.Lat, &p.Coordinate.Lon,
		)
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

	stmts.selectByFullname, err = db.Prepare(`
		SELECT id, osm_id,
			country, city, street, housenumber, name,
			fullname,
			coordinate[0], coordinate[1]
		FROM place
		WHERE tsv_fullname @@ to_tsquery($1)
		LIMIT $2
	`)
	if err != nil {
		return nil, fmt.Errorf("faied to prepare selectByFullname statement: %v", err)
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
