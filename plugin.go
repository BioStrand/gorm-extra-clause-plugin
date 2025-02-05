package gormextraclauseplugin

import "gorm.io/gorm"

// ExtraClausePlugin support plugin that not supported clause by gorm
type ExtraClausePlugin struct{}

// Name return plugin name
func (e *ExtraClausePlugin) Name() string {
	return "ExtraClausePlugin"
}

// Initialize register BuildClauses
func (e *ExtraClausePlugin) Initialize(db *gorm.DB) error {
	additionalClauses := []string{"WITH", "UNION", "INTERSECT", "EXCEPT"}
	db.Callback().Query().Clauses = append(db.Callback().Query().Clauses, additionalClauses...)
	db.Callback().Row().Clauses = append(db.Callback().Row().Clauses, additionalClauses...)
	db.Callback().Update().Clauses = append(db.Callback().Update().Clauses, additionalClauses...)

	return nil
}

// New create new ExtraClausePlugin
//
//	// example
//	db.Use(extraClausePlugin.New())
func New() *ExtraClausePlugin {
	return &ExtraClausePlugin{}
}
